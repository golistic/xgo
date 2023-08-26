/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xsql

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/golistic/xgo/xstrings"
)

var reDSNPassword = regexp.MustCompile(`^(.*):[^/]+?(@.*)$`)
var dsnPasswordMask = "********"

var reDSN = regexp.MustCompile(`(.*?)(?::(.*?))?@(\w+)\((.*?)\)(?:/([^?]+))?/?(\?)?(.*)?`)

type DataSource struct {
	Driver   string
	User     string
	Password string
	Protocol string
	Address  string
	Schema   string
	Options  url.Values
}

// String returns a simplified representation of d. Simplified means that
// only username, protocol, address, and schema (if available) are included. The password is
// obfuscated.
func (d *DataSource) String() string {

	n := fmt.Sprintf("%s:%s@%s(%s)", d.User, dsnPasswordMask, d.Protocol, d.Address)
	if d.Schema != "" {
		n += "/" + d.Schema
	} else {
		n += "/"
	}

	return n
}

func (d *DataSource) Format() string {

	var query string

	if len(d.Options) != 0 {
		query = "?" + d.Options.Encode()
	}

	n := fmt.Sprintf("%s:%s@%s(%s)",
		d.User, d.Password, d.Protocol, d.Address)
	if d.Schema != "" {
		n += "/" + d.Schema
	} else {
		n += "/"
	}

	if query != "" {
		n += query
	}

	return n
}

// ParseDSN parsers the name as a data source name (DSN).
func ParseDSN(name string) (*DataSource, error) {
	errMsg := "invalid data source name (%w)"

	m := reDSN.FindAllStringSubmatch(name, -1)
	if m == nil {
		return nil, fmt.Errorf(errMsg, fmt.Errorf("could not parse"))
	}

	protocol := strings.ToLower(m[0][3])
	if !(protocol == "unix" || protocol == "tcp") {
		return nil, fmt.Errorf(errMsg, fmt.Errorf("unsupported protocol '%s'", m[0][3]))
	}

	if err := validateSchema(m[0][5]); err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	ds := &DataSource{
		User:     m[0][1],
		Password: m[0][2],
		Protocol: protocol,
		Address:  m[0][4],
		Schema:   m[0][5],
		Options:  url.Values{},
	}

	if xstrings.SliceHas(m[0], "?") {
		var err error
		ds.Options, err = url.ParseQuery(m[0][len(m[0])-1])
		if err != nil {
			return nil, fmt.Errorf(errMsg, fmt.Errorf("could not parse query part"))
		}
	}

	return ds, nil
}

// ReplaceDSNDatabase takes dsn and replaces the database name. It
// returns the new Data Source Name.
func ReplaceDSNDatabase(dsn string, name string) (string, error) {
	ds, err := ParseDSN(dsn)
	if err != nil {
		return "", err
	}

	ds.Schema = name

	newDSN := ds.Format()
	if _, err := ParseDSN(newDSN); err != nil {
		return "", err
	}

	return newDSN, nil
}

// SetDSNOptions sets parameters for the given dsn.
// This function will parse dsn, add options, and return a new formatted DSN.
func SetDSNOptions(dsn string, options map[string]string) (string, error) {
	ds, err := ParseDSN(dsn)
	if err != nil {
		return "", err
	}

	for k, v := range options {
		ds.Options.Set(k, v)
	}

	return ds.Format(), nil
}

// MaskPasswordInDSN masks the password within the MySQL data source name dsn. This
// function is usually used when displaying or logging the DSN.
//
// When password is empty (not provided) the mask is added anyway.
// When the DSN is something that was not a DSN, the mask itself is returned to
// prevent possible mistakes.
func MaskPasswordInDSN(dsn string) string {
	var res = dsn

	if reDSNPassword.MatchString(dsn) {
		res = reDSNPassword.ReplaceAllString(dsn, `$1:`+dsnPasswordMask+`$2`)
	} else {
		if strings.Contains(dsn, ":@") {
			res = strings.Replace(dsn, ":@", ":"+dsnPasswordMask+"@", 1)
		} else {
			res = strings.Replace(dsn, "@", ":"+dsnPasswordMask+"@", 1)
		}
	}

	if res == dsn {
		return dsnPasswordMask
	}

	return res
}

func validateSchema(name string) error {
	if strings.ContainsAny(name, " \t\n\r\f\v") {
		return fmt.Errorf("schema contains whitespace")
	}

	return nil
}

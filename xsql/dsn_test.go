/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xsql_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/golistic/xgo/xsql"
	"github.com/golistic/xgo/xt"
)

func TestDataSource_String(t *testing.T) {
	t.Run("password is obfuscated", func(t *testing.T) {
		var dsn = "u:pwd@tcp(127.0.0.1:3306)/schemaName?useTLS=true"
		var exp = "u:********@tcp(127.0.0.1:3306)/schemaName"

		got, err := xsql.ParseDSN(dsn)
		xt.OK(t, err)
		xt.Eq(t, exp, got.String())
	})

	t.Run("without schema", func(t *testing.T) {
		var dsn = "u:pwd@tcp(127.0.0.1:3306)?useTLS=true"
		var exp = "u:********@tcp(127.0.0.1:3306)/"

		got, err := xsql.ParseDSN(dsn)
		xt.OK(t, err)
		xt.Eq(t, exp, got.String())
	})
}

func TestDataSource_Format(t *testing.T) {
	t.Run("with schema", func(t *testing.T) {
		var dsn = "u:pwd@tcp(127.0.0.1:3306)/schemaName?useTLS=true&parseTime=true"
		var exp = "u:pwd@tcp(127.0.0.1:3306)/schemaName?parseTime=true&useTLS=true"

		got, err := xsql.ParseDSN(dsn)
		xt.OK(t, err)
		xt.Eq(t, exp, got.Format())
	})

	t.Run("without schema", func(t *testing.T) {
		var dsn = "u:pwd@tcp(127.0.0.1:3306)?useTLS=true&parseTime=true"
		var exp = "u:pwd@tcp(127.0.0.1:3306)/?parseTime=true&useTLS=true"

		got, err := xsql.ParseDSN(dsn)
		xt.OK(t, err)
		xt.Eq(t, exp, got.Format())
	})
}

func TestParseDSN(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := xsql.ParseDSN("not a dsn")
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (could not parse)", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "could not parse", err.Error())
	})

	t.Run("unsupported protocol", func(t *testing.T) {
		_, err := xsql.ParseDSN("u:p@udp(127.0.0.1)/")
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (unsupported protocol 'udp')", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "unsupported protocol 'udp'", err.Error())
	})

	t.Run("invalid query string", func(t *testing.T) {
		_, err := xsql.ParseDSN("u:p@tcp(127.0.0.1)/?bad;query") // semicolon considered invalid
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (could not parse query part)", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "could not parse query part", err.Error())
	})
}

func TestReplaceDSNDatabase(t *testing.T) {

	var baseDSN = "u:pwd@tcp(127.0.0.1:3306)/"

	t.Run("invalid", func(t *testing.T) {
		_, err := xsql.ReplaceDSNDatabase("not a dsn", "something")
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (could not parse)", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "could not parse", err.Error())
	})

	t.Run("schema name makes DSN invalid", func(t *testing.T) {
		_, err := xsql.ReplaceDSNDatabase("u:p@tcp(127.0.0.1)/schemaName?useTLS=true", "foo bar")
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (schema contains whitespace)", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "schema contains whitespace", err.Error())
	})

	t.Run("remove database name from DSN", func(t *testing.T) {
		exp := baseDSN

		dsn, err := xsql.ReplaceDSNDatabase(baseDSN+"foo", "")
		xt.OK(t, err)
		xt.Eq(t, exp, dsn)
	})

	t.Run("using DSN with database name", func(t *testing.T) {
		got := baseDSN + "bar"
		exp := baseDSN + "foo"

		dsn, err := xsql.ReplaceDSNDatabase(got, "foo")

		xt.OK(t, err)
		xt.Eq(t, exp, dsn)
	})
}

func TestSetDSNOptions(t *testing.T) {

	var baseDSN = "u:pwd@tcp(127.0.0.1:3306)/"

	t.Run("set some parameter", func(t *testing.T) {
		dsn, err := xsql.SetDSNOptions(baseDSN, map[string]string{
			"fooBar": "1234",
		})
		xt.OK(t, err)
		xt.Eq(t, "u:pwd@tcp(127.0.0.1:3306)/?fooBar=1234", dsn)
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := xsql.SetDSNOptions("not a dsn", map[string]string{
			"fooBar": "1234",
		})
		xt.KO(t, err)
		xt.Eq(t, "invalid data source name (could not parse)", err.Error())

		err = errors.Unwrap(err)
		xt.Eq(t, "could not parse", err.Error())
	})
}

func TestMaskPasswordInDSN(t *testing.T) {

	var baseDSN = "u:pwd@tcp(127.0.0.1:3306)/"
	var mask = "********"

	t.Run("mask password", func(t *testing.T) {
		got := baseDSN
		exp := strings.Replace(got, ":pwd", ":"+mask, 1)

		xt.Eq(t, exp, xsql.MaskPasswordInDSN(got))
	})

	t.Run("dsn got empty password", func(t *testing.T) {
		got := strings.Replace(baseDSN, ":pwd", ":", 1)
		exp := strings.Replace(got, ":@", ":"+mask+"@", 1)

		xt.Eq(t, exp, xsql.MaskPasswordInDSN(got))
	})

	t.Run("dsn got no password", func(t *testing.T) {
		got := strings.Replace(baseDSN, ":pwd", "", 1)
		exp := strings.Replace(got, "@", ":"+mask+"@", 1)

		xt.Eq(t, exp, xsql.MaskPasswordInDSN(got))
	})

	t.Run("something not DSN", func(t *testing.T) {
		xt.Eq(t, mask, xsql.MaskPasswordInDSN("foobar"))
	})
}

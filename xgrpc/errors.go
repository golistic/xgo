/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xgrpc

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrServerUnavailable = errors.New("gRPC server not available")
var ErrServiceUnavailable = errors.New("gRPC service not available")

// ErrorFromRPC cleans up errors returned by the proto package. Errors with
// codes.Unavailable will have ErrServerUnavailable wrapped into the error.
func ErrorFromRPC(err error) error {

	st, ok := status.FromError(err)
	if !ok {
		return errors.New(strings.TrimSpace(strings.TrimPrefix(err.Error(), "proto:")))
	}

	m := st.Message()

	if st.Code() == codes.Unavailable {
		m := st.Message()
		if m[len(m)-1] == '"' {
			m = m[strings.Index(m, "\"")+1 : len(m)-1]
		}
		m = strings.TrimSpace(strings.TrimPrefix(m, "transport:"))
		m = strings.TrimSpace(strings.TrimPrefix(m, "Error"))

		return fmt.Errorf("%w (%s)", ErrServerUnavailable, fmt.Errorf(m))
	}

	m = strings.TrimSpace(strings.TrimPrefix(m, "proto:"))

	return fmt.Errorf(m)
}

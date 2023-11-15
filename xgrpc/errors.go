/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xgrpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrServerUnavailable = errors.New("gRPC server not available")
var ErrServiceUnavailable = errors.New("gRPC service not available")

func ErrorFromRPC(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	if st.Code() == codes.Unavailable {
		return ErrServerUnavailable
	}

	return fmt.Errorf(st.Message())
}

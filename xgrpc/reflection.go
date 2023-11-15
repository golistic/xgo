/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xgrpc

import (
	"context"

	"google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection/grpc_reflection_v1"
)

// CheckServiceAvailability connects using address and checks if the gRPC symbol is
// available. The serviceSymbol must be of the format `<package>.<service>`.
//
// When the method is not available, error ErrServiceUnavailable is returned.
//
// The opts argument can be used pass options for the grpc.Dial function.
func CheckServiceAvailability(address, serviceSymbol string, opts ...grpc.DialOption) error {

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return ErrorFromRPC(err)
	}

	client := reflection.NewServerReflectionClient(conn)

	stream, err := client.ServerReflectionInfo(context.Background())
	if err != nil {
		return ErrorFromRPC(err)
	}

	err = stream.Send(&reflection.ServerReflectionRequest{
		Host:           address,
		MessageRequest: &reflection.ServerReflectionRequest_ListServices{},
	})
	if err != nil {
		return ErrorFromRPC(err)
	}

	res, err := stream.Recv()
	if err != nil {
		return ErrorFromRPC(err)
	}

	for _, s := range res.GetListServicesResponse().GetService() {
		if s.Name == serviceSymbol {
			return nil
		}
	}

	return ErrServiceUnavailable
}

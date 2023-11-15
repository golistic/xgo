/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xgrpc_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/golistic/xgo/xgrpc"
	"github.com/golistic/xgo/xgrpc/testprotos/services/v1"
	"github.com/golistic/xgo/xnet"
	"github.com/golistic/xgo/xt"
)

func TestCheckServiceAvailability(t *testing.T) {
	portAAA, err := xnet.GetTCPPort("")
	xt.OK(t, err)
	portBBB, err := xnet.GetTCPPort("")
	xt.OK(t, err)
	portZZZ, err := xnet.GetTCPPort("")
	xt.OK(t, err)

	addrAAA := fmt.Sprintf("127.0.0.1:%d", portAAA)
	addrBBB := fmt.Sprintf("127.0.0.1:%d", portBBB)
	addrZZZ := fmt.Sprintf("127.0.0.1:%d", portZZZ)

	serverAAA := grpc.NewServer()
	reflection.Register(serverAAA)
	services.RegisterAAAServiceServer(serverAAA, &AAAServiceServer{})
	stopAAA := make(chan struct{}, 1)
	var errAAA error
	go func() { errAAA = startServer(serverAAA, addrAAA, stopAAA) }()
	defer func() { stopAAA <- struct{}{} }()

	serverBBB := grpc.NewServer()
	reflection.Register(serverBBB)
	services.RegisterBBBServiceServer(serverBBB, &BBBServiceServer{})
	stopBBB := make(chan struct{}, 1)
	var errBBB error
	go func() { errBBB = startServer(serverBBB, addrBBB, stopBBB) }()
	defer func() { stopBBB <- struct{}{} }()

	// service without reflection
	serverZZZ := grpc.NewServer()
	services.RegisterBBBServiceServer(serverZZZ, &BBBServiceServer{})
	stopZZZ := make(chan struct{}, 1)
	var errZZZ error
	go func() { errZZZ = startServer(serverZZZ, addrZZZ, stopZZZ) }()
	defer func() { stopZZZ <- struct{}{} }()

	time.Sleep(time.Second)
	xt.OK(t, errAAA, "serverAAA")
	xt.OK(t, errBBB, "serverBBB")
	xt.OK(t, errZZZ, "serverZZZ")

	t.Run("serverAAA has AAAService", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrAAA, "services.AAAService",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.OK(t, err)
	})

	t.Run("serverAAA does not have BogusService", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrAAA, "services.BogusService",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.Assert(t, errors.Is(err, xgrpc.ErrServiceUnavailable))
	})

	t.Run("serverAAA does not not registered BBBService", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrAAA, "services.BBBService",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.Assert(t, errors.Is(err, xgrpc.ErrServiceUnavailable))
	})

	t.Run("serverBBBB has BBBService", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrBBB, "services.BBBService",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.OK(t, err)
	})

	t.Run("use unavailable gRPC server", func(t *testing.T) {
		port, err := xnet.GetTCPPort("")
		xt.OK(t, err)

		err = xgrpc.CheckServiceAvailability(fmt.Sprintf("127.0.0.1:%d", port),
			"some.Service",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.Assert(t, errors.Is(err, xgrpc.ErrServerUnavailable))
	})

	t.Run("does not work without reflection", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrZZZ, "services.BBBService",
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		xt.KO(t, err)
		xt.Eq(t, "unknown service grpc.reflection.v1.ServerReflection", err.Error())
	})

	t.Run("without explicit credentials fails", func(t *testing.T) {
		err := xgrpc.CheckServiceAvailability(addrZZZ, "services.BBBService")
		xt.KO(t, err)
		xt.Assert(t, strings.HasPrefix(err.Error(), "grpc: no transport security set"))
	})
}

type AAAServiceServer struct {
	services.UnimplementedAAAServiceServer
}

func (A AAAServiceServer) Method1(ctx context.Context, request *services.Method1Request) (*services.Method1Reply, error) {
	return &services.Method1Reply{Ok: true}, nil
}

type BBBServiceServer struct {
	services.UnimplementedBBBServiceServer
}

func (B BBBServiceServer) MethodB(ctx context.Context, request *services.MethodBRequest) (*services.MethodBReply, error) {
	return &services.MethodBReply{Ok: true}, nil
}

func startServer(server *grpc.Server, addr string, stop chan struct{}) error {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var serveErr error
	go func() {
		serveErr = server.Serve(listen)
	}()

	time.Sleep(1 * time.Second)

	if serveErr != nil {
		_ = listen.Close()
		return err
	}

	<-stop
	server.Stop()
	_ = listen.Close()

	return nil
}

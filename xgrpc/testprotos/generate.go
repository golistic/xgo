//go:generate protoc -I=. --go_out=services/v1 --go-grpc_out=services/v1 service_aaa.proto service_bbb.proto

package testprotos

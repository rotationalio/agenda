package api

//go:generate protoc -I=$GOPATH/src/github.com/rotationalio/agenda/proto --go_out=. --go_opt=module=github.com/rotationalio/agenda/pkg/api/v1 --go-grpc_out=. --go-grpc_opt=module=github.com/rotationalio/agenda/pkg/api/v1 v1/agenda.proto

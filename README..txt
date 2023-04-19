
protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto 
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/alperergul/gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/alperergul/gRPC greet/proto/dummy.proto
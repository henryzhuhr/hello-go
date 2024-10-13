module github.com/henryzhuhr/hello-go

go 1.23.0

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.67.1

require (
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.35.1
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)

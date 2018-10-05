Checkout Code:
go get -u github.com/mhsundar/grpcmath

Build Proto File:
protoc -I math/ -I${GOPATH}/src --go_out=plugins=grpc:math math/math.proto

Compile Server:
go build -i -v -o bin/server github.com/mhsundar/grpcmath/server

Compile Client:
go build -i -v -o bin/client github.com/mhsundar/grpcmath/client

Launch Server:
./bin/server

Launch Client:
./bin/client

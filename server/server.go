// main server file

package main

import (
	"log"
	"net"

	"github.com/mhsundar/grpcmath/math"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const port = ":9000"

// main works as the entry point for starting the math service
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	math.RegisterMathServer(s, new(MathService))
	log.Println("Starting server on port " + port)
	s.Serve(lis)

}

// MathService implements the Math interface
type MathService struct{}

// Add adds tow integers and return a float result
func (c *MathService) Add(ctx context.Context, params *math.Parameters) (*math.Result, error) {
	result := math.Add(params.Param1, params.Param2)
	return &math.Result{Result: result}, nil
}

// Subtract substracts two integers and return a float result
func (c *MathService) Subtract(ctx context.Context, params *math.Parameters) (*math.Result, error) {
	result := math.Subtract(params.Param1, params.Param2)
	return &math.Result{Result: result}, nil
}

// Multiply multiplies two integers and return a float result
func (c *MathService) Multiply(ctx context.Context, params *math.Parameters) (*math.Result, error) {
	result := math.Multiply(params.Param1, params.Param2)
	return &math.Result{Result: result}, nil
}

// Divide divides two integers and return a float result
func (c *MathService) Divide(ctx context.Context, params *math.Parameters) (*math.Result, error) {
	result, err := math.Divide(params.Param1, params.Param2)
	return &math.Result{Result: result}, err
}

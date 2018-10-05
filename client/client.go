package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/mhsundar/grpcmath/math"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":9000"

// main works as the main entry point of the client cli
func main() {

	newline := "\n"
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Server Address [Defaults to localhost:9000] : ")
	address, _ := reader.ReadString('\n')
	add := strings.TrimSuffix(address, newline)
	if strings.Compare(add, "") == 0 {
		add = "localhost:9000"
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(add, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// close connection later
	defer conn.Close()
	client := math.NewMathClient(conn)

	for {
		fmt.Println("\n1. Add\n2. Subtract\n3. Multiply\n4. Divide\n5. Exit")
		fmt.Printf("Choose Operation: ")
		operation, _ := reader.ReadString('\n')
		op, _ := strconv.Atoi(strings.TrimSuffix(operation, newline))

		var loperand, roperand int

		if op >= 1 && op <= 4 {
			fmt.Printf("Enter Operand1: ")
			param1, _ := reader.ReadString('\n')
			loperand, _ = strconv.Atoi(strings.TrimSuffix(param1, newline))

			fmt.Printf("Enter Operand2: ")
			param2, _ := reader.ReadString('\n')
			roperand, _ = strconv.Atoi(strings.TrimSuffix(param2, newline))
		}

		switch op {
		case 1:

			res, err := client.Add(context.Background(),
				&math.Parameters{Param1: int32(loperand),
					Param2: int32(roperand)})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Result)

		case 2:
			res, err := client.Subtract(context.Background(),
				&math.Parameters{Param1: int32(loperand),
					Param2: int32(roperand)})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Result)

		case 3:
			res, err := client.Multiply(context.Background(),
				&math.Parameters{Param1: int32(loperand),
					Param2: int32(roperand)})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Result)

		case 4:
			res, err := client.Divide(context.Background(),
				&math.Parameters{Param1: int32(loperand),
					Param2: int32(roperand)})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Result)

		case 5:
			fmt.Println("Exiting")
			os.Exit(1)

		default:
			fmt.Println("Operation must be one of the following (add, subtract, multiply, divide)")
		}
	}
}

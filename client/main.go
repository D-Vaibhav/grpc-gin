package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vaibhav/grpc_gin/protos"
	"google.golang.org/grpc"
)

const toBase = 10
const bitSize = 64
const port = ":8000"

func main() {
	// dialing for server connection
	clientConnection, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// establish client
	client := protos.NewSubtractDivideClient(clientConnection)

	// for API end-points we're using gin framework
	engine := gin.Default()

	engine.GET("/difference/:numberFirst/:numberSecond", func(ctx *gin.Context) {
		// getting data from request
		numberFirstAsString, numberSecondAsString := ctx.Param("numberFirst"), ctx.Param("numberSecond")

		// parsing uri data
		numberFirst, err := strconv.ParseUint(numberFirstAsString, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter for first number"})
		}
		numberSecond, err := strconv.ParseUint(numberSecondAsString, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter for second number"})
		}

		// assigning data to the request struct
		clientRequest := &protos.Request{NumberFirst: int64(numberFirst), NumberSecond: int64(numberSecond)}

		// calling server side implemented client interface functions
		responseFromServer, err := client.CalculateDifference(ctx, clientRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(responseFromServer.CalculatedAnswer),
		})
	})

	engine.GET("/product/:numberFirst/:numberSecond", func(ctx *gin.Context) {
		numberFirstAsString, numberSecondAsString := ctx.Param("numberFirst"), ctx.Param("numberSecond")

		numberFirst, err := strconv.ParseUint(numberFirstAsString, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter for first number"})
		}
		numberSecond, err := strconv.ParseUint(numberSecondAsString, toBase, bitSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter for second number"})
		}

		clientRequest := &protos.Request{NumberFirst: int64(numberFirst), NumberSecond: int64(numberSecond)}

		responseFromServer, err := client.CalculateProduct(ctx, clientRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(responseFromServer.CalculatedAnswer),
		})
	})

	// running the client engine
	err = engine.Run(port)
	if err != nil {
		log.Fatalf("failed to run Server: %v", err)
	}

	fmt.Println("client server is running")
}

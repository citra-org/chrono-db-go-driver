package main

import (
	"fmt"
	"net/http"

	"github.com/citra-org/chrono-db-go-driver/client"
	"github.com/gin-gonic/gin"
)

var dbClient *client.Client
var dbName string

func main() {
	uri := "itlg://admin:D!EO$H2i!MbIuZy8@127.0.0.1:3141/test11"
	var err error

	dbClient, dbName, err = client.Connect(uri)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer dbClient.Close()

	r := gin.Default()
	// r.GET("/c", handleCreate)
	r.POST("/w/:stream", handleWrite)
	r.GET("/r/:stream", handleRead)
	r.GET("/cs/:stream", handleCreateStream)
	r.GET("/ds/:stream", handleDeleteStream)

	fmt.Println("Server listening on port 3000")
	err = r.Run(":3002")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// func handleCreate(c *gin.Context) {
// 	err := dbClient.CreateChrono(dbName)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating record: %s", err)})
// 		return
// 	}
// 	c.String(http.StatusOK, "Create operation successful")
// }

func handleCreateStream(c *gin.Context) {
	stream := c.Param("stream")
	err := dbClient.CreateStream(dbName, stream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating record: %s", err)})
		return
	}
	c.String(http.StatusOK, "Create operation successful")
}

func handleDeleteStream(c *gin.Context) {
	stream := c.Param("stream")
	err := dbClient.DeleteStream(dbName, stream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting record: %s", err)})
		return
	}
	c.String(http.StatusOK, "Delete operation successful")
}

func handleWrite(c *gin.Context) {
	stream := c.Param("stream")

	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err)})
		return
	}

	err := dbClient.WriteEvent(dbName, stream, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error writing data: %s", err)})
		return
	}
	c.String(http.StatusOK, "Write operation successful")
}

func handleRead(c *gin.Context) {
	stream := c.Param("stream")
	response, err := dbClient.Read(dbName, stream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error reading data: %s", err)})
		return
	}
	c.String(http.StatusOK, response)
}

package main

import (
	"go-api/configs"
	"go-api/internal/initialize"
	"go-api/internal/routers"
	"os"
)

var (
	variables configs.Variables = configs.Variables{}
)

func main() {
	variables.LoadEnv()
	r := routers.NewRouter()
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	var (
		kafkaInitialize = initialize.NewProducerKafka(os.Getenv("BOOTSTRAP_SERVERS"))
	)
	kafkaInitialize.Consumer()
}

// func pong(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }

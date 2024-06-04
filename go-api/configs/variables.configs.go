package configs

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Variables struct {
	Hello string `env:"HELLO" envDefault:"Hello"`
}

func (v *Variables) LoadEnv() {
	envMode, ok := os.LookupEnv("GO_ENV")
	if !ok {
		log.Fatalf("Error loading .env file")
		os.Exit(0)
	}
	projectName := regexp.MustCompile(`^(.*` + "go-api" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	switch envMode {
	case "development":
		godotenv.Load(string(rootPath) + "/.dev.env")
	case "production":
		godotenv.Load(string(rootPath) + "/.prod.env")
	default:
		log.Fatalf("Error loading .env file")
		os.Exit(0)
	}
	v.Hello = os.Getenv("GREETING")
	println(v.Hello)
}

package lib

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(config string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	return os.Getenv(config)
}

package environments

import (
	"os"

	"github.com/joho/godotenv"
)

func Environments(env string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	environment := os.Getenv(env)
	return environment
}

package cmd

import (
	"fmt"

	"github.com/bookpanda/minio-api/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
}

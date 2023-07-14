package helpers

import (
	"fmt"
	"os"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}

	return fmt.Sprintf(":%s", port)
}

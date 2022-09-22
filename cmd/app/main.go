// $ go run cmd/app/main.go
// $ curl http://localhost:8080
package main

import (
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Warn().Err(err)
	}
}

// $ go run cmd/app/main.go
// $ curl http://localhost:8080
package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/ch-random/random-launcher-backend/configs"
	"github.com/ch-random/random-launcher-backend/delivery/httpserver"
	"github.com/ch-random/random-launcher-backend/repository/pscale"
	"github.com/ch-random/random-launcher-backend/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from `.env`
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load environment variables:", err)
	}

	_, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Println("connection to Firebase failed:", err)
	}

	db := pscale.GetDB()

	articleRepo := pscale.NewPscaleArticleRepository(db)
	userRepo := pscale.NewPscaleUserRepository(db)
	au := usecase.NewArticleUsecase(articleRepo, userRepo, configs.TIMEOUT)

	e := httpserver.NewArticleHandler(db, au)

	port := os.Getenv("PORT")
	if port == "" {
		port = configs.PORT
	}
	log.Println("listening on", port)
	addr := ":" + port
	if err := e.Start(addr); err != nil {
		log.Println("failed to serve HTTP:", err)
	}
}

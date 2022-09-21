// $ go run cmd/app/main.go
// $ curl http://localhost:8080
package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/ch-random/random-launcher-backend/cmd"
	"github.com/ch-random/random-launcher-backend/config"
	"github.com/ch-random/random-launcher-backend/delivery/httpserver"
	"github.com/ch-random/random-launcher-backend/repository/pscale"
	"github.com/ch-random/random-launcher-backend/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := cmd.Init(); err != nil {
		log.Println(err)
	}

	// Load environment variables from `.env`
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load environment variables:", err)
	}

	_, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Println("failed to connect to Firebase:", err)
	}

	db, err := pscale.GetDB()
	if err != nil {
		log.Fatalln("failed to connect to PlanetScale:", err)
	}

	ur := pscale.NewUserRepository(db)
	ar := pscale.NewArticleRepository(db)
	agc := pscale.NewArticleGameContentRepository(db)
	aor := pscale.NewArticleOwnerRepository(db)
	atr := pscale.NewArticleTagRepository(db)
	acr := pscale.NewArticleCommentRepository(db)
	aiur := pscale.NewArticleImageURLRepository(db)
	timeout := config.Timeout
	au := usecase.NewArticleUsecase(
		ur,
		ar,
		agc,
		aor,
		atr,
		acr,
		aiur,
		timeout,
	)

	e := httpserver.NewHandler(db, au)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}
	log.Println("listening on", port)
	addr := ":" + port
	if err := e.Start(addr); err != nil {
		log.Println("failed to serve HTTP:", err)
	}
}

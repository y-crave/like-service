package main

import (
	"like-service/internal/config"
	"like-service/internal/controller/like_controller"
	"like-service/internal/repository/like_repo"
	"like-service/internal/service"
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	likeRepo := likerepo.NewLikeRepo(db)
	likeService := service.NewLikeService(likeRepo)
	likeController := likecontroller.NewLikeController(likeService)

	router := mux.NewRouter()
	likeController.RegisterRoutes(router)

	log.Printf("Server started on port %d", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.HTTPPort), router))
}

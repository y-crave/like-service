package likecontroller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"like-service/internal/service"
	"net/http"
	"github.com/google/uuid"
)

type likeController struct {
	svc service.LikeService
}

func NewLikeController(svc service.LikeService) *likeController {
	return &likeController{
		svc: svc,
	}
}

func (c *likeController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/like-service/reaction", c.SetReaction).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/like-service/reactions", c.GetReactionsByUser).Methods(http.MethodGet)
}

func (c *likeController) SetReaction(w http.ResponseWriter, r *http.Request) {

	var reaction JSONReaction
	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		// Пока просто заглушки ошибок
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	domainReaction, err := toDomainReaction(reaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.svc.SetReaction(r.Context(), &domainReaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *likeController) GetReactionsByUser(w http.ResponseWriter, r *http.Request) {
	
}

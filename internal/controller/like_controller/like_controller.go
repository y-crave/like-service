package likecontroller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"like-service/internal/service"
	"net/http"
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
	router.HandleFunc("/api/v1/like", c.SetLike).Methods(http.MethodPost)
}

func (c *likeController) SetLike(w http.ResponseWriter, r *http.Request) {

	var like JSONLike
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		// Пока просто заглушки ошибок
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	domainLike, err := toDomainLike(like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.svc.SetLike(r.Context(), domainLike); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

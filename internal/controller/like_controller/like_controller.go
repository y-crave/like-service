package likecontroller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"like-service/internal/service"
	"net/http"
)

const (
	DEFAULT_REACTIONS_LIMIT = 20
	MAX_REACTIONS_LIMIT     = 1000
	DEFAULT_OFFSET          = 0
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

	//TODO id пользователя будет браться из токена в куках.
	userID, err := parseUUIDQuery(r, "userID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit := parseIntQuery(r, "limit", DEFAULT_REACTIONS_LIMIT)
	if limit <= 0 || limit > MAX_REACTIONS_LIMIT {
		//TODO Можно сделать: Если limit больше max_limit можно сделать его default_limit
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	offset := parseIntQuery(r, "offset", DEFAULT_OFFSET)
	if offset < 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	reactions, err := c.svc.GetReactionsByUserID(r.Context(), userID, limit, offset)
	if err != nil {
		//TODO Обработать ошибку.
		http.Error(w, err.Error(), 500)
	}

	jsonReactions := make([]JSONReaction, 0, len(reactions))

	for i := 0; i < len(reactions); i++ {
		jsonReactions[i] = toJSONReaction(&reactions[i])
	}

	var response JSONReactionsByUser
	response.Reactions = jsonReactions
	response.Limit = limit
	response.Offset = offset

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		//TODO Это стоит хотябы залогировать.
	}
}

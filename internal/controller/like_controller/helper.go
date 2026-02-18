package likecontroller

import (
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func parseIntQuery(r *http.Request, key string, def int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return def
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return n
}

func parseUUIDQuery(r *http.Request, key string) (uuid.UUID, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		//TODO Нужно продумать ошибку.
		return uuid.UUID{}, nil
	}
	id, err := uuid.Parse(val)
	if err != nil {
		//TODO Тоже самое.
		return uuid.UUID{}, err
	}
	return id, nil
}

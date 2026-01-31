package book

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TommyLearning/bookBackend/internal/logger"
	"github.com/TommyLearning/bookBackend/internal/response"
	"github.com/google/uuid"
)

type Storer interface {
	Create(context.Context, *Record) (*Record, error)
	FindById(context.Context, uuid.UUID) (*Record, error)
	FindAll(context.Context) ([]*Record, error)
	UpdateById(context.Context, uuid.UUID, *Record) error
	DeleteById(context.Context, uuid.UUID) error
}

type Handler struct {
	store Storer
}

func NewHandler(store Storer) *Handler {
	return &Handler{store: store}
}

func Save(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("post bi100")

		var requestBody SaveRequestBody
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			log.Error("failed to decode request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := requestBody.Validate()
		if err != nil {
			log.Error("failed to validate request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(ctx, n); err != nil {
			log.Error("failed to create bi100", "error", err)
			var dbErr *response.CustomError
			if errors.As(err, &dbErr) {
				w.WriteHeader(dbErr.HttpStatusCode())
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func Filter(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("get all news")
		//n, err := ns.FindAll(ctx)
		//if err != nil {
		//	log.Error("failed to get all news", "error", err)
		//	var dbErr *news.CustomError
		//	if errors.As(err, &dbErr) {
		//		w.WriteHeader(dbErr.HttpStatusCode())
		//		return
		//	}
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//allNewsResponse := AllNewsResponse{News: n}
		//if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
		//	log.Error("failed to encode response", "error", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
	}
}

func Get(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("get news by id")
		//newsID := r.PathValue("news_id")
		//newsUUID, err := uuid.Parse(newsID)
		//if err != nil {
		//	log.Error("failed to parse news id", "error", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		//n, err := ns.FindById(ctx, newsUUID)
		//if err != nil {
		//	log.Error("failed to get news by id", "error", err)
		//	var dbErr *news.CustomError
		//	if errors.As(err, &dbErr) {
		//		w.WriteHeader(dbErr.HttpStatusCode())
		//		return
		//	}
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//
		//if err := json.NewEncoder(w).Encode(n); err != nil {
		//	log.Error("failed to encode response", "error", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
	}
}

func Update(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("update news by id")

		//var newsReqBody NewsPostReqBody
		//if err := json.NewDecoder(r.Body).Decode(&newsReqBody); err != nil {
		//	log.Error("failed to decode the request", "error", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		//
		//n, err := newsReqBody.Validate()
		//if err != nil {
		//	log.Error("failed to validate request body", "error", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte(err.Error()))
		//	return
		//}
		//
		//if err2 := ns.UpdateById(ctx, n.Id, n); err2 != nil {
		//	log.Error("failed to update news by id", "error", err2)
		//	var dbErr *news.CustomError
		//	if errors.As(err, &dbErr) {
		//		w.WriteHeader(dbErr.HttpStatusCode())
		//		return
		//	}
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
	}
}

func Delete(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("delete news by id")

		//newsID := r.PathValue("news_id")
		//newsUUID, err := uuid.Parse(newsID)
		//if err != nil {
		//	log.Error("failed to parse news id", "error", err)
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		//if err := ns.DeleteById(ctx, newsUUID); err != nil {
		//	log.Error("failed to delete news by id", "error", err)
		//	var dbErr *news.CustomError
		//	if errors.As(err, &dbErr) {
		//		w.WriteHeader(dbErr.HttpStatusCode())
		//		return
		//	}
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		w.WriteHeader(http.StatusNoContent)
	}
}

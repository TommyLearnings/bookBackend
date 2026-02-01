package book

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/TommyLearning/bookBackend/internal/logger"
	"github.com/TommyLearning/bookBackend/internal/response"
)

type Storer interface {
	Create(context.Context, *Record) (*Record, error)
	FindById(context.Context, int) (*Record, error)
	FindAll(context.Context) ([]*Record, error)
	UpdateById(context.Context, int, *Record) error
	DeleteById(context.Context, int) error
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
				response.JSON(w, dbErr.HttpStatusCode(), dbErr.Error(), nil)
				return
			}

			response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
			return
		}
		response.JSON(w, http.StatusCreated, "成功新增", nil)
	}
}

func Filter(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)

		log.Info("get all news")
		n, err := ns.FindAll(ctx)

		if err != nil {
			log.Error("failed to get all news", "error", err)
			var dbErr *response.CustomError

			if errors.As(err, &dbErr) {
				response.JSON(w, dbErr.HttpStatusCode(), dbErr.Error(), nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
			return
		}

		response.JSON(w, http.StatusOK, "成功取得新聞列表", n)
	}
}

func Get(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("get bi100 by id")

		bi100ID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "id型別異常", nil)
			return
		}
		n, err := ns.FindById(ctx, bi100ID)
		if err != nil {
			log.Error("failed to get bi100 by id", "error", err)
			var dbErr *response.CustomError
			if errors.As(err, &dbErr) {
				response.JSON(w, dbErr.HttpStatusCode(), dbErr.Error(), nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
			return
		}

		//if err := json.NewEncoder(w).Encode(n); err != nil {
		//	log.Error("failed to encode response", "error", err)
		//	response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
		//	return
		//}
		response.JSON(w, http.StatusOK, "成功取得書籍", n)
	}
}

func Update(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("update bi100 by id")

		var requestBody SaveRequestBody
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			log.Error("failed to decode the request", "error", err)
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

		if err := ns.UpdateById(ctx, n.Id, n); err != nil {
			log.Error("failed to update bi100 by id", "error", err)
			var dbErr *response.CustomError
			if errors.As(err, &dbErr) {
				response.JSON(w, dbErr.HttpStatusCode(), dbErr.Error(), nil)
				return
			}
			response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
			return
		}

		response.JSON(w, http.StatusCreated, "成功更新", nil)
	}
}

func Delete(ns Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info("delete bi100 by id")

		bi100Id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "id型別異常", nil)
			return
		}

		if err := ns.DeleteById(ctx, bi100Id); err != nil {
			log.Error("failed to delete bi100 by id", "error", err)
			var dbErr *response.CustomError
			if errors.As(err, &dbErr) {
				w.WriteHeader(dbErr.HttpStatusCode())
				return
			}
			response.JSON(w, http.StatusInternalServerError, "伺服器內部錯誤", nil)
			return
		}
		response.JSON(w, http.StatusCreated, "成功刪除", nil)
	}
}

package router

import (
	"net/http"

	"github.com/TommyLearning/bookBackend/internal/book"
)

type Dependencies struct {
	BookHandler *book.Handler
	// 未來可以加入更多模組的 handler
}

func New(deps Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	// 註冊各模組的路由
	book.RegisterRoutes(mux, deps.BookHandler)

	// 健康檢查端點
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}

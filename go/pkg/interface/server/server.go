package server

import (
	"fmt"
	"net/http"

	"l-semi-chat/pkg/domain/logger"

	"github.com/gorilla/mux"
)

type server struct {
	Addr   string
	Port   string
	Router *mux.Router
}

// Server server
type Server interface {
	Serve()
	Handle(string, http.HandlerFunc)
}

// NewServer serverの作成
func NewServer(addr, port string) Server {
	return &server{
		Addr:   addr,
		Port:   port,
		Router: mux.NewRouter(),
	}
}

func (s *server) Serve() {
	logger.Info(fmt.Sprintf("Starting server http://%s:%s", s.Addr, s.Port))
	http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Addr, s.Port),
		s.Router,
	)
}

func (s *server) Handle(endpoint string, apiFunc http.HandlerFunc) {
	s.Router.HandleFunc(endpoint, httpSetting(apiFunc))
}

func httpSetting(apiFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*") // client server
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin")
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")

		apiFunc(writer, request)
	}
}

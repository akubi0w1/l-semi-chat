package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	Addr string
	Port string
}

// Server server
type Server interface {
	Serve()
	Handle(string, http.HandlerFunc)
}

// NewServer serverの作成
func NewServer(addr, port string) Server {
	return &server{
		Addr: addr,
		Port: port,
	}
}

func (s *server) Serve() {
	log.Println("Server running...")
	http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Addr, s.Port),
		nil,
	)
}

func (s *server) Handle(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc))
}

func httpMethod(apiFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")

		apiFunc(writer, request)
	}
}

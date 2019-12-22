package server

import (
	"fmt"
	"l-semi-chat/pkg/interface/server/response"
	"log"
	"net/http"
)

type server struct {
	Addr string
	Port string
}

type Server interface {
	Serve()
	Get(string, http.HandlerFunc)
	Post(string, http.HandlerFunc)
	Put(string, http.HandlerFunc)
	Delete(string, http.HandlerFunc)
}

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

func (s *server) Get(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodGet))
}

func (s *server) Post(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodPost))
}

func (s *server) Put(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodPut))
}

func (s *server) Delete(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodDelete))
}

func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
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

		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			response.MethodNotAllowed(writer, "Method Not Allowed")
			return
		}

		apiFunc(writer, request)
	}
}

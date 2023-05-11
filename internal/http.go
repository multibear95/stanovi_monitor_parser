package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	FlatsInfo *[]FlatInfo
}

func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		FlatsInfo: &[]FlatInfo{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/getFlats", server.getFlats).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *httpServer) getFlats(w http.ResponseWriter, r *http.Request) {
	var url GetMessagesByUrlReq
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := ParseMessages(url.Url)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

type GetMessagesByUrlReq struct {
	Url string `json:"url"`
}

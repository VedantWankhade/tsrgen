package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vedantwankhade/tsrgen/gateway-service/config"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
)

func (s *server) getPage(w http.ResponseWriter, r *http.Request) {
	// TODO)) cors for prod
	if config.GetAppRunMode() == "dev" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var pageReq domain.PageGetReq
	err := json.NewDecoder(r.Body).Decode(&pageReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	page, err := s.app.GetPage(pageReq.Id, pageReq.ConfluenceInstance, pageReq.ConfluenceUsername, pageReq.ConfluenceToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("get page request failed: %v", err), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(&page)
	w.Write(b)
	// w.Write([]byte("page got: " + page.ID + " " + page.Title))
}

func (s *server) createPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var pageReq domain.PageCreateReq
	err := json.NewDecoder(r.Body).Decode(&pageReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	page, err := s.app.CreatePage(pageReq.HTMLContent, pageReq.Title, pageReq.ParentPageId, pageReq.SpaceId, pageReq.ConfluenceInstance, pageReq.ConfluenceUsername, pageReq.ConfluenceToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("get page request failed: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "page created: %#v", page)
}

func (s *server) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", s.getPage)
	mux.HandleFunc("/create", s.createPage)
	return mux
}

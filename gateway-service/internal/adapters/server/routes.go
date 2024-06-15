package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
)

func corsMiddleware(next http.Handler) http.Handler {
	// TODO)) cors for prod
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *server) saveEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var saveReq domain.DBPageSaveReq
	err := json.NewDecoder(r.Body).Decode(&saveReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	id, err := s.app.SaveEntry(saveReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("save entry request failed: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) getEntries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: GET)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	entries, err := s.app.GetEntries()
	if err != nil {
		http.Error(w, fmt.Sprintf("get entries request failed: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(entries)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) getHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var statsReq domain.StatReq
	err := json.NewDecoder(r.Body).Decode(&statsReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	htmlRes, err := s.app.GetHTML(statsReq.Issues)
	if err != nil {
		http.Error(w, fmt.Sprintf("get issues request failed: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(htmlRes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) getIssues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var issuesReq domain.IssuesReq
	err := json.NewDecoder(r.Body).Decode(&issuesReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	issues, err := s.app.GetIssues(issuesReq.Jql, issuesReq.JiraInstance, issuesReq.JiraUsername, issuesReq.JiraToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("get issues request failed: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(issues)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) getPage(w http.ResponseWriter, r *http.Request) {
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
	err = json.NewEncoder(w).Encode(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
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
	err = json.NewEncoder(w).Encode(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) generate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s not supported (allowed: POST)", r.Method), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var generateReq domain.GenerateReq
	err := json.NewDecoder(r.Body).Decode(&generateReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse the request: %v", err), http.StatusBadRequest)
		return
	}
	pageUrl, err := s.app.Generate(generateReq.AtlassianInstance, generateReq.AtlassianUsername, generateReq.AtlassianToken, generateReq.JQL, generateReq.Title, generateReq.ParentId, generateReq.SpaceId)
	if err != nil {
		http.Error(w, fmt.Sprintf("generate page request failed: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(pageUrl)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding json: err: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/page", s.getPage)
	mux.HandleFunc("/page/create", s.createPage)
	mux.HandleFunc("/issues", s.getIssues)
	mux.HandleFunc("/html", s.getHTML)
	mux.HandleFunc("/generate", s.generate)
	mux.HandleFunc("/save", s.saveEntry)
	mux.HandleFunc("/entries", s.getEntries)
	return mux
}

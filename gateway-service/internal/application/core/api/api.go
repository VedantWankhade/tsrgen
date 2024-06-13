package api

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/ports"
)

type application struct {
	confluenceClient ports.ConfluenceClientPort
	jiraClient       ports.JiraClientPort
}

func (a *application) CreatePage(content, title, parentId, spaceId, confluenceInstance, username, token string) (*domain.Page, error) {
	return a.confluenceClient.CreatePage(content, title, parentId, spaceId, confluenceInstance, username, token)
}

func (a *application) GetPage(pageId int, confluenceInstance, username, token string) (*domain.Page, error) {
	return a.confluenceClient.GetPage(pageId, confluenceInstance, username, token)
}

func (a *application) GetIssues(jql, confluenceInstance, username, token string) ([]*domain.Issue, error) {
	return a.jiraClient.GetIssues(jql, confluenceInstance, username, token)
}

func NewApplication(confluenceClient ports.ConfluenceClientPort, jiraClient ports.JiraClientPort) *application {
	return &application{confluenceClient: confluenceClient, jiraClient: jiraClient}
}

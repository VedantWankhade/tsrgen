package api

import (
	"fmt"
	"log"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/ports"
)

type application struct {
	confluenceClient ports.ConfluenceClientPort
	jiraClient       ports.JiraClientPort
	statsClient      ports.StatsClientPort
	dbClient         ports.DBClientPort
}

func (a *application) GetHTML(issues []*domain.Issue) (string, error) {
	return a.statsClient.GetHTML(issues)
}

func (a *application) SaveEntry(page domain.DBPageSaveReq) (int, error) {
	return a.dbClient.SaveEntry(page)
}

func (a *application) GetEntries() ([]*domain.Entry, error) {
	return a.dbClient.GetEntries()
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

func (a *application) Generate(atlassianInstance, atlassianUsername, atlassianToken, jql, title, parentId, spaceId string) (string, error) {
	issues, err := a.GetIssues(jql, atlassianInstance, atlassianUsername, atlassianToken)
	if err != nil {
		return "", fmt.Errorf("error fetching the issues: %w", err)
	}
	html, err := a.GetHTML(issues)
	if err != nil {
		return "", fmt.Errorf("error fetching the html: %w", err)
	}
	page, err := a.CreatePage(html, title, parentId, spaceId, atlassianInstance, atlassianUsername, atlassianToken)
	if err != nil {
		return "", fmt.Errorf("error generating page: %w", err)
	}
	pageLink := fmt.Sprintf("https://%s/wiki/pages/viewinfo.action?pageId=%s", atlassianInstance, page.ID)
	saveReq := domain.DBPageSaveReq{
		Id:    page.ID,
		Title: page.Title,
		Link:  pageLink,
	}
	_, err = a.SaveEntry(saveReq)
	if err != nil {
		log.Printf("error saving entry to db: %v\n", err)
	}
	return pageLink, nil
}

func NewApplication(confluenceClient ports.ConfluenceClientPort, jiraClient ports.JiraClientPort, statsClient ports.StatsClientPort, dbClient ports.DBClientPort) *application {
	return &application{confluenceClient: confluenceClient, jiraClient: jiraClient, statsClient: statsClient, dbClient: dbClient}
}

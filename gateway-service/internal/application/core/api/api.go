package api

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/ports"
)

type application struct {
	confluenceClient ports.ConfluenceClientPort
}

func (a *application) CreatePage(content, title, parentId, spaceId, confluenceInstance, username, token string) (*domain.Page, error) {
	return a.confluenceClient.CreatePage(content, title, parentId, spaceId, confluenceInstance, username, token)
}

func (a *application) GetPage(pageId int, confluenceInstance, username, token string) (*domain.Page, error) {
	return a.confluenceClient.GetPage(pageId, confluenceInstance, username, token)
}

func NewApplication(confluenceClient ports.ConfluenceClientPort) *application {
	return &application{confluenceClient: confluenceClient}
}

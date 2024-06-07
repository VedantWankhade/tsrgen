package clients

import (
	"context"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/services/confluence"
	"google.golang.org/grpc"
)

type confluenceClient struct {
	Client confluence.ConfluenceClient
}

func NewConfluenceClient(conn *grpc.ClientConn) *confluenceClient {
	return &confluenceClient{Client: confluence.NewConfluenceClient(conn)}
}

func (c *confluenceClient) GetPage(pageId int, confluenceInstance, username, token string) (*domain.Page, error) {
	pageReq := confluence.PageReq{
		Id:                 int64(pageId),
		ConfluenceInstance: confluenceInstance,
		ConfleunceUsername: username,
		ConfluenceToken:    token,
	}
	pageRes, err := c.Client.GetPage(context.Background(), &pageReq)
	if err != nil {
		return nil, err
	}
	page := domain.Page{
		ID:    pageRes.GetId(),
		Title: pageRes.GetTitle(),
		Body: domain.Body{
			HTML: map[string]string{
				"representation": "storage",
				"value":          pageRes.Html,
			},
		},
	}
	return &page, nil
}

func (c *confluenceClient) CreatePage(content, title, parentId, spaceId, confluenceInstance, username, token string) (*domain.Page, error) {
	pageReq := confluence.PageCreateReq{
		ParentPageId:       parentId,
		SpaceId:            spaceId,
		Title:              title,
		ConfluenceInstance: confluenceInstance,
		ConfluenceToken:    token,
		HtmlContent:        content,
		ConfleunceUsername: username,
	}
	pageRes, err := c.Client.CreatePage(context.Background(), &pageReq)
	if err != nil {
		return nil, err
	}
	page := domain.Page{
		ID:    pageRes.GetId(),
		Title: pageRes.GetTitle(),
	}
	return &page, nil
}

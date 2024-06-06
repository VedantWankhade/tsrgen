package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vedantwankhade/tsrgen/confluence-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/confluence-service/pkg/util"
)

type application struct{}

func (a *application) GetPageFromID(pageId int, confluenceInstance, username, token string) (*domain.Page, error) {
	creds := fmt.Sprintf("%s:%s", username, token)
	bearerToken := base64.StdEncoding.EncodeToString([]byte(creds))
	url := fmt.Sprintf("https://%s/wiki/api/v2/pages/%d", confluenceInstance, pageId)

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Authorization"] = "Basic " + bearerToken

	params := make(map[string]string)
	params["body-format"] = "storage"

	body, err := util.MakeHTTPRequest(http.MethodGet, url, headers, nil, params)
	if err != nil {
		return nil, fmt.Errorf("request to get page %d failed %w", pageId, err)
	}
	defer body.Close()

	var page domain.Page
	err = json.NewDecoder(body).Decode(&page)
	if err != nil {
		return nil, fmt.Errorf("error decoding the response from get page %d request %w", pageId, err)
	}
	return &page, nil
}

func NewApplication() *application {
	return &application{}
}

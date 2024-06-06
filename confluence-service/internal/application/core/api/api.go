package core

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vedantwankhade/tsrgen/confluence-service/pkg/util"
)

type resJson struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  struct {
		Storage map[string]string `json:"storage"`
	} `json:"body"`
}

func GetPageFromID(pageId int, confluenceInstance, username, token string) []byte {
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
		log.Fatalf("Request to get page %d failed %v", pageId, err)
	}
	defer body.Close()
	var obj resJson
	json.NewDecoder(body).Decode(&obj)
	fmt.Println(obj.Body.Storage["value"])
	return nil
}

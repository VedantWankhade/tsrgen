package domain

type GenerateReq struct {
	AtlassianInstance string `json:"atlassianInstance"`
	AtlassianUsername string `json:"atlassianUsername"`
	AtlassianToken    string `json:"atlassianToken"`
	JQL               string `json:"jql"`
	Title             string `json:"title"`
	ParentId          string `json:"parentId"`
	SpaceId           string `json:"spaceId"`
}

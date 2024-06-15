package domain

type Page struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	SpaceID string `json:"spaceId"`
	Body    Body   `json:"body"`
}

type Body struct {
	HTML map[string]string `json:"storage"`
}

type PageGetReq struct {
	Id                 int    `json:"id"`
	ConfluenceInstance string `json:"confluenceInstance"`
	ConfluenceUsername string `json:"confluenceUsername"`
	ConfluenceToken    string `json:"confluenceToken"`
}

type PageCreateReq struct {
	Title              string `json:"title"`
	SpaceId            string `json:"spaceId"`
	HTMLContent        string `json:"htmlContent"`
	ParentPageId       string `json:"parentPageId"`
	ConfluenceInstance string `json:"confluenceInstance"`
	ConfluenceUsername string `json:"confluenceUsername"`
	ConfluenceToken    string `json:"confluenceToken"`
}

type DBPageSaveReq struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

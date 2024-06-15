package domain

type Entry struct {
	EntryId   int    `json:"entryId"`
	PageId    string `json:"pageId"`
	PageTitle string `json:"pageTitle"`
	PageLink  string `json:"pageLink"`
}

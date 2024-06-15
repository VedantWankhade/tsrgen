package domain

type Page struct {
	Id    string `json:"id"`
	Title string `json:"tite"`
	Link  string `json:"link"`
}

type Entry struct {
	EntryId   int `gorm:"primaryKey"`
	PageId    string
	PageTitle string
	PageLink  string
}

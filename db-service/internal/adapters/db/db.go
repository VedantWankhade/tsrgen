package db

import (
	"fmt"

	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBAdapter struct {
	db *gorm.DB
}

type entry struct {
	EntryId   int `gorm:"primaryKey"`
	PageId    string
	PageTitle string
	PageLink  string
}

func openDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func NewDBAdapter(dsn string) (*DBAdapter, error) {
	dbConn, err := openDB(dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}
	return &DBAdapter{
		db: dbConn,
	}, nil
}

func (a *DBAdapter) save(entryItem entry) error {
	tx := a.db.Create(&entryItem)
	return tx.Error
}

func (a *DBAdapter) Save(page domain.Page) error {
	a.db.AutoMigrate(&entry{})
	item := entry{
		PageId:    page.Id,
		PageTitle: page.Title,
		PageLink:  page.Link,
	}
	err := a.save(item)
	return err
}

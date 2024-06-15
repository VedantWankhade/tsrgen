package db

import (
	"fmt"
	"time"

	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBAdapter struct {
	db *gorm.DB
}

func openDB(dsn string) (*gorm.DB, error) {
	tries := 5
	var err error
	var conn *gorm.DB
	for {
		if tries == 0 {
			return nil, fmt.Errorf("could not connect to the database after 5 tries: %w", err)
		}
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		tries--
		if err != nil {
			fmt.Println("DB not up yet. Trying again after 2 seconds...")
			time.Sleep(2 * time.Second)
		} else {
			return conn, err
		}
	}
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

func (a *DBAdapter) save(entryItem domain.Entry) (int, error) {
	tx := a.db.Create(&entryItem)
	return entryItem.EntryId, tx.Error
}

func (a *DBAdapter) Save(page domain.Page) (int, error) {
	a.db.AutoMigrate(&domain.Entry{})
	item := domain.Entry{
		PageId:    page.Id,
		PageTitle: page.Title,
		PageLink:  page.Link,
	}
	return a.save(item)
}

func (a *DBAdapter) GetAll() ([]*domain.Entry, error) {
	var entries []*domain.Entry
	tx := a.db.Find(&entries)
	return entries, tx.Error
}

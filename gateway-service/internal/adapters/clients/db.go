package clients

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/services/db"
	"google.golang.org/grpc"
)

type dbClient struct {
	Client db.DBClient
}

func NewDBClient(conn *grpc.ClientConn) *dbClient {
	return &dbClient{Client: db.NewDBClient(conn)}
}

func (c *dbClient) SaveEntry(page domain.DBPageSaveReq) (int, error) {
	saveReq := db.EntrySaveReq{
		PageId:    page.Id,
		PageTitle: page.Title,
		PageLink:  page.Link,
	}
	saveRes, err := c.Client.SaveEntry(context.Background(), &saveReq)
	if err != nil {
		return -1, err
	}
	id, err := strconv.Atoi(saveRes.GetEntryId())
	if err != nil {
		return -1, fmt.Errorf("invalid return type of entry id: %w", err)
	}
	return id, nil
}

func (c *dbClient) GetEntries() ([]*domain.Entry, error) {
	entriesRes, err := c.Client.GetEntries(context.Background(), &db.None{})
	if err != nil {
		return nil, err
	}
	var entries []*domain.Entry
	for _, entry := range entriesRes.Entries {
		entries = append(entries, &domain.Entry{
			EntryId:   int(entry.EntryId),
			PageId:    entry.PageId,
			PageTitle: entry.PageTitle,
			PageLink:  entry.PageLink,
		})
	}
	return entries, nil
}

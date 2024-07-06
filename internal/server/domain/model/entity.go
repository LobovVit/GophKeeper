package model

import (
	"encoding/json"
	"time"

	pb "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

type Entity struct {
	ID        int64
	UserID    int64
	Data      []byte
	Metadata  MetadataEntity
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type MetadataEntity struct {
	Name        string
	Description string
	Type        string
}

type CreateEntityRequest struct {
	UserID      int64
	Data        []byte
	Metadata    MetadataEntity
	AccessToken string
}

func GetListEntity(data []Entity) ([]*pb.Entity, error) {
	items := make([]*pb.Entity, len(data))
	for i := range data {
		jsonMetadata, err := json.Marshal(data[i].Metadata)
		if err != nil {
			return items, err
		}
		created := utils.ConvertTimeToTimestamp(data[i].CreatedAt)
		updated := utils.ConvertTimeToTimestamp(data[i].UpdatedAt)
		items[i] = &pb.Entity{Id: data[i].ID, UserId: data[i].UserID, Data: data[i].Data,
			Metadata: string(jsonMetadata), CreatedAt: created, UpdatedAt: updated}
	}
	return items, nil
}

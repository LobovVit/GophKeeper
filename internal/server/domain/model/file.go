package model

import (
	"time"

	pb "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

type File struct {
	ID        int64
	UserID    int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type FileRequest struct {
	UserID      int64
	Name        string
	AccessToken string
}

func GetListFile(binary []File) []*pb.Binary {
	items := make([]*pb.Binary, len(binary))
	for i := range binary {
		created := utils.ConvertTimeToTimestamp(binary[i].CreatedAt)
		items[i] = &pb.Binary{Id: binary[i].ID, Name: binary[i].Name, CreatedAt: created}
	}
	return items
}

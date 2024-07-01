package model

import (
	"database/sql"

	pb "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
	DeletedAt timestamp.Timestamp
}

type GetAllUsers struct {
	ID        int64
	Username  string
	Password  string
	DeletedAt sql.NullTime
}

type UserRequest struct {
	Username string
	Password string
}

func GetUserData(data *User) *pb.User {
	return &pb.User{
		UserId:    data.ID,
		Username:  data.Username,
		CreatedAt: &data.CreatedAt,
		UpdatedAt: &data.UpdatedAt,
		DeletedAt: &data.DeletedAt,
	}
}

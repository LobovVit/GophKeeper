package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	customErrors "github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
)

type User struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Registration(ctx context.Context, user *model.UserRequest) (*model.User, error) {
	registeredUser := &model.User{}
	const sqlText = "INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING user_id, username"
	if err := u.db.QueryRowContext(ctx, sqlText,
		user.Username,
		user.Password,
		time.Now(),
	).Scan(&registeredUser.ID, &registeredUser.Username); err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (u *User) Authentication(ctx context.Context, userRequest *model.UserRequest) (*model.User, error) {
	authenticatedUser := &model.User{}
	const sqlText = "SELECT user_id, username FROM users WHERE username=$1 and password=$2 and deleted_at IS NULL"
	err := u.db.QueryRowContext(ctx, sqlText,
		userRequest.Username, userRequest.Password).Scan(
		&authenticatedUser.ID,
		&authenticatedUser.Username,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrWrongUsernameOrPassword
		} else {
			return nil, err
		}
	}
	return authenticatedUser, nil
}

func (u *User) UserExists(ctx context.Context, username string) (bool, error) {
	var exists bool
	const sqlText = "SELECT EXISTS(SELECT 1 FROM users where username = $1)"
	row := u.db.QueryRowContext(ctx, sqlText, username)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

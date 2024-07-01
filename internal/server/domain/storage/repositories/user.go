package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	custom_errors "github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
)

type Constructor interface {
	Registration(data *model.UserRequest) error
	Authentication(user *model.UserRequest) (bool, error)
	UserExists(username string) (bool, error)
}

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
			return nil, custom_errors.ErrWrongUsernameOrPassword
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

func (u *User) UserList(ctx context.Context) ([]model.GetAllUsers, error) {
	users := []model.GetAllUsers{}
	const sqlText = "SELECT username, deleted_at FROM users"
	rows, err := u.db.QueryContext(ctx, sqlText)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, custom_errors.ErrRecordNotFound
		} else {
			return users, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		user := model.GetAllUsers{}

		err = rows.Scan(&user.Username, &user.DeletedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) Block(ctx context.Context, username string) (int64, error) {
	var id int64
	const sqlText = "UPDATE users SET deleted_at = $1 where username = $2 RETURNING user_id"
	if err := u.db.QueryRowContext(ctx, sqlText,
		time.Now(),
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (u *User) Unblock(ctx context.Context, username string) (int64, error) {
	var id int64
	const sqlText = "UPDATE users SET deleted_at = null where username = $1 RETURNING user_id"
	if err := u.db.QueryRowContext(ctx, sqlText,
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (u *User) GetUserID(ctx context.Context, username string) (int64, error) {
	var id int64
	const sqlText = "SELECT user_id FROM users where username = $1"
	if err := u.db.QueryRowContext(ctx, sqlText,
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

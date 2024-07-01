package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	custom_errors "github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

const lengthToken = 32

type TokenRepository interface {
	Create(user *model.User) (string, error)
}

type Token struct {
	db *sql.DB
}

func NewTokenRepo(db *sql.DB) *Token {
	return &Token{
		db: db,
	}
}

func (t *Token) Create(ctx context.Context, userID int64, lifetime time.Duration) (*model.Token, error) {
	token := &model.Token{}
	accessToken := utils.NewToken(lengthToken)
	currentTime := time.Now()
	const sqlText = "INSERT INTO access_token (access_token, user_id, created_at, end_date_at) VALUES ($1, $2, $3, $4) RETURNING access_token, user_id, created_at, end_date_at"
	if err := t.db.QueryRowContext(ctx, sqlText,
		accessToken,
		userID,
		currentTime,
		currentTime.Add(lifetime),
	).Scan(&token.AccessToken, &token.UserID, &token.CreatedAt, &token.EndDateAt); err != nil {
		return nil, err
	}
	return token, nil
}

func (t *Token) Validate(endDate time.Time) bool {
	now := time.Now().Format("01/02/2006 15:04:05")
	end := endDate.Format("01/02/2006 15:04:05")
	valid := end > now
	return valid
}

func (t *Token) GetEndDateToken(ctx context.Context, accessToken string) (time.Time, error) {
	var end time.Time
	const sqlText = "SELECT end_date_at FROM access_token where access_token = $1"
	if err := t.db.QueryRowContext(ctx, sqlText,
		accessToken,
	).Scan(&end); err != nil {
		return end, err
	}
	return end, nil
}

func (t *Token) Block(ctx context.Context, accessToken string) (string, error) {
	var token string
	const sqlText = "UPDATE access_token SET end_date_at = $1 where access_token = $2 RETURNING access_token"
	if err := t.db.QueryRowContext(ctx, sqlText,
		time.Now(),
		accessToken,
	).Scan(&token); err != nil {
		return token, err
	}
	return token, nil
}

func (t *Token) BlockAllTokenUser(ctx context.Context, userID int64) (string, error) {
	var token string
	const sqlText = "UPDATE access_token SET end_date_at = $1 where user_id = $2 RETURNING access_token"
	if err := t.db.QueryRowContext(ctx, sqlText,
		time.Now(),
		userID,
	).Scan(&token); err != nil {
		return token, err
	}
	return token, nil
}

func (t *Token) GetList(ctx context.Context, userID int64) ([]model.Token, error) {
	tokens := []model.Token{}
	const sqlText = "SELECT access_token, user_id, created_at, end_date_at FROM access_token where user_id = $1"
	rows, err := t.db.QueryContext(ctx, sqlText,
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tokens, custom_errors.ErrRecordNotFound
		} else {
			return tokens, err
		}
	}
	if err = rows.Err(); err != nil {
		return tokens, err
	}
	defer rows.Close()
	for rows.Next() {
		token := model.Token{}
		err = rows.Scan(&token.AccessToken, &token.UserID, &token.CreatedAt, &token.EndDateAt)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

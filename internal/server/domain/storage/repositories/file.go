package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	customErrors "github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
)

type Files struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) *Files {
	return &Files{
		db: db,
	}
}

func (f *Files) UploadFile(ctx context.Context, binaryRequest *model.FileRequest) (*model.File, error) {
	binary := &model.File{}
	const sqlText = "INSERT INTO file (user_id, name, created_at) VALUES ($1, $2, $3) RETURNING file_id, name"
	if err := f.db.QueryRowContext(ctx, sqlText,
		binaryRequest.UserID,
		binaryRequest.Name,
		time.Now(),
	).Scan(&binary.ID, &binary.Name); err != nil {
		return nil, err
	}
	return binary, nil
}

func (f *Files) GetListFile(ctx context.Context, userID int64) ([]model.File, error) {
	listFile := []model.File{}
	const sqlText = "SELECT file_id, user_id, name, created_at FROM file where user_id = $1 and deleted_at IS NULL"
	rows, err := f.db.QueryContext(ctx, sqlText, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		binary := model.File{}
		err = rows.Scan(&binary.ID, &binary.UserID, &binary.Name, &binary.CreatedAt)
		if err != nil {
			return nil, err
		}
		listFile = append(listFile, binary)
	}
	return listFile, nil
}

func (f *Files) FileExists(ctx context.Context, binaryRequest *model.FileRequest) (bool, error) {
	var exists bool
	const sqlText = "SELECT EXISTS(SELECT 1 FROM file where file.user_id = $1 and file.name = $2 and file.deleted_at IS NULL)"
	row := f.db.QueryRowContext(ctx, sqlText,
		binaryRequest.UserID,
		binaryRequest.Name)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (f *Files) DeleteFile(ctx context.Context, binaryRequest *model.FileRequest) (int64, error) {
	var id int64
	const sqlText = "UPDATE file SET deleted_at = $1 where file.user_id = $2 and file.name = $3 RETURNING file_id"
	if err := f.db.QueryRowContext(ctx, sqlText,
		time.Now(),
		binaryRequest.UserID,
		binaryRequest.Name,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

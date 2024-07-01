package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/LobovVit/GophKeeper/internal/server/domain/model"
	custom_errors "github.com/LobovVit/GophKeeper/internal/server/domain/storage/errors"
)

type Entity struct {
	db *sql.DB
}

func NewEntityRepo(db *sql.DB) *Entity {
	return &Entity{
		db: db,
	}
}

func (e *Entity) Create(ctx context.Context, entityRequest *model.CreateEntityRequest) (int64, error) {
	var id int64
	metadata := model.MetadataEntity{Name: entityRequest.Metadata.Name, Description: entityRequest.Metadata.Description, Type: entityRequest.Metadata.Type}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return 0, err
	}
	const sqlText = "INSERT INTO entity (user_id, data, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING entity_id"
	if err = e.db.QueryRowContext(ctx,
		sqlText,
		entityRequest.UserID,
		entityRequest.Data,
		jsonMetadata,
		time.Now(),
		time.Now(),
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (e *Entity) GetList(ctx context.Context, userID int64, typeEntity string) ([]model.Entity, error) {
	entities := []model.Entity{}
	const sqlText = "SELECT entity_id, user_id, data, metadata, created_at, updated_at FROM entity where user_id = $1 and metadata->>'Type' = $2 and deleted_at IS NULL"
	rows, err := e.db.QueryContext(ctx, sqlText, userID, typeEntity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities, custom_errors.ErrRecordNotFound
		} else {
			return entities, err
		}
	}
	if err = rows.Err(); err != nil {
		return entities, err
	}
	defer rows.Close()
	for rows.Next() {
		entity := model.Entity{}
		entityMetadata := model.MetadataEntity{}

		var jsonEntity string
		err = rows.Scan(&entity.ID, &entity.UserID, &entity.Data, &jsonEntity, &entity.CreatedAt, &entity.UpdatedAt)
		if err != nil {
			return entities, err
		}

		err = json.Unmarshal([]byte(jsonEntity), &entityMetadata)
		if err != nil {
			return entities, err
		}
		entity.Metadata = entityMetadata
		entities = append(entities, entity)
	}
	return entities, nil
}

func (e *Entity) Exists(ctx context.Context, entityRequest *model.CreateEntityRequest) (bool, error) {
	var exists bool
	const sqlText = "SELECT EXISTS(SELECT 1 FROM entity where entity.user_id = $1 and entity.metadata->>'Name' = $2 and entity.metadata->>'Type' = $3 and entity.deleted_at IS NULL)"
	row := e.db.QueryRowContext(ctx, sqlText,
		entityRequest.UserID, entityRequest.Metadata.Name, entityRequest.Metadata.Type)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (e *Entity) Delete(ctx context.Context, userID int64, name string, typeEntity string) (int64, error) {
	var id int64
	const sqlText = "UPDATE entity SET deleted_at = $1 where entity.user_id = $2 and entity.metadata->>'Name' = $3 and entity.metadata->>'Type' = $4 RETURNING entity_id"
	if err := e.db.QueryRowContext(ctx, sqlText,
		time.Now(),
		userID,
		name,
		typeEntity,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (e *Entity) Update(ctx context.Context, userID int64, name string, typeEntity string, data []byte) (int64, error) {
	var id int64
	const sqlText = "UPDATE entity SET data = $1, updated_at = $2 where entity.user_id = $3 and entity.metadata->>'Name' = $4 and entity.metadata->>'Type' = $5 RETURNING entity_id"
	if err := e.db.QueryRowContext(ctx, sqlText,
		data,
		time.Now(),
		userID,
		name,
		typeEntity,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

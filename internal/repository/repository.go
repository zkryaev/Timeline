package repository

import (
	"context"
	"fmt"
	"time"
	"timeline/internal/config"
	"timeline/internal/entity"
	"timeline/internal/repository/database/postgres"
	"timeline/internal/repository/models"
)

type Database interface {
	Open() error
	Close()
	Repository
}

type Repository interface {
	UserRepository
	OrgRepository
}

type UserRepository interface {
	UserSave(ctx context.Context, user *models.UserRegisterModel) (int, error)
	UserByEmail(ctx context.Context, email string) (*entity.User, error)
	UserByID(ctx context.Context, id int) (*entity.User, error)
	UserGetMetaInfo(ctx context.Context, email string) (*models.MetaInfo, error)
	UserSaveCode(ctx context.Context, code string, user_id int) error
	UserCode(ctx context.Context, code string, user_id int) (time.Time, error)
	UserActivateAccount(ctx context.Context, user_id int) error
	UserIsExist(ctx context.Context, email string) (int, error)
}

type OrgRepository interface {
	OrgSave(ctx context.Context, org *models.OrgRegisterModel, cityName string) (int, error)
	OrgByEmail(ctx context.Context, email string) (*entity.Organization, error)
	OrgByID(ctx context.Context, id int) (*entity.Organization, error)
	OrgGetMetaInfo(ctx context.Context, email string) (*models.MetaInfo, error)
	OrgSaveCode(ctx context.Context, code string, org_id int) error
	OrgCode(ctx context.Context, code string, org_id int) (time.Time, error)
	OrgActivateAccount(ctx context.Context, user_id int) error
	OrgIsExist(ctx context.Context, email string) (int, error)
}

// Паттерн фабричный метод, для гибкого создания БД
func GetDB(name string, cfg config.Database) (Database, error) {
	switch name {
	case "postgres":
		return postgres.New(cfg), nil
	default:
		return nil, fmt.Errorf("unexpected db name")
	}
}
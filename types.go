package gormrepository

import "gorm.io/gorm"

type Repository interface {
	FindAll(target interface{}, preloads ...string) error
	FindBatch(target interface{}, limit, offset int, orderBy, orderSort *string, groupBy *string, preloads ...string) (count int64, err error)

	FindWhere(target interface{}, condition string, preloads ...string) error
	FindWhereBatch(target interface{}, condition string, limit, offset int, preloads ...string) error

	FindByField(target interface{}, field string, value interface{}, preloads ...string) error
	FindByFields(target interface{}, fields map[string]interface{}, preloads ...string) error

	FindByFieldBatch(target interface{}, field string, value interface{}, limit, offset int, preloads ...string) error
	FindByFieldsBatch(target interface{}, fields map[string]interface{}, limit, offset int, preloads ...string) error

	FindOneByField(target interface{}, field string, value interface{}, preloads ...string) error
	FindOneByFields(target interface{}, fields map[string]interface{}, preloads ...string) error

	FindOneByID(target interface{}, id interface{}, preloads ...string) error

	Create(target interface{}) error
	Save(target interface{}) error
	Delete(target interface{}) error

	DB() *gorm.DB
	DBWithPreloads(preloads []string) *gorm.DB
	HandleError(res *gorm.DB) error
	HandleOneError(res *gorm.DB) error
}

type TransactionRepository interface {
	Repository
	CreateTx(target interface{}, tx *gorm.DB) error
	SaveTx(target interface{}, tx *gorm.DB) error
	DeleteTx(target interface{}, tx *gorm.DB) error
}

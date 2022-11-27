package gormrepository

import (
	"fmt"

	"gorm.io/gorm"
)

type gormRepository struct {
	db           *gorm.DB
	defaultJoins []string
}

func NewGormRepository(db *gorm.DB, defaultJoins ...string) *gormRepository {
	return &gormRepository{db: db, defaultJoins: defaultJoins}
}

func (r *gormRepository) DB() *gorm.DB {
	return r.DBWithPreloads(nil)
}

func (r *gormRepository) DBWithPreloads(preloads []string) *gorm.DB {
	conn := r.db

	for _, join := range r.defaultJoins {
		conn = conn.Joins(join)
	}

	for _, preload := range r.defaultJoins {
		conn = conn.Preload(preload)
	}

	return conn
}

func (r *gormRepository) FindAll(target interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindBatch(target interface{}, limit, offset int, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindWhere(target interface{}, condition string, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(condition).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindWhereBatch(target interface{}, condition string, limit, offset int, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(condition).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindByField(target interface{}, field string, value interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%s = ?", field), value).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindByFields(target interface{}, fields map[string]interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fields).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindByFieldBatch(target interface{}, field string, value interface{}, limit, offset int, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%s = ?", field), value).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindByFieldsBatch(target interface{}, fields map[string]interface{}, limit, offset int, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fields).
		Limit(limit).
		Offset(offset).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindOneByField(target interface{}, field string, value interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%s = ?", field), value).
		First(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindOneByFields(target interface{}, fields map[string]interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fields).
		First(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindOneByID(target interface{}, id interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where("id = ?", id).
		First(target)

	return r.HandleError(res)
}

func (r *gormRepository) Create(target interface{}) error {
	res := r.DB().Create(target)

	return r.HandleError(res)
}

func (r *gormRepository) CreateTx(target interface{}, tx *gorm.DB) error {
	res := tx.Create(target)

	return r.HandleError(res)
}

func (r *gormRepository) Save(target interface{}) error {
	res := r.DB().Save(target)

	return r.HandleError(res)
}

func (r *gormRepository) SaveTx(target interface{}, tx *gorm.DB) error {
	res := tx.Save(target)

	return r.HandleError(res)
}

func (r *gormRepository) Delete(target interface{}) error {
	res := r.DB().Delete(target)

	return r.HandleError(res)
}

func (r *gormRepository) DeleteTx(target interface{}, tx *gorm.DB) error {
	res := tx.Delete(target)

	return r.HandleError(res)
}

// handle error
func (r *gormRepository) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("error: %w", res.Error)

		return err
	}

	return nil
}

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
	return r.DBWithPreloads(nil).Debug()
}

func (r *gormRepository) DBWithPreloads(preloads []string) *gorm.DB {
	conn := r.db.Debug()

	for _, join := range r.defaultJoins {
		conn = conn.Joins(join)
	}

	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}

	return conn
}

func (r *gormRepository) FindByRaw(target interface{}, query string) error {
	res := r.DB().Raw(query).
		Scan(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindAll(target interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindBatch(target interface{}, limit, offset int, search, orderBy, groupBy string, preloads ...string) (count int64, err error) {
	query := r.DBWithPreloads(preloads)

	if search != "" {
		query = query.Where(search)
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	if groupBy != "" {
		query = query.Group(groupBy)
	}

	res := query.
		Limit(limit).
		Offset(offset).
		Find(target).
		Count(&count)

	return count, r.HandleError(res)
}

func (r *gormRepository) FindWhere(target interface{}, condition string, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(condition).
		Find(target)

	return r.HandleError(res)
}

func (r *gormRepository) FindWhereBatch(target interface{}, condition string, limit, offset int, search, orderBy, groupBy string, preloads ...string) (count int64, err error) {
	query := r.DBWithPreloads(preloads)

	if search != "" {
		query = query.Where(search)
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	if groupBy != "" {
		query = query.Group(groupBy)
	}

	res := query.Where(condition).
		Limit(limit).
		Offset(offset).
		Find(target).
		Count(&count)

	return count, r.HandleError(res)
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

func (r *gormRepository) FindByFieldBatch(target interface{}, field string, value interface{}, limit, offset int, search, orderBy, groupBy string, preloads ...string) (count int64, err error) {
	query := r.DBWithPreloads(preloads)

	if search != "" {
		query = query.Where(search)
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	if groupBy != "" {
		query = query.Group(groupBy)
	}

	res := query.Where(fmt.Sprintf("%s = ?", field), value).
		Limit(limit).
		Offset(offset).
		Find(target).
		Count(&count)

	return count, r.HandleError(res)
}

func (r *gormRepository) FindByFieldsBatch(target interface{}, fields map[string]interface{}, limit, offset int, search, orderBy, groupBy string, preloads ...string) (count int64, err error) {
	query := r.DBWithPreloads(preloads)

	if search != "" {
		query = query.Where(search)
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	if groupBy != "" {
		query = query.Group(groupBy)
	}

	res := query.Where(fields).
		Limit(limit).
		Offset(offset).
		Find(target).
		Count(&count)

	return count, r.HandleError(res)
}

func (r *gormRepository) FindOneByField(target interface{}, field string, value interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%s = ?", field), value).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneLastByField(target interface{}, field string, value interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fmt.Sprintf("%s = ?", field), value).
		Last(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneByFields(target interface{}, fields map[string]interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(fields).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneByID(target interface{}, id interface{}, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where("id = ?", id).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneByCondition(target interface{}, condition string, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(condition).
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) FindOneLastByCondition(target interface{}, condition string, preloads ...string) error {
	res := r.DBWithPreloads(preloads).
		Where(condition).
		Order("id desc").
		First(target)

	return r.HandleOneError(res)
}

func (r *gormRepository) Create(target interface{}) error {
	res := r.DB().Create(target)

	return r.HandleError(res)
}

func (r *gormRepository) CreateTx(target interface{}, tx *gorm.DB) error {
	res := tx.Debug().Create(target)

	return r.HandleError(res)
}

func (r *gormRepository) Save(target interface{}) error {
	res := r.DB().Save(target)

	return r.HandleError(res)
}

func (r *gormRepository) SaveTx(target interface{}, tx *gorm.DB) error {
	res := tx.Debug().Save(target)

	return r.HandleError(res)
}

func (r *gormRepository) UpdateOrCreateTx(target interface{}, attributes map[string]interface{}, values map[string]interface{}, tx *gorm.DB) error {
	res := tx.Debug().Where(attributes).Assign(values).FirstOrCreate(target)

	return r.HandleError(res)
}

func (r *gormRepository) Delete(target interface{}) error {
	res := r.DB().Delete(target)

	return r.HandleError(res)
}

func (r *gormRepository) DeleteTx(target interface{}, tx *gorm.DB) error {
	res := tx.Debug().Delete(target)

	return r.HandleError(res)
}

func (r *gormRepository) DeleteTxByCondition(target interface{}, condition string, tx *gorm.DB) error {
	res := tx.Debug().Where(condition).Delete(target)

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

func (r *gormRepository) HandleOneError(res *gorm.DB) error {
	if err := r.HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return ErrRecordNotFound
	}

	return nil
}

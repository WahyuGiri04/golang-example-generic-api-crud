package baseService

import (
	"golang-example-generic-api-crud/config"
	baseModel "golang-example-generic-api-crud/model/base"
	"math"

	"gorm.io/gorm"
)


type BaseService[T any] struct {
	DB *gorm.DB
}

func NewBaseService[T any](db *gorm.DB) *BaseService[T] {
	return &BaseService[T]{DB: config.DB}
}

func (s *BaseService[T]) Create(entity *T) error {
	return s.DB.Create(entity).Error
}

func (s * BaseService[T]) GetAll(entities *[]T) error {
	return s.DB.Find(entities).Error
}

func (s *BaseService[T]) GetById(id uint, entity *T) error{
	return s.DB.First(entity, id).Error
}

func (s *BaseService[T]) Update(id uint, entity *T) error{
	return s.DB.Model(entity).Where("id = ?", id).Updates(entity).Error
}

func (s *BaseService[T]) Delete(id uint) error{
	return s.DB.Delete(new(T), id).Error
}

func (s *BaseService[T]) GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error){
	var totalRows int64
	s.DB.Model(new(T)).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	offset := (page - 1) * pageSize

	err := s.DB.Limit(pageSize).Offset(offset).Find(entities).Error
	if err != nil {
		return baseModel.Pagination{}, err
	}

	return baseModel.Pagination{
		Page: page,
		PageSize: pageSize,
		TotalRows: totalRows,
		TotalPages: totalPages,
		Data: entities,
	}, nil
}
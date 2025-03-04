package service

import (
	"golang-example-generic-api-crud/model"
	baseService "golang-example-generic-api-crud/service/base"
)

type RoleService struct {
	*baseService.BaseService[model.Role]
}

func NewRoleService() *RoleService {
	return &RoleService{
		BaseService: baseService.NewBaseService[model.Role](nil),
	}
}

// func (s *RoleService) GetByName(name string) ([]model.Role, error) {
// 	var roles []model.Role
// 	err := s.DB.Where("name LIKE ? ", "%"+name+"%").Find(&roles).Error
// 	return roles, err
// }

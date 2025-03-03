package service

import (
	"golang-example-generic-api-crud/model"
	baseService "golang-example-generic-api-crud/service/base"
)


type RoleService struct{
	*baseService.BaseService[model.Role]
}

func NewRoleService() *RoleService {
	return &RoleService{
		BaseService: baseService.NewBaseService[model.Role](nil),
	}
}

func (s *RoleService) GetByName(name string) (model.Role, error) {
	var role model.Role
	err := s.DB.Where("name = ?", name).First(&role).Error
	return role, err
}
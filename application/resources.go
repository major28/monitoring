package application

import (
	"github.com/jinzhu/gorm"
)

type IResources interface {
	GetDbManager() *gorm.DB
}

// ResourceCollection contains resources
type ResourceCollection struct {
	DbManager *gorm.DB
}

func (s *ResourceCollection) GetDbManager() *gorm.DB {
	return s.DbManager
}

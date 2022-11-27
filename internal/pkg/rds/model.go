package rds

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

//Model common model
type Model struct {
	ID        string                `json:"id"`
	CreatedAt time.Time             `json:"createdAt" example:"2021-04-28T05:22:10.916052Z"`
	UpdatedAt time.Time             `json:"updatedAt" example:"2021-04-28T05:22:10.916052Z"`
	IsDel     soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag"`
	Version   int64                 `json:"version"`
}

//NewModel init model
func NewModel(id string) Model {
	return Model{ID: id, Version: 1}
}

//BeforeUpdate update version
func (m Model) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.Where("version=?", m.Version)
	db.Statement.SetColumn("version", m.Version+1)
	return nil
}

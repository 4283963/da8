package models

import (
	"time"
)

type JadeBracelet struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null;column:name" json:"name" binding:"required"`
	Material     string    `gorm:"type:varchar(50);not null;column:material" json:"material" binding:"required"`
	Translucency float64   `gorm:"type:decimal(5,2);not null;column:translucency" json:"translucency" binding:"required,min=0,max=100"`
	Fineness     float64   `gorm:"type:decimal(5,2);not null;column:fineness" json:"fineness" binding:"required,min=0,max=100"`
	BeadCount    int       `gorm:"type:int;not null;default:17;column:bead_count" json:"bead_count" binding:"required,min=1"`
	Score        float64   `gorm:"type:decimal(5,2);column:score" json:"score"`
	Grade        string    `gorm:"type:varchar(20);column:grade" json:"grade"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (JadeBracelet) TableName() string {
	return "jade_bracelets"
}

type CreateBraceletRequest struct {
	Name         string  `json:"name" binding:"required"`
	Material     string  `json:"material" binding:"required"`
	Translucency float64 `json:"translucency" binding:"required,min=0,max=100"`
	Fineness     float64 `json:"fineness" binding:"required,min=0,max=100"`
	BeadCount    int     `json:"bead_count" binding:"required,min=1"`
}

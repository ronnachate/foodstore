package domain

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Lastname  string    `gorm:"type:varchar(100)"`
	Telephone string    `gorm:"type:varchar(30)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

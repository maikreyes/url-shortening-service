package domain

import (
	"time"

	"github.com/google/uuid"
)

type ApiResponseWithotStats struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"-"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	Url       string    `gorm:"type:text;not null" json:"url"`
	ShortCode string    `gorm:"size:20;uniqueIndex;not null" json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

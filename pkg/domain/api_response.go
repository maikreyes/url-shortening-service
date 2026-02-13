package domain

import (
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ApiResponse struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID      uuid.UUID `gorm:"type:char(36);not null;index" json:"-"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	Url         string    `gorm:"type:text;not null" json:"url"`
	ShortCode   string    `gorm:"size:20;not null" json:"shortCode"`
	IsWebhook   bool      `gorm:"default:false" json:"isWebhook"`
	AccessCount int       `gorm:"default:0" json:"accessCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (ApiResponse) TableName() string {
	if table := strings.TrimSpace(os.Getenv("URL_TABLE_NAME")); table != "" {
		return table
	}
	return "api_responses"
}

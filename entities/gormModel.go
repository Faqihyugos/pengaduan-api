package entities

import "time"

type GormModel struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Created_At time.Time `json:"created_at,omitempty"`
	Updated_At time.Time `json:"updated_at,omitempty"`
	Deleted_At time.Time `json:"deleted_at,omitempty"`
}

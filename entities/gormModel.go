package entities

import "time"

type GormModel struct {
	ID         uint      `gorm:"primaryKey" json:"id,omiempty" form:"id,omitempty"`
	Created_At time.Time `json:"created_at,omitempty" form:"created_at,omitempty"`
	Updated_At time.Time `json:"updated_at,omitempty" form:"updated_at,omitempty"`
	Deleted_At time.Time `json:"deleted_at,omitempty" form:"deleted_at,omitempty"`
}

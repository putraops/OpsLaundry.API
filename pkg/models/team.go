package models

import (
	"time"
)

type Team struct {
	Id             string     `gorm:"type:varchar(50);primary_key" json:"id"`
	IsActive       *bool      `gorm:"type:bool" json:"is_active"`
	IsLocked       *bool      `gorm:"type:bool" json:"is_locked"`
	IsDefault      *bool      `gorm:"type:bool" json:"is_default"`
	OwnerId        string     `gorm:"type:varchar(50)" json:"owner_id"`
	CreatedAt      *time.Time `gorm:"type:timestamp" json:"created_at" time_format:"2006-01-02" time_utc:"0" `
	CreatedBy      string     `gorm:"type:varchar(50)" json:"created_by"`
	UpdatedAt      *time.Time `gorm:"type:timestamp" json:"updated_at" time_format:"2006-01-02" time_utc:"0"`
	UpdatedBy      string     `gorm:"type:varchar(50)" json:"updated_by"`
	SubmittedAt    *time.Time `gorm:"type:timestamp;default:null" json:"submitted_at" time_format:"2006-01-02" time_utc:"0"`
	SubmittedBy    string     `gorm:"type:varchar(50)" json:"submitted_by"`
	ApprovedAt     *time.Time `gorm:"type:timestamp;default:null" json:"approved_at" time_format:"2006-01-02" time_utc:"0"`
	ApprovedBy     string     `gorm:"type:varchar(50)" json:"approved_by"`
	EntityId       string     `gorm:"type:varchar(50);null" json:"entity_id"`
	OrganizationId string     `gorm:"type:varchar(50);null" json:"organization_id"`
	Name           string     `gorm:"type:varchar(50);unique" json:"name"`
	Description    string     `gorm:"type:text" json:"description"`

	Organization Organization `gorm:"foreignkey:OrganizationId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"organization"`
}

func (Team) TableName() string {
	return "team"
}

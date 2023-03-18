package models

import (
	"errors"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
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

	Name        string `gorm:"type:text" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Organization Organization `gorm:"foreignkey:OrganizationId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}

func (r *ProductCategory) BeforeCreate(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsAdmin && !user.IsSystemAdmin {
		return errors.New(constants.NoPermissionCreateMessage)
	}

	r.Id = uuid.New().String()
	r.IsActive = &isActive
	r.IsDefault = &isDefault
	r.IsLocked = &isLocked
	r.OwnerId = user.UserId
	r.CreatedBy = user.UserId
	r.CreatedAt = &dateNow
	r.OrganizationId = user.OrganizationId
	return
}

func (r *ProductCategory) BeforeSave(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsAdmin && !user.IsSystemAdmin {
		return errors.New(constants.NoPermissionUpdateMessage)
	}

	r.UpdatedAt = &dateNow
	r.UpdatedBy = user.UserId
	return
}

func (r *ProductCategory) BeforeDelete(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsSystemAdmin && !user.IsAdmin {
		return errors.New(constants.NoPermissionDeleteMessage)
	}

	if user.IsAdmin && user.OrganizationId != r.OrganizationId {
		return errors.New(constants.NoPermissionDeleteAnotherOrganizationMessage)
	}

	return
}

package models

import (
	"errors"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductDetail struct {
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
	ItemName       string     `gorm:"type:text" json:"item_name"`
	ProductId      string     `gorm:"type:varchar(50)" json:"product_id"`
	ServiceTypeId  string     `gorm:"type:text" json:"service_type_id"`
	UomId          string     `gorm:"type:text" json:"uom_id"`

	Organization Organization `gorm:"foreignkey:OrganizationId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Product      Product      `gorm:"foreignkey:ProductId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	ServiceType  ServiceType  `gorm:"foreignkey:ServiceTypeId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Uom          Uom          `gorm:"foreignkey:UomId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`

	// ProductTypeId  string     `gorm:"type:text" json:"product_type_id"`
	// ProductType  ProductType  `gorm:"foreignkey:ProductTypeId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

func (ProductDetail) TableName() string {
	return "product_detail"
}

func (r *ProductDetail) BeforeCreate(db *gorm.DB) (err error) {
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

func (r *ProductDetail) BeforeSave(db *gorm.DB) (err error) {
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

func (r *ProductDetail) BeforeDelete(db *gorm.DB) (err error) {
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

package models

import (
	"errors"
	"fmt"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Uom struct {
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
	UomName        string     `gorm:"type:varchar(50);unique" json:"uom_name"`
	UomCode        string     `gorm:"type:text" json:"uom_code"`
	UomSymbol      string     `gorm:"type:text" json:"uom_symbol"`

	Organization Organization `gorm:"foreignkey:OrganizationId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

func (Uom) TableName() string {
	return "uom"
}

func (r *Uom) BeforeCreate(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsSystemAdmin {
		return errors.New(fmt.Sprintf("%v. Hanya bisa ditambah oleh System Administrator", constants.NoPermissionCreateMessage))
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

func (r *Uom) BeforeSave(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsSystemAdmin {
		return errors.New(fmt.Sprintf("%v. Hanya bisa diubah oleh System Administrator", constants.NoPermissionUpdateMessage))
	}

	r.UpdatedAt = &dateNow
	r.UpdatedBy = user.UserId
	return
}

func (r *Uom) BeforeDelete(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if !user.IsSystemAdmin {
		return errors.New(fmt.Sprintf("%v. Hanya bisa dihapus oleh System Administrator", constants.NoPermissionDeleteMessage))
	}

	return
}

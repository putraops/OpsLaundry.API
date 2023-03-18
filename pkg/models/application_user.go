package models

import (
	"context"
	"errors"
	"fmt"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationUser struct {
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
	OrganizationId string     `gorm:"type:varchar(50);null" json:"organization_id"`
	TenantId       string     `gorm:"type:varchar(50);null" json:"tenant_id"`

	FbUid             string `gorm:"type:text" json:"fb_uid"`
	FirstName         string `gorm:"type:varchar(50)" json:"first_name"`
	LastName          string `gorm:"type:varchar(50)" json:"last_name"`
	Username          string `gorm:"type:varchar(50)" json:"username"`
	Title             string `gorm:"type:varchar(50)" json:"title"`
	Password          string `gorm:"->;<-;" json:"-"`
	Address           string `gorm:"type:text" json:"address"`
	Phone             string `gorm:"type:varchar(15)" json:"phone"`
	Email             string `gorm:"type:varchar(255)" json:"email"`
	TotalPoint        int    `gorm:"type:int;default:0" json:"total_point"`
	IsEmailVerified   bool   `gorm:"type:bool" json:"is_email_verified"`
	IsPhoneVerified   bool   `gorm:"type:bool" json:"is_phone_verified"`
	IsSystemAdmin     bool   `gorm:"type:bool" json:"is_system_admin"`
	IsAdmin           bool   `gorm:"type:bool" json:"is_admin"`
	UserType          int    `gorm:"type:int" json:"user_type"`
	Gender            bool   `gorm:"type:bool" json:"gender"`
	Filepath          string `gorm:"type:varchar(200)" json:"filepath"`
	FilepathThumbnail string `gorm:"type:varchar(200)" json:"filepath_thumbnail"`
	Filename          string `gorm:"type:varchar(200)" json:"filename"`
	Extension         string `gorm:"type:varchar(10)" json:"extension"`
	Size              string `gorm:"type:varchar(100)" json:"size"`
	Description       string `gorm:"type:text" json:"description"`
	Token             string `gorm:"-" json:"token,omitempty"`

	Organization Organization `gorm:"foreignkey:OrganizationId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

func (ApplicationUser) TableName() string {
	return "application_user"
}

func (r *ApplicationUser) BeforeCreate(db *gorm.DB) (err error) {
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

func (r *ApplicationUser) BeforeSave(db *gorm.DB) (err error) {
	user, err := utils.GetUserIdentity(db.Statement.Context)
	if err != nil {
		return err
	}

	if r.OwnerId != user.UserId && !user.IsAdmin && !user.IsSystemAdmin {
		return errors.New(constants.NoPermissionUpdateMessage)
	}

	r.UpdatedAt = &dateNow
	r.UpdatedBy = user.UserId
	return
}

func (u *ApplicationUser) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind")
	// if u.MemberShip == "" {
	// 	u.MemberShip = "user"
	// }
	return
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

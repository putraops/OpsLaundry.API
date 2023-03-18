package repository

import (
	"encoding/json"
	"errors"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationUserRepository interface {
	Create(context *gin.Context, record models.ApplicationUser) (interface{}, error)
	Update(context *gin.Context, record models.ApplicationUser) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetProfile(context *gin.Context) (interface{}, error)
	GetUserByEmail(email string) (interface{}, error)
}

type applicationUserRepository struct {
	DB *gorm.DB
}

// NewApplicationUserRepository is creates a new instance of ApplicationUserRepository
func NewApplicationUserRepository(db *gorm.DB) ApplicationUserRepository {
	return &applicationUserRepository{
		DB: db,
	}
}

func (r applicationUserRepository) Create(context *gin.Context, record models.ApplicationUser) (interface{}, error) {
	if err := r.DB.WithContext(context.Request.Context()).Create(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r applicationUserRepository) Update(context *gin.Context, record models.ApplicationUser) (interface{}, error) {
	var obj models.ApplicationUser

	if err := r.DB.First(&obj, "id = ?", record.Id).Error; err != nil {
		return nil, err
	}

	mapResult, err := utils.MapFields(obj, record)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(mapResult, &obj); err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(context.Request.Context()).Save(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r applicationUserRepository) GetAll(context *gin.Context) (interface{}, error) {
	var records []views.ApplicationUser
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r applicationUserRepository) GetProfile(context *gin.Context) (interface{}, error) {
	userIdentity, err := utils.GetUserIdentity(context.Request.Context())
	if err != nil {
		return nil, err
	}

	var record views.ApplicationUser
	if err := r.DB.First(&record, "id = ?", userIdentity.UserId).Error; err != nil {
		return nil, err
	}

	if *record.IsLocked {
		return nil, errors.New(constants.UserLockedMessage)
	}

	return record, nil
}

func (r applicationUserRepository) GetUserByEmail(email string) (interface{}, error) {
	var record views.ApplicationUser
	if err := r.DB.Where("LOWER(email) = ?", strings.ToLower(email)).Take(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

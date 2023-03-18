package services

import (
	"errors"
	"net/http"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/dto"
	"opslaundry/pkg/models"
	"opslaundry/pkg/repository"
	"opslaundry/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

// -- This is user contract
type ApplicationUserService interface {
	Register(context *gin.Context, dto dto.RegisterDto) (interface{}, error)
	Update(context *gin.Context, record models.ApplicationUser) (interface{}, error)
	GetAll(context *gin.Context) (interface{}, error)
	GetProfile(context *gin.Context) (interface{}, error)
	// TestFirebase(context context.Context) *helper.StandartResult
	// Register(context *gin.Context, dto dto.ApplicationUserRegisterDto) helper.StandartResult
	// // Update(user dto.ApplicationUserUpdateDto) models.ApplicationUser
	// // UpdateProfile(dtoRecord dto.ApplicationUserDescriptionDto) helper.Result
	// // ChangePhone(recordDto dto.ChangePhoneDto) helper.Result
	// ChangePassword(recordDto dto.ChangePasswordDto) helper.StandartResult
	// // RecoverPassword(recordId string, oldPassword string) helper.Result

	// Lookup(request helper.ReactSelectRequest) helper.StandartResult
	// // UserProfile(recordId string) models.ApplicationUser

	// GetPagination(context *gin.Context, request commons.DataTableRequest) interface{}
	// GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult
	// GetViewAll(filter map[string]interface{}) helper.StandartResult
	// GetById(id string) helper.StandartResult
	// GetViewById(id string) helper.StandartResult

	// DeleteById(id string) helper.StandartResult
}

type applicationUserService struct {
	jwtService                JWTService
	applicationUserRepository repository.ApplicationUserRepository
}

func NewApplicationUserService(db *gorm.DB) ApplicationUserService {
	return &applicationUserService{
		applicationUserRepository: repository.NewApplicationUserRepository(db),
	}
}

func (r applicationUserService) Register(context *gin.Context, dto dto.RegisterDto) (interface{}, error) {
	var record models.ApplicationUser
	//smapping.FillStruct(&newRecord, smapping.MapFields(&record))
	// if dto.UserType == 1 {
	// 	newRecord.Password = helper.HashAndSalt([]byte(dto.Password))
	// }
	// newRecord.OrganizationId = dto.OrganizationId

	// r := token.Claims.(jwt.MapClaims)

	// fmt.Println(record)
	//fmt.Println(newRecord)
	// fmt.Println("in")
	//return helper.StandartResult{Success: true, Message: "", Data: newRecord}

	// if authHeader := context.GetHeader("Authorization"); authHeader != "" {
	// 	// return helper.StandartResult{Success: false, Message: fmt.Sprintf("%v,", err.Error()), Data: nil}
	// 	user = r.jwtService.GetUserByToken(authHeader)
	// }

	// return r.applicationUserRepository.CreateOrUpdate(context, record)

	// r.DB.Set("identity", userIdentity)

	if err := smapping.FillStruct(&record, smapping.MapFields(&dto)); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return nil, err
	}

	record.Password = utils.HashAndSalt([]byte(record.Password))
	return r.applicationUserRepository.Create(context, record)
}

func (r applicationUserService) Update(context *gin.Context, record models.ApplicationUser) (interface{}, error) {
	result, err := r.applicationUserRepository.Update(context, record)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.NoDataFound)
		}
		return nil, err
	}
	return result, nil
}

func (r applicationUserService) GetAll(context *gin.Context) (interface{}, error) {
	return r.applicationUserRepository.GetAll(context)
}

func (r applicationUserService) GetProfile(context *gin.Context) (interface{}, error) {
	return r.applicationUserRepository.GetProfile(context)
}

// // func (r applicationUserService) UpdateProfile(dtoRecord dto.ApplicationUserDescriptionDto) helper.Result {
// // 	return service.applicationUserRepository.UpdateProfile(dtoRecord)
// // }

// func (r applicationUserService) GetPagination(context *gin.Context, request commons.DataTableRequest) interface{} {
// 	return r.applicationUserRepository.GetPagination(context, request)
// }

// func (r applicationUserService) GetAll(context *gin.Context, filter map[string]interface{}) helper.StandartResult {
// 	return r.applicationUserRepository.GetAll(context, filter)
// }

// func (r applicationUserService) GetViewAll(filter map[string]interface{}) helper.StandartResult {
// 	return r.applicationUserRepository.GetViewAll(filter)
// }

// func (r applicationUserService) GetById(recordId string) helper.StandartResult {
// 	return r.applicationUserRepository.GetById(recordId)
// }

// func (r applicationUserService) GetViewById(recordId string) helper.StandartResult {
// 	return r.applicationUserRepository.GetViewById(recordId)
// }

// func (r applicationUserService) DeleteById(id string) helper.StandartResult {
// 	return r.applicationUserRepository.DeleteById(id)
// }

// func (r applicationUserService) ChangePassword(dto dto.ChangePasswordDto) helper.StandartResult {
// 	res := r.applicationUserRepository.GetViewById(dto.Id)
// 	if !res.Success {
// 		return res
// 	}

// 	if v, ok := res.Data.(entity_view_models.EntityApplicationUserView); ok {
// 		comparedPassword := comparePassword(v.Password, []byte(dto.OldPassword))
// 		if (v.Id == dto.Id) && comparedPassword {
// 			success, message := r.applicationUserRepository.SetPasswordById(v.Id, helper.HashAndSalt([]byte(dto.NewPassword)))
// 			if !success {
// 				return helper.StandartResult{Success: false, Message: message, Data: nil}
// 			}
// 		} else {
// 			return helper.StandartResult{Success: false, Message: helper.PasswordNotMatchMessage, Data: nil}
// 		}
// 	}

// 	//-- Update Token
// 	var user = entity_view_models.EntityApplicationUserView{}
// 	user = res.Data.(entity_view_models.EntityApplicationUserView)
// 	generatedToken := r.jwtService.GenerateToken(user, dto.NewPassword)
// 	user.Token = generatedToken
// 	return helper.StandartResult{Success: true, Message: helper.ChangePasswordSuccessMessage, Data: user}
// }

// // func setUserContext(jwtService JWTService, context *gin.Context) *gin.Context {
// // 	if authHeader := context.GetHeader("Authorization"); authHeader != "" {
// // 		var user = jwtService.GetUserByToken(authHeader)
// // 		context.Set("user", user)
// // 		return context
// // 	}
// // 	return context
// // }

// // func (r applicationUserService) ChangePhone(recordDto dto.ChangePhoneDto) helper.Result {
// // 	res := service.applicationUserRepository.GetById(recordDto.Id)
// // 	if !res.Status {
// // 		return res
// // 	}

// // 	var userData = (res.Data).(models.ApplicationUser)
// // 	userData.Password = "" //-- Set to empty :: use for restrict change password
// // 	userData.Phone = recordDto.Phone

// // 	_ = service.applicationUserRepository.Update(userData)
// // 	return helper.StandartResult(true, "Ok", helper.EmptyObj{})
// // }

// // func (r applicationUserService) RecoverPassword(recordId string, oldPassword string) helper.Result {
// // 	return service.applicationUserRepository.RecoverPassword(recordId, oldPassword)
// // }

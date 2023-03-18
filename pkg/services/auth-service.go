package services

import (
	"errors"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/dto"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/repository"
	"opslaundry/pkg/utils"

	"gorm.io/gorm"
)

type AuthService interface {
	VerifyCredential(dto dto.LoginDto) (interface{}, error)
}

type authService struct {
	DB                        *gorm.DB
	jwtService                JWTService
	applicationUserRepository repository.ApplicationUserRepository
}

func NewAuthService(db *gorm.DB, jwtService JWTService) AuthService {
	return &authService{
		DB:                        db,
		jwtService:                jwtService,
		applicationUserRepository: repository.NewApplicationUserRepository(db),
	}
}

func (r authService) VerifyCredential(dto dto.LoginDto) (interface{}, error) {
	obj, err := r.applicationUserRepository.GetUserByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotFoundMessage)
		}
		return obj, err
	}

	user := (obj).(views.ApplicationUser)
	if *user.IsActive == false {
		return nil, errors.New(constants.UserNotActiveMessage)
	}

	if *user.IsLocked {
		return nil, errors.New(constants.UserLockedMessage)
	}

	comparedPassword := utils.ComparePassword(user.Password, []byte(dto.Password))
	if (user.Email == dto.Email) && !comparedPassword {
		return nil, errors.New(constants.PasswordNotMatchMessage)
	}

	generatedToken := r.jwtService.GenerateToken(user, dto.Password)
	user.Token = "Bearer " + generatedToken

	return user, nil
}

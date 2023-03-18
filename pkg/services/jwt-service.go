package services

import (
	// "opslaundry/models/entity_view_models"
	"opslaundry/pkg/commons"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"fmt"
	// "time"
)

// -- JWTService is What is JWT can do
type JWTService interface {
	GenerateToken(user views.ApplicationUser, originPassword string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByToken(token string) utils.UserIdentity
}

type jwtCustomClaim struct {
	UserId         string `json:"user_id"`
	IsAdmin        bool   `json:"is_admin"`
	IsSystemAdmin  bool   `json:"is_system_admin"`
	IsLocked       bool   `json:"is_locked"`
	OrganizationId string `json:"organization_id"`
	TenantId       string `json:"tenant_id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	// *auth.UserRecord
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
	logger    commons.OpsLogger
}

// -- New Instance JWT Service
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "opslaundry",
		secretKey: getSecretKey(),
		logger:    commons.NewLogger(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRETKEY")
	if secretKey != "" {
		secretKey = "createdbyputraops"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(user views.ApplicationUser, originPassword string) string {
	claims := &jwtCustomClaim{
		user.Id,
		user.IsAdmin,
		user.IsSystemAdmin,
		*user.IsLocked,
		user.OrganizationId,
		user.TenantId,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (c *jwtService) GetUserByToken(token string) utils.UserIdentity {
	var identity utils.UserIdentity
	aToken, err := c.ValidateToken(token)
	if err != nil {
		log.Error("jwt-service.go:GetUserByToken")
		log.Error(err.Error())
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	identity.Token = token
	identity.UserId = fmt.Sprintf("%v", claims["user_id"])
	identity.IsSystemAdmin = false
	identity.IsAdmin = false
	if claims["is_system_admin"].(bool) == true {
		identity.IsSystemAdmin = true
		identity.IsAdmin = true
	} else if claims["is_admin"].(bool) == true {
		identity.IsAdmin = true
	}
	identity.OrganizationId = fmt.Sprintf("%v", claims["organization_id"])
	identity.TenantId = fmt.Sprintf("%v", claims["tenant_id"])
	identity.Username = fmt.Sprintf("%v", claims["username"])
	identity.FirstName = fmt.Sprintf("%v", claims["first_name"])
	identity.LastName = fmt.Sprintf("%v", claims["last_name"])
	identity.Email = fmt.Sprintf("%v", claims["email"])
	return identity
}

package utils

import (
	"context"
	"fmt"
	"log"
	"opslaundry/pkg/constants"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserIdentity struct {
	UserId         string
	IsAdmin        bool
	IsSystemAdmin  bool
	IsLocked       bool
	OrganizationId string
	TenantId       string
	Username       string
	FirstName      string
	LastName       string
	Email          string
	Token          string
}

func GetUserIdentity(context context.Context) (*UserIdentity, error) {
	ginContext := context.Value(constants.GIN_CONTEXT_KEY).(*gin.Context)
	if ginContext == nil {
		return nil, fmt.Errorf("could not retrieve gin.Context")
	}

	userIdentity := ginContext.Value("USER_IDENTITY").(UserIdentity)
	return &userIdentity, nil
}

func ErrorToArray(err string) []string {
	return strings.Split(err, "\n")
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

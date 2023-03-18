package middleware

import (
	"context"
	"opslaundry/pkg/constants"
	"opslaundry/pkg/services"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": constants.NoTokenFound})
			return
		}

		token := strings.Split(authorizationHeader, " ")
		if token[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, "Token Key is invalid!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token Key is not valid"})
			return
		}

		validate, err := jwtService.ValidateToken(token[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
			return
		}

		if !validate.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token is invalid"})
			return
		}

		userIdentity := jwtService.GetUserByToken(token[1])
		if userIdentity.IsLocked {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": constants.UserLockedMessage})
			return
		}

		c.Set(constants.USER_IDENTITY, userIdentity)
		ctx := context.WithValue(c.Request.Context(), constants.GIN_CONTEXT_KEY, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// func ClaimToken(jwtService services.JWTService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			response := helper.Result(false, helper.NoTokenFound, nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}
// 		token, err := jwtService.ValidateToken(authHeader)

// 		//print(token)
// 		if !token.Valid {
// 			log.Println(err)
// 			response := helper.Result(false, "Token is not valid", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			// } else {
// 			// 	claims := token.Claims.(jwt.MapClaims)
// 			// 	log.Println("Claim[user_id]: ", claims["claim_user"])
// 		}
// 	}
// }

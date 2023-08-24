package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/user/usecase"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

func AuthMiddleware(auth JWTAuthentication, userUseCase usecase.UserUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"error": err.Error()}
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", errorMessage)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userUseCase.GetUserById(userId)
		if err != nil {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// userRole := int(claim["role"].(float64))

		// if userRole != 0 || userRole != 1 {
		// 	if userRole != user.Role {
		// 		response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 		return
		// 	}
		// 	response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		ctx.Set("currentUser", user)
	}
}

func AuthAdminMiddleware(auth JWTAuthentication, userUseCase usecase.UserUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userUseCase.GetUserById(userId)
		if err != nil {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userRole := int(claim["role"].(float64))

		if userRole != 0 && userRole != user.Role {
			response := customresponse.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("currentUser", user)
	}
}

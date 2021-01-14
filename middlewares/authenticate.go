package middlewares

import (
	"document-service-gin/helpers"
	"document-service-gin/models"
	"document-service-gin/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// Authenticate fetches user details from token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		requiredToken := c.GetHeader("Authorization")
		requiredToken = strings.Replace(requiredToken, "Bearer ", "", -1)

		if len(requiredToken) == 0 {
			helpers.ResponseWithError(c, 403, "Please login to your account")
			c.Abort()
			return
		}

		UserId, CompanyId, err := services.DecodeToken(requiredToken)
		var dataUser models.User
		user_id,err := strconv.ParseInt(UserId, 10, 64)
		company_id,err := strconv.ParseInt(CompanyId, 10, 64)
		dataUser.User_id = int32(user_id)
		dataUser.Company_id = int32(company_id)

		if err != nil {
			helpers.ResponseWithError(c, 500, "Something went wrong giving you access")
			c.Abort()
			return
		}

		c.Set("User", dataUser)

		c.Next()
	}
}
package folders

import (
	"document-service-gin/db"
	"document-service-gin/forms"
	helper "document-service-gin/helpers"
	"document-service-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (f *FolderController) DeleteFolderById(c *gin.Context)  {
	user := c.MustGet("User").(models.User)

	if user.User_id == 0 {
		helper.ResponseWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	var data forms.DeleteFolderCommand
	fmt.Println(data)
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Provide relevant fields"})
		c.Abort() // abort the request
		return
	}

	err := folderModel.DeleteFolderById(data.Id)
	if err != nil && err.Error() == "not found" {
		helper.ResponseWithError(c, 404, err.Error())
		c.Abort()
		return
	}else if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	redisClient := db.Initialize()
	keyForRedis := "folder/" + string(data.Id)
	err = redisClient.DelKey(c, keyForRedis)
	if err != nil {
		helper.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	helper.ResponseSuccess(c, 200, "Success delete document", nil)
}
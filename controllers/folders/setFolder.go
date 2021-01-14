package folders

import (
	"document-service-gin/forms"
	helper "document-service-gin/helpers"
	"document-service-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (f *FolderController) SetFolder(c *gin.Context)  {
	user := c.MustGet("User").(models.User)

	if user.User_id == 0 {
		helper.ResponseWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	var data forms.SetFolderCommand
	fmt.Println(data)
	if data.Owner_id == 0 {
		data.Owner_id = user.User_id
	}
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Provide relevant fields"})
		c.Abort() // abort the request
		return
	}

	err := folderModel.SetFolder(data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	result, err := folderModel.GetFolderById(string(data.Id))
	if err != nil {
		helper.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	helper.ResponseSuccess(c, 201, "Success set folder", result)
}
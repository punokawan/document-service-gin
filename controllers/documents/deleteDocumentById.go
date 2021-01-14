package documents

import (
	"document-service-gin/forms"
	helper "document-service-gin/helpers"
	"document-service-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (d *DocumentController) DeleteDocumentById(c *gin.Context)  {
	user := c.MustGet("User").(models.User)

	if user.User_id == 0 {
		helper.ResponseWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	var data forms.DeleteDocumentCommand
	fmt.Println(data)
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Provide relevant fields"})
		c.Abort() // abort the request
		return
	}

	err := documentModel.DeleteDocumentById(data.Id)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	helper.ResponseSuccess(c, 200, "Success delete document", nil)
}
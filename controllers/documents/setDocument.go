package documents

import (
	"document-service-gin/forms"
	helper "document-service-gin/helpers"
	"document-service-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (d *DocumentController) SetDocument(c *gin.Context)  {
	user := c.MustGet("User").(models.User)

	if user.User_id == 0 {
		helper.ResponseWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	var data forms.SetDocumentCommand
	fmt.Println(data)
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Provide relevant fields"})
		c.Abort() // abort the request
		return
	}

	err := documentModel.SetDocument(data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	result, err := documentModel.GetDocumentById(string(data.Id))
	if err != nil {
		helper.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	helper.ResponseSuccess(c, 201, "Success set document", result)
}
package folders

import (
	"document-service-gin/db"
	"document-service-gin/helpers"
	"document-service-gin/models"
	_ "document-service-gin/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func (f *FolderController) GetDocumentsByFolderId(c *gin.Context)  {
	documentId := c.Param("id")
	isValidUUID := helpers.IsValidUUID(documentId)
	if !isValidUUID {
		helpers.ResponseWithError(c, 400, "id required and must an UUID")
		c.Abort()
		return
	}

	keyForRedis := "folder/" + string(documentId)

	dataFromRedis := []models.Document{}
	val,err := db.Client.Get(c, keyForRedis).Result()
	if err != redis.Nil {
		err = json.Unmarshal([]byte(val), &dataFromRedis)
		if err != nil {
			message := "error unmarshall: "+ err.Error()
			helpers.ResponseWithError(c, 500, message)
			c.Abort()
			return
		}
		helpers.ResponseSuccess(c, 200, "Success get list from redis documents folder", dataFromRedis)
		c.Abort()
		return
	}

	result, err := folderModel.GetDocumentByFolderId(string(documentId))
	if err != nil {
		helpers.ResponseWithError(c, 404, err.Error())
		c.Abort()
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Client.Set(c, keyForRedis, data,0).Err()
	if err != nil {
		fmt.Println(err)
	}

	helpers.ResponseSuccess(c, 200, "Success get list documents folder", result)
}
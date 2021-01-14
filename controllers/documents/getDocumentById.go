package documents

import (
	"document-service-gin/db"
	"document-service-gin/helpers"
	"document-service-gin/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func (d *DocumentController) GetDocumentById(c *gin.Context)  {
	documentId := c.Param("id")
	isValidUUID := helpers.IsValidUUID(documentId)
	if !isValidUUID {
		helpers.ResponseWithError(c, 400, "id required and must an UUID")
		c.Abort()
		return
	}

	redisClient := db.Initialize()
	keyForRedis := "document/" + string(documentId)

	dataFromRedis := models.Document{}
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

	result, err := documentModel.GetDocumentById(string(documentId))
	if err != nil {
		helpers.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	err = redisClient.SetKey(c, keyForRedis, result, 0)
	if err != nil {
		helpers.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	helpers.ResponseSuccess(c, 200, "Success get list documents folder", result)
}
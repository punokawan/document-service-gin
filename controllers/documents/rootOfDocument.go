package documents

import (
	"document-service-gin/db"
	"document-service-gin/helpers"
	"document-service-gin/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func (d *DocumentController) GetRootOfDocument(c *gin.Context)  {
	user := c.MustGet("User").(models.User)

	if user.User_id == 0 {
		helpers.ResponseWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	redisClient := db.Initialize()
	keyForRedis := "root/" + strconv.Itoa(int(user.Company_id))

	dataFromRedis := []models.RootDocument{}
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

	folders, err := folderModel.GetFolderByOwnerId(string(user.Company_id))
	if err != nil {
		helpers.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	documents, err := documentModel.GetDocumentByOwnerId(string(user.Company_id))
	if err != nil {
		helpers.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	result := append(folders, documents...)

	err = redisClient.SetKey(c, keyForRedis, result, 0)
	if err != nil {
		helpers.ResponseWithError(c, 500, err.Error())
		c.Abort()
		return
	}

	helpers.ResponseSuccess(c, 200, "Success get list documents folder", result)
}
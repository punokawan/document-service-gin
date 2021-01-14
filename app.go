package main

import (
	"document-service-gin/controllers/documents"
	"document-service-gin/controllers/folders"
	"document-service-gin/db"
	"document-service-gin/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

}

func main() {
	pong, err := db.Client.Ping(db.Client.Context()).Result()
	fmt.Println(pong, err)

	// Init gin router
	router := gin.Default()

	// Define the document controller
	document := new(documents.DocumentController)
	// Define the folder controller
	folder := new(folders.FolderController)

	// Its great to version your API's
	basePath := router.Group("/document-service")
	basePath.Use(middlewares.Authenticate())
	basePath.GET("/", document.GetRootOfDocument)
	{

		documentRoute := basePath.Group("/document")
		{
			documentRoute.GET("/:id", document.GetDocumentById)
			documentRoute.POST("/", document.SetDocument)
			documentRoute.DELETE("/", document.DeleteDocumentById)
		}

		folderRoute := basePath.Group("/folder")
		{
			folderRoute.GET("/:id", folder.GetDocumentsByFolderId)
			folderRoute.POST("/", folder.SetFolder)
			folderRoute.DELETE("/", folder.DeleteFolderById)
		}
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// Init our server
	router.Run(":5000")
}

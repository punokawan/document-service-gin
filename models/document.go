package models

import (
	"document-service-gin/forms"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// Document defines document structure
type Document struct {
	Id		string `json:"id" bson:"id"`
	Name	string `json:"name" bson:"name"`
	Type	string `json:"type" bson:"type"`
	Folder_id string `json:"folder_id" bson:"folder_id"`
	Content map[string] []map[string] string	`json:"content" bson:"content"`
	Owner_id	int32 `json:"owner_id" bson:"owner_id"`
	Share	[]int32 `json:"share" bson:"share"`
	Company_id int32 `json:"company_id" bson:"company_id"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"'`
	Is_public bool `json:"is_public" bson:"is_public"`
}

// DocumentModel defines the model structure
type DocumentModel struct {}

// GetDocumentById handles fetching user by email
func (d *DocumentModel) GetDocumentById(Id string) ( document Document, err error) {
	// Connect to the document collection
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "documents")
	// Assign result to error object while saving document
	err = collection.Find(bson.M{"id": Id}).One(&document)
	if err != nil {
		log.Fatal("Mongo collection find fail: ", err)
	}
	return document, err
}

// SetDocument handles create or update document
func (d *DocumentModel) SetDocument(data forms.SetDocumentCommand) error {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "documents")
	update := bson.M{
			"id":     data.Id,
			"name":    data.Name,
			"type": "document",
			"folder_id": data.Folder_id,
			"content": data.Content,
			"owner_id": data.Owner_id,
			"share": data.Share,
			"company_id": data.Company_id,
			"timestamp": data.Timestamp,
			"is_public": data.Is_public,
		}
	selector := bson.M{"id": data.Id}
	_, err := collection.Upsert(selector, update)

	return err
}

// GetDocumentById handles fetching user by email
func (d *DocumentModel) DeleteDocumentById(Id string) ( err error) {
	// Connect to the document collection
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "documents")
	// Assign result to error object while saving document
	selector := bson.M{"id": Id}
	err = collection.Remove(selector)

	return err
}

func (d *DocumentModel) GetDocumentByOwnerId(CompanyId string) (folder []RootDocument, err error) {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "documents")
	query := bson.M{
		"$or": []interface{}{
			bson.M{"is_public": true},
			bson.M{"company_id": CompanyId},
		},
	}
	err = collection.Find(query).All(&folder)
	if err != nil {
		log.Fatal("Mongo collection find fail: ", err)
	}
	return folder, err
}
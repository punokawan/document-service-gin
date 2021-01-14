package models

import (
	"document-service-gin/forms"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// Folder defines document structure
type Folder struct {
	Id		string `json:"id" bson:"id"`
	Name	string `json:"name" bson:"name"`
	Type	string `json:"type" bson:"type"`
	Is_public bool `json:"is_public" bson:"is_public"`
	Owner_id	int32 `json:"owner_id" bson:"owner_id"`
	Share	[]int32 `json:"share" bson:"share"`
	Company_id int32 `json:"company_id" bson:"company_id"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
}

type FolderModel struct {}

func (f *FolderModel) GetDocumentByFolderId(Id string) (document []Document, err error) {
	// Connect to the document collection
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "documents")
	// Assign result to error object while saving document
	err = collection.Find(bson.M{"folder_id": Id}).All(&document)
	if err != nil {
		log.Fatal("Mongo collection find fail: ", err)
	}
	return document, err
}

func (f *FolderModel) GetFolderById(Id string) (folder []Folder, err error) {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "folders")
	err = collection.Find(bson.M{"id": Id}).All(&folder)
	if err != nil {
		log.Fatal("Mongo collection find fail: ", err)
	}
	return folder, err
}

// SetFolder handles create or update folder
func (f *FolderModel) SetFolder(data forms.SetFolderCommand) error {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "folders")
	update := bson.M{
		"id":     data.Id,
		"name":    data.Name,
		"type": "folder",
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

func (f *FolderModel) DeleteFolderById(Id string) (err error) {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "folders")
	selector := bson.M{"id": Id}
	err = collection.Remove(selector)

	return err
}

func (f *FolderModel) GetFolderByOwnerId(CompanyId string) (folder []RootDocument, err error) {
	collection := DbConnect.Use(os.Getenv("DATABASE_NAME"), "folders")
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
package forms

type SetDocumentCommand struct {
	Id		string `json:"id" bson:"id" binding:"required"`
	Name	string `json:"name" bson:"name"`
	Type	string `json:"type" bson:"type"`
	Folder_id string `json:"folder_id" bson:"folder_id"`
	Content map[string] []map[string] string	`json:"content" bson:"content"`
	Owner_id	int32 `json:"owner_id" bson:"owner_id"`
	Share	[]int32 `json:"share" bson:"share"`
	Company_id int32 `json:"company_id" bson:"company_id"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
	Is_public bool `json:"is_public" bson:"is_public"`
}

type DeleteDocumentCommand struct {
	Id		string `json:"id" bson:"id" binding:"required"`
}
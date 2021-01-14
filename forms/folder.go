package forms

type SetFolderCommand struct {
	Id		string `json:"id" bson:"id" binding:"required"`
	Name	string `json:"name" bson:"name"`
	Type	string `json:"type" bson:"type"`
	Owner_id	int32 `json:"owner_id" bson:"owner_id"`
	Share	[]int32 `json:"share" bson:"share"`
	Company_id int32 `json:"company_id" bson:"company_id"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
	Is_public bool `json:"is_public" bson:"is_public"`
}

type DeleteFolderCommand struct {
	Id		string `json:"id" bson:"id" binding:"required"`
}
package models

// RootDocument defines document structure
type RootDocument struct {
	Id		string `json:"id" bson:"id"`
	Name	string `json:"name" bson:"name"`
	Type	string `json:"type" bson:"type"`
	Is_public bool `json:"is_public" bson:"is_public"`
	Owner_id	int32 `json:"owner_id" bson:"owner_id"`
	Share	[]int32 `json:"share" bson:"share"`
	Company_id int32 `json:"company_id" bson:"company_id"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
}
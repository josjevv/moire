package db

import "gopkg.in/mgo.v2/bson"

const (
	ASSET      = "asset"
	PENDING    = "pending"
	PROCESSING = "processing"
	READY      = "ready"
	LOST       = "lost"
)

type Asset struct {
	History       []string
	Id            bson.ObjectId     `json:"_id,omitempty" bson:"_id",omitempty`
	Size          int64             `json:"size" bson:"size"`
	CreatedOn     int64             `json:"created_on" bson:"created_on"`
	Name          string            `json:"name" bson:"name"`
	FileType      string            `json:"file_type" bson:"file_type"`
	MimeType      string            `json:"mime_type" bson:"mime_type"`
	Bucket        string            `json:"bucket" bson:"bucket"`
	Path          string            `json:"path" bson:"path"`
	Collection    string            `json:"collection" bson:"collection"`
	ThumbnailPath string            `json:"thumbnail_path" bson:"thumbnail_path"`
	Status        string            `json:"status" bson:"status"`
	Thumbnails    map[string]string `json:"-" bson:"thumbnails"`
}

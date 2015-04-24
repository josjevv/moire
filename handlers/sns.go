package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/bulletind/moire/db"
	"gopkg.in/simversity/gottp.v2"
	"gopkg.in/simversity/gottp.v2/utils"
)

type SNS struct {
	gottp.BaseHandler
}

func (self *SNS) Post(request *gottp.Request) {
	var errs *[]error

	n := snsNotice{}
	request.ConvertArguments(&n)

	errs = utils.Validate(&n)
	if len(*errs) > 0 {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			ConcatenateErrors(errs),
		})

		return
	}

	msg := snsMessage{}
	utils.Decoder([]byte(n.Message), &msg)

	errs = utils.Validate(&msg)
	if len(*errs) > 0 {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			ConcatenateErrors(errs),
		})

		return
	}

	record := msg.Records[0]

	key, err := url.QueryUnescape(record.S3.Object.Key)
	if err != nil {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			"Cannot escape Path",
		})

		return
	}

	if !strings.HasPrefix(key, UploadPrefix) {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			"Skipping path as its not meant to be monitored.",
		})

		return
	}

	doc := db.M{
		"size":   record.S3.Object.Size,
		"status": db.READY,
	}

	conn := getConn()
	asset := assetReady(conn, key, record.S3.Bucket.Name, db.M{"$set": doc})

	if strings.HasPrefix(asset.MimeType, VideoFile) {
		thumbPath := videoThumbnail(asset.Bucket, asset.Path)
		uploadFile(thumbPath)
	} else if strings.HasPrefix(asset.MimeType, ImageFile) {
		thumbPath := imageThumbnail(asset.Bucket, asset.Path)
		uploadFile(thumbPath)
	}

	request.Write("asset " + asset.Id.Hex() + " marked as ready")
}

type snsMessage struct {
	Records []struct {
		S3 struct {
			Bucket struct {
				Name string `json:"name" required:"true"`
			} `json:"bucket" required:"true"`
			Object struct {
				Key  string `json:"key" required:"true"`
				Size int    `json:"size" required:"true"`
			} `json:"object" required:"true"`
		}
	} `json:"Records" required:"true"`
}

type snsNotice struct {
	Type      string `json:"Type"`
	MessageId string `json:"MessageId" required:"true"`
	TopicArn  string `json:"TopicArn"`
	Subject   string `json:"Subject"`
	Message   string `json:"Message" required:"true"`
	Timestamp string `json:"Timestamp"`
	Signature string `json:"Signature" required:"true"`
}

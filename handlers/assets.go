package handlers

import (
	"net/http"
	"strings"

	"github.com/bulletind/moire/db"

	"gopkg.in/simversity/gottp.v3"
	"gopkg.in/simversity/gottp.v3/utils"
)

type Assets struct {
	gottp.BaseHandler
}

type assetArgs struct {
	fileType   string
	Name       string `json:"name" required:"true"`
	MimeType   string `json:"mime_type" required:"true"`
	Collection string `json:"collection"`
}

func (self *Assets) Post(request *gottp.Request) {
	args := assetArgs{}

	request.ConvertArguments(&args)

	err := utils.Validate(&args)
	if len(*err) > 0 {
		request.Raise(gottp.HttpError{
			http.StatusBadRequest,
			ConcatenateErrors(err),
		})

		return
	}

	fileType := PlainFile

	if strings.HasPrefix(args.MimeType, ImageFile) {
		fileType = ImageFile
	} else if strings.HasPrefix(args.MimeType, VideoFile) {
		fileType = VideoFile
	} else if strings.HasPrefix(args.MimeType, AudioFile) {
		fileType = AudioFile
	}

	args.fileType = fileType

	conn := getConn()
	asset := createAsset(conn, &args)

	_id := asset.Id.Hex()

	request.Write(db.M{
		"_id":        _id,
		"url":        "/assets/" + _id,
		"upload_url": getSignedUploadURL(asset.Bucket, asset.Path, asset.MimeType),
	})

	return
}

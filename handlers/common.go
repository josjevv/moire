package handlers

import (
	"github.com/bulletind/moire/config"
	"github.com/bulletind/moire/db"
)

const (
	AudioFile = "audio"
	ImageFile = "image"
	VideoFile = "video"
	PlainFile = "plain"
)

func getConn() *db.MConn {
	return db.GetConn(
		config.Settings.Moire.DBName,
		config.Settings.Moire.DBAddress,
		config.Settings.Moire.DBUsername,
		config.Settings.Moire.DBPassword,
	)
}
package config

import (
	"github.com/ortizdavid/go-nopain/conversion"
)

func ListenAddr() string {
	return GetEnv("APP_HOST")+":"+GetEnv("APP_PORT")
}

func ItemsPerPage() int {
	return conversion.StringToInt(GetEnv("APP_ITEMS_PER_PAGE"))
}

func UploadImagePath() string {
	return GetEnv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	return GetEnv("UPLOAD_DOCUMENT_PATH")
}

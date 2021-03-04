package helpers

import "github.com/lithammer/shortuuid/v3"

func AddFilename(mimeType string) (filename string) {
	u := shortuuid.New()

	switch mimeType {
	case "image/png":
		return u + ".png"
	case "image/jpeg":
		return u + ".jpg"
	case "image/jpg":
		return u + ".jpg"
	default:
		return ""
	}
}

package autoload

import (
	"mime"

	mimedb "github.com/sniperkit/mimedb/pkg"
)

func init() {
	for ext, mimeType := range mimedb.DB {
		if err := mime.AddExtensionType("."+ext, mimeType.ContentType); err != nil {
			panic(err)
		}
	}
}

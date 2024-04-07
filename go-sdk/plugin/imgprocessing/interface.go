package imgprocessing

import (
	"github.com/thailemeetai/go-sdk/go-sdk/sdkcm"
	"mime/multipart"
)

type ImgProcessing interface {
	// call img processing service to resize img and upload to s3
	Resize(file *multipart.FileHeader, folder string, longEdge int, quality int) (*sdkcm.Image, error)

	ResizeFile(filePath string, folder string, longEdge int, quality int) (*sdkcm.Image, error)
}

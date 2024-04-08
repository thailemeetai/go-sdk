package imgprocessing

import "github.com/thailemeetai/go-sdk/sdkcm"

type Response struct {
	sdkcm.AppError
	Data *sdkcm.Image `json:"data"`
}

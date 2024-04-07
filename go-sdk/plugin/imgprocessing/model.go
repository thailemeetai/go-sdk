package imgprocessing

import "github.com/thailemeetai/go-sdk/go-sdk/sdkcm"

type Response struct {
	sdkcm.AppError
	Data *sdkcm.Image `json:"data"`
}

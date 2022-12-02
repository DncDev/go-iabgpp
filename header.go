package gpp

import (
	"github.com/DncDev/iabgpp/base64"
)

func DecodeHeader(spec string) Header {

	out := base64.Decode(spec)

	h := Header{}

	h.Decode(out)

	return h

}

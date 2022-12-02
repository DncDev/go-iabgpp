package gpp

import (
	"errors"
	"github.com/DncDev/go-iabgpp/privacy"
	iabtcf "github.com/SirDataFR/iabtcfv2"
	"strings"
)

/*Decodes a GPP Privacy String
DBABMA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA
Add other privacy sub string handling here ie uspv1 etc
*/
func Decode(spec string) (*Block, error) {

	parts := strings.Split(spec, "~")

	blk := &Block{}

	if len(parts) < 2 {

		return blk, errors.New("Bad Data")

	}

	header := DecodeHeader(parts[0])

	blk.Header = header

	sections := parts[1:] //the rest of the items after the header

	if blk.HasSection("tcfeuv2") {

		res, err := iabtcf.Decode(blk.GetSection("tcfeuv2", sections))

		if err == nil {

			blk.TcfeuV2 = res
		}

	}

	if blk.HasSection("tcfca") {

		res, err := iabtcf.Decode(blk.GetSection("tcfca", sections))

		if err == nil {

			blk.TcfCA = res
		}

	}

	if blk.HasSection("uspv1") {

		uspv1 := privacy.UspV1{}

		str := blk.GetSection("uspv1", sections)

		res, err := uspv1.Decode(str)

		if err == nil {
			blk.UspV1 = &res
		}

	}

	if blk.HasSection("usnat") {

		usnat := privacy.UsNat{}
		res, err := usnat.Decode(blk.GetSection("usnat", sections))

		if err == nil {
			blk.UsNat = &res
		}

	}

	return blk, nil

}

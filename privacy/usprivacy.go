package privacy

import (
	"errors"
	"strconv"
)

//https://github.com/InteractiveAdvertisingBureau/USPrivacy/blob/master/CCPA/US%20Privacy%20String.md

const (
	NA = iota
	Yes
	No
)

type UspV1 struct {
	Version    int
	Notice     int
	OptOutSale int
	LSPA       int
}

func (u UspV1) Decode(str string) (UspV1, error) {

	if len(str) < 4 {

		return u, errors.New("Invalid Length")
	}

	for i, c := range str {

		cc := string(c)

		if i == 0 {
			u.Version, _ = strconv.Atoi(cc)
			continue
		}

		switch i {
		case 1:
			u.Notice = valToInt(cc)
			break
		case 2:
			u.OptOutSale = valToInt(cc)
			break
		case 3:
			u.LSPA = valToInt(cc)
			break

		}

	}

	return u, nil

}

func valToInt(val string) int {

	switch val {
	case "N":
		return No
		break
	case "Y":
		return Yes
		break
	case "-":
		return NA
		break
	}

	return NA
}

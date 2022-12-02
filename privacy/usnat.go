package privacy

import (
	"fmt"
	"github.com/DncDev/go-iabgpp/base64"
	"strconv"
)

//https://github.com/InteractiveAdvertisingBureau/USPrivacy/blob/master/CCPA/US%20Privacy%20String.md

type UsNat struct {
	Version                             int64
	SharingNotice                       int64
	SaleOptOutNotice                    int64
	SharingOptOutNotice                 int64
	TargetedAdvertisingOptOutNotice     int64
	SensitiveDataProcessingOptOutNotice int64
	SensitiveDataLimitUseNotice         int64
	SaleOptOut                          int64
	SharingOptOut                       int64
	TargetedAdvertisingOptOut           int64
	SensitiveDataProcessing             int //bitfield
	KnownChildSensitiveDataConsents     int //bitfield
	PersonalDataConsents                int64
	MspaCoveredTransaction              int64
	MspaOptOutOptionMode                int64
	MspaServiceProviderMode             int64
}

/*
Array of Number	Consists of two datapoints: a fixed length Integer(16) that denotes the length and a bitfield with that specific length.

Please note: Although the API reads/writes to fields (length + bitfield), it will only output the IDs from the bitfield via JS APIs.
*/

func (u UsNat) Decode(str string) (UsNat, error) {

	out := base64.Decode(str)

	fmt.Println("NAT", out)

	if val, err := strconv.ParseInt(out[0:6], 2, 6); err == nil {
		u.Version = val
	}

	if val, err := strconv.ParseInt(out[6:8], 2, 2); err == nil {
		u.SharingNotice = val
	}

	if val, err := strconv.ParseInt(out[8:10], 2, 2); err == nil {
		u.SaleOptOutNotice = val
	}

	if val, err := strconv.ParseInt(out[10:12], 2, 2); err == nil {
		u.SharingOptOutNotice = val
	}

	if val, err := strconv.ParseInt(out[12:14], 2, 2); err == nil {
		u.TargetedAdvertisingOptOutNotice = val
	}

	if val, err := strconv.ParseInt(out[14:16], 2, 2); err == nil {
		u.SensitiveDataProcessingOptOutNotice = val
	}

	if val, err := strconv.ParseInt(out[16:18], 2, 2); err == nil {
		u.SensitiveDataLimitUseNotice = val
	}

	if val, err := strconv.ParseInt(out[18:20], 2, 2); err == nil {
		u.SaleOptOut = val
	}

	if val, err := strconv.ParseInt(out[20:22], 2, 2); err == nil {
		u.SharingOptOut = val
	}

	if val, err := strconv.ParseInt(out[22:24], 2, 2); err == nil {
		u.TargetedAdvertisingOptOut = val
	}

	fmt.Println(out[24:])

	return u, nil

}

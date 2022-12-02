package gpp

//https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Sections/Section%20Information.md
//Section id's

const (
	TcfeuV1 = iota + 1
	TcfeuV2
	Gppheader
	Gppintegrity
	TcfCA
	UspV1
	UsNat
	UsCA
	UsVA
	UsCO
	UsUT
	UsCT
)

var sections = map[string]int{
	"tcfeuv1": TcfeuV1,
	"tcfeuv2": TcfeuV2,
	"tcfca":   TcfCA,
	"uspv1":   UspV1,
	"usnat":   UsNat,
}

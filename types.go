package gpp

import (
	"fmt"
	"github.com/DncDev/go-iabgpp/base64"
	"github.com/DncDev/go-iabgpp/field"
	"github.com/DncDev/go-iabgpp/privacy"
	iabtcf "github.com/SirDataFR/iabtcfv2"

	"strconv"
	"strings"
)

type Block struct {
	Header  Header
	TcfeuV1 *Section
	TcfeuV2 *iabtcf.TCData
	TcfCA   *iabtcf.TCData
	UspV1   *privacy.UspV1
	UsNat   *privacy.UsNat
	UsCA    *Section
	UsVA    *Section
	UsCO    *Section
	UsUT    *Section
	UsCT    *Section
}

func (b *Block) Init() {

	b.TcfeuV1 = nil
	b.TcfeuV2 = nil
	b.TcfCA = nil
	b.UspV1 = nil
	b.UsNat = nil
	b.UsCA = nil
	b.UsVA = nil
	b.UsCO = nil
	b.UsUT = nil
	b.UsCT = nil

}

//blank section until support is added
type Section struct {
}

type Header struct {
	Type     int64
	Version  int64
	Sections field.FibonacciRange
}

func (h *Header) Encode() {

	id := fmt.Sprintf("%06b", h.Type)
	version := fmt.Sprintf("%06b", h.Version)

	fmt.Println(id, version)

}

//check whether a block has a section

func (b *Block) GetSection(item string, check []string) string {

	sec := int64(0)

	if val, ok := sections[item]; ok {

		for ct, i := range b.Header.Sections.Items {

			sec = i.Start + sec
			if sec == int64(val) {
				if len(check) >= ct {
					return check[ct]
				}
			}

		}

	}

	return ""

}

func (b *Block) HasSection(item string) bool {

	sec := int64(0)

	if val, ok := sections[item]; ok {

		for _, i := range b.Header.Sections.Items {

			sec = i.Start + sec

			if sec == int64(val) {

				return true
			}

		}

	}

	return false

}

func (h *Header) Decode(header string) {

	blocks := base64.ChunkString(header, 6)

	if len(blocks) > 2 {

		//header
		if id, err := strconv.ParseInt(blocks[0], 2, 64); err == nil {
			h.Type = id
		}

		if vers, err := strconv.ParseInt(blocks[1], 2, 64); err == nil {
			h.Version = vers
		}

		sections := strings.Join(blocks[2:], "")

		sections = strings.TrimRight(sections, "0")

		ranges, err := field.DecodeFibonacciRange(sections)

		if err != nil {
		}

		h.Sections = ranges //this decodes to the numeric type ie 2,6

	}

}

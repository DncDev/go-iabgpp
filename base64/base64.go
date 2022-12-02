package base64

import (
	"strings"
)

var chars map[string]string
var bins map[string]string

func init() {

	chars = make(map[string]string)

	chars["A"] = "000000"
	chars["Q"] = "010000"
	chars["g"] = "100000"
	chars["w"] = "110000"
	chars["B"] = "000001"
	chars["R"] = "010001"
	chars["h"] = "100001"
	chars["x"] = "110001"
	chars["C"] = "000010"
	chars["S"] = "010010"
	chars["i"] = "100010"
	chars["y"] = "110010"
	chars["D"] = "000011"
	chars["T"] = "010011"
	chars["j"] = "100011"
	chars["z"] = "110011"
	chars["E"] = "000100"
	chars["U"] = "010100"
	chars["k"] = "100100"
	chars["0"] = "110100"
	chars["F"] = "000101"
	chars["V"] = "010101"
	chars["l"] = "100101"
	chars["1"] = "110101"
	chars["G"] = "000110"
	chars["W"] = "010110"
	chars["m"] = "100110"
	chars["2"] = "110110"
	chars["H"] = "000111"
	chars["X"] = "010111"
	chars["n"] = "100111"
	chars["3"] = "110111"
	chars["I"] = "001000"
	chars["Y"] = "011000"
	chars["o"] = "101000"
	chars["4"] = "111000"
	chars["J"] = "001001"
	chars["Z"] = "011001"
	chars["p"] = "101001"
	chars["5"] = "111001"
	chars["K"] = "001010"
	chars["a"] = "011010"
	chars["q"] = "101010"
	chars["6"] = "111010"
	chars["L"] = "001011"
	chars["b"] = "011011"
	chars["r"] = "101011"
	chars["7"] = "111011"
	chars["M"] = "001100"
	chars["c"] = "011100"
	chars["s"] = "101100"
	chars["8"] = "111100"
	chars["N"] = "001101"
	chars["d"] = "011101"
	chars["t"] = "101101"
	chars["9"] = "111101"
	chars["O"] = "001110"
	chars["e"] = "011110"
	chars["u"] = "101110"
	chars["."] = "111110"
	chars["P"] = "001111"
	chars["f"] = "011111"
	chars["v"] = "101111"
	chars["/"] = "111111"

	bins = make(map[string]string)

	bins["000000"] = "A"
	bins["010000"] = "Q"
	bins["100000"] = "g"
	bins["110000"] = "w"
	bins["000001"] = "B"
	bins["010001"] = "R"
	bins["100001"] = "h"
	bins["110001"] = "x"
	bins["000010"] = "C"
	bins["010010"] = "S"
	bins["100010"] = "i"
	bins["110010"] = "y"
	bins["000011"] = "D"
	bins["010011"] = "T"
	bins["100011"] = "j"
	bins["110011"] = "z"
	bins["000100"] = "E"
	bins["010100"] = "U"
	bins["100100"] = "k"
	bins["110100"] = "0"
	bins["000101"] = "F"
	bins["010101"] = "V"
	bins["100101"] = "l"
	bins["110101"] = "1"
	bins["000110"] = "G"
	bins["010110"] = "W"
	bins["100110"] = "m"
	bins["110110"] = "2"
	bins["000111"] = "H"
	bins["010111"] = "X"
	bins["100111"] = "n"
	bins["110111"] = "3"
	bins["001000"] = "I"
	bins["011000"] = "Y"
	bins["101000"] = "o"
	bins["111000"] = "4"
	bins["001001"] = "J"
	bins["011001"] = "Z"
	bins["101001"] = "p"
	bins["111001"] = "5"
	bins["001010"] = "K"
	bins["011010"] = "a"
	bins["101010"] = "q"
	bins["111010"] = "6"
	bins["001011"] = "L"
	bins["011011"] = "b"
	bins["101011"] = "r"
	bins["111011"] = "7"
	bins["001100"] = "M"
	bins["011100"] = "c"
	bins["101100"] = "s"
	bins["111100"] = "8"
	bins["001101"] = "N"
	bins["011101"] = "d"
	bins["101101"] = "t"
	bins["111101"] = "9"
	bins["001110"] = "O"
	bins["011110"] = "e"
	bins["101110"] = "u"
	bins["111110"] = "+"
	bins["001111"] = "P"
	bins["011111"] = "f"
	bins["101111"] = "v"
	bins["111111"] = "/"

}

func Decode(str string) string {
	items := strings.Split(str, "")
	var out []string

	for _, i := range items {
		if val, ok := chars[i]; ok {

			out = append(out, val)
		}
	}

	return strings.Join(out, "")

}

func Encode(items string) string {

	return EncodeSlice([]string{items})

}

func EncodeSlice(items []string) string {

	all := strings.Join(items, "")

	parts := ChunkString(all, 6)

	var out []string

	for _, i := range parts {
		if val, ok := bins[i]; ok {

			out = append(out, val)
		}
	}

	return strings.Join(out, "")

}

func ChunkString(s string, n int) []string {
	var ss []string
	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}
	ss = append(ss, rightPad2Len(s, "0", 6))
	return ss
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

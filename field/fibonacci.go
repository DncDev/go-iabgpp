package field

import (
	"errors"
	"strconv"
	"strings"
)

//intentionally removed the 0 leading chars
var fibonacci = []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811}

type FibonacciRange struct {
	Amount int64
	Items  []FibonacciRangeItem
}

type FibonacciRangeItem struct {
	Type  bool
	Start int64
}

func DecodeFibonacci(wrd string) int64 {

	seq := strings.TrimSuffix(wrd, "1")

	fib := fibonacci[1:]

	list := strings.Split(seq, "")

	num := int64(0)

	for ind, val := range list {

		if val != "0" {

			num = num + fib[ind]

		}

	}

	return num

}

//https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Core/Consent%20String%20Specification.md#discrete-sections-

func DecodeFibonacciRange(str string) (FibonacciRange, error) {

	rng := FibonacciRange{}

	if len(str) < 15 {
		//min length that's possible - TODO confirm this length

		return rng, errors.New("Too Short")
	}

	amount := str[:12] //representing the amount of items to follow

	var amt int64
	var err error

	if amt, err = strconv.ParseInt(amount, 2, 12); err == nil {

	}

	rng.Amount = amt

	items := str[12:]
	typ := items[0:1]
	news := items[1:] //pull of the first one we know it's a bool

	for i := 1; i < len(news)+1; i++ {

		seg := news[0:i]

		//if true we found a fibonacci number
		if len(seg) > 1 && seg[len(seg)-2:] == "11" {

			//next char is potentially a bool for the next item check if it's there for the next section if its singele type 0
			if len(news) >= len(seg)+1 && typ == "0" {
				//set the new type and then set news to the remainder
				fibnum := seg[0:i]
				typb, _ := strconv.ParseBool(typ)

				it := FibonacciRangeItem{}
				it.Type = typb
				it.Start = DecodeFibonacci(fibnum)
				rng.Items = append(rng.Items, it)

				typ = news[i : i+1]

				news = news[i+1:]
				i = 1 //star over

			} else if typ == "1" {
				fibnum := seg[0:i]
				news = news[i:]
				typb, _ := strconv.ParseBool(typ)
				it := FibonacciRangeItem{typb, DecodeFibonacci(fibnum)}
				rng.Items = append(rng.Items, it)
				i = 1
			} else if typ == "0" { //0 and doesn't continue

				fibnum := seg[0:i]
				typb, _ := strconv.ParseBool(typ)
				it := FibonacciRangeItem{}
				it.Type = typb
				it.Start = DecodeFibonacci(fibnum)
				rng.Items = append(rng.Items, it)
			}

		}

	}

	return rng, nil

}

func EncodeFibonacci(num int64) string {

	var out []int
	var code []string

	fib, ind := maxFib(num)

	out = append(out, ind)

	rem := num - fib

	for rem > 0 {

		fib, ind = maxFib(rem)

		rem = rem - fib

		out = append(out, ind)

	}

	code = make([]string, out[0])

	for i := 0; i < out[0]; i++ {

		code[i] = "0"

	}

	for _, vv := range out {
		code[vv-1] = "1"
	}

	code = append(code, "1")

	return strings.Join(code, "")

}

func maxFib(num int64) (int64, int) {

	maxfib := 0

	for i, f := range fibonacci {

		if num < f {

			maxfib = i - 1 //one step back in the sequence
			return fibonacci[maxfib], maxfib

		}

	}

	return 0, 0

}

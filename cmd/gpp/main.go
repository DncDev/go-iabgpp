package main

import (
	"flag"
	gpp "github.com/DncDev/go-iabgpp"
	"github.com/davecgh/go-spew/spew"
)

var input string

func main() {

	flag.StringVar(&input, "gpp", "", "Gpp String Ex: BDACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN")
	flag.Parse()

	var out *gpp.Block

	if len(input) < 1 {

		out, _ = gpp.Decode("BDACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN")

	} else {

		out, _ = gpp.Decode(input)
	}

	spew.Dump(out)

	//gpp.Decode("DBACMMA~CPXxRfAPXxRfAAfKABENB~BbbbGxsbG2w.MA")
	//gpp.Decode("BDACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN")

}

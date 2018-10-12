package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/b4fun/adblockdomain"
)

var (
	flagWithException *bool
	flagDecodeWithB64 *bool
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		abort()
	}

	inputUrlOrFile := flag.Arg(0)
	input, err := tryUrl(inputUrlOrFile)
	if err != nil {
		input, err = tryFile(inputUrlOrFile)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "io: %v\n", err)
		abort()
	}
	defer input.Close()

	var reader io.Reader = input

	if *flagDecodeWithB64 {
		reader = base64.NewDecoder(base64.StdEncoding, reader)
	}

	var domains []string
	if *flagWithException {
		domains, err = adblockdomain.ParseExceptionFromReader(reader)
	} else {
		domains, err = adblockdomain.ParseFromReader(reader)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		abort()
	}

	for _, domain := range domains {
		fmt.Println(domain)
	}
}

func init() {
	flagWithException = flag.Bool("e", false, "show exception domains")
	flagDecodeWithB64 = flag.Bool("b64", false, "decode content as base64 first")
}

func abort() {
	os.Exit(-1)
}

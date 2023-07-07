package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	gham "gh-action-multiline"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	const usage = `Usage: gh-action-multiline <name>

$ your_command | gh-action-multiline OUTPUT_OR_ENV_NAME >> "$GITHUB_OUTPUT"
$ gh-action-multiline -version`

	versionFlag := flag.Bool("version", false, "print the version of this program")
	byteSizeFlag := flag.Int("bytesize", gham.ByteSizeFromGitHubDoc, "specify byte size for delimiter")

	flag.Usage = func() {
		// https://github.com/golang/go/issues/57059#issuecomment-1336036866
		fmt.Printf("%s", usage+"\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *versionFlag {
		revision := commit[:7]
		fmt.Printf("%s\n", "gh-action-multiline"+" "+version+" "+"("+revision+") # "+date)
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	name := os.Args[1]

	nr := gham.DefaultNormalizer()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf(err.Error())
	}

	normalized, err := nr.Normalize(name, string(input), *byteSizeFlag)

	if err != nil {
		log.Fatalf(err.Error())
	}

	println(normalized)
}

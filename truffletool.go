package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

var excludedPaths = flag.StringSlice("exclude", []string{}, "list of path names to exclude")

type Truffle struct {
	Branch       string
	Commit       string
	CommitHash   string
	Date         string
	Path         string
	PrintDiff    string
	Diff         string
	Reason       string
	StringsFound []string
}

func (t *Truffle) Print() {
	fmt.Printf("\n")
	fmt.Printf("**************************************\n")
	fmt.Printf("*** %s\n", t.Reason)
	fmt.Printf("**************************************\n")
	fmt.Printf("[*] commit......: %s\n", t.CommitHash)
	fmt.Printf("[*] branch......: %s\n", t.Branch)
	fmt.Printf("[*] file........: %s\n", t.Path)
	fmt.Printf("[*] violations..: %d\n", len(t.StringsFound))
	for _, s := range t.StringsFound {
		if len(s) > 80 {
			s = s[:80] + fmt.Sprintf("...(%d characters truncated)", len(s)-80)
		}
		fmt.Println(s)
	}
}

func excluded(t Truffle) bool {
	for _, s := range *excludedPaths {
		if t.Path == s {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	// decode json from standard input
	reader := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(reader)

	for decoder.More() {
		var t Truffle
		// decode an array value (Message)
		err := decoder.Decode(&t)
		if err != nil {
			log.Fatal("fatal error: %s\n", err)
		}

		if excluded(t) {
			continue
		}

		t.Print()
	}
}

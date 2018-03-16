package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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

func main() {
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

		t.Print()
	}
}

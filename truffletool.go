package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

type Violation struct {
	File     string
	Branches map[string]int
	Commits  map[string]int
	Strings  map[string]int
	Reasons  map[string]int
}

func NewViolation(file string) Violation {
	return Violation{
		File:     file,
		Branches: map[string]int{},
		Commits:  map[string]int{},
		Reasons:  map[string]int{},
		Strings:  map[string]int{},
	}
}

func (v Violation) AddTruffle(t Truffle) Violation {
	v.Commits[t.CommitHash] += 1
	v.Branches[t.Branch] += 1
	v.Reasons[t.Reason] += 1
	for _, s := range t.StringsFound {
		v.Strings[s] += 1
	}
	return v
}

func commitDigest(hash map[string]int) string {
	for k, _ := range hash {
		if len(hash) == 1 {
			return k
		}
		return fmt.Sprintf("%s (and %d more)", k, len(hash)-1)
	}
	return ""
}

func digest(hash map[string]int) string {
	s := make([]string, 0, len(hash))
	for k, v := range hash {
		s = append(s, fmt.Sprintf("%s (%d)", k, v))
	}
	return strings.Join(s, ", ")
}

func (v *Violation) Print() {
	fmt.Printf("[*] file.....: %s\n", v.File)
	fmt.Printf("[*] commits..: %s\n", commitDigest(v.Commits))
	fmt.Printf("[*] branches.: %s\n", digest(v.Branches))
	fmt.Printf("[*] reasons..: %s\n", digest(v.Reasons))
	for s, _ := range v.Strings {
		if len(s) > 80 {
			s = s[:80] + fmt.Sprintf("...(%d characters truncated)", len(s)-80)
		}
		fmt.Println(s)
	}
	fmt.Printf("---\n")
}

func excluded(t Truffle) bool {
	_, filename := filepath.Split(t.Path)
	for _, s := range *excludedPaths {
		if t.Path == s || filename == s {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	// store violations
	violations := map[string]Violation{}

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

		v, ok := violations[t.Path]
		if !ok {
			v = NewViolation(t.Path)
		}
		violations[t.Path] = v.AddTruffle(t)
	}

	for _, v := range violations {
		v.Print()
	}
}

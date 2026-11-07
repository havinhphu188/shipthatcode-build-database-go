package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The token kinds, spelled exactly as they are printed.
const (
	KEYWORD    = "KEYWORD"
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	STRING     = "STRING"
	OPERATOR   = "OPERATOR"
	SYMBOL     = "SYMBOL"
)

type Tok struct {
	Kind string
	Text string
}

var KEYWORDS = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "INSERT": true, "INTO": true, "VALUES": true,
	"CREATE": true, "TABLE": true, "UPDATE": true, "SET": true, "DELETE": true, "AND": true,
	"OR": true, "NOT": true, "IS": true, "NULL": true, "INTEGER": true, "TEXT": true, "BEGIN": true,
	"COMMIT": true, "ROLLBACK": true, "ORDER": true, "BY": true, "LIMIT": true, "OFFSET": true,
	"COUNT": true, "SUM": true, "AVG": true, "MIN": true, "MAX": true, "ASC": true, "DESC": true,
	"INDEX": true, "ON": true,
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Split one line of SQL into tokens. err is "" on success, or the message that follows "ERR "
// when the line cannot be tokenized.
func tokenize(line string) (toks []Tok, err string) {
	n := len(line)
	i := 0
	for i < n {
		c := line[i]
		if c == ' ' || c == '\t' || c == '\r' {
			i++
		} else if strings.IndexByte("(),;", c) >= 0 {
			toks = append(toks, Tok{SYMBOL, string(c)})
			i++
		} else {
			// TODO: comments (-- to end of line), 'strings' and "quoted identifiers" with a doubled
			// quote as an escape, numbers (and the malformed ones), words (keyword or identifier),
			// and operators (two-character ones first). Anything else is an error.
			return nil, "unexpected character: " + string(c)
		}
	}
	return toks, ""
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		toks, err := tokenize(line)
		if err != "" {
			fmt.Fprintln(w, "ERR "+err)
			continue
		}
		for _, t := range toks {
			fmt.Fprintln(w, t.Kind+" "+t.Text)
		}
		fmt.Fprintln(w, "END")
	}
}

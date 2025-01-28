package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rgeoghegan/tabulate"
)

// authored by: Leo
// inspired by: https://gist.github.com/djoreilly/57ef65723bc8ad7ecdb899c5b8aaca47

func main() {
	in_chain, eof := false, false
	headers, table := []string{}, [][]string{}

	stdin := bufio.NewReader(os.Stdin)
	// iptables comment format: /* Comment */
	comment_re := regexp.MustCompile("/\\*.*/")

	for {
		line, _ := stdin.ReadString('\n')
		if len(line) == 0 {
			eof = true
		}
		line = strings.TrimSpace(line)

		if line == "" {
			if in_chain {
				headers = append(headers, "extra")
				text, _ := tabulate.Tabulate(
					table,
					&tabulate.Layout{Headers: headers, Format: tabulate.SimpleFormat},
				)
				fmt.Println(text)
				headers = []string{}
				table = [][]string{}
				in_chain = false
			}
			if eof {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "Chain") {
			fmt.Println(line)
			continue
		}
		if strings.HasPrefix(line, "pkts") || strings.HasPrefix(line, "target") {
			headers = strings.Fields(line)
			in_chain = true
			continue
		}
		if in_chain {
			parts := strings.Fields(line)
			begin := parts[:len(headers)]
			extra := strings.Join(parts[len(headers):], " ")
			extra = comment_re.ReplaceAllString(extra, "")  // strip comments
			table = append(table, append(begin, extra))
		}
	}
}
package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/solver"
	"github.com/jyggen/advent-of-go/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type regexPart struct {
	id int
	pattern string
	replace []int
	resolved bool
}

func main() {
	p1, p2, err := solver.SolveFromFile(os.Stdin, SolvePart1, SolvePart2)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func SolvePart1(input string) (string, error) {
	r, msgs := parse(input)
	sum := 0

	for _, m := range msgs {
		if r.MatchString(m) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

func SolvePart2(input string) (string, error) {
	input = strings.Replace(input, "8: 42", "8: 42 | 42 (42)+", 1)
	input = strings.Replace(input, "11: 42 31", "11: 42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31", 1)

	r, msgs := parse(input)
	sum := 0

	for _, m := range msgs {
		if r.MatchString(m) {
			sum++
		}
	}

	return strconv.Itoa(sum), nil
}

func build(parts map[int]*regexPart, index int, history []int) string {
	for _, h := range history {
		if h == index {
			fmt.Println("loop @", index, history)
			//return "+"
		}
	}

	if !parts[index].resolved {
		replace := make([]interface{}, len(parts[index].replace))

		for k, v := range parts[index].replace {
			replace[k] = build(parts, v, append(history, index))
		}

		parts[index].pattern = fmt.Sprintf(parts[index].pattern, replace...)
		parts[index].resolved = true
	}

	return parts[index].pattern
}

func parse(input string) (*regexp.Regexp, []string) {
	phase := 1
	msgs := make([]string, 0)
	parts := make(map[int]*regexPart, 0)

	for _, v := range utils.ToStringSlice(input, "\n") {
		if v == "" {
			phase++
		} else if phase == 1 {
			fields := strings.Fields(v)
			or := false
			part := &regexPart{
				pattern: "",
			}

			for _, f := range fields {
				if strings.HasSuffix(f, ":") {
					part.id, _ = strconv.Atoi(f[0:len(f)-1])
				} else if strings.HasPrefix(f, "\"") {
					part.pattern += f[1:len(f)-1]
				} else if f == "|" {
					part.pattern += "|"
					or = true
				} else {
					if strings.HasPrefix(f, "(") {
						part.pattern += "(?:"
						f = f[1:]
					}

					suffix := ""

					if strings.HasSuffix(f, ")+") {
						suffix = ")+"
						f = f[0:len(f)-2]
					}

					num, _ := strconv.Atoi(f)
					part.pattern += "%s" + suffix
					part.replace = append(part.replace, num)
				}
			}

			if or {
				part.pattern = "(?:" + part.pattern + ")"
			}

			parts[part.id] = part
		} else {
			msgs = append(msgs, v)
		}
	}

	b := "^" + build(parts, 0, make([]int, 0, len(parts))) + "$"

	return regexp.MustCompile(b), msgs
}

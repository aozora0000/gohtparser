package gohtparser

import (
	"fmt"
	"regexp"
	"strings"
)

var is_multiline_regex = regexp.MustCompile(`\\\\$`)
var is_comment_regex = regexp.MustCompile(`^#`)
var is_directive_regex = regexp.MustCompile(`^[^#\<]`)
var is_block_regex = regexp.MustCompile(`^<[^/].*>$`)
var block_regex = regexp.MustCompile(`(?:[\s|<]")([^<>"]+)(?:"[\s|>])|([^<>\s]+)`)
var directive_regex = regexp.MustCompile(`"(?:\\.|[^\\"])*"|\S+`)

func isWhiteLine(val string) bool {
	return val == ""
}

func isMultiLine(val string) bool {
	return is_multiline_regex.MatchString(val)
}

func IsComment(val string) bool {
	return is_comment_regex.MatchString(val)
}

func isDirective(val string) bool {
	return is_directive_regex.MatchString(val)
}

func isBlock(val string) bool {
	return is_block_regex.MatchString(val)
}

func isBlockEnd(val string, name string) bool {
	if name == "" {
		name = `[^\s>]+`
	}
	fmt.Println("name: " + `^</` + name + `>$`)
	fmt.Println("val: " + val)
	fmt.Printf("match: %s\n", regexp.MustCompile(`^</`+name+`>$`).MatchString(val))
	fmt.Println()
	return regexp.MustCompile(`^</` + name + `>$`).MatchString(val)
}

func blockRegex(line string) []string {
	var block []string
	if block_regex.MatchString(line) {
		for _, f := range block_regex.FindAllString(line, -1) {
			f = strings.TrimSpace(f)
			f = strings.Trim(f, "<>")
			block = append(block, f)
		}
	}
	return block
}

func directiveRegex(line string) []string {
	var directives []string
	if directive_regex.MatchString(line) {
		for _, f := range block_regex.FindAllString(line, -1) {
			f = strings.TrimSpace(f)
			if f != "" {
				directives = append(directives, f)
			}

		}
	}
	return directives
}

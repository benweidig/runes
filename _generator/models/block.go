package models

import (
	"regexp"
	"strings"
)

var blockPattern = regexp.MustCompile("([0-9A-F]+)\\.\\.([0-9A-F]+); (.*)")

type Block struct {
	codePointRange CodePointRange
	name           string
}

func NewBlock(line string) Block {
	match := blockPattern.FindStringSubmatch(line)

	r := NewCodePointRange(match[1], match[2])
	return Block{
		codePointRange: r,
		name:           match[3],
	}
}

func (b Block) Range() CodePointRange {
	return b.codePointRange
}

func (b Block) Name() string {
	return b.name
}

var packageNameRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

func (b Block) PackageName() string {
	sanitized := packageNameRegex.ReplaceAllString(b.Name(), "")
	return strings.ToLower(sanitized)
}

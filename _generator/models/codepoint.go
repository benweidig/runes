package models

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CodePoint struct {
	point          int64
	pointStr       string
	paddedPointStr string
	categoryAbbr   string
	category       string
	name           string
	altName        string
}

func NewCodePoint(csv []string) CodePoint {
	pointStr := csv[0]
	point, _ := strconv.ParseInt(pointStr, 16, 64)

	var paddedPointStr string
	if len(pointStr) > 4 {
		paddedPointStr = fmt.Sprintf("%08s", strings.ToLower(pointStr))
	} else {
		paddedPointStr = fmt.Sprintf("%04s", strings.ToLower(pointStr))
	}
	return CodePoint{
		point:          point,
		pointStr:       pointStr,
		paddedPointStr: paddedPointStr,
		categoryAbbr:   csv[2],
		category:       Categories[csv[2]],
		name:           csv[1],
		altName:        csv[10],
	}
}

func (c CodePoint) Point() int64 {
	return c.point
}

func (c CodePoint) PointString() string {
	return c.pointStr
}

func (c CodePoint) Category() string {
	return c.category
}
func (c CodePoint) Name() string {
	if len(c.altName) > 0 {
		return fmt.Sprintf("%s / %s", c.name, c.altName)
	}
	if len(c.name) > 0 {
		return c.name
	}

	return fmt.Sprintf("U%s", c.pointStr)
}

func (c CodePoint) String() string {
	var s string
	var prefix string
	if len(c.pointStr) > 4 {
		prefix = `"\U`
	} else {
		prefix = `"\u`
	}
	s, _ = strconv.Unquote(prefix + c.paddedPointStr + `"`)

	if c.categoryAbbr == "Cc" {
		return ""
	}
	return s
}

var variableNameSanitizer = regexp.MustCompile("[^a-zA-Z0-9]+")

func (c CodePoint) VariableName() string {
	variableName := c.name
	if strings.HasPrefix(variableName, "<") || len(variableName) == 0 {
		variableName = c.altName
	}

	if len(variableName) == 0 {
		return fmt.Sprintf("U%s", c.pointStr)
	}

	variableName = variableNameSanitizer.ReplaceAllString(variableName, " ")
	variableName = strings.ToLower(variableName)
	variableName = strings.Title(variableName)
	variableName = strings.Replace(variableName, " ", "", -1)

	return variableName
}

func (c CodePoint) Escaped() string {
	var prefix string
	if len(c.pointStr) > 4 {
		prefix = "\\U"
	} else {
		prefix = "\\u"
	}

	return fmt.Sprintf("%s%s", prefix, c.paddedPointStr)
}

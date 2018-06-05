package models

import (
	"fmt"
	"strconv"
)

type CodePointRange struct {
	start    int64
	end      int64
	strValue string
}

func NewCodePointRange(start, end string) CodePointRange {
	startInt, _ := strconv.ParseInt(start, 16, 64)
	endInt, _ := strconv.ParseInt(end, 16, 64)
	return CodePointRange{
		start:    startInt,
		end:      endInt,
		strValue: fmt.Sprintf("%s..%s", start, end),
	}
}

func (r CodePointRange) Start() int64 {
	return r.start
}

func (r CodePointRange) End() int64 {
	return r.end
}

func (r CodePointRange) String() string {
	return r.strValue
}

func (r CodePointRange) IsIn(value int64) bool {
	return value >= r.start && value <= r.end
}

func (r CodePointRange) IsInStr(value string) bool {
	i, _ := strconv.ParseInt(value, 16, 64)
	return r.IsIn(i)
}

package models

type TemplateData struct {
	Block      Block
	CodePoints []CodePoint
}

func NewTemplateData(block Block, codepoints []CodePoint) TemplateData {
	return TemplateData{
		Block:      block,
		CodePoints: codepoints,
	}
}

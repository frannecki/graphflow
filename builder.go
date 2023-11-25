package graphflow

type metaData struct {
	inputFieldName string
	outputFieldName string
}

type Builder struct {
	meta *metaData
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (builder *Builder) Build() {}

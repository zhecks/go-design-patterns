package builder

type Builder interface {
	buildHeader()
	buildBody()
	buildFooter()
	getResult() string
}

type Director struct {
	build Builder
}

func (director *Director) construct() string {
	director.build.buildHeader()
	director.build.buildBody()
	director.build.buildFooter()
	return director.build.getResult()
}

type XmlBuilder struct {
	data string
}

func (builder *XmlBuilder) buildHeader() {
	builder.data += "XmlHeader"
}

func (builder *XmlBuilder) buildBody() {
	builder.data += "XmlBody"
}

func (builder *XmlBuilder) buildFooter() {
	builder.data += "XmlFooter"
}

func (builder *XmlBuilder) getResult() string {
	return builder.data
}

type TxtBuilder struct {
	data string
}

func (builder *TxtBuilder) buildHeader() {
	builder.data += "TxtHeader"
}

func (builder *TxtBuilder) buildBody() {
	builder.data += "TxtBody"
}

func (builder *TxtBuilder) buildFooter() {
	builder.data += "TxtFooter"
}

func (builder *TxtBuilder) getResult() string {
	return builder.data
}

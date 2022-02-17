package extractors

type Extractor interface {
	Extract(fromDir string)
}

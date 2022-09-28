package updatestrategy

// this update strategy will be used for unknown types, quality update logic & sellin update logic both are default
type defaultUpdateStrategy struct {
	defaultSellinUpdateStrategy
	defaultQualityUpdateStrategy
}

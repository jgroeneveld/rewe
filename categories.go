package rewe

type CategoryInfo struct {
	Product    string
	Categories []string
}

type CategoryFetcher interface {
	Fetch(productName string) (CategoryInfo, error)
}

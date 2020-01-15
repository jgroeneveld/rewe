package rewe

type CategoryInfo struct {
	Product    string   `json:"product"`
	Categories []string `json:"categories"`
}

type CategoryFetcher interface {
	Fetch(productName string) (CategoryInfo, error)
}

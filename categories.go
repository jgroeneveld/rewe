package rewe

type CategoryInfo struct {
	Product    string   `json:"product"`
	Categories []string `json:"categories"`
}

//go:generate mockgen -source=categories.go -package=rewe -destination mock_categories_test.go
type CategoryFetcher interface {
	Fetch(productName string) (CategoryInfo, error)
}

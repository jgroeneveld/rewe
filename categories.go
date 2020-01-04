package rewe

type Categories []string

type CategoriesFetcher interface {
	Fetch(productName string) (Categories, error)
}

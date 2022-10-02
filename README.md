# Rewe

This app analyzes product names from the Rewe online catalog and fetches the categories from their website.

This can be used to analyze in what categories you spend your money.

## Usage

```shell
bin/rewe [command] --help # prints help
bin/rewe bill ./Rechnung.pdf # identify products and their categories. See below for individual steps.
bin/rewe read-bill ./Rechnung.pdf # Parses the PDF and identifies all products and their costs
bin/rewe fetch-categories "Apfelsaft" # Looks for the Apfelsaft product on the rewe websites and extracts the categories printing them in json
```

## Development

### Prerequisites


- Install [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- Install [golangci-lint](https://github.com/golangci/golangci-lint) for linting. There is a `.golangci.yml` file for configuration.
- run `scripts/install-pre-commit` to install *build/all* as pre-commit hook.

### Building

Just use the go tools or the convenient build scripts located under `./build`

```bash
build/all # lint, test, build

build/lint # just lint
build/test # just run the tests

build/cli # build the cli to ./bin/rewe
```

### vscode Linting

```
"go.lintTool": "golangci-lint",
"go.lintFlags": [
    "--fast"
],
```

## Method

We can obtain the categories from the product pages. 
The categories are a tree. 
To obtain the product page we need to use the search with the name of the product.
The product names are unqiue enough for our use case.

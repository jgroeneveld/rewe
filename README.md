# Rewe

This app analyzes product names from the Rewe online catalog and fetches the categories from their website.

This can be used to analyze in what categories you spend your money.

## Development

build scripts are located under `./build`

```bash
build/all # lint, test, build

build/lint # just lint
build/test # just run the tests

build/cli # build the cli to bin/rewe
```

## Method

We can obtain the categories from the product pages. 
The categories are a tree. 
To obtain the product page we need to use the search with the name of the product.
The product names are unqiue enough for our use case.

## Use cases when we can use the json data in search only

- [X] ScenarioA: Enter the the name of a product and receive the categories
    - [X] Provide cli `rewe categories --product "REWE Bio Apfelsaft naturtrüb 1l"`
    - [X] Client that fetches the search page
    - [X] Parser for the search page
    - [X] CategoriesFetcher combining parser and client
    - [X] Log an error when the search contains more than one product
    - [X] Print result to stdout

- [X] ScenarioB: CLI supports json output
    - [X] Provide cli output `--json`
    - [X] Provide file writer that takes result  struct and writes to stdout
    - [X] Existing CLI uses it instead of stdout when `out` is specified
    
- [ ] ScenarioC: Enter list of product names and return categories
    - [ ] -> ScenarioA
    - [ ] Provide cli `rewe categories --products ??` (stdin? filein? pipe?)
    - [ ] CategoriesByNames that calls CategoriesByName for all and combines results
    - [ ] Print result to stdout or to file depending on cli parameters
    - [ ] Make CategoriesByNames async for performance

## JSON Format

```json
{
  "name": "REWE Bio Apfelsaft naturtrüb 1l",
  "url": "https://shop.cli.de/p/cli-bio-apfelsaft-naturtrueb-1l/254615",
  "categories": ["Getränke", "Soft Drinks", "Fruchtsäfte & Nektare", "Äpfel"]
}
```

# Rewe

This app analyzes product names from the Rewe online catalog and fetches the categories from their website.

This can be used to analyze in what categories you spend your money.

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

- [ ] ScenarioB: CLI supports json output
    - [ ] Provide cli output into file `--out file.json`
    - [ ] Provide file writer that takes result  struct and write into file as json
    - [ ] Existing CLI uses it instead of stdout when `out` is specified
    
- [ ] ScenarioC: Enter list of product names and return categories
    - [ ] -> ScenarioC
    - [ ] Provide cli `rewe categories --products ??` (stdin? filein? pipe?)
    - [ ] CategoriesByNames that calls CategoriesByName for all and combines results
    - [ ] Print result to stdout or to file depending on cli parameters
    - [ ] Make CategoriesByNames async for performance

## Use cases when we have to go page by page

- [ ] ScenarioA: Enter the link of a product and recieve the categories
    - [ ] Provide cli `rewe categories --product-url https://shop.rewe.de/p/rewe-bio-apfelsaft-naturtrueb-1l/254615`
    - [ ] Client that can fetch the html of a product
    - [ ] Product Page parser providing the categories
    - [ ] ProductPageService that implements the task

- [ ] ScenarioB: CLI supports json output
    - [ ] Provide cli output into file `--out file.json`
    - [ ] Provide file writer that takes result  struct and write into file as json
    - [ ] Existing CLI uses it

- [ ] ScenarioC: Enter the the name of a product and receive the categories
    - [ ] -> ScenarioA
    - [ ] Provide cli `rewe categories --product "REWE Bio Apfelsaft naturtrüb 1l"`
    - [ ] Client that fetches the search page
    - [ ] Parser for the search page
    - [ ] SearchPageService combining parser and client
    - [ ] CategoriesByName that uses the SearchPageService and ProductPageService
    
- [ ] ScenarioD: Enter list of product names and return categories
    - [ ] -> ScenarioC
    - [ ] Provide cli `rewe categories --products ??` (stdin? filein? pipe?)
    - [ ] CategoriesByNames that calls CategoriesByName for all
    - [ ] Make CategoriesByNames async for performance
    
## JSON Format

```json
{
  "name": "REWE Bio Apfelsaft naturtrüb 1l",
  "url": "https://shop.cli.de/p/cli-bio-apfelsaft-naturtrueb-1l/254615",
  "categories": ["Getränke", "Soft Drinks", "Fruchtsäfte & Nektare", "Äpfel"]
}
```

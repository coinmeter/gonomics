gonomics
========
![test](https://github.com/dariubs/gonomics/actions/workflows/test.yml/badge.svg)

Go library for interacting with the  [nomics API](https://nomics.com).

install
-------

```bash
go get github.com/dariubs/gonomics
```

usage
-----

```go
package main

import (
	"log"

	nomics "github.com/dariubs/gonomics"
)

func main() {
	nb := nomics.NewNomics("your-api-key")
	ctr := nomics.CurrenciesTickerRequest{PerPage: "100"}

	data, err := nb.CurrenciesTicker(&ctr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)

}

```

by
---

- [Dariush Abbasi](https://github.com/dariubs)

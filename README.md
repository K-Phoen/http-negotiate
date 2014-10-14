negotiate
=========

Content negotiation HTTP middleware for Go applications.

It's a simple wrapper around my [negotiation library](https://github.com/K-Phoen/negotiation)
which implements the `http.Handler` interface.

## Usage

Here is a ready to use example with [Negroni](https://github.com/codegangsta/negroni):

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/codegangsta/negroni"
  "github.com/K-Phoen/http-negotiate/negotiate"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
    fmt.Fprintf(w, "The negotiated format is: " + w.Header().Get("Content-Type"))
  })

  n := negroni.Classic()
  n.Use(negotiate.FormatNegotiator([]string{"application/json", "application/xml"}))
  n.UseHandler(mux)
  n.Run(":3000")
}
```

## ToDo

  * provide tools to negotiate other things (language for instance)
  * write tests

## License

This library is released under the MIT License. See the bundled LICENSE file for
details.

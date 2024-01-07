A very simple redirect package that works with gorilla/mux.

It pulls from a redirects.yml, supports different status codes and wildcards.

Add this to your main.go router:

// main.go
package main

import (
    "https://github.com/lojol-git/redirects/"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.Use(redirects.Run)
    // rest of your code
}

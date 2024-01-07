A very simple redirect package that works with gorilla/mux.

It pulls from a redirects.yml, supports different status codes and wildcards.

**Structure a "redirects.yml" file like this:**
redirects:

- from: "/old-path/(.\*)"  
  to: "/new-path?new=$1"  
  status: 301
- from: /another-old-path  
  to: /another-new-path  
  status: 301

**Add this to your main.go router:**

// main.go  
package main

import (  
"log"  
"https://github.com/lojol-git/redirects/"  
"github.com/gorilla/mux"  
)

func main() {  
 err := redirects.Load("redirects.yml")  
 if err != nil {  
 log.Fatalf("Failed to load redirects: %v", err)  
}

r := mux.NewRouter()
r.Use(redirects.Run)

// rest of your code

}

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
&nbsp;&nbsp;&nbsp;&nbsp;"log"  
&nbsp;&nbsp;&nbsp;&nbsp;"https://github.com/lojol-git/redirects/"  
&nbsp;&nbsp;&nbsp;&nbsp;"github.com/gorilla/mux"  
)

func main() {  
&nbsp;&nbsp;&nbsp;&nbsp;err := redirects.Load("redirects.yml")  
&nbsp;&nbsp;&nbsp;&nbsp;if err != nil {  
&nbsp;&nbsp;&nbsp;&nbsp;log.Fatalf("Failed to load redirects: %v", err)  
}

r := mux.NewRouter()
r.Use(redirects.Run)

// rest of your code

}

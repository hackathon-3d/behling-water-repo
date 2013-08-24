//App ID:	673350062694474
//App Secret:	736400b3b54625da28e9a250500e5ca8

package whosin

import (
	"fmt"
	"net/http"
)

type Inviter struct {
	Name string `json:"name"`
}

type Joiner struct {
	Name string `json:"name"`
}

type Outing struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Inviter  Inviter  `json:"inviter"`
	Joiners  []Joiner `json:"joiners"`
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login/facebook", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

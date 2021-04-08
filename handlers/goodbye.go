package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// func (gh *Goodbye) serveHTTP(rw http.ResponseWriter, r *http.Request) {
//YOU FORGET CAPITAL letter
func (gh *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "goodbye")
	fmt.Fprintf(rw, "<h1>bismillah</h1><h2>alhamdulillah</h2>\n")
	gh.l.Println("log bismillah...")
	dbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	fmt.Printf("data: %s", string(dbyte))

}

// func (gh *Goodbye) serveHTTP(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(rw, "goodbye")
// 	fmt.Fprintf(rw, "<h1>bismillah</h1><h2>alhamdulillah</h2>\n")
// 	gh.l.Println("log bismillah...")
// 	dbyte, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 	}
// 	fmt.Printf("data: %s", string(dbyte))

// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// // ノーマル
	// dump, err := httputil.DumpRequest(r, true)
	// if err != nil {
	// 	http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(string(dump))
	// fmt.Fprintf(w, "<html><body>hello</body></html>")

	// // クッキー使用
	// w.Header().Add("Set-Cookie", "VISIT=TRUE")
	// if _, ok := r.Header["Cookie"]; ok {
	// 	// クッキーがあるということは一度来たことがある人
	// 	fmt.Fprintf(w, "<html><body>２回目以降</body></html>")
	// } else {
	// 	fmt.Fprintf(w, "<html><body>１回目</body></html>")
	// }

	// Cache-Control実験
	w.Header().Add("Cache-Control", "max-age=180")
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening : 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

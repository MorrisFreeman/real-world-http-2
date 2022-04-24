package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// ノーマル
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		// クッキーがあるということは一度来たことがある人
		fmt.Fprintf(w, "<html><body>２回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>１回目</body></html>")
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
}

func chacheHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "max-age=180")
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Disposition", "attachment; filename=filename.txt")
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "hello")
}

func thanksHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, `<html>
	<meta http-equiv="refresh" content="0;URL=./download">
	<body>thanks</body>
	</html>`)
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", cookieHandler)
	http.HandleFunc("/cache", chacheHandler)
	http.HandleFunc("/slow", slowHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/thanks", thanksHandler)
	log.Println("start http listening : 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

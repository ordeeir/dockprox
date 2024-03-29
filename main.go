package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var lport string
var dip string
var sslcert string
var sslkey string

func main() {

	lport = "80"
	dip = "91.107.130.30"
	sslcert = os.Getenv("SSL_CERT")
	sslkey = os.Getenv("SSL_KEY")

	if lport == "" {
		lport = os.Args[1]
	}

	if dip == "" {
		dip = os.Args[2]
	}

	if sslcert == "" {
		//sslcert = os.Args[3]
	}

	if sslkey == "" {
		//sslkey = os.Args[4]
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Forward /")
	//	b, _ := os.ReadFile("index.html")
	//	fmt.Fprint(w, string(b))
	//})

	// Anything we don't do in Go, we pass to the old platform
	http.HandleFunc("/", wordpress) //  ex. /wordpress/xyxshxpxchxy/photo3235

	fmt.Println("Listening on localhost:" + lport)
	fmt.Println("Forward to:" + dip)

	// Start the server
	if lport == "443" {

		http.ListenAndServeTLS("0.0.0.0:"+lport, sslcert, sslkey, nil)

	} else {
		http.ListenAndServe("0.0.0.0:"+lport, nil)

	}

}

func wordpress(w http.ResponseWriter, r *http.Request) {

	// change the request host to match the target
	vars := strings.Split(r.URL.Path, "/")

	str := vars[2]

	str = strings.Replace(str, "xch", "4", -1)
	str = strings.Replace(str, "xsh", "6", -1)
	str = strings.Replace(str, "xha", "8", -1)
	str = strings.Replace(str, "xse", "0", -1)
	str = strings.Replace(str, "xy", "1", -1)
	str = strings.Replace(str, "xd", "2", -1)
	str = strings.Replace(str, "xs", "3", -1)
	str = strings.Replace(str, "xp", "5", -1)
	str = strings.Replace(str, "xh", "7", -1)
	str = strings.Replace(str, "xn", "9", -1)

	fmt.Println("Path: " + r.URL.Path + " Forward to: " + "http://" + dip + ":" + str)

	u, _ := url.Parse("http://" + dip + ":80")
	pro := httputil.NewSingleHostReverseProxy(u)

	//r.URL.Path = "/"

	pro.ServeHTTP(w, r)
}

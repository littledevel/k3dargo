package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

const RESPONSE = `
<html>
<title>Demo</title>
<body>
<h2> Demo API </h2>
</body>
</html>
`

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a request for /")
	fmt.Fprintf(w, "%s", RESPONSE)
}

func APICallHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting a request for %s\n", r.URL.Path)
	ip, err := getIP(r)
	if err != nil {
		fmt.Fprintf(w, "Not found")
	}
	fmt.Fprintf(w, "%s", ip)
}

func Serve(address string) {
	http.HandleFunc("/whoami", APICallHandler)
	http.HandleFunc("/", HTTPHandler)
	http.ListenAndServe((address), nil)
}

func getIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-REAL-IP")
	networkIP := net.ParseIP(ip)
	if networkIP != nil {
		return ip, nil
	}

	ipList := r.Header.Get("X-FORWARDED-FOR")
	ips := strings.Split(ipList, ",")
	for _, ip := range ips {
		networkIP := net.ParseIP(ip)
		if networkIP != nil {
			return ip, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	networkIP = net.ParseIP(ip)
	if networkIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No ip found")

}

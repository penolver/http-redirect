package main
import (
    "net/http"
    "flag"
    "fmt"
    "log"
)

func redirect(w http.ResponseWriter, req *http.Request) {

  var proto string

  if *httpsOff == true {
    proto = "http://"
  }else{
    proto = "https://"
  }
  //log.Printf("%s requested %s on %s from %s\n", req.RemoteAddr, req.Method, req.URL, req.Host)
  log.Println("Redirecting request from",req.RemoteAddr,"to",proto + req.Host + ":"+*redirectToPort + req.URL.String())

  http.Redirect(w, req,
            proto + req.Host + ":"+*redirectToPort + req.URL.String(),
            http.StatusMovedPermanently)

}

var redirectToPort *string
var httpsOff *bool

func main() {

    fmt.Println()
    fmt.Println("***************************************************")
    fmt.Println("                HTTP Redirector   ")
    fmt.Println("    Redirects HTTP to HTTP(S) or another port   ")
    fmt.Println("***************************************************")
    fmt.Println()
    fmt.Println("Defaulting to redirecting port 80 (HTTP) to 443 (HTTPS).")
    fmt.Println("to use alternatives, use -h flag for options")
    fmt.Println("use CTRL+C to kill")
    fmt.Println()

    listenPort := flag.String("l", "80", "Port to listen on")
    redirectToPort = flag.String("r", "443", "Port to redirect to")
    httpsOff = flag.Bool("n", false, "if redirect (destination) port is NOT HTTPS (i.e. do not prefix with https)")
    help := flag.Bool("h", false, "help")

    flag.Parse()

    if *help == true { flag.PrintDefaults() }

    log.Println("Starting server on port",*listenPort,"and redirecting to port",*redirectToPort)

    // redirect every http request to https
    http.ListenAndServe(":"+*listenPort, http.HandlerFunc(redirect))
}

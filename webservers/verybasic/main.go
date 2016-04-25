package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This webserver is running. You are currently in %q.", html.EscapeString(r.URL.Path))
	})
	print("\033[H\033[2J")
	fmt.Println("\nWebserver running on port 9000 \n")
	fmt.Println("The IP addresses are: ")
	fmt.Println("127.0.0.1")
	ipaddys()
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func ipaddys() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/micro/mdns"
)

func main() {

	serviceString := flag.String("service", "", "specify one service")
	servicePath := flag.String("servicepath", "", "specify path to service file")
	queryDelay := flag.Duration("delay", 100, "specify query delay (>= 100 ms)")
	async := flag.Bool("async", false, "Sync/Async. Async out of order results but faster. Set to true/false")
	flag.Parse()

	// check high level flags
	if *serviceString == "" && *servicePath == "" {
		fmt.Fprintf(os.Stderr, "%s\n", "Need a service or a file with service lines")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *serviceString != "" && *servicePath != "" {
		fmt.Fprintf(os.Stderr, "%s\n", "Need a service -OR- a file with service lines")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *servicePath != "" {
		if _, err := os.Stat(*servicePath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s\n", "File not found. Check file path")
			flag.PrintDefaults()
			os.Exit(3)
		}
	}

	if *queryDelay < 100 {
		fmt.Fprintf(os.Stderr, "%s", "Delay too low")
		flag.PrintDefaults()
		os.Exit(4)
	}

	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 8)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("[+] Response: %v\n", entry)
		}
	}()

	// Asked to run from file
	if *servicePath != "" {

		// Read service entries and look them up
		f, _ := os.Open(*servicePath)
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {

			line := scanner.Text()
			srvSl := strings.Split(line, ":")

			if len(srvSl) < 2 {
				fmt.Printf("[!] Skipping ... Check line format %s \n", line)
				continue
			}

			serviceStr := srvSl[0]
			servicePort := srvSl[1]

			fmt.Printf("[i] Looking up service %s (port:%s) \n", serviceStr, servicePort)

			// Start the lookup (async).
			// Warning: number of open files on OS may be an issue
			time.Sleep(*queryDelay)

			// Async vs. Sync
			if *async == true {
				go querySrv(serviceStr, entriesCh)
			} else {
				querySrv(serviceStr, entriesCh)
			}
		}
		fmt.Println("All done")

	} else {
		// Asked to run directly

		fmt.Println("[i] Looking up service : ", *serviceString)
		fmt.Println("[i] Ctrl-C when done listening for responses\n")
		querySrv(*serviceString, entriesCh)

		wait()
		fmt.Println("All done")
	}

	close(entriesCh)
}

func querySrv(service string, entriesCh chan *mdns.ServiceEntry) {
	err := mdns.Lookup(service, entriesCh)
	if err != nil {
		fmt.Println(err)
	}
}

func wait() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v2"
	osfamily "github.com/Ullaakut/nmap/v2/pkg/osfamilies"
)

func main() {
	// Simple example
	simpleExample()

	// count_hosts_by_os
	countHostsByOs()

	// service_detection
	serviceDetection()

	// spoof_and_decoys
	spoofAndDecoys()
}

/********************************************
			Simple example
*********************************************/
func simpleExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("google.com", "facebook.com", "youtube.com"),
		nmap.WithPorts("80,443,843"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}

/********************************************
			countHostsByOs
*********************************************/
func countHostsByOs() {
	// Equivalent to
	// nmap -F -O 192.168.0.0/24
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("192.168.0.0/24"),
		nmap.WithFastMode(),
		nmap.WithOSDetection(),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("nmap scan failed: %v", err)
	}

	countByOS(result)
}

func countByOS(result *nmap.Run) {
	var (
		linux, windows int
	)

	// Count the number of each OS for all hosts.
	for _, host := range result.Hosts {
		for _, match := range host.OS.Matches {
			for _, class := range match.Classes {
				switch class.OSFamily() {
				case osfamily.Linux:
					linux++
				case osfamily.Windows:
					windows++
				}
			}

		}
	}

	fmt.Printf("Discovered %d linux hosts and %d windows hosts out of %d total up hosts.\n", linux, windows, result.Stats.Hosts.Up)
}

/********************************************
			serviceDetection
*********************************************/
func serviceDetection() {
	// Equivalent to
	// nmap -sV -T4 192.168.0.0/24 with a filter to remove non-RTSP ports.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("192.168.0.0/24"),
		nmap.WithPorts("554", "8554"),
		nmap.WithServiceInfo(),
		nmap.WithTimingTemplate(nmap.TimingAggressive),
		// Filter out ports that are not RTSP
		nmap.WithFilterPort(func(p nmap.Port) bool {
			return p.Service.Name == "rtsp"
		}),
		// Filter out hosts that don't have any open ports
		nmap.WithFilterHost(func(h nmap.Host) bool {
			// Filter out hosts with no open ports.
			for idx := range h.Ports {
				if h.Ports[idx].Status() == "open" {
					return true
				}
			}

			return false
		}),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("nmap scan failed: %v", err)
	}

	for _, host := range result.Hosts {
		fmt.Printf("Host %s\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d open with RTSP service\n", port.ID)
		}
	}
}

/********************************************
			spoofAndDecoys
*********************************************/
func spoofAndDecoys() {
	// Equivalent to
	// nmap -sS 192.168.0.10 \
	// -D 192.168.0.2,192.168.0.3,192.168.0.4,192.168.0.5,192.168.0.6,ME,192.168.0.8 \
	// 192.168.0.72`.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("192.168.0.72"),
		nmap.WithSpoofIPAddress("192.168.0.10"),
		nmap.WithDecoys(
			"192.168.0.2",
			"192.168.0.3",
			"192.168.0.4",
			"192.168.0.5",
			"192.168.0.6",
			"ME",
			"192.168.0.8",
		),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("nmap scan failed: %v", err)
	}

	printResults(result)
}

func printResults(result *nmap.Run) {
	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}

package main

import (
	"ipclinux/setupcheck/input"
	"ipclinux/setupcheck/router"
	"log"
	"net"
	"sync"
)

func main() {

	ifaces := router.GetInterfaces()

	opt, err := input.SelectInterface(len(ifaces))

	if err != nil {
		log.Fatal(err)
	}

	selectedIface := router.PickSelectedInterface(opt, ifaces)

	var wg sync.WaitGroup
	for _, selectedIface := range selectedIface {
		wg.Add(1)
		// Start up a scan on each interface.
		go func(iface net.Interface) {
			defer wg.Done()
			if err := router.Scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(selectedIface)
	}
	wg.Wait()
}

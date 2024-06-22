package lan

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/grandcat/zeroconf"
)

var (
	addresses = make(map[string]net.IP)
	requests  = make(map[string]chan net.IP)
	m         = sync.Mutex{}
)

func setAddress(name string, address net.IP) {
	nameparts := strings.Split(name, "_")
	name = nameparts[len(nameparts)-1]
	m.Lock()
	addresses[name] = address

	if request, ok := requests[name]; ok {
		request <- address
		close(request)
	}

	m.Unlock()
}

func RequestAddress(name string) chan net.IP {
	ch := make(chan net.IP, 1)
	m.Lock()

	if address, ok := addresses[name]; ok {
		ch <- address
		close(ch)
	} else {
		requests[name] = ch
		go func(n string) {
			time.Sleep(10 * time.Second)
			cancelRequest(n)
		}(name)
	}

	m.Unlock()
	return ch
}

func cancelRequest(name string) {
	m.Lock()

	if request, ok := requests[name]; ok {
		request <- nil
		close(request)
		delete(requests, name)
	}

	m.Unlock()
}

func init() {
	go func() {
		resolver, err := zeroconf.NewResolver(nil)
		if err != nil {
			panic(fmt.Errorf("unable to initialize zeroconf resolver: %v", err))
		}

		entries := make(chan *zeroconf.ServiceEntry)
		go func(results <-chan *zeroconf.ServiceEntry) {
			for entry := range results {
				setAddress(entry.Instance, entry.AddrIPv4[0])
			}
		}(entries)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err = resolver.Browse(ctx, "_dyson_mqtt._tcp", "local.", entries)
		if err != nil {
			log.Fatalln("Failed to browse:", err.Error())
		}

		<-ctx.Done()
	}()
}

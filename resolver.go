package cast

import (
	"log"
	"net"

	bonjour "github.com/oleksandr/bonjour"
)

type Resolver struct {
	bonjourResolver *bonjour.Resolver
}

func NewResolver(iface *net.Interface) (*Resolver, error) {
	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		return nil, err
	}

	return &Resolver{resolver}, nil
}

func (resolver *Resolver) Browse(chromecasts chan<- *Chromecast) error {
	serviceEntries := make(chan *bonjour.ServiceEntry)

	go func(chromecasts chan<- *Chromecast, serviceEntries <-chan *bonjour.ServiceEntry) {
		for entry := range serviceEntries {
			chromecasts <- &Chromecast{entry}
		}
		close(chromecasts)
	}(chromecasts, serviceEntries)

	err := resolver.bonjourResolver.Browse("_googlecast._tcp", "local.", serviceEntries)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
		return err
	}

	return nil
}

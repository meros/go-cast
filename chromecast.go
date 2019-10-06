package cast

import (
	"crypto/tls"
	"log"
	"strconv"

	bonjour "github.com/oleksandr/bonjour"
)

type Chromecast struct {
	serviceEntry *bonjour.ServiceEntry
}

func (chromecast *Chromecast) Name() string {
	return chromecast.serviceEntry.ServiceInstanceName()
}

func (chromecast *Chromecast) Connect() (*Channel, error) {
	conn, err := tls.Dial(
		"tcp",
		chromecast.serviceEntry.AddrIPv4.String()+":"+strconv.Itoa(chromecast.serviceEntry.Port),
		&tls.Config{InsecureSkipVerify: true})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Channel{conn}, nil
}

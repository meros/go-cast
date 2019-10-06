package cast

import (
	"log"
	"os"
	"testing"
)

func TestBrowse(t *testing.T) {
	resolver, err := NewResolver(nil)
	if err != nil {
		log.Println("Failed to initialize resolver:", err.Error())
		os.Exit(1)
	}

	results := make(chan *Chromecast)

	go func(results chan *Chromecast) {
		for chromecast := range results {
			channel, err := chromecast.Connect()
			if err != nil {
				log.Println(err)
				continue
			}

			channel.Connect()
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Got one %s!\n", chromecast.Name())
		}
		log.Println("Done...")

	}(results)

	err = resolver.Browse(results)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
	}

	select {}

}

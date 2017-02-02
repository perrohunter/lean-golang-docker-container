package main

import (
	"log"
	"net/http"
	"io"
	"os"
	"github.com/nats-io/go-nats"
)

func main() {

	response, err := http.Get("http://nats:8222/varz")
	if err != nil {
		log.Println(err)
	} else {
		defer response.Body.Close()
		asd, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Println(err)
		}
		log.Println("nats-x")
		log.Println(asd)
	}
	response, err = http.Get("http://nats-x:8222/varz")
	if err != nil {
		log.Println(err)
	} else {
		defer response.Body.Close()
		asd, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Println(err)
		}
		log.Println("nats")
		log.Println(asd)
	}

	xnc, err := nats.Connect("nats://nats:4222")

	if err != nil {
		log.Println("Fail 1")
		log.Fatal(err)
	}
	nc, err := nats.NewEncodedConn(xnc, nats.JSON_ENCODER)

	if err != nil {
		log.Println("Fail 2")
		log.Fatal(err)
	}
	log.Println(nc)

}

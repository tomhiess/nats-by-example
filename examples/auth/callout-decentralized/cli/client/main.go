package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		natsURL   string
		credsFile string
		user      string
		pass      string
	)

	flag.StringVar(&natsURL, "url", nats.DefaultURL, "NATS URL")
	flag.StringVar(&credsFile, "creds", "", "NATS credentials file")
	flag.StringVar(&user, "user", "", "NATS user")
	flag.StringVar(&pass, "pass", "", "NATS password")

	flag.Parse()

	nc, err := nats.Connect(
		natsURL,
		nats.UserCredentials(credsFile),
		nats.UserInfo(user, pass),
	)
	if err != nil {
		return err
	}
	defer nc.Drain()

	log.Printf("%s connected to %s", user, nc.ConnectedUrl())

	_, err = nc.Subscribe(">", func(msg *nats.Msg) {

	})

	if err != nil {
		log.Printf("could not subscribe: %s", err)
		return err
	}

	for nc.LastError() == nil {
		time.Sleep(time.Second)
		log.Printf("waiting for revocation")
		log.Printf("is connected %v", nc.IsConnected())
	}

	log.Printf("%s", nc.LastError())

	return nil
}

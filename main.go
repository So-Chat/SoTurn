package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pion/logging"
	"github.com/pion/turn/v5"
)

var showVersion bool

func main() {
	applicationVersion := 0.1

	flag.BoolVar(&showVersion, "version", false, "Check application version")
	publicIP := flag.String("public-ip", "", "IP Address that TURN can be contacted by.")
	port := flag.Int("port", 3478, "Listening port.")
	realm := flag.String("realm", "", "Realm")
	jwtSecret := flag.String("jwt", "", "JWT key for generating keys")
	flag.Parse()

	if showVersion {
		log.Println("Application version:", applicationVersion)
		os.Exit(0)
	}

	if len(*publicIP) == 0 {
		log.Fatalf("'publicIP' is required")
	} else if len(*jwtSecret) == 0 {
		log.Fatalf("JWT secret not provided")
	}

	updListener, err := net.ListenPacket("udp4", net.JoinHostPort("0.0.0.0", strconv.Itoa(*port)))
	if err != nil {
		log.Panicf("Failed: %v", err)
	}

	logger := logging.NewDefaultLeveledLoggerForScope("lt-creds", logging.LogLevelTrace, os.Stdout)

	server, err := turn.NewServer(turn.ServerConfig{
		Realm:       *realm,
		AuthHandler: turn.LongTermTURNRESTAuthHandler(*jwtSecret, logger),
		PacketConnConfigs: []turn.PacketConnConfig{
			{
				PacketConn: updListener,
				RelayAddressGenerator: &turn.RelayAddressGeneratorStatic{
					RelayAddress: net.ParseIP(*publicIP),
					Address:      "0.0.0.0",
				},
			},
		},
	})
	if err != nil {
		log.Panic(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	if err = server.Close(); err != nil {
		log.Panic(err)
	}

}

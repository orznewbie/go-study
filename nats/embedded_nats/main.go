package main

import (
	"flag"
	"fmt"

	"github.com/nats-io/nats-server/v2/server"
)

var (
	c = flag.String("c", "/etc/nats/server.conf", "config file")
)

func main() {
	flag.Parse()

	opts := &server.Options{}
	if err := opts.ProcessConfigFile(*c); err != nil {
		panic(err)
	}
	opts.JetStream = true

	srv, err := server.NewServer(opts)
	if err != nil {
		panic(fmt.Errorf("failed to new nats server: %v", err))
	}

	srv.ConfigureLogger()
	srv.Start()

	//if !srv.ReadyForConnections(30 * time.Second) {
	//	panic(fmt.Errorf("failed to connect to NATS server failed after 30s"))
	//}

	//go func() {
	//	fmt.Println("waiting nats server to start for 7s")
	//	time.Sleep(7 * time.Second)
	//	conn, err := nats.Connect("nats://127.0.0.1:4222", nats.Name("kine using bucket: kine"))
	//	if err != nil {
	//		panic(fmt.Errorf("failed to connect to NATS server: %w", err))
	//	}
	//	js, err := conn.JetStream()
	//	if err != nil {
	//		panic(fmt.Errorf("failed to get JetStream context: %w", err))
	//	}
	//	_, err = js.KeyValue("kine")
	//	if err == nats.ErrBucketNotFound {
	//		fmt.Println("bucket kine not found")
	//	} else if err != nil {
	//		fmt.Printf("lookup kine kv bucket: %v", err)
	//	}
	//}()

	srv.WaitForShutdown()
}

package nats

import (
	"fmt"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestStreamInfo(t *testing.T) {
	conn, err := nats.Connect("nats://192.168.3.11:4222")
	if err != nil {
		t.Fatal(err)
	}
	js, err := conn.JetStream(nats.MaxWait(5 * time.Second))
	if err != nil {
		t.Fatal(err)
	}

	info, err := js.StreamInfo("KV_kine")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
}

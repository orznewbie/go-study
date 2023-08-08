package nats

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/shengdoushi/base58"
)

func TestKeyValue(t *testing.T) {
	conn, err := nats.Connect("nats://192.168.3.11:4222")
	if err != nil {
		t.Fatal(err)
	}
	js, err := conn.JetStream()
	if err != nil {
		t.Fatal(err)
	}

	bucket, err := js.KeyValue("kine")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bucket.Bucket())
}

func TestKVGet(t *testing.T) {
	conn, err := nats.Connect("nats://192.168.3.11:4222")
	if err != nil {
		t.Fatal(err)
	}
	js, err := conn.JetStream(nats.MaxWait(5 * time.Second))
	if err != nil {
		t.Fatal(err)
	}

	bucket, err := js.KeyValue("kine")
	if err != nil {
		t.Fatal(err)
	}

	key, err := encode("/registry/leases/kube-node-lease/vone")
	if err != nil {
		t.Fatal(err)
	}
	entry, err := bucket.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("bucket=%s,key=%s,value=%s\n", entry.Bucket(), entry.Key(), entry.Value())
}

var (
	keyAlphabet = base58.BitcoinAlphabet
)

func encode(key string) (retKey string, e error) {
	var parts []string
	for _, part := range strings.Split(strings.TrimPrefix(key, "/"), "/") {
		if part == ">" || part == "*" {
			parts = append(parts, part)
			continue
		}
		parts = append(parts, base58.Encode([]byte(part), keyAlphabet))

	}
	if len(parts) == 0 {
		return "", nats.ErrInvalidKey
	}
	return strings.Join(parts, "."), nil
}

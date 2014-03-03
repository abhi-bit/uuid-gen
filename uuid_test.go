package uuid

import (
    "testing"
)

func TestUUIDv4(t *testing.T) {
    uuid, err := GenUUIDv4()
    if err != nil {
        t.Fatalf("GenUUIDv4 error %s", err)
    }
    t.Logf("uuid[%s]\n", uuid)
}

func TestUUIDRand(t *testing.T) {
    uuid, err := GenUUIDRand()
    if err != nil {
        t.Fatalf("GenUUIDRand error %s", err)
    }
    t.Logf("uuid[%s]\n", uuid)
}

func BenchmarkUUIDv4(b *testing.B) {
    m := make(map[string]int, 100)
    for i:= 0; i < b.N; i++ {
        uuid, err := GenUUIDv4()
        if err != nil {
            b.Fatalf("GenUUIDv4 error %s", err)
        }
        b.StopTimer()
        c := m[uuid]
        if c > 0 {
            b.Fatalf("duplicate uuid[%s] count %d", uuid, c)
        }
        m[uuid] = c + 1
        b.StartTimer()
    }
}

func BenchmarkUUIDRand(b *testing.B) {
    m := make(map[string]int, 1000)
    for i:= 0; i < b.N; i++ {
        uuid, err := GenUUIDRand()
        if err != nil {
            b.Fatalf("GenUUIDRand error %s", err)
        }
        b.StopTimer()
        c := m[uuid]
        if c > 0 {
            b.Fatalf("duplicate uuid[%s] count %d", uuid, c)
        }
        m[uuid] = c + 1
        b.StartTimer()
    }
}

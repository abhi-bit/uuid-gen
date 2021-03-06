package uuid

import (
    "os"
    "fmt"
    "log"
    "crypto/rand"
)

func GenUUIDv4() (string, error) {
    uuid := make([]byte, 16)
    n, err := rand.Read(uuid)
    if n != len(uuid) || err != nil {
        return "", err
    }

    //Version 4 RFC4122
    uuid[6] = (uuid[6] & 0x0f) | 0x40
    uuid[8] = (uuid[8] & 0x3f) | 0x80

    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), err
}

func GenUUIDRand() (string, error) {
    file, err := os.Open("/dev/urandom")
    defer file.Close()
    if err != nil {
        log.Fatal("Unable to open /dev/urandom")
    }

    uuid := make([]byte, 16)
    file.Read(uuid)

    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), err
}

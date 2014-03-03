package uuid

import (
    "encoding/hex"
    "crypto/rand"
)

func GenUUID() (string, error) {
    uuid := make([]byte, 16)
    n, err := rand.Read(uuid)
    if n != len(uuid) || err != nil {
        return "", err
    }

    //Version 4 RFC4122
    uuid[8] = 0x80
    uuid[4] = 0x40

    return hex.EncodeToString(uuid), nil
}

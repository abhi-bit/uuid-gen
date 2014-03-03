package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    uuid "github.com/abhi-bit/uuid-gen"
)

func UUIDServer(w http.ResponseWriter, req *http.Request) {
    values := req.URL.Query()
    var count int
    var err error
    fmt.Println("Values:", values)
    if values != nil {
        cnt := values.Get("count")
        if cnt != "" {
            count, err = strconv.Atoi(cnt)
            if err != nil {
                log.Fatal(w, "{\"Error\":\"%s\"}\n", err)
                return
            }
        }
    }
    for ; count > 0; count -= 1 {
        uuid, err := uuid.GenUUID()
        if err != nil {
            fmt.Fprintf(w, "{\"Error\":\"%s\"}\n", err)
        } else {
            fmt.Fprintf(w, "{\"uuid\":[\"%s\"]}\n", uuid)
        }
    }
}

func main() {
    http.HandleFunc("/_uuid", UUIDServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

package main

import (
    "fmt"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "testing"
)

func TestAPIPost(t *testing.T) {
    ts := httptest.NewServer(setupServer())
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

    values := map[string]string{"roman": "MCMXCIV"}
    jsonValue, _ := json.Marshal(values)
    resp, err := http.Post(fmt.Sprintf("%s/roman", ts.URL), "application/json", bytes.NewBuffer(jsonValue))
    // resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
    }

    // Assert that the "content-type" header is actually set
    val, _ := resp.Header["Content-Type"]

    if val[0] != "application/json; charset=utf-8" {
        t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
    }
}

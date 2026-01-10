package main

import (
	"fmt"
	"io"
	"scope-guard"
)

func main() {
	raw := "https://example.com"

	u, err := scopeguard.ValidateURL(raw)
	if err != nil {
		panic(err)
	}

	_, err = scopeguard.ResolveAndCheck(u.Hostname())
	if err != nil {
		panic(err)
	}

	client := scopeguard.NewSafeClient()
	resp, err := client.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("OK:", len(body))
}

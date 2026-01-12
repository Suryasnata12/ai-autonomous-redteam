package ownership

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func VerifyHTTP(domain, token string) error {
	url := fmt.Sprintf("https://%s/.well-known/redforge.txt", domain)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	expected := fmt.Sprintf("redforge-verification=%s", token)

	if strings.TrimSpace(string(body)) == expected {
		return nil
	}
	return ErrVerificationFailed
}

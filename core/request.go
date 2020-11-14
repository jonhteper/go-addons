package core

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

// WARNING: Esta función permite conectarse a dominios con certificado autofirmado.
// Utilizar solo en dominios de confianza.
//
// SimplePostRequest realiza una request con el método POST con autorización
//tipo Bearer token.
//
// El parámetro `body` debe ser escrito con "formato JSON", por ejemplo:
//
// response, err := SimplePostRequest("https://example.com", "my.awesome.Token", `{"param":"value"}`)
//
func SimplePostRequest(url, token, body string) (*http.Response, error) {
	data := []byte(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client.Do(req)
}

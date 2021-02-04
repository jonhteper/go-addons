package addons

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

// WARNING: Esta función permite conectarse a dominios con certificado autofirmado.
// Utilizar solo en dominios de confianza.
//
// PostRequest realiza una request con el método POST con autorización tipo Bearer
// token.
//
// EXAMPLE:
//	body := struct {
//		Message string `json:"message"`
//	}{ Message: "Hola Mundo" }
//	data, _ := json.Marshal(body)
//
//  response, err := PostRequest("https://example.com", "my.awesome.Token", data)
//
func PostRequest(url, token string, body []byte) (*http.Response, error) {
	return request(http.MethodPost, url, token, body)
}

// WARNING: Esta función permite conectarse a dominios con certificado autofirmado.
// Utilizar solo en dominios de confianza.
//
// PutRequest realiza una request con el método PUT con autorización tipo Bearer
// token.
//
// EXAMPLE:
//	body := struct {
//		Message string `json:"message"`
//	}{ Message: "Hola Mundo" }
//	data, _ := json.Marshal(body)
//
//  response, err := PutRequest("https://example.com", "my.awesome.Token", data)
//
func PutRequest(url, token string, body []byte) (*http.Response, error) {
	return request(http.MethodPut, url, token, body)
}

// WARNING: Esta función permite conectarse a dominios con certificado autofirmado.
// Utilizar solo en dominios de confianza.
//
// DeleteRequest realiza una request con el método PUT con autorización tipo Bearer
// token.
//
// EXAMPLE:
//	body := struct {
//		ID string `json:"id"`
//	}{ ID: "MyElementID" }
//	data, _ := json.Marshal(body)
//
//  response, err := DeleteRequest("https://example.com", "my.awesome.Token", data)
//
func DeleteRequest(url, token string, body []byte) (*http.Response, error) {
	return request(http.MethodDelete, url, token, body)
}

func request(method, url, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
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

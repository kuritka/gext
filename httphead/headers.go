package httphead

import "net/http"

func HeaderAsMap(header http.Header) map[string]string{
	result := make(map[string]string)
	for k, v := range header {
		values := ""
		for _, headerValue := range v {
			values += headerValue + " "
		}
		result[k] = values
	}
	return result
}

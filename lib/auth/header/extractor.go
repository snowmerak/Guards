package header

import "net/http"

func ExtractFromHttp(request *http.Request, config *Config) (Headers, error) {
	h := map[string]string{}
	for _, key := range config.Requires {
		v := request.Header.Values(key)
		if len(v) == 0 {
			return nil, ErrMissingHeader
		}
		h[key] = v[0]
	}
	return h, nil
}

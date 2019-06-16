package onion

import (
	"os"
	"strings"
)

// NewEnvLayer create new layer using the whitelist of environment values.
func NewEnvLayer(separator string, whiteList ...string) Layer {
	data := make(map[string]interface{})
	for i := range whiteList {
		if s := os.Getenv(whiteList[i]); s != "" {
			data[strings.ToLower(whiteList[i])] = s
		}
	}

	return NewMapLayerSeparator(data, separator)
}

// NewEnvLayerPrefix create new env layer, with all values with the same prefix
func NewEnvLayerPrefix(separator string, prefix string) Layer {
	data := make(map[string]interface{})
	pf := strings.ToUpper(prefix) + separator
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, pf) {
			k := strings.Trim(strings.Split(env, "=")[0], "\t\n ")
			ck := strings.ToLower(strings.TrimPrefix(k, pf))
			data[ck] = os.Getenv(k)
		}
	}

	return NewMapLayerSeparator(data, separator)
}

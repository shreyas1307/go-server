package utilities

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

// ResponseBodyUnmarshal unmarshals the response body into the provided data struct.
//
// - b is the byte slice containing the response body
//
// - dt is the destination struct where the unmarshalled data will be stored
//
// - requestPath is the path of the request for logging purposes
//
// Returns true if unmarshalling is successful, false otherwise.
func ResponseBodyUnmarshal(b []byte, dt any, requestPath string) bool {
	if err := json.Unmarshal(b, &dt); err != nil {
		LogError("error unmarshalling response body", err, "request", requestPath)
		return false
	}
	return true
}

// LogError logs an error message with additional context.
//
// - msg is the error message
//
// - err is the error to log
//
// - key and val are key:value for additional context
func LogError(msg string, err error, key string, val string) {
	log.Error().AnErr(msg, err).Str(key, val).Send()
}

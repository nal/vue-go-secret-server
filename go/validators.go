/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func validateSecret(secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("secret is required")
	}

	return secret, nil
}

func validateExpireAfterViews(expireAfterViews string) (int32, error) {
	if len(expireAfterViews) == 0 {
		return 0, errors.New("expireAfterViews is required")
	}

	// string, converting to int32
	i64, err := strconv.ParseInt(expireAfterViews, 10, 32)

	if err != nil {
		return 0, errors.New("expireAfterViews should be int32 value")
	}

	i32 := int32(i64)
	if i32 <= 0 {
		return 0, errors.New("expireAfterViews should be positive int32 value")
	}

	return i32, nil
}

func validateInt32Value(possibleInt32 string, paramName string) (int32, error) {
	if len(possibleInt32) == 0 {
		return 0, fmt.Errorf("%s is required", paramName)
	}

	// string, converting to int32
	i64, err := strconv.ParseInt(possibleInt32, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("%s should be int32 value", paramName)
	}

	i32 := int32(i64)
	if i32 <= 0 {
		return 0, fmt.Errorf("%s should be positive int32 value", paramName)
	}

	return i32, nil
}

func validateHash(hash string) (string, error) {
	if len(hash) == 0 {
		return "", errors.New("hash is required")
	}

	matched, _ := regexp.MatchString("^[a-z0-9]{64}$", hash)
	if !matched {
		return "", errors.New("hash has invalid format")
	}

	return hash, nil
}

func renderInternalError(w http.ResponseWriter, logString string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("{'error':1,'error_code':'internal error'}"))
}

func renderValidationError(w http.ResponseWriter, logString string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("{\"error\":1,\"error_message\":\"%s\"}", logString)))
}

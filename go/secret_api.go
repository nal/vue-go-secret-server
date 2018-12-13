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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Index serves index page
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

// AddSecret validates input variables and tries to generate unique URL
// and returns it in JSON format
func AddSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// First parse request to get required form data
	r.ParseForm()

	// Collect and validate required data
	// secret
	secret, err := validateSecret(r.FormValue("secret"))
	if err != nil {
		renderValidationError(w, fmt.Sprint(err))
		return
	}

	// expireAfterViews
	expireAfterViews, err := validateInt32Value(r.FormValue("expireAfterViews"), "expireAfterViews")
	if err != nil {
		renderValidationError(w, fmt.Sprint(err))
		return
	}

	// expireAfter
	expireAfter, err := validateInt32Value(r.FormValue("expireAfter"), "expireAfter")
	if err != nil {
		renderValidationError(w, fmt.Sprint(err))
		return
	}

	// CreatedAt
	timeNow := time.Now()

	// ExpiresAt
	timeExpires := timeNow.Add(time.Minute * time.Duration(int64(expireAfter)))

	// Hash
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(secret)))

	log.Printf("secret = %s, expireAfterViews = %d, expireAfter = %d, time = %s", secret, expireAfterViews, expireAfter, timeNow)

	if err := addSecretStore(hash, secret, expireAfterViews, expireAfter); err != nil {
		renderInternalError(w, fmt.Sprintf("Failed to store secret: %s", err))
		return
	}

	secretData := &secretStruct{
		Hash:           hash,
		SecretText:     secret,
		CreatedAt:      timeNow,
		ExpiresAt:      timeExpires,
		RemainingViews: expireAfterViews}
	secretDataJSON, _ := json.Marshal(secretData)

	w.WriteHeader(http.StatusOK)
	w.Write(secretDataJSON)
}

// addSecretStore stores secret in Redis DB. Returns error if failed to store data in Redis DB.
func addSecretStore(hash string, secret string, expireAfterViews int32, expireAfter int32) error {
	redisConn, err := Redis()
	if err != nil {
		return err
	}
	defer redisConn.Close()

	counterViews := strings.Join([]string{hash, "counter"}, "-")
	redisConn.Send("MULTI")
	// Let website user overwrite settings with each subsequent request
	redisConn.Send("DEL", hash)
	redisConn.Send("DEL", counterViews)
	redisConn.Send("SET", hash, secret)
	redisConn.Send("EXPIRE", hash, expireAfter*60)
	redisConn.Send("SET", counterViews, expireAfterViews)
	redisConn.Send("EXPIRE", counterViews, expireAfter*60)
	if _, err := redisConn.Do("EXEC"); err != nil {
		return err
	}
	return nil
}

// GetSecretByHash searches for URL in database and returns secret generated by AddSecret
// in JSON format. If secret not found then it returns error message.
func GetSecretByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

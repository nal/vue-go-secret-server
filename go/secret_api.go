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
	"encoding/json"
	"log"
	"net/http"
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
	w.WriteHeader(http.StatusOK)

	// First parse request to get required form data
	r.ParseForm()

	// Collect required data
	// TODO: validation
	// TODO: incorrect request params in app.js, cant' decode it

	// 2018/12/09 17:40:50 POST /v1/secret map[{"secret":"qwerty","expireAfterViews":"11","expireAfter":"22"}:[]] AddSecret 129.53µs
	// VS
	// 2018/12/09 17:41:11 POST /v1/secret map[secret:[11] expireAfterViews:[22] expireAfter:[33]] AddSecret 83.518µs

	secret := r.FormValue("secret")
	expireAfterViews := r.FormValue("expireAfterViews")
	expireAfter := r.FormValue("expireAfter")

	time := time.Now()

	log.Printf("secret = %s, expireAfterViews = %s, expireAfter = %s, time = %s", secret, expireAfterViews, expireAfter, time)

	resD := &Secret{
		Hash:           "1234567890123456789012345678901234567890", // TODO: sha256 sum
		SecretText:     secret,
		CreatedAt:      time, // TODO time.Time,
		ExpiresAt:      time, // TODO add ExpiresAt to current time
		RemainingViews: 123}
	resB, _ := json.Marshal(resD)

	w.Write(resB)
}

// GetSecretByHash searches for URL in database and returns secret generated by AddSecret
// in JSON format. If secret not found then it returns error message.
func GetSecretByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
)

// ChallengeCreate function returns Comment with given id in json format
func ChallengeCreate(w http.ResponseWriter, r *http.Request) {
	// 1. Decode Request into Challenge struct

	// 2. Connect to database

	// 3. Store comment entity in database

	// 4. Cast Challenge to JSON

	// 5. Send response
	_, _ = fmt.Fprint(w, "Hello, World!\n")
}

// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ChallengeCreate function returns Comment with given id in json format
func ChallengeCreate(w http.ResponseWriter, r *http.Request) {
	// 1. Decode Request into Challenge struct
	var challenge Challenge
	err := json.NewDecoder(r.Body).Decode(&challenge)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
	}
	challenge.Created = time.Now()

	// 2. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)

	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
		return
	}
	// 3. Store comment entity in database
	key := datastore.IncompleteKey(EntityName, nil)
	key, err = client.Put(ctx, key, &challenge)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
		return
	}

	// 4. Cast Challenge to JSON
	challenge.ChallengeId = key.ID
	byteArray, err := json.Marshal(challenge)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 5. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))

}

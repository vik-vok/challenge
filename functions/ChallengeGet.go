package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ChallengeGet function returns Challenges list with given user id in JSON format
func ChallengeGet(w http.ResponseWriter, r *http.Request) {
	// 1. Write ID from request into struct d
	var d struct {
		ReceiverUserId string `json:"receiverUserId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		arr := r.URL.Query()["receiverUserId"]
		if len(arr) < 1 {
			_, _ = fmt.Fprint(w, "Error While Parsing Request Body!\n URL: "+r.URL.String())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		d.ReceiverUserId = arr[0]
		if d.ReceiverUserId == "" {
			_, _ = fmt.Fprint(w, "Error While Parsing Request Body!\n URL: "+r.URL.String())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// 2. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. Get data
	var challenges []Challenge
	query := datastore.NewQuery(EntityName).Filter("ReceiverUserId =", d.ReceiverUserId)
	ids, err := client.GetAll(ctx, query, &challenges)

	if challenges == nil {
		challenges = []Challenge{}
		ids = []*datastore.Key{}
	}
	// Iterate over objects and append ID-s
	for i := range challenges {
		challenges[i].ChallengeId = ids[i].ID
	}

	// 4. Cast Comment to JSON
	byteArray, err := json.Marshal(challenges)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))
}

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
		ReceiverUserId int64 `json:"receiverUserId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		_, _ = fmt.Fprint(w, "Error While Parsing Request Body!")
		return
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

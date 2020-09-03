package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

var Message struct {
	ReceiverUserId string  `json:"receiverUserId"`
	Score          float32 `json:"score"`
}

func ChallengeUpdate(ctx context.Context, m PubSubMessage) {

	data := string(m.Data) // Automatically decoded from base64.

	message := &Message
	err := json.Unmarshal([]byte(data), message)
	if err != nil {
		log.Printf("Error While Parsing Request Body!")
		return
	}

	// 2. Connect to database
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		return
	}

	// 3. Get data
	var challenges []Challenge
	query := datastore.NewQuery(EntityName).Filter("ReceiverUserId =", message.ReceiverUserId).Filter("Score =", message.Score)
	ids, err := client.GetAll(ctx, query, &challenges)
	accomplished := 0
	for i := range challenges {
		if challenges[i].Accomplished {
			continue
		}
		if message.Score >= challenges[i].Score {
			challenges[i].Accomplished = true
			challengeKey := datastore.IDKey(EntityName, ids[i].ID, nil)
			if _, err := client.Put(ctx, challengeKey, &challenges[i]); err != nil {
				log.Fatalf("tx.Put: %v", err)
			}
			accomplished++
		}

	}
	log.Printf("%d\n", accomplished)
	return
}

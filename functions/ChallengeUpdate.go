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
	ReceiverUserId  string  `json:"receiverUserId"`
	OriginalVoiceId string  `json:"originalVoiceId"`
	Score           float32 `json:"score"`
}

func ChallengeUpdate(ctx context.Context, m PubSubMessage) error {

	data := string(m.Data) // Automatically decoded from base64.
	log.Printf("%s\n", data)
	message := &Message
	err := json.Unmarshal([]byte(data), message)
	log.Printf(message.ReceiverUserId)
	log.Printf(message.OriginalVoiceId)
	log.Printf("%.6f", message.Score)
	if err != nil {
		log.Printf("Error While Parsing Request Body!")
		return err
	}

	//2. Connect to database
	ctxDb := context.Background()
	client, err := datastore.NewClient(ctxDb, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		return err
	}

	// 3. Get data
	var challenges []Challenge

	query := datastore.NewQuery(EntityName).Filter("ReceiverUserId =", message.ReceiverUserId).Filter("OriginalVoiceId =", message.OriginalVoiceId)
	ids, err := client.GetAll(ctxDb, query, &challenges)
	accomplished := 0
	log.Printf("%d\n", len(challenges))
	for i := range challenges {
		fmt.Println(challenges[i])
		if challenges[i].Accomplished {
			continue
		}
		if message.Score >= challenges[i].Score {
			challenges[i].Accomplished = true
			challengeKey := datastore.IDKey(EntityName, ids[i].ID, nil)
			if _, err := client.Put(ctxDb, challengeKey, &challenges[i]); err != nil {
				log.Fatalf("tx.Put: %v", err)
			}
			accomplished++
		}

	}
	log.Printf("%d\n", accomplished)
	return nil
}

package p

import "time"

// Comment is good
type Challenge struct {
	ChallengeId     int64     `json:"challengeId" datastore:"-" `
	OriginalVoiceId int64     `json:"originalVoiceId"`
	SenderUserId    string    `json:"senderUserId"`
	ReceiverUserId  string    `json:"receiverUserId"`
	Score           int64     `json:"score"`
	Accomplished    bool      `json:"accomplished"`
	Created         time.Time `json:"created"`
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Challenge"

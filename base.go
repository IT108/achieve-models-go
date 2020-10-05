package achieve_models_go

import (
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Request struct {
	self   RequestInterface
	Sender string
	User   string
	GateId string
	Data   string
}

type RequestInterface interface {
	AcceptReq(kafkaMsg *kafka.Message)
	Serialize()
}

func (receiver *Request) AcceptReq(kafkaMsg *kafka.Message) {
	err := json.Unmarshal(kafkaMsg.Value, receiver)
	if err != nil {
		panic(err)
	}
}

func (receiver *Request) Serialize() []byte {
	data, _ := json.Marshal(receiver)
	return data
}

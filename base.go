package achieve_models_go

import (
	"encoding/json"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Request struct {
	self      RequestInterface
	RequestId string `json:"request_id"`
	Method    string `json:"method"`
	Sender    string `json:"sender"`
	User      string `json:"user"`
	GateId    string `json:"gate_id"`
	Data      string `json:"data"`
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

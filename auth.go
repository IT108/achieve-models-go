package achieve_models_go

const (
	AUTH_REGISTER_KEY = "register"
)

type RegisterRequest struct {
	Request
	Email string
	Username string
	Password string
}

//func (receiver *RegisterRequest) AcceptReq(kafkaMsg *kafka.Message) {
//	err := json.Unmarshal(kafkaMsg.Value, receiver)
//	if err != nil {
//		panic(err)
//	}
//}
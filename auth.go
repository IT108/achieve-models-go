package achieve_models_go

const (
	AUTH_REGISTER_KEY = "register"
	AUTH_ISREGISTERED_KEY = "isreg"
	AUTH_IS_USER_REGISTERED_KEY = "isreg_u"
	AUTH_IS_EMAIL_REGISTERED_KEY = "isreg_e"
	AUTH_AUTHENTICATE_KEY = "authenticate"
	AUTH_AUTHORIZE_KEY = "authorize"
)

type RegisterRequest struct {
	Request
	Email    string
	Username string
	Password string
}

type RegisterResponse struct {
	Request
	ResponseCode int
	Error        string
}

type IsEmailRegisteredRequest struct {
	Request
	Email    string
}

type IsEmailResponse struct {
	Request
	ResponseCode         int
	IsEmailRegistered    bool
	Error                string
}

type IsUserRegisteredRequest struct {
	Request
	Username string
}

type IsUserRegisteredResponse struct {
	Request
	ResponseCode         int
	IsUsernameRegistered bool
	Error                string
}

type IsRegisteredRequest struct {
	Request
	Email    string
	Username string
}

type IsRegisteredResponse struct {
	Request
	ResponseCode         int
	IsEmailRegistered    bool
	IsUsernameRegistered bool
	Error                string
}

type AuthenticateRequest struct {
	Request
	Username string
	Password string
}

type AuthenticateResponse struct {
	Request
	ResponseCode int
	Error        string
}

type AuthorizeRequest struct {
	Request
	Email string
}

type AuthorizeResponse struct {
	Request
	ResponseCode int
	Roles        []string
	Error        string
}

//func (receiver *RegisterRequest) AcceptReq(kafkaMsg *kafka.Message) {
//	err := json.Unmarshal(kafkaMsg.Value, receiver)
//	if err != nil {
//		panic(err)
//	}
//}

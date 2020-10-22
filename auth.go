package achieve_models_go

const (
	AUTH_REGISTER_KEY            = "register"
	AUTH_ISREGISTERED_KEY        = "isreg"
	AUTH_IS_USER_REGISTERED_KEY  = "isreg_u"
	AUTH_IS_EMAIL_REGISTERED_KEY = "isreg_e"
	AUTH_AUTHENTICATE_KEY        = "authenticate"
	AUTH_AUTHORIZE_KEY           = "authorize"
	AUTH_REFRESH_KEY             = "auth_refresh"
)

type RegisterRequest struct {
	Request  `json:"request"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Request      `json:"request"`
	ResponseCode int    `json:"response_code"`
	Error        string `json:"error"`
}

type IsEmailRegisteredRequest struct {
	Request `json:"request"`
	Email   string `json:"email"`
}

type IsEmailResponse struct {
	Request           `json:"request"`
	IsEmailRegistered bool   `json:"is_email_registered"`
	ResponseCode      int    `json:"response_code"`
	Error             string `json:"error"`
}

type IsUserRegisteredRequest struct {
	Request  `json:"request"`
	Username string `json:"username"`
}

type IsUserRegisteredResponse struct {
	Request              `json:"request"`
	ResponseCode         int    `json:"response_code"`
	IsUsernameRegistered bool   `json:"is_username_registered"`
	Error                string `json:"error"`
}

type IsRegisteredRequest struct {
	Request  `json:"request"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type IsRegisteredResponse struct {
	Request              `json:"request"`
	ResponseCode         int    `json:"response_code"`
	IsEmailRegistered    bool   `json:"is_email_registered"`
	IsUsernameRegistered bool   `json:"is_username_registered"`
	Error                string `json:"error"`
}

type AuthenticateRequest struct {
	Request  `json:"request"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	Request      `json:"request"`
	ResponseCode int    `json:"response_code"`
	Error        string `json:"error"`
}

type AuthorizeRequest struct {
	Request `json:"request"`
	Email   string `json:"email"`
}

type AuthorizeResponse struct {
	Request      `json:"request"`
	ResponseCode int      `json:"response_code"`
	Roles        []string `json:"roles"`
	Error        string   `json:"error"`
}

//func (receiver *RegisterRequest) AcceptReq(kafkaMsg *kafka.Message) {
//	err := json.Unmarshal(kafkaMsg.Value, receiver)
//	if err != nil {
//		panic(err)
//	}
//}

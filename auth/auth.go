package auth

import (
	"github.com/IT108/achieve-models-go"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AUTH_REGISTER_KEY     = "register"
	AUTH_SIGNIN_KEY       = "signin"
	AUTH_AUTHENTICATE_KEY = "authenticate"
	AUTH_AUTHORIZE_KEY    = "authorize"
	AUTH_REFRESH_KEY      = "auth_refresh"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthToken struct {
	TokenString    string    `json:"token_string"`
	ExpirationTime time.Time `json:"expiration_time"`
}

type SigninRequest struct {
	achieve_models_go.Request `json:"request"`
	Username                  string `json:"username"`
	Password                  string `json:"password"`
}

type SigninResponse struct {
	achieve_models_go.Request `json:"request"`
	Token                     AuthToken `json:"token"`
	ResponseCode              int       `json:"response_code"`
	Error                     string    `json:"error"`
}

type RegisterRequest struct {
	achieve_models_go.Request `json:"request"`
	Email                     string `json:"email"`
	Username                  string `json:"username"`
	Password                  string `json:"password"`
}

type RegisterResponse struct {
	achieve_models_go.Request `json:"request"`
	ResponseCode              int    `json:"response_code"`
	Error                     string `json:"error"`
}

type AuthenticateRequest struct {
	achieve_models_go.Request `json:"request"`
	Username                  string `json:"username"`
	Password                  string `json:"password"`
}

type AuthenticateResponse struct {
	achieve_models_go.Request `json:"request"`
	ResponseCode              int    `json:"response_code"`
	Error                     string `json:"error"`
}

type AuthorizeRequest struct {
	achieve_models_go.Request `json:"request"`
	Email                     string `json:"email"`
}

type AuthorizeResponse struct {
	achieve_models_go.Request `json:"request"`
	ResponseCode              int      `json:"response_code"`
	Roles                     []string `json:"roles"`
	Error                     string   `json:"error"`
}

//func (receiver *RegisterRequest) AcceptReq(kafkaMsg *kafka.Message) {
//	err := json.Unmarshal(kafkaMsg.Value, receiver)
//	if err != nil {
//		panic(err)
//	}
//}

package auth

import "github.com/IT108/achieve-models-go"

const (
	AUTH_ISREGISTERED_KEY        = "isreg"
	AUTH_IS_USER_REGISTERED_KEY  = "isreg_u"
	AUTH_IS_EMAIL_REGISTERED_KEY = "isreg_e"
)

type IsEmailRegisteredRequest struct {
	achieve_models_go.Request `json:"request"`
	Email                     string `json:"email"`
}

type IsEmailResponse struct {
	achieve_models_go.Request `json:"request"`
	IsEmailRegistered         bool   `json:"is_email_registered"`
	ResponseCode              int    `json:"response_code"`
	Error                     string `json:"error"`
}

type IsUserRegisteredRequest struct {
	achieve_models_go.Request `json:"request"`
	Username                  string `json:"username"`
}

type IsUserRegisteredResponse struct {
	achieve_models_go.Request `json:"request"`
	ResponseCode              int    `json:"response_code"`
	IsUsernameRegistered      bool   `json:"is_username_registered"`
	Error                     string `json:"error"`
}

type IsRegisteredRequest struct {
	achieve_models_go.Request `json:"request"`
	Email                     string `json:"email"`
	Username                  string `json:"username"`
}

type IsRegisteredResponse struct {
	achieve_models_go.Request `json:"request"`
	ResponseCode              int    `json:"response_code"`
	IsEmailRegistered         bool   `json:"is_email_registered"`
	IsUsernameRegistered      bool   `json:"is_username_registered"`
	Error                     string `json:"error"`
}

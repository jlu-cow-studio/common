package http_struct

type UserLoginState string

const (
	NotLogged    UserLoginState = "Not logged"
	LoggedIn     UserLoginState = "Logged in"
	InvalidToken UserLoginState = "Invalid tokne"
)

type UserTokenInfo struct {
	Token    string `json:"token"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     int    `json:"role:"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`

	LoginState UserLoginState `json:"login_state"`
}

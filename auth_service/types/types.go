package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Res       string `json:"res"`
	SessionId string `json:"session_id"`
	Msg       string `json:"msg"`
}

type UsersEntity struct {
	Id       int
	Email    string
	Password string
}

type AuthRequest struct {
	SessionId string `json:"session_id"`
}

type LogoutRequest struct {
	SessionId string `json:"session_id"`
}

type UsersSessionsEntity struct {
	Id        int
	UserId    int
	SessionId string
}

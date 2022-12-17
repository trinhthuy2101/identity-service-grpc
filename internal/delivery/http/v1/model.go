package v1

type AuthenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	ID       uint32 `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	FullName string `json:"fullName,omitempty"`
	Address  string `json:"address,omitempty"`
}

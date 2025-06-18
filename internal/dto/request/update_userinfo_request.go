package request

type UpdateUserInfoRequest struct {
	Uuid      string `json:"uuid"`
	Nickname  string `json:"nickname"`
	Signature string `json:"signature"`
	Avatar    string `json:"avatar"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
}

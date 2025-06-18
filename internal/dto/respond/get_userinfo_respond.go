package respond

type GetUserInfoRespond struct {
	Uuid      string `json:"uuid"`
	Nickname  string `json:"nickname"`
	Telephone string `json:"telephone"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	Gender    int8   `json:"gender"`
	Birthday  string `json:"birthday"`
	Status    int8   `json:"status"`
	IsAdmin   int8   `json:"is_admin"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

package respond

type RegisterRespond struct {
	Uuid      string `json:"uuid"`
	Avatar    string `json:"avatar"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Gender    int8   `json:"gender"`
	IsAdmin   int8   `json:"is_admin"`
	Nickname  string `json:"nickname"`
	Signature string `json:"signature"`
	Status    int8   `json:"status"`
	Telephone string `json:"telephone"`
}

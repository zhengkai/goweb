package handle

type LoginJSON struct {
	Success bool   `json:"success"`
	Uid     int64  `json:"uid,omitempty"`
	Error   string `json:"error,omitempty"`
}

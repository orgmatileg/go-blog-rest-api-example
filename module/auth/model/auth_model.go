package model

type Auth struct {
	UserID       string `json:"user_id"`
	NamaLengkap  string `json:"nama_lengkap"`
	PhotoProfile string `json:"photo_profile"`
	Token        string `json:"token"`
}

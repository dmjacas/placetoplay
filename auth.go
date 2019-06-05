package placetopay

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

// Auth structure
type Auth struct {
	Login   string `json:"login"`
	Nonce   string `json:"nonce"`
	Seed    string `json:"seed"`
	TranKey string `json:"tranKey"`
}

/*
// AuthBody structure
type AuthBody struct {
	ConnectionType string `json:"connectionType"`
	Endpoint       string `json:"endpoint"`
	Login          string `json:"login"`
	TranKey        string `json:"tranKey"`
}*/

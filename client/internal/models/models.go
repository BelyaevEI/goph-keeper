package models

type (
	// Data sent by the user
	RegistrationData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	RespRegistrationData struct {
		Token  string `json:"token"`
		UserID uint32 `json:"userid"`
	}
)

var CharSet = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM0987654321"

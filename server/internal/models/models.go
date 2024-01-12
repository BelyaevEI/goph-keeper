package models

type (
	// Structure sent by the user
	RegistrationData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	// Responsing structure
	RespRegistrationData struct {
		Token  string `json:"token"`
		UserID uint32 `json:"userid"`
	}
)

var CharSet = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM0987654321"

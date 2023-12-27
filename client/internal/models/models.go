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

	// Structure for work with login/password data
	LRdata struct {
		UserID   uint32 `json:"userid"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Service  string `json:"service"`
		Note     string `json:"note"`
	}

	// Structure for work with texts data
	Textsdata struct {
		UserID  uint32 `json:"userid"`
		Text    string `json:"text"`
		Service string `json:"service"`
		Note    string `json:"note"`
	}

	// Structure for work with bin data
	Binarydata struct {
		UserID  uint32 `json:"userid"`
		Bin     string `json:"bin"`
		Service string `json:"service"`
		Note    string `json:"note"`
	}

	// Structure for work with bank data
	Bankdata struct {
		UserID   uint32 `json:"userid"`
		Fullname string `json:"fullname"`
		Number   string `json:"number"`
		Date     string `json:"date"`
		Cvc      int    `json:"cvc"`
		Bankname string `json:"bankname"`
		Note     string `json:"note"`
	}
)

var CharSet = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM0987654321"

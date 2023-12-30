package passwordsmodels

import (
	"bytes"
	"encoding/binary"
)

// Structure for work with login/password data
type LRdata struct {
	UserID   uint32 `json:"userid"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Service  string `json:"service"`
	Note     string `json:"note"`
}

// MarshalBinary implement the interface encoding.BinaryMarshaler.
func (l *LRdata) MarshalBinary() ([]byte, error) {

	// Create buffer for binary data
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, l.UserID)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, l.Login)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, l.Password)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, l.Service)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, l.Note)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implement the interface encoding.BinaryUnmarshaler.
func (l *LRdata) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &l.UserID)
	if err != nil {
		return err
	}

	login, err := readString(buf)
	if err != nil {
		return err
	}
	l.Login = login

	password, err := readString(buf)
	if err != nil {
		return err
	}
	l.Password = password

	service, err := readString(buf)
	if err != nil {
		return err
	}
	l.Service = service

	note, err := readString(buf)
	if err != nil {
		return err
	}
	l.Note = note

	return nil
}

// writeString searializing data to binary
func writeString(buf *bytes.Buffer, str string) error {

	err := binary.Write(buf, binary.LittleEndian, uint8(len(str)))
	if err != nil {
		return err
	}

	err = binary.Write(buf, binary.LittleEndian, []byte(str))
	if err != nil {
		return err
	}

	return nil
}

// readString desiarilizing data from binary
func readString(buf *bytes.Reader) (string, error) {

	var length uint8
	err := binary.Read(buf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	err = binary.Read(buf, binary.LittleEndian, &strBytes)
	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}

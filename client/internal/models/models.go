package models

import (
	"bytes"
	"encoding/binary"
)

type (
	RegistrationData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	RespRegistrationData struct {
		Token  string `json:"token"`
		UserID uint32 `json:"userid"`
	}

	Data struct {
		UserID   uint32 `json:"userid"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Text     string `json:"text"`
		Bin      string `json:"bin"`
		Fullname string `json:"fullname"`
		Number   string `json:"number"`
		Date     string `json:"date"`
		Cvc      int    `json:"cvc"`
		Service  string `json:"service"`
		Note     string `json:"note"`
	}
)

func (data *Data) MarshalBinary() ([]byte, error) {

	// Create buffer for binary data
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, data.UserID)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Login)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Password)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Text)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Bin)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Fullname)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Number)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Date)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(data.Cvc))
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Service)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, data.Note)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implement the interface encoding.BinaryUnmarshaler.
func (d *Data) UnmarshalBinary(data []byte) error {

	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &d.UserID)
	if err != nil {
		return err
	}

	login, err := readString(buf)
	if err != nil {
		return err
	}
	d.Login = login

	password, err := readString(buf)
	if err != nil {
		return err
	}
	d.Password = password

	text, err := readString(buf)
	if err != nil {
		return err
	}
	d.Text = text

	bin, err := readString(buf)
	if err != nil {
		return err
	}
	d.Bin = bin

	fullname, err := readString(buf)
	if err != nil {
		return err
	}
	d.Fullname = fullname

	number, err := readString(buf)
	if err != nil {
		return err
	}
	d.Number = number

	date, err := readString(buf)
	if err != nil {
		return err
	}
	d.Date = date

	err = binary.Read(buf, binary.LittleEndian, &d.Cvc)
	if err != nil {
		return err
	}

	service, err := readString(buf)
	if err != nil {
		return err
	}
	d.Service = service

	note, err := readString(buf)
	if err != nil {
		return err
	}
	d.Note = note

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

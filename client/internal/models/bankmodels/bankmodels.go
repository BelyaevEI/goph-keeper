package bankmodels

import (
	"bytes"
	"encoding/binary"
)

// Structure for work with bank data
type Bankdata struct {
	UserID   uint32 `json:"userid"`
	Fullname string `json:"fullname"`
	Number   string `json:"number"`
	Date     string `json:"date"`
	Cvc      int    `json:"cvc"`
	Bankname string `json:"bankname"`
	Note     string `json:"note"`
}

// MarshalBinary implement interface encoding.BinaryMarshaler.
func (b *Bankdata) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, b.UserID)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Fullname)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Number)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Date)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(b.Cvc))
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Bankname)
	if err != nil {
		return nil, err
	}

	err = writeString(buf, b.Note)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implement interface encoding.BinaryUnmarshaler.
func (b *Bankdata) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &b.UserID)
	if err != nil {
		return err
	}

	fullname, err := readString(buf)
	if err != nil {
		return err
	}
	b.Fullname = fullname

	number, err := readString(buf)
	if err != nil {
		return err
	}
	b.Number = number

	date, err := readString(buf)
	if err != nil {
		return err
	}
	b.Date = date

	err = binary.Read(buf, binary.LittleEndian, &b.Cvc)
	if err != nil {
		return err
	}

	bankname, err := readString(buf)
	if err != nil {
		return err
	}
	b.Bankname = bankname

	note, err := readString(buf)
	if err != nil {
		return err
	}
	b.Note = note

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

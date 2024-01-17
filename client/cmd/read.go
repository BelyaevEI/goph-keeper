package cmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read user data",
	Long:  `read user data(bank, text, login/password, binary)`,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			err  error
			data models.Data
		)

		// Getting flag for saving user data
		flag, _ := cmd.Flags().GetString("data")
		data.UserID = User.UserID

		// Need adding check input data!
		data.Service, _ = cmd.Flags().GetString("service")
		data, err = readData(flag, data)
		if err != nil {
			fmt.Println("reading data is failed: ", err)
		}

		fmt.Println("reading data is success: ", data)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	createCmd.Flags().String("data", "", "data reading option(capital letter)")

	createCmd.Flags().String("service", "", "reading user service")
}

func readData(flag string, data models.Data) (models.Data, error) {
	var responseData models.Data

	if len(flag) == 0 {
		return models.Data{}, errors.New("empty flag data")
	}

	client := &http.Client{}
	reqBody := new(bytes.Buffer)

	err := binary.Write(reqBody, binary.LittleEndian, &data)
	if err != nil {
		return models.Data{}, err
	}

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/data/read", reqBody)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Add(flag, flag)
	request.Header.Add("UserID", strconv.Itoa(int(User.UserID)))
	request.Header.Add("Token", User.Token)

	resp, err := client.Do(request)
	if err != nil {
		return models.Data{}, err
	}

	if resp.StatusCode != http.StatusOK {
		er := "geting response with status code: " + strconv.Itoa(resp.StatusCode)
		return models.Data{}, errors.New(er)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Data{}, err
	}

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &responseData); err != nil {
		return models.Data{}, err
	}
	return responseData, nil
}

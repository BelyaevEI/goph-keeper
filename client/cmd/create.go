package cmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/spf13/cobra"
)

// This command sends the data to be saved to the server.
// It is necessary to pass a flag that defines the format of the transmitted data.
var createCmd = &cobra.Command{
	Use:   "save",
	Short: "save user data",
	Long:  `save user data(bank, text, login/password, binary)`,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			err  error
			data models.Data
		)

		// Getting flag for saving user data
		flag, _ := cmd.Flags().GetString("data")
		data.UserID = User.UserID

		// Need adding check input data!
		switch flag {
		case "Login":

			data.Login, _ = cmd.Flags().GetString("login")
			data.Password, _ = cmd.Flags().GetString("password")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = createData(flag, data)

		case "Text":

			data.Text, _ = cmd.Flags().GetString("text")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = createData(flag, data)
		case "Bin":

			data.Bin, _ = cmd.Flags().GetString("bin")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = createData(flag, data)

		case "Bank":

			data.Fullname, _ = cmd.Flags().GetString("fullname")
			data.Date, _ = cmd.Flags().GetString("date")
			data.Cvc, _ = cmd.Flags().GetInt("cvc")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = createData(flag, data)

		}

		if err != nil {
			fmt.Println("saving data is failed: ", err)
		}

		fmt.Println("saving data is success")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("data", "", "data saving option(capital letter)")

	createCmd.Flags().String("login", "", "saving user login")
	createCmd.Flags().String("password", "", "saving user password")
	createCmd.Flags().String("text", "", "saving user text")
	createCmd.Flags().String("fullname", "", "saving user fullname")
	createCmd.Flags().String("date", "", "saving user date")
	createCmd.Flags().Int("cvc", 0, "saving user cvc")
	createCmd.Flags().String("service", "", "saving user service")
	createCmd.Flags().String("note", "", "saving user note")
}

func createData(flag string, data models.Data) error {

	if len(flag) == 0 {
		return errors.New("empty flag data")
	}

	client := &http.Client{}
	reqBody := new(bytes.Buffer)

	err := binary.Write(reqBody, binary.LittleEndian, &data)
	if err != nil {
		return err
	}

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/data/create", reqBody)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Add(flag, flag)
	request.Header.Add("UserID", strconv.Itoa(int(User.UserID)))
	request.Header.Add("Token", User.Token)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		er := "geting response with status code: " + strconv.Itoa(resp.StatusCode)
		return errors.New(er)
	}

	return nil
}

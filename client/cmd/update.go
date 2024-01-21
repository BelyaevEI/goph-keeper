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

// This command transmits the data for updating to the servers.
// It is necessary to pass a flag that defines the format of the transmitted data.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update user data",
	Long:  `update user data(bank, text, login/password, binary)`,
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

			err = updateData(flag, data)

		case "Text":

			data.Text, _ = cmd.Flags().GetString("text")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = updateData(flag, data)

		case "Bin":

			data.Bin, _ = cmd.Flags().GetString("bin")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = updateData(flag, data)

		case "Bank":

			data.Fullname, _ = cmd.Flags().GetString("fullname")
			data.Date, _ = cmd.Flags().GetString("date")
			data.Cvc, _ = cmd.Flags().GetInt("cvc")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = updateData(flag, data)

		}

		if err != nil {
			fmt.Println("updating data is failed: ", err)
		}

		fmt.Println("updating data is success")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	createCmd.Flags().String("data", "", "data updating option(capital letter)")

	createCmd.Flags().String("login", "", "updating user login")
	createCmd.Flags().String("password", "", "updating user password")
	createCmd.Flags().String("text", "", "updating user text")
	createCmd.Flags().String("fullname", "", "updating user fullname")
	createCmd.Flags().String("date", "", "updating user date")
	createCmd.Flags().Int("cvc", 0, "updating user cvc")
	createCmd.Flags().String("service", "", "updating user service")
	createCmd.Flags().String("note", "", "updating user note")
}

func updateData(flag string, data models.Data) error {

	if len(flag) == 0 {
		return errors.New("empty flag data")
	}

	client := &http.Client{}
	reqBody := new(bytes.Buffer)

	err := binary.Write(reqBody, binary.LittleEndian, &data)
	if err != nil {
		return err
	}

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/data/up", reqBody)
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

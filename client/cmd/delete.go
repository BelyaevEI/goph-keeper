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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete user data",
	Long:  `delete user data(bank, text, login/password, binary)`,
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

			err = deleteData(flag, data)

		case "Text":

			data.Text, _ = cmd.Flags().GetString("text")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = deleteData(flag, data)
		case "Bin":

			data.Bin, _ = cmd.Flags().GetString("bin")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = deleteData(flag, data)

		case "Bank":

			data.Fullname, _ = cmd.Flags().GetString("fullname")
			data.Date, _ = cmd.Flags().GetString("date")
			data.Cvc, _ = cmd.Flags().GetInt("cvc")
			data.Service, _ = cmd.Flags().GetString("service")
			data.Note, _ = cmd.Flags().GetString("note")

			err = deleteData(flag, data)

		}

		if err != nil {
			fmt.Println("delete data is failed: ", err)
		}

		fmt.Println("delete data is success")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("data", "", "data deleting option(capital letter)")

	createCmd.Flags().String("login", "", "deletinguser login")
	createCmd.Flags().String("password", "", "deleting user password")
	createCmd.Flags().String("text", "", "deleting user text")
	createCmd.Flags().String("fullname", "", "deleting user fullname")
	createCmd.Flags().String("date", "", "deleting user date")
	createCmd.Flags().Int("cvc", 0, "deleting user cvc")
	createCmd.Flags().String("service", "", "deleting user service")
	createCmd.Flags().String("note", "", "deleting user note")
}

func deleteData(flag string, data models.Data) error {

	if len(flag) == 0 {
		return errors.New("empty flag data")
	}

	client := &http.Client{}
	reqBody := new(bytes.Buffer)

	err := binary.Write(reqBody, binary.LittleEndian, &data)
	if err != nil {
		return err
	}

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/data/del", reqBody)
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

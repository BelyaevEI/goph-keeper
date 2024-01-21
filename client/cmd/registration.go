package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/spf13/cobra"
)

// Registration of a new user is carried out by transferring
// a username and password to the command using flags
var regCmd = &cobra.Command{
	Use:   "registration",
	Short: "registration new user",
	Long:  `registration new user`,
	Run: func(cmd *cobra.Command, args []string) {

		// Getting login and password for registation new user
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		err := registationNewUser(login, password)
		if err != nil {
			fmt.Println("registration is failed: ", err)
		}

		fmt.Println("registration new user is success")
	},
}

func init() {
	rootCmd.AddCommand(regCmd)
	regCmd.Flags().String("login", "", "login new user")
	regCmd.Flags().String("password", "", "password new user")
}

func registationNewUser(login, password string) error {

	var (
		reginfo     models.RegistrationData
		respRegData models.RespRegistrationData
	)

	if len(login) == 0 || len(password) == 0 {
		return errors.New("empty login/password")
	}

	client := &http.Client{}

	reginfo.Login, reginfo.Password = login, password
	body, _ := json.Marshal(reginfo)
	requestBody := strings.NewReader(string(body))

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/user/registration", requestBody)
	request.Header.Set("Content-Type", "json/application")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		er := "geting response with status code: " + strconv.Itoa(resp.StatusCode)
		return errors.New(er)
	}

	respBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	json.Unmarshal(respBody, &respRegData)

	User.UserID, User.Token = respRegData.UserID, respRegData.Token
	return nil
}

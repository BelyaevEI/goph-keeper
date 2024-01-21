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

// This command has two flags, login and password.
// These flags are used to send data to the function for user authorization to the server.
var authCmd = &cobra.Command{
	Use:   "authorization",
	Short: "authorization user in application",
	Long:  `authorization user in application`,
	Run: func(cmd *cobra.Command, args []string) {

		// Getting login and password for registation new user
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		err := authorizationUser(login, password)
		if err != nil {
			fmt.Println("authorization is failed: ", err)
		}

		fmt.Println("authorization user is success")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	regCmd.Flags().String("login", "", "user login")
	regCmd.Flags().String("password", "", "user password")
}

func authorizationUser(login, password string) error {

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

	request, _ := http.NewRequest(http.MethodPost, UrlService+"/api/user/authorization", requestBody)
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

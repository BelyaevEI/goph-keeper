package cmd

import (
	"fmt"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/spf13/cobra"
)

var (
	Version    string
	BuildTime  string
	UrlService string = "localhost:8080"
	User       models.RespRegistrationData
)

// getinfoCmd represents the getinfo command
var getinfoCmd = &cobra.Command{
	Use:   "getinfo",
	Short: "information about the application build",
	Long:  `information about the application build`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getInfo())
	},
}

func init() {
	rootCmd.AddCommand(getinfoCmd)
}

func getInfo() string {
	return fmt.Sprintf("Build version: %s\nBuild date: %s\n", Version, BuildTime)
}

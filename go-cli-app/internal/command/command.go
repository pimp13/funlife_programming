package command

import (
	"errors"

	"github.com/spf13/cobra"
)

// Errors
var (
	ENV_API_KEY_NOT_SET = errors.New("CURRENCY_API_KEY enviroment variable is not set")
)

var apiKey string

var rootCmd = &cobra.Command{
	Use:   "currency",
	Short: "Currenct cmd cli apps",
	Long:  "This is my currency cmd CLI app project",
}

func Execute(key string) error {
	if key == "" {
		return ENV_API_KEY_NOT_SET
	}
	apiKey = key
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

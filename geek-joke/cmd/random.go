/*
Copyright © 2026 DevPouyaGh <pimp.puma.13@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random geek joke",
	Long:  `Get a random geek joke`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeRandomCmd(cmd, args)
	},
}

// args: cmd *cobra.Command, args []string
func executeRandomCmd(_ *cobra.Command, _ []string) error {
	// fire off a request to get a random geek joke
	url := "https://geek-jokes.sameerkumar.website/api?format=json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var jokeResponse JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&jokeResponse); err != nil {
		return err
	}

	fmt.Println("🤓🤡 This Random Geek Joke")
	fmt.Printf("%s\n", jokeResponse.Joke)

	return nil
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

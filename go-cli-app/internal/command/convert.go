package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Response struct {
	Success bool `json:"success"`
	Query   struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount float64
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
}

var (
	from   string
	to     string
	amount float64
)

// Errors
var (
	NO_SUCCESS_FROM_REQUEST = errors.New("failed to get response success is false")
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert from one currency to another",
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeConvertCmd(cmd, args)
	},
}

func executeConvertCmd(cmd *cobra.Command, args []string) error {

	// fire off a request to convert the currency
	url := fmt.Sprintf(
		"https://api.exchangerate.host/convert?access_key=%s&from=%s&to=%s&amount=%f",
		apiKey,
		from,
		to,
		amount,
	)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return NO_SUCCESS_FROM_REQUEST
	}

	fmt.Printf("%.2f %s is equal to %.2f %s\n",
		response.Query.Amount, response.Query.From, response.Result, response.Query.To)

	return nil
}

func init() {
	convertCmd.Flags().StringVarP(&from, "from", "f", "USD", "The currency to convert from")
	convertCmd.Flags().StringVarP(&to, "to", "t", "IRR", "The currency to convert to")
	convertCmd.Flags().Float64VarP(&amount, "amount", "a", 1, "The amount to convert")
	rootCmd.AddCommand(convertCmd)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this command to list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		url := URL + "/task"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		var prettyJSON map[string]interface{}
		err = json.Unmarshal(body, &prettyJSON)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		// Pretty-print JSON
		formattedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON:", err)
			return
		}

		fmt.Println("API Response:")
		fmt.Println(string(formattedJSON))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

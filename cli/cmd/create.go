package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var bodyFile string

var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new Task in the Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		url := URL + "/task"

		// Read JSON file
		fileContent, err := os.ReadFile(bodyFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		// Validate JSON format
		var requestBody map[string]interface{}
		if err := json.Unmarshal(fileContent, &requestBody); err != nil {
			fmt.Println("Invalid JSON format in file:", err)
			return
		}

		// Send HTTP request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(fileContent))
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// Read API response
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
	createCmd.Flags().StringVarP(&bodyFile, "bodyfile", "f", "", "Path to JSON file containing request body")
	createCmd.MarkFlagRequired("bodyfile")
	rootCmd.AddCommand(createCmd)
}

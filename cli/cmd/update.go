package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var updatetaskID string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task by providing its task ID",
	Run: func(cmd *cobra.Command, args []string) {
		if updatetaskID == "" {
			fmt.Println("Error: --taskid flag is required")
			return
		}

		url := fmt.Sprintf("%s/task?taskId=%s", URL, updatetaskID)

		req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
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

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Task deleted successfully!")
		} else {
			fmt.Printf("Failed to delete task (Status: %d)\n", resp.StatusCode)
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
	updateCmd.Flags().StringVarP(&updatetaskID, "taskid", "", "", "Task ID of the task to delete")
	rootCmd.AddCommand(updateCmd)
}

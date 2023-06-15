package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	uhttp "github.com/bestchains/bc-cli/pkg/utils/http"
	"github.com/spf13/cobra"
)

const (
	chatApiPath = "%s/chat"
)

func Chat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat [args]",
		Short: "Chat with assistant",
		RunE: func(cmd *cobra.Command, args []string) error {
			question, err := cmd.Flags().GetString("question")
			if err != nil {
				return err
			}
			postValue := url.Values{}
			postValue.Add("question", question)
			// call assistant api to generate unit tests
			resp, err := uhttp.Do(
				fmt.Sprintf(chatApiPath, server),
				http.MethodPost,
				map[string]string{
					"Content-Type": "application/x-www-form-urlencoded",
				},
				[]byte(postValue.Encode()),
			)
			if err != nil {
				return err
			}

			var vresp = make(map[string]any)
			err = json.Unmarshal(resp, &vresp)
			if err != nil {
				return err
			}
			fmt.Println("Answer from assistant:")
			fmt.Println(vresp["answer"])

			return nil
		},
	}

	cmd.Flags().StringP("question", "q", "", "question to ask assistant")
	cmd.MarkFlagRequired("question")

	return cmd
}

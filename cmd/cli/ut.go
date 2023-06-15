package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	uhttp "github.com/bestchains/bc-cli/pkg/utils/http"
	"github.com/bjwswang/assistant/pkg/assistant"
	"github.com/spf13/cobra"
)

const (
	unitTestApiPath = "%s/ut"
)

func GenUnitTests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ut [args]",
		Short: "Generate unit tests for source code",
		Long:  `Generate unit tests for source code with the help of assistant.You must provide the file path of the source code with flag --file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// read file from flag --file
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}
			// read the content from local file whose path is `file`
			fileContent, err := os.ReadFile(file)
			if err != nil {
				return err
			}
			// encode fiele content to base64
			fileContentBase64 := base64.StdEncoding.EncodeToString(fileContent)
			postValue := url.Values{}
			postValue.Add("code", fileContentBase64)
			// call assistant api to generate unit tests
			resp, err := uhttp.Do(
				fmt.Sprintf(unitTestApiPath, server),
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
			fmt.Println(vresp[assistant.ChainUnitTestOutputKeys[0]])

			return nil
		},
	}

	cmd.Flags().String("file", "", "file to generate unit tests")

	cmd.MarkFlagRequired("file")

	return cmd
}

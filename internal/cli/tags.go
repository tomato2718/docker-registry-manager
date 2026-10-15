package cli

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/tomato2718/docker-registry-manager/internal/httpclient"
)

var tagsCmd = &cobra.Command{
	Use:   "tags <registry url> <name>",
	Short: "Show tags of an images",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := httpclient.Tags(args[0], args[1])
		sort.Strings(tags)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, tag := range tags {
			fmt.Println(tag)
		}

	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tomato2718/docker-registry-manager/internal/httpclient"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <registry url> <image:tag>",
	Short: "Delete a tagged image",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name, tag, found := strings.Cut(args[1], ":")
		if !found || name == "" || tag == "" {
			fmt.Println("Image must be in the form <image:tag>.")
			os.Exit(1)
		}
		err := httpclient.Delete(args[0], name, tag)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Deleted " + args[1])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

package cmd

import (
	"context"
	"fmt"
	"github.com/bavix/ghmd/internal/app"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ghmd",
	Run: func(cmd *cobra.Command, args []string) {
		input, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return
		}

		replacer := app.New(app.WithUserReplacer())

		fmt.Println(string(replacer.Replace(input)))
	},
}

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

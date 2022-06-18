package main

import (
	"context"

	"github.com/leonardosaid/stc/accounts/cmd/commands"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rootCmd := &cobra.Command{Use: "account"}
	rootCmd.AddCommand(commands.NewServerCommand(ctx))
	_ = rootCmd.Execute()
}

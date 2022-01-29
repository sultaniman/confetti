package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sultaniman/confetti/util"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test",
	Long:  "Test",
	RunE: func(cmd *cobra.Command, args []string) error {
		password := "MySecretPassword"
		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			return err
		}

		fmt.Println("HASHED PASSWORD:", hashedPassword)

		// Comparing the password with the hash
		return util.CheckPassword(hashedPassword, password)
	},
}

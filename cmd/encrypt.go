/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-cli/internal/crypto"
	"todo-cli/internal/utils"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use: "encrypt [key] [data]",
	Short: "Encrypt",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := []byte(args[0])
		key = utils.AdjustKey(key)

		data := []byte(args[1])

		encrypted, err := crypto.Encrypt(data, key)
		if err != nil {
			fmt.Println("Error during encryption:", err)
			return
		}
		fmt.Println("Encrypted text:", encrypted)

		decrypted, err := crypto.Decrypt(encrypted, key)
		if err != nil {
			fmt.Println("Error during decryption:", err)
			return
		}
		fmt.Println("Decrypted text:", string(decrypted))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

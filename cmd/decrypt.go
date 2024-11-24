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

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use: "decrypt key <key_value> data <data_value>",
	Short: "Decrypt",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := []byte(args[0])
		key = utils.AdjustKey(key)

		data := args[1]

		decrypted, err := crypto.Decrypt(data, key)
		if err != nil {
			fmt.Println("Error during decryptiom", err)
			return
		}
		fmt.Println("Decrypted text:", string(decrypted))
	},
}
func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

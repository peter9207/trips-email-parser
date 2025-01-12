/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/peter9207/trips-email-parser/email"
	"github.com/spf13/cobra"
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("email called")
	},
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "generate html from the email",
	Long:  `A longer description that spans multiple lines and likely contains examples to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			return
		}

		filename := args[0]

		reader, err := os.Open(filename)
		defer func() {
			reader.Close()
		}()

		if err != nil {
			panic(err)
		}

		e, err := email.Parse(reader) // returns Email struct and error
		if err != nil {
			// handle error
			panic(err)
		}
		fmt.Println(e.HTMLBody)
		output, err := os.Create("output.html")
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := output.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = output.Write([]byte(e.HTMLBody))
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
	emailCmd.AddCommand(htmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

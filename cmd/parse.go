/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/peter9207/trips-email-parser/email"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called")

		if len(args) < 1 {
			cmd.Help()
			return
		}

		filename := args[0]

		reader, err := os.Open(filename)

		if err != nil {
			panic(err)
		}

		e, err := email.Parse(reader) // returns Email struct and error
		if err != nil {
			// handle error
			panic(err)
		}

		fmt.Println(e.String())
		// fmt.Println(email.Subject)
		// fmt.Println(email.From)
		// fmt.Println(email.To)
		// fmt.Println(email.HTMLBody)

	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// type Email struct {
// 	Header mail.Header

// 	Subject    string
// 	Sender     *mail.Address
// 	From       []*mail.Address
// 	ReplyTo    []*mail.Address
// 	To         []*mail.Address
// 	Cc         []*mail.Address
// 	Bcc        []*mail.Address
// 	Date       time.Time
// 	MessageID  string
// 	InReplyTo  []string
// 	References []string

// 	ResentFrom      []*mail.Address
// 	ResentSender    *mail.Address
// 	ResentTo        []*mail.Address
// 	ResentDate      time.Time
// 	ResentCc        []*mail.Address
// 	ResentBcc       []*mail.Address
// 	ResentMessageID string

// 	ContentType string
// 	Content io.Reader

// 	HTMLBody string
// 	TextBody string

// 	Attachments   []Attachment
// 	EmbeddedFiles []EmbeddedFile
// }

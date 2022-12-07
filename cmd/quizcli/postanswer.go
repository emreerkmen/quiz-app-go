package main

import (
	"fmt"
	"quiz-app/quiz-api/pkg/quizcli"
	"github.com/spf13/cobra"
)

var postanswerCmd = &cobra.Command{
    Use:   "post-answer",
    Aliases: []string{"pa"},
    Short:  "Make a quiz",
    Args:  cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {

		quizID:=args[0]
		userID:=args[1]
		selectedAnswers:=args[2]

        res := quizcli.PostAnswer(quizID,userID,selectedAnswers,beautify)
        fmt.Println(res)
    },
}

func init() {
	postanswerCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Print result as beautiful json")
    rootCmd.AddCommand(postanswerCmd)
}
package main

import (
	"fmt"
	"os"

	"github.com/lucsky/cuid"
	"github.com/spf13/cobra"
)

func main() {
	// コマンドを定義
	var rootCmd = &cobra.Command{
		Use:   "cuidgen",
		Short: "generate a collision-resistant identifier",
		Long:  "generate a collision-resistant identifier",
		Run: func(cmd *cobra.Command, args []string) {
			// cuidを生成して表示
			fmt.Println(cuid.New())
		},
	}

	// コマンドを実行
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

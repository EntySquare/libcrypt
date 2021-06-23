package command

import (
	"EntySquare/libcrypt"
	"bufio"
	"encoding/hex"
	"github.com/spf13/cobra"
)

const (
	CRYPT_IV = "yueliyangzi"
)

func Run(appName string) {
	var rootCmd = &cobra.Command{
		Use: appName,
	}
	var entCmd = &cobra.Command{
		Use:   "encode [message]",
		Short: "encode",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			cmd.Println("please input key for content")
			input := bufio.NewReader(cmd.InOrStdin())
			bytes, err := input.ReadBytes('\n')
			if err != nil {
				panic("error")
			}
			originData := []byte(args[0])
			key := bytes
			data := libcrypt.AesEncrypt(key, []byte(CRYPT_IV), originData)
			strData := hex.EncodeToString(data)
			println(strData)
		},
	}
	var deCmd = &cobra.Command{
		Use:   "decode [message]",
		Short: "decode",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			cmd.Println("please input key for content")
			input := bufio.NewReader(cmd.InOrStdin())
			bytes, err := input.ReadBytes('\n')
			if err != nil {
				panic("error")
			}

			strData := args[0]
			key := bytes
			encrData, _ := hex.DecodeString(strData)
			data := libcrypt.AesDecrypt(key, []byte(CRYPT_IV), encrData)
			println(string(data))
		},
	}
	//var webCmd = &cobra.Command{
	//	Use:   "web",
	//	Short: "web mining data Chia..."}
	//var webChiaCmd = &cobra.Command{
	//	Use:   "chia",
	//	Short: "web mining data Chia..."}
	//webCmd.AddCommand(webChiaCmd)
	//webChiaCmd.AddCommand(ChiaWebService)
	//
	//// check duplicated data or insufficient sized data
	//var storageCheckCmd = &cobra.Command{
	//	Use:   "check",
	//	Short: "check duplicated data or insufficient sized data"}
	//// check cache
	//storageCheckCmd.AddCommand(checkCacheCmd)
	//// check plot
	//storageCheckCmd.AddCommand(checkPlotCmd)
	////fmv cmd
	////fmvCmd.AddCommand(fmvStartCmd)
	//fmvCmd.AddCommand(fmvStopCmd)

	rootCmd.AddCommand(entCmd)

	rootCmd.AddCommand(deCmd)
	//rootCmd.AddCommand(fmvCmd)
	rootCmd.Execute()
}

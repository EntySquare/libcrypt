package command

import (
	"EntySquare/libcrypt"
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
		Use:   "encode",
		Short: "encode",
		Run: func(cmd *cobra.Command, args []string) {
			println(args[0])
			println(args[1])
			originData := []byte(args[1])
			key := []byte(args[0])
			data := libcrypt.AesEncrypt(key, []byte(CRYPT_IV), originData)
			strData := hex.EncodeToString(data)
			println(strData)
		},
	}
	var deCmd = &cobra.Command{
		Use:   "decode",
		Short: "decode",
		Run: func(cmd *cobra.Command, args []string) {
			strData := args[1]
			key := []byte(args[0])
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

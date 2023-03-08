package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"reddit-sync/src/reddit"
)

// downloadAssetCmd represents the downloadAsset command
var downloadAssetCmd = &cobra.Command{
	Use:   "downloadAsset",
	Short: "Download an asset from its source URL to a designated local filepath.",
	Run: func(cmd *cobra.Command, args []string) {
		targetPath, err := cmd.Flags().GetString("target-path")
		if err != nil {
			log.Fatalln(err)
		}

		sourceUrl, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalln(err)
		}

		n, err := reddit.DownloadAsset(sourceUrl, targetPath)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(n)
	},
}

func init() {
	rootCmd.AddCommand(downloadAssetCmd)

	downloadAssetCmd.Flags().StringP("url", "u", "", "Source URL to download asset from.")
	downloadAssetCmd.MarkFlagRequired("url")

	downloadAssetCmd.Flags().StringP("target-path", "t", "", "Target path to download asset to.")
	downloadAssetCmd.MarkFlagRequired("target-path")
}

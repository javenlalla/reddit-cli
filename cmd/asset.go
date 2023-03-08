package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"reddit-sync/src/reddit"
)

var audioUrl string
var audioFilename string

// assetCmd represents the asset command
var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Sync a Reddit media asset (Post image or thumbnail, Subreddit banner, etc.) by its source URL.",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatal(err)
		}

		reddit.SyncAsset(url, audioUrl, audioFilename)
	},
}

func init() {
	syncCmd.AddCommand(assetCmd)

	assetCmd.Flags().Int("port", 1138, "Port to run Application server on")
	viper.BindPFlag("port", assetCmd.Flags().Lookup("port"))

	assetCmd.Flags().StringP("url", "u", "", "Source URL for target media asset.")
	assetCmd.MarkFlagRequired("url")

	assetCmd.Flags().StringVar(&audioUrl, "audio-url", "", "Source URL for media's audio file, if any.")
	assetCmd.Flags().StringVar(&audioFilename, "audio-filename", "", "Destination filename for audio file.")
}

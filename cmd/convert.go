package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yterajima/go-sitemap"
)

var (
sitemapURL string
)


var ConvertCmd = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"c", "convert"},
	Short:   "Convert sitemap to text file in vegeta format.",
	Long:    "Convert sitemap to text file in vegeta format.",
	Run: func(cmd *cobra.Command, args []string) {

	smap, err := sitemap.Get("http://www.e2esound.com/sitemap.xml", nil)
	if err != nil {
		fmt.Println(err)
	}

	// Print URL in sitemap.xml
	for _, URL := range smap.URL {
		fmt.Println(URL.Loc)
	}
}

func init() {
	// https://t.me/octaveqa_bot
	ConvertCmd.Flags().StringVarP(&sitemapURL, "url", "u", "", "Sitemap URL.")
	RootCmd.AddCommand(ChatbotCmd)
}

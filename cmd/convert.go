package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yterajima/go-sitemap"
)

var (
	sitemapURL string
	outputDir  string
)

var ConvertCmd = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"c", "convert"},
	Short:   "Convert sitemap to text file in vegeta format.",
	Long:    "Convert sitemap to text file in vegeta format.",
	Run: func(cmd *cobra.Command, args []string) {
		if sitemapURL == "" {
			log.Fatal("You must provide a sitemap url parameter.")
		}
		smap, err := sitemap.Get(sitemapURL, nil)
		if err != nil {
			fmt.Println(err)
		}
		// Print URL in sitemap.xml
		for _, URL := range smap.URL {
			fmt.Println(URL.Loc)
		}
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&outputDir, "out", "o", ".", "Output directory.")
	ConvertCmd.Flags().StringVarP(&sitemapURL, "url", "u", "", "Sitemap URL.")
	RootCmd.AddCommand(ConvertCmd)
}

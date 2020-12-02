package cmd

import (
   	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	log        = logrus.New()
)

var options struct {
	verbose     bool
	debug       bool
	logLevel    string
	generateDoc bool
}

// RootCmd is the root command for ovh-qa
var RootCmd = &cobra.Command{
	Use:   "sitemap2vegeta",
	Short: "Convert a sitemap to a vegeta target urls list",
	Long:  `Convert a sitemap to a vegeta target urls list.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if options.generateDoc {
		err := doc.GenMarkdownTree(RootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
		// todo: generate changelog and todo from code
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.BoolVarP(&options.generateDoc, "generate-doc", "g", false, "generate documentation.")
	flags.BoolVarP(&options.debug, "debug", "d", false, "debug mode.")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output.")
}

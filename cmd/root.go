package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "App",
		Short: "App is Rest API",
		Long:  `App is a Rest API created with golang.`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.env)")
	rootCmd.PersistentFlags().StringP("server-port", "p", "7000", "Server port listening")
	rootCmd.PersistentFlags().BoolP("log-json", "j", false, "Log format json (default is false)")
	rootCmd.PersistentFlags().StringP("log-level", "l", "error", "Log level (default is error)")

	viper.BindPFlag("REST_API_SERVER_PORT", rootCmd.PersistentFlags().Lookup("server-port"))
	viper.BindPFlag("REST_API_LOG_JSON", rootCmd.PersistentFlags().Lookup("log-json"))
	viper.BindPFlag("REST_API_LOG_LEVEL", rootCmd.PersistentFlags().Lookup("log-level"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		ex, err := os.Executable()
		cobra.CheckErr(err)
		exePath := filepath.Dir(ex)
		viper.AddConfigPath(exePath)
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	} else {
		log.Errorf("Error: %v", err)
	}

}

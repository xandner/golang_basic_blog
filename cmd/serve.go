package cmd

import (
	"blog/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving app",
	Long:  "app is running",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	configs := configSet()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"app-name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func configSet() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic("Config not found")
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("config error: ,%v", err)
	}
	return configs
}

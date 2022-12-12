package main

import (
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/web3ssc/ss-go-common/log"
	"github.com/web3ssc/ss-go-we-chat-gpt/config"
	"github.com/web3ssc/ss-go-we-chat-gpt/router"
	"github.com/web3ssc/ss-go-we-chat-gpt/we_chat"
	"net/http"
	"os"
	"time"
)

var (
	cfg = pflag.StringP("profile", "c", "", "")
	nickname = pflag.StringP("nickname", "n", "", "")
	token    = pflag.StringP("token", "t", "", "")
)

func main() {
	pflag.Parse()

	config.Init(*cfg)

	config.SetValue(*nickname, *token)

	g := gin.New()

	router.Load(
		g,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Errorf(err, "the router has no response")
			os.Exit(0)
		}
		log.Info("the router has been deployed successfully")
		weChat.StartWeChat()
	}()

	log.Infof("the port: %s", viper.GetString("port"))
	log.Info(http.ListenAndServe(viper.GetString("port"), g).Error())

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		url := viper.GetString("url") + viper.GetString("port") + "/api/common/checkHealth"
		log.Infof("url: %s", url)
		resp, err := client.Get(url)

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router error")
}

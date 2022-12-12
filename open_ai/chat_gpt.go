package open_ai

import (
	"fmt"
	"github.com/goccy/go-json"
	baseUtil "github.com/web3ssc/ss-go-common/base_util"
	"github.com/web3ssc/ss-go-common/log"
	"github.com/web3ssc/ss-go-we-chat-gpt/common"
	"github.com/web3ssc/ss-go-we-chat-gpt/config"
	"strings"
)

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int                      `json:"created"`
	Model   string                   `json:"model"`
	Choices []map[string]interface{} `json:"choices"`
	Usage   map[string]interface{}   `json:"usage"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}


func Completions(msg string) (*string, error) {
	requestBody := ChatGPTRequestBody{
		Model:            "text-davinci-003",
		Prompt:           msg,
		MaxTokens:        4000,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		log.Errorf(err, "json.Marshal error")
		return nil, err
	}
	log.Infof("request openai json string : %v", string(requestData))

	url := "https://api.openai.com/v1/completions"
	auth :=  fmt.Sprintf("Bearer %s", config.Token)
	callResp, err := common.HttpClient(url, baseUtil.POST, requestData, true, auth)
	if err != nil {
		log.Errorf(err, "baseUtil.HttpClientSub error")
		return nil, err
	}
	log.Infof("callResp: %v", string(callResp))

	gptResponseBody := &ChatGPTResponseBody{}
	if err = json.Unmarshal(callResp, gptResponseBody); err != nil {
		log.Errorf(err, "json.Unmarshal error")
		return nil, err
	}
	var reply string
	if len(gptResponseBody.Choices) > 0 {
		for _, v := range gptResponseBody.Choices {
			reply = v["text"].(string)
			break
		}
	}
	result := strings.TrimSpace(reply)
	return &result, nil
}

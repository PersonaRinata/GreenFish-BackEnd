package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
)

type Menu struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	SafeMode string    `json:"safe_mode"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Resp struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   Usage     `json:"usage"`
	Choices []Choices `json:"choices"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
	PreTokenCount    int `json:"pre_token_count"`
	PreTotal         int `json:"pre_total"`
	AdjustTotal      int `json:"adjust_total"`
	FinalTotal       int `json:"final_total"`
}
type Choices struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

func GetGptMessage(msg ...string) string {
	m := Menu{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				"user",
				"我要你一直扮演专业的医生解答我的一些医疗问题，对于不确定的问题你可以选择不回答，以下是我的一些身体信息",
			},
		},
		SafeMode: "false",
	}
	for _, v := range msg {
		m.Messages = append(m.Messages, Message{"user", v})
	}
	m1, _ := json.Marshal(m)
	fmt.Println(string(m1))
	client1 := req.C()
	r, err := client1.R().SetBody(m).SetHeader("Content-Type", "application/json").SetHeader("Authorization", "Bearer "+"fk210314-SgSKCEJRaFf5wzjUJt4h6mGwJvjhz6Do").
		Post("https://oa.api2d.net/v1/chat/completions")
	fmt.Println(r)
	s, _ := r.ToString()
	var resp Resp
	err = json.Unmarshal([]byte(s), &resp)
	if err != nil {
		return "对于这个问题我不能给出准确的回答"
	}
	return resp.Choices[0].Message.Content
}

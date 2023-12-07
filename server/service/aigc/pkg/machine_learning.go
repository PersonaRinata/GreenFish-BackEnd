package pkg

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

type answser struct {
	Result string `json:"result"`
}

func DoctorRecommend(content string) (error, string) {
	url := "http://127.0.0.1:5000/predict"
	m := map[string]string{
		"input": content,
	}
	client := req.C()
	r, err := client.R().SetFormData(m).SetHeader("Content-Type", "multipart/form-data").
		Post(url)
	if err != nil {
		fmt.Println(err)
	}
	var a answser
	sonic.Unmarshal(r.Bytes(), &a)
	return nil, a.Result
}

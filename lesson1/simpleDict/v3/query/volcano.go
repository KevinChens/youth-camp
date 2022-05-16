package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// tips:结构体字段首字母必须大写，否则无法解析
// 火山翻译
type DictRequest struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type DictResponse struct {
	Words []struct {
		Source  int    `json:"source"`
		Text    string `json:"text"`
		PosList []struct {
			Type      int `json:"type"`
			Phonetics []struct {
				Type int    `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text     string `json:"text"`
				Examples []struct {
					Type      int `json:"type"`
					Sentences []struct {
						Text      string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases  []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

func WithVolcano(word string) {
	client := &http.Client{}
	// 构造请求头
	request := DictRequest{Text: word, Language: "en"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	//var data = strings.NewReader(`{"text":"main\n","language":"en"}`)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDGi0kQrSWyYDe9WX7j8&_signature=_02B4Z6wo00001shAlegAAIDCSEJvq55XUt7IQJFAANB2Y8xWaqs2fop8vk02uXF2toi-9P7gEvogkw07q4BNVcV6g-pgohptz.OauDlz.71a.i9FVpwpMlbCTle5KEvLVD5W0U3KbQWWGZU-75", data)
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16519321977718926; ttcid=9ae700ffe7c34d1a9fb723c486d7226839; s_v_web_id=verify_174080b263720677d6065cb149f52d84; _tea_utm_cache_2018=undefined; i18next=zh-CN; tt_scid=qiuJMaps6jUIO3eMhWvW8hgzSSNUTcwFTP5.8wRHYUMCHfFCQYpgf-V1GPHZKzjZ52a7; _dd_s=logs=1&id=083f71d0-1ff6-463b-95bb-0b57843ebb36&created=1652157474637&expire=1652158448710")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=en&target_language=zh&text=main")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 读取响应
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 防御式编程
	if resp.StatusCode != 200 {
		log.Fatal("error StatusCode:", resp.StatusCode, "body:", string(bodyText))
	}

	// 返回值处理
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*********火山翻译***********")
	fmt.Println(word, "UK:", dictResponse.Words[0].PosList[0].Phonetics[0].Text, "US:",
		dictResponse.Words[0].PosList[0].Phonetics[1].Text)
	for _, item := range dictResponse.Words[0].PosList[0].Explanations {
		fmt.Println(item.Text)
	}
}

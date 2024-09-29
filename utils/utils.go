package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	apikey = os.Getenv("API_KEY")
	url    = "https://api.apilayer.com/exchangerates_data/"
)

func GetSymbols() string {
	symbols := requestSymbols()
	return symbols
}

func GetConvertedValue(to, from, amount string) float64 {
	data, err := requestData(to, from, amount)
	if err != nil {
		return 0
	}

	result := data["result"].(float64)
	return result
}

func GetRate(to, from string) float64 {
	data, err := requestData(to, from, "1")
	if err != nil {
		return 0
	}

	info, ok := data["info"].(map[string]interface{})
	if !ok {
		return 0
	}
	result, ok := info["rate"].(float64)
	if !ok {
		return 0
	}
	return result
}

func requestSymbols() string {
	reqUrl := fmt.Sprintf("%ssymbols", url)
	req, err := http.NewRequest("GET", reqUrl, nil)
	client := &http.Client{}
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	req.Header.Set("apikey", apikey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(body)
}

func requestData(to, from, amount string) (map[string]interface{}, error) {
	reqUrl := fmt.Sprintf("%sconvert?to=%s&from=%s&amount=%s", url, to, from, amount)
	req, err := http.NewRequest("GET", reqUrl, nil)
	client := &http.Client{}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	req.Header.Set("apikey", apikey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var data map[string]interface{}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func isSymbol(symbol string, symbolsString string) bool {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(symbolsString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	symbols, ok := data["symbols"].(map[string]interface{})
	if !ok {
		fmt.Println("Error retrieving symbols")
		return false
	}

	if _, ok = symbols[symbol]; ok {
		return true
	} else {
		return false
	}
}

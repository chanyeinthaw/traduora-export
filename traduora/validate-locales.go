package traduora

import (
	"encoding/json"
	"fmt"
	"github.com/chanyeinthaw/traduora-export/config"
	"io"
	"net/http"
	"os"
)

type translations struct {
	Data []translation `json:"data"`
}

type translation struct {
	Locale struct {
		Code string `json:"code"`
	} `json:"locale"`
}

func ValidateLocales() {
	resp := sendGetTranslationsRequest()
	locales := getAvailableLocales(resp)

	for _, l := range config.Locales() {
		_, ok := locales[l]
		if !ok {
			fmt.Printf("Invalid locale code %s\n", l)
			os.Exit(1)
		}
	}
}

func getAvailableLocales(resp *http.Response) (availableLocales map[string]bool) {
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)

	parsedBody := &translations{}
	err := json.Unmarshal(body, parsedBody)

	if err != nil {
		fmt.Print(err)
		fmt.Println("Error reading locales response")
		os.Exit(1)
	}

	availableLocales = make(map[string]bool)
	for _, t := range parsedBody.Data {
		availableLocales[t.Locale.Code] = true
	}

	return
}

func sendGetTranslationsRequest() (resp *http.Response) {
	url := config.ApiURL(fmt.Sprintf("/api/v1/projects/%s/translations", config.ProjectId()))
	req, _ := http.NewRequest("GET", url, nil)
	authenticateRequests(req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Unable to get available locales")
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Unable to get available locales")
		os.Exit(1)
	}

	return
}

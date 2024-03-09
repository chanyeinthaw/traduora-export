package traduora

import (
	"fmt"
	"github.com/chanyeinthaw/traduora-export/config"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Translation struct {
	locale string
	data   string
}

func ExportTranslations() {
	numLocales := len(config.Locales())
	ch := make(chan bool, numLocales)
	for _, l := range config.Locales() {
		go exportTranslation(l, ch)
	}

	for i := 0; i < numLocales; i++ {
		<-ch
	}

	return
}

func exportTranslation(locale string, result chan bool) {
	reqURL := config.ApiURL(fmt.Sprintf("/api/v1/projects/%s/exports", config.ProjectId()))

	params := url.Values{}
	params.Set("locale", locale)
	params.Set("format", "jsonnested")

	u, _ := url.Parse(reqURL)
	u.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	authenticateRequests(req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Unable to export locale %s\n", locale)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unable to export locale %s\n", locale)
		os.Exit(1)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)

	err = os.MkdirAll(config.OutputDir(), 0755)
	if err != nil {
		fmt.Println("Error creating output directory")
		os.Exit(1)
	}

	outFile, err := os.OpenFile(fmt.Sprintf("%s/%s.json", config.OutputDir(), locale), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file to write")
		os.Exit(1)
	}
	defer func() {
		_ = outFile.Close()
	}()

	_, err = outFile.WriteString(string(body))
	if err != nil {
		fmt.Printf("Error writing to file %s.json", locale)
		os.Exit(1)
	}

	result <- true
}

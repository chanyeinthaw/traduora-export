package traduora

import (
	"fmt"
	"github.com/chanyeinthaw/traduora-export/config"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type Translation struct {
	locale string
	data   string
}

func ExportTranslations() {
	fmt.Println("Exporting locales ...")
	wg := sync.WaitGroup{}
	for _, l := range config.Locales() {
		wg.Add(1)
		l := l
		go func() {
			exportTranslation(l)
			wg.Done()
		}()
	}

	wg.Wait()
}

func exportTranslation(locale string) {
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

	ext, content := prepareOutput(body)
	outFile, err := os.OpenFile(fmt.Sprintf("%s/%s.%s", config.OutputDir(), locale, ext), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file to write")
		os.Exit(1)
	}
	defer func() {
		_ = outFile.Close()
	}()

	_, err = outFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file %s.json", locale)
		os.Exit(1)
	}
}

func prepareOutput(body []byte) (ext string, content string) {
	ext = "json"
	content = string(body)

	switch config.OutputFormat() {
	case "ts":
		ext = "ts"
		content = fmt.Sprintf("export default %s", content)
		break
	}

	return
}

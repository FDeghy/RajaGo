package raja

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func GetApiKey() (string, error) {
	var jsName, apiKey string

	resp, err := Client.Get(BASE_URL)
	if err != nil {
		return "", fmt.Errorf("get base error: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read base error: %w", err)
	}
	defer resp.Body.Close()

	re, err := regexp.Compile("<script src=\"([^\"]+)\" type=\"module\">")
	if err != nil {
		return "", fmt.Errorf("<script> regex compile error: %w", err)
	}
	if match := re.FindAllStringSubmatch(string(body), -1); match != nil {
		jsName = match[len(match)-1][1] // find the last js in raja base
	} else {
		return "", errors.New("js not found")
	}

	resp, err = Client.Get(BASE_URL + "/" + jsName)
	if err != nil {
		return "", fmt.Errorf("get js file error: %w", err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read js file error: %w", err)
	}
	defer resp.Body.Close()

	re, err = regexp.Compile("\"api-key\":\"([^\"]+)\"")
	if err != nil {
		return "", fmt.Errorf("\"api-key\" regex compile error: %w", err)
	}
	if match := re.FindStringSubmatch(string(body)); match != nil {
		apiKey = match[1]
	} else {
		return "", errors.New("api-key not found")
	}

	return apiKey, nil
}

func GetTrainList(tinfo TrainInfo, apiKey string, password []byte) (TrainList, error) {
	jsonTrains := new(TrainList)
	q := tinfo.encode()
	query, err := encrypt(q, password)
	if err != nil {
		return *jsonTrains, err
	}

	req, err := http.NewRequest("GET", SERVICE_URL+"/Api/ServiceProvider/TrainListEq", nil)
	if err != nil {
		return *jsonTrains, err
	}

	req.Header.Set("api-key", apiKey)
	req.Header.Set("User-Agent", USER_AGENT)
	params := req.URL.Query()
	params.Set("q", query)
	req.URL.RawQuery = params.Encode()

	resp, err := Client.Do(req)
	if err != nil {
		return *jsonTrains, fmt.Errorf("get train list error: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return *jsonTrains, err
	}
	err = json.Unmarshal(data, &jsonTrains)
	if err != nil {
		return *jsonTrains, fmt.Errorf("decode json error: %w", err)
	}

	return *jsonTrains, nil
}

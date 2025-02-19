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

	req, _ := http.NewRequest(
		http.MethodGet,
		BASE_URL,
		nil,
	)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:135.0) Gecko/20100101 Firefox/135.0")

	resp, err := Client.Do(req)
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

	req, _ = http.NewRequest(
		http.MethodGet,
		BASE_URL+"/"+jsName,
		nil,
	)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:135.0) Gecko/20100101 Firefox/135.0")

	for fail := 0; ; fail++ {
		resp, err = Client.Do(req)
		if err != nil {
			if fail > 3 {
				return "", fmt.Errorf("get js file error: %w", err)
			}
			continue
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			if fail > 3 {
				return "", fmt.Errorf("read js file error: %w", err)
			}
			continue
		}
		defer resp.Body.Close()

		re, err = regexp.Compile("\"api-key\":\"([^\"]+)\"")
		if err != nil {
			return "", fmt.Errorf("\"api-key\" regex compile error: %w", err)
		}
		if match := re.FindStringSubmatch(string(body)); match != nil {
			apiKey = match[1]
			break
		} else {
			if fail > 3 {
				return "", errors.New("api-key not found")
			}
			continue
		}
	}
	return apiKey, nil
}

func GetTrainList(query Query, opt *GetTrainListOpt) (*TrainList, error) {
	if opt == nil {
		ak, err := GetApiKey()
		if err != nil {
			return nil, err
		}
		opt = &GetTrainListOpt{
			ApiKey:     ak,
			HttpClient: &Client,
		}
	}

	req, err := http.NewRequest("GET", SERVICE_URL+"/Api/ServiceProvider/TrainListEq", nil)
	if err != nil {
		return nil, ErrGetTrains
	}

	req.Header.Set("api-key", opt.ApiKey)
	req.Header.Set("User-Agent", USER_AGENT)
	params := req.URL.Query()
	params.Set("q", query.String())
	req.URL.RawQuery = params.Encode()

	resp, err := Client.Do(req)
	if err != nil {
		return nil, ErrGetTrains
	}
	if resp.StatusCode == 400 {
		return nil, ErrTrainsNotFound
	} else if resp.StatusCode != 200 {
		return nil, ErrBadStatus
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrGetTrains
	}
	jsonTrains := &TrainList{}
	err = json.Unmarshal(data, &jsonTrains)
	if err != nil {
		return nil, ErrGetTrainsDecode
	}

	return jsonTrains, nil
}

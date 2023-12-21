package raja

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)

func UpdateStationsJson() error {
	resp, err := Client.Get(BASE_URL + "/assets/File/station.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	fd, err := os.Create("./stations.json")
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func LoadStations() error {
	data, _ := os.ReadFile("./stations.json")
	err := json.Unmarshal(data, &Stations)
	if err != nil {
		return err
	}
	return nil
}

func FindStation(name string) (Station, error) {
	for _, i := range Stations {
		if strings.EqualFold(name, i.EnglishName) {
			return i, nil
		}
	}
	return Station{}, errors.New("station \"" + name + "\" not found")
}

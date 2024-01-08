package raja

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func GetStations() (*Stations, error) {
	resp, err := Client.Get(BASE_URL + "/assets/File/station.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	stations := &Stations{}
	err = json.Unmarshal(body, stations)
	if err != nil {
		return nil, err
	}
	return stations, nil
}

// english name
func (sts Stations) FindStationID(name string) (int, error) {
	for _, i := range sts {
		if strings.EqualFold(name, i.EnglishName) {
			return i.Id, nil
		}
	}
	return -1, errors.New("station \"" + name + "\" not found")
}

// by id
func (sts Stations) GetPersianName(id int) (string, error) {
	for _, i := range sts {
		if id == i.Id {
			return i.PersianName, nil
		}
	}
	return "", fmt.Errorf("station \"%d\" not found", id)
}

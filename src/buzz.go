package buzz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type BuzzWord struct {
	Keyword string `json:"keyword"`
	Detail  string `json:"detail"`
}

var (
	BuzzFilePath = path.Join(HomePath(), ".buzzwords.json")
)

func HomePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func SaveBuzzwords(buzzwords []BuzzWord, filepath string) error {
	buzzString, err := json.Marshal(buzzwords)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, buzzString, 0666)
	return err
}

func LoadBuzzwords(filepath string) ([]BuzzWord, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ret []BuzzWord
	err = json.Unmarshal(content, &ret)

	return ret, err
}

func AppendBuzzword(buzzword BuzzWord) ([]BuzzWord, error) {
	buzzwords, err := LoadBuzzwords(BuzzFilePath)
	fmt.Println(buzzwords)
	if err != nil {
		return nil, err
	}

	buzzwords = append(buzzwords, buzzword)
	err = SaveBuzzwords(buzzwords, BuzzFilePath)
	return buzzwords, err
}

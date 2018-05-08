package buzz

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"fmt"
)

type BuzzWord struct {
	Keyword string `json:"keyword"`
	Detail  string `json:"detail"`
}

var (
	BuzzFilePath = path.Join(HomePath(), ".buzzwords.json")
)

// Get user homepath
func HomePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

// Storge buzzwords
func SaveBuzzwords(buzzwords []BuzzWord, filepath string) error {
	buzzString, err := json.Marshal(buzzwords)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, buzzString, 0666)
	return err
}

// Load all buzzwords from local
func LoadBuzzwords(filepath string) ([]BuzzWord, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil || len(content) == 0 {
		return make([]BuzzWord, 0), err
	}

	var ret []BuzzWord
	err = json.Unmarshal(content, &ret)

	return ret, err
}

// List all local buzzwords
func ListBuzzwords() ([]BuzzWord, error) {
	buzzwords, err := LoadBuzzwords(BuzzFilePath)
	return buzzwords, err
}

func ShowListBuzzwords() {
	buzzwords, err := ListBuzzwords()
	if err != nil {
		panic(err)
	}

	for _, buzzword := range buzzwords {
		fmt.Printf("%s:\t%s\n", buzzword.Keyword, buzzword.Detail)
	}
}

// Append buzzword to database
func AppendBuzzword(buzzword BuzzWord) ([]BuzzWord, error) {
	buzzwords, err := LoadBuzzwords(BuzzFilePath)
	if err != nil {
		return nil, err
	}

	buzzwords = append(buzzwords, buzzword)
	err = SaveBuzzwords(buzzwords, BuzzFilePath)
	return buzzwords, err
}

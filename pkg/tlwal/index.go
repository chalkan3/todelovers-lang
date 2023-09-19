package wal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type KeyIndex struct {
	fileName string
	Index    map[string]int64 `json:"index"`
}

func NewKeyIndex() *KeyIndex {
	return &KeyIndex{
		Index:    make(map[string]int64),
		fileName: "data/index/idx.json",
	}
}

func (k *KeyIndex) Load() *KeyIndex {
	if _, err := os.Stat(k.fileName); err == nil {
		data, err := ioutil.ReadFile(k.fileName)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(data, k); err != nil {
			log.Fatal(err)
		}
	} else if os.IsNotExist(err) {
		k.Index = make(map[string]int64)
	} else {
		log.Fatal(err)
	}
	return k
}

func (k *KeyIndex) Save() {
	data, err := json.Marshal(k)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(k.fileName, data, 0644); err != nil {
		log.Fatal(err)
	}
}

package wal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type TLWAL interface {
	CreateLogFile() (*os.File, error)
	Write(record *Record, unique bool) error
	Read(filename string) ([]*Record, error)
	ReadAll() ([]*Record, error)
}

type tLWAL struct {
	fileName       string
	maxLogFileSize int64
	indexes        *KeyIndex
}

func NewTLWAL() TLWAL {
	return &tLWAL{
		fileName:       fmt.Sprintf("data/wal-%s.log", time.Now().Format("2006-01-02-150405")),
		maxLogFileSize: 1 * 1024 * 1024, //1 MB
		indexes:        NewKeyIndex().Load(),
	}
}

func (wal *tLWAL) CreateLogFile() (*os.File, error) {
	file, err := os.Create(wal.fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (wal *tLWAL) write(record *Record) error {
	fileInfo, err := os.Stat(wal.fileName)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if fileInfo != nil && fileInfo.Size() > wal.maxLogFileSize {
		newLogFile, err := wal.CreateLogFile()
		if err != nil {
			return err
		}
		wal.fileName = newLogFile.Name()
		newLogFile.Close()
	}

	record.Timestamp = time.Now()

	file, err := os.OpenFile(wal.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(record); err != nil {
		return err
	}

	index := NewKeyIndex().Load()
	index.Index[record.Key()] = fileInfo.Size()
	index.Save()

	return nil
}

func (wal *tLWAL) Write(record *Record, unique bool) error {
	if unique {
		if _, exists := wal.indexes.Index[record.Key()]; exists {
			fmt.Println("=====================O registro com a chave 1 já existe.")
			return nil
		}
	}
	return wal.write(record)
}
func (wal *tLWAL) Read(filename string) ([]*Record, error) {
	var records []*Record

	file, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var record *Record
		if err := decoder.Decode(&record); err != nil {
			return records, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (wal *tLWAL) ReadAll() ([]*Record, error) {
	var allRecords []*Record

	files, err := filepath.Glob("data/wal-*.log")
	if err != nil {
		return allRecords, err
	}

	for _, file := range files {
		records, err := wal.Read(file)
		if err != nil {
			return allRecords, err
		}
		allRecords = append(allRecords, records...)
	}

	return allRecords, nil
}

package main

import (
	"bytes"
	"io/ioutil"
)

func getData() []byte {
	data, _ := ioutil.ReadFile("assets/nvi.json")
	//Correção de erro: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	return bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

}

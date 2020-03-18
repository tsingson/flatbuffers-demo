package main

import (
	"io/ioutil"
)

// WriteToFile  write []byte to file
func WriteToFile(c []byte, filename string) error {
	// 将指定内容写入到文件中

	err := ioutil.WriteFile(filename, c, 0666)
	return err
}

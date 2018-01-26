package util

import "os"

//FileExists is ファイルの有無を確認する
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

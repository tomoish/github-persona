package funcs

import (
    "bytes"
    "os"
)


// 画像をファイルに保存する
func saveImageToFile(filename string, imgBytes []byte) error {
    outputFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    _, err = outputFile.Write(imgBytes)
    if err != nil {
        return err
    }

    return nil
}
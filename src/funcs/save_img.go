package funcs


import (
    "os"

)

// SaveImage は指定されたファイルパスに画像データを保存します。


// SaveImage は指定されたファイルパスに画像データを保存します。
// filename には保存先のファイルパスを、imgBytes には画像データを渡します。
func SaveImage(filename string, imgBytes []byte) error {
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

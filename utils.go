package main


import (
"io"
"bufio"
"log"
)
func IsInSlice[T any](str T, arr []T) bool {
for _,s := range arr {
  if str == s {
  return true 
}
  
}
return false 
}


func ByteFromMegaFile(file io.Reader) ([]byte, error) {

        reader := bufio.NewReader(file)

        finalByteArr := make([]byte, 0, 2048*1000)

        for {
                soloByte, err := reader.ReadByte()
                if err != nil {
                        log.Println(err)
                        break
                }

                finalByteArr = append(finalByteArr, soloByte)
        }

        return finalByteArr, nil

}
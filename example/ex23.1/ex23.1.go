package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "data.txt"

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n') //한 줄을 읽는다

	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
	}

	line, err = ReadFile(filename)
	if err != nil {
		fmt.Println("파일 읽기에 실패했습니다.", err)
		return
	}

	fmt.Println("파일내용:", line)
}

/*

PS C:\Users\조재모\goprojects\example\ex23.1> .\ex23.1.exe
파일내용: This is WriteFile

PS C:\Users\조재모\goprojects\example\ex23.1> ls




Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----      2023-02-03  오전 11:23             18 data.txt
-a----      2023-02-03  오전 11:23        1999360 ex23.1.exe
-a----      2023-02-03  오전 11:22            946 ex23.1.go
-a----      2023-02-03  오전 11:21             34 go.mod


PS C:\Users\조재모\goprojects\example\ex23.1> cat .\data.txt
This is WriteFile

*/

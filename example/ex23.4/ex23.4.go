package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split((bufio.ScanWords)) // 한단어씩

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pod:%d err:%w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)

		var numError *strconv.NumError
		// num 관련 에러로 변경
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("3 abc")
}

/*
PS C:\Users\조재모\goprojects\example\ex23.4> .\ex23.4.exe
369
failed to readNextInt(), pod:2 err:failed to convert word to int, word:abc, err:strconv.Atoi: parsing "abc": invalid syntax
NumberError strconv.Atoi: parsing "abc": invalid syntax

*/

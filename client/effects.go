package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

// 让mac说一句话
func say(word string) {
	cmd := exec.Command("say", word)
	if err := cmd.Run(); err != nil {

	}
}

// 抽卡动画
func animation() {
	clearScreen()
	hideCursor()
	assets := AssetNames()
	sort.Strings(assets)
	for _, file := range assets {
		raw, err := Asset(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s", raw)
		time.Sleep(10000000)
		fmt.Printf("\033[0;0H")
	}
	showCursor()
}

// 清空屏幕
func clearScreen() {
	fmt.Printf("\033[2J\033[0;0H")
}

// 隐藏光标
func hideCursor() {
	fmt.Printf("\033[?25l")
}

// 显示光标
func showCursor() {
	fmt.Printf("\033[?25h")
}

// 开始屏幕
func startScreen() {
	clearScreen()
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("*         __________________                                                                  *\n")
	fmt.Printf("*        |     ☆ ☆ ★ ☆ ☆    |                                                                 *\n")
	fmt.Printf("*        |                  |                                                                 *\n")
	fmt.Printf("*        |     1 1  2 2 2   |          EncryptCard                                            *\n")
	fmt.Printf("*        |    1  1     2    |                                                                 *\n")
	fmt.Printf("*        |   1 1 1    2     |          v0.0.1                                                 *\n")
	fmt.Printf("*        |       1   2      |                                                                 *\n")
	fmt.Printf("*        |       1  2 2 2   |          Use Proof-of-Work digging card                         *\n")
	fmt.Printf("*        |                  |          with Distributed Game Archives System                  *\n")
	fmt.Printf("*        |       Answer     |                                                                 *\n")
	fmt.Printf("*        |                  |          Project: https://github.com/bydmm/encryptcard          *\n")
	fmt.Printf("*        |   A T K   D E F  |                                                                 *\n")
	fmt.Printf("*        |    9 9     9 9   |                                                                 *\n")
	fmt.Printf("*        |__________________|          Concurrency: %d                                         *\n", maxConcurrency)
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
	time.Sleep(1 * time.Second)
}

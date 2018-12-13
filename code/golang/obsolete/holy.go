package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// 类型定义
type Celsius float64    // 攝氏溫度
type Fahrenheit float64 // 華氏溫度
const (
	AbsoluteZeroC Celsius = -273.15 // 絶對零度
	FreezingC     Celsius = 0       // 結冰點溫度
	BoilingC      Celsius = 100     // 沸水溫度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// 命令行参数解析
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	// 数组迭代
	medals := []string{"gold", "silver", "bronze"}
	for i := range medals {
		fmt.Println(medals[i])
	}

	fmt.Println(CToF(32))
}

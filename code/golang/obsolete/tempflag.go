package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Celsius float64
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

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *celsiusFlag) String() string {
	c := float64(f.Celsius)
	return strconv.FormatFloat(c, 'E', -1, 64)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

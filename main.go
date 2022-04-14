package main

import (
	"fmt"
	"time"
)

//i5 := 500
var i5 int = 500

func main() {
	fmt.Println("hello NiMiL")
	fmt.Println(time.Now())

	//明示的な定義
	//var 変数名 型　＝　値
	var i int = 100
	fmt.Println(i)
	var s string = "test string"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "test string2"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "test string3"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//変数名 := 値

	i4 := 400 //最初に定義した値の型を同的に位置付ける
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	println(i5) //明示的な変数定義だとbody outsideでも出力できる
}

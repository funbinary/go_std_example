package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Example1() {
	var buf bytes.Buffer

	// 序列化
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(123)
	if err != nil {
		panic(err)
	}

	// 反序列化
	decoder := gob.NewDecoder(&buf)
	var i int
	err = decoder.Decode(&i)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)

	//Output:
	//123
}

type Person struct {
	Name string
	Age  int
}

func Example2() {
	var buf bytes.Buffer

	// 序列化
	encoder := gob.NewEncoder(&buf)
	p := Person{Name: "Alice", Age: 20}
	err := encoder.Encode(p)
	if err != nil {
		panic(err)
	}

	// 反序列化
	decoder := gob.NewDecoder(&buf)
	var p2 Person
	err = decoder.Decode(&p2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p2)

	//Output:
	//{Alice 20}
}

func Example3() {
	var buf bytes.Buffer

	// 注册数据类型
	gob.Register(Person{})

	// 序列化
	encoder := gob.NewEncoder(&buf)
	p := Person{Name: "Alice", Age: 20}
	err := encoder.Encode(p)
	if err != nil {
		panic(err)
	}

	// 反序列化
	decoder := gob.NewDecoder(&buf)
	var p2 Person
	err = decoder.Decode(&p2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p2)

	//Output:
	//{Alice 20}
}

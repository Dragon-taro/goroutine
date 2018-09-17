package main

import (
	"fmt"
	"time"
)

type Task struct {
	name        string
	processTime int
}

func (t *Task) ProcessTask(c chan string) {
	fmt.Printf("Start %s\n", t.name)
	for i := 0; i < t.processTime; i++ {
		time.Sleep(1 * time.Millisecond)
	}

	c <- "End " + t.name + "!\n"
}

var TaskA = Task{name: "TaskA", processTime: 1000}
var TaskB = Task{name: "TaskB", processTime: 3000}
var TaskC = Task{name: "TaskC", processTime: 2000}

func Parallel() {
	c := make(chan string)
	go TaskA.ProcessTask(c)
	go TaskB.ProcessTask(c)
	go TaskC.ProcessTask(c)

	// 入る順番はgoroutineが終わった順
	// channelはベルトコンベアのイメージ
	m1, m2, m3 := <-c, <-c, <-c

	// 全部の<-cを受け取るまで処理がブロックされてこの処理は実行されない
	fmt.Println(m1, m2, m3)

	// 以下だと順次<-cを受け取って表示される
	// m1 := <-c
	// fmt.Println(m1)
	// m2 := <-c
	// fmt.Println(m2)
	// m3 := <-c
	// fmt.Println(m3)
}

func main() {
	Parallel()
}

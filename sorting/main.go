package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StudentSlice []Student

func (ss StudentSlice) Len() int {
	return len(ss)
}

func (ss StudentSlice) Less(i, j int) bool {
	return ss[i].Score > ss[j].Score
}

func (ss StudentSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func randomScore() float64 {
	v := fmt.Sprintf("%.2f", rand.Float64()*100)
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		return s
	}
	return 0
}

func main() {
	var studentSlice StudentSlice
	for i := 0; i < 10; i++ {
		s := Student{
			Name:  fmt.Sprintf("Student-%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: randomScore(),
		}
		studentSlice = append(studentSlice, s)
	}

	fmt.Println("---Before sorting by scores---")
	for _, s := range studentSlice {
		fmt.Println(s)
	}

	sort.Sort(studentSlice)

	fmt.Println("---After sorting by scores---")
	for _, s := range studentSlice {
		fmt.Println(s)
	}
}

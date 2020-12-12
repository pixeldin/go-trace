package mdl

import "fmt"

// init param
type Param struct {
	Concurrency  uint64
	TotalQuest   uint64
	Url string
	Header  array
	Body    string
}

/*
	implement of flag.Value, signature:
	String() string
	Set(string) error
 */
type array[]string

func (a *array) String() string {
	return fmt.Sprint(*a)
}

func (a *array) Set(s string) error {
	*a = append(*a, s)

	return nil
}


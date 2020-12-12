package com

func AssertErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AssertWithCondition(pan bool, msg string) {
	if pan {
		panic(msg)
	}
}

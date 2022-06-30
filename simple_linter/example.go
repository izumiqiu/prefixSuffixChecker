package main

type Context struct {
	val int
}
type aa struct {
	Context Context
}

func add(a int, b int) int {
	return a + b
}

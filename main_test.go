package main

import (
	"regexp"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("bodyFrom", func() {
		g.It("Should join a slice into a single string", func() {
			g.Assert(bodyFrom([]string{"no", "hello", "world"})).Equal("hello world")
		})
	})

	g.Describe("randomString", func() {
		g.It("Should generate a random string of the passed in length", func() {
			g.Assert(len(randomString(10))).Equal(10)
			g.Assert(len(randomString(20))).Equal(20)
			g.Assert(len(randomString(30))).Equal(30)
			g.Assert(len(randomString(100))).Equal(100)
		})

		g.It("Should only use capital letters", func() {
			allCaps := func(str string) bool {
				matched, _ := regexp.MatchString("^[^a-z]*$", str)
				return matched
			}

			g.Assert(allCaps(randomString(10))).Equal(true)
			g.Assert(allCaps(randomString(30))).Equal(true)
		})
	})

	g.Describe("randInt", func() {
		g.It("Should generate a random int between passed in parameters", func() {
			tups := [][]int{{10, 20}, {30, 50}, {65, 90}}

			between := func(n, min, max int) bool {
				return min <= n && n <= max
			}

			for _, v := range tups {
				min, max := v[0], v[1]
				g.Assert(between(randInt(min, max), min, max)).Equal(true)
			}

		})
	})

}

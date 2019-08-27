package spider

import (
	"fmt"
	"regexp"
)

func (g *GameInfo) ParseFenxi(body string) {
	s := NewSpider()
	body := s.Fetch(g.fenxiurl)
	//regexp.MustCompile(``)
	fmt.Println(body)
}

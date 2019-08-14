package bilibili

import (
	"fmt"
	"learningnotes/reptile/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://www.bilibili.com/bangumi/play/ss[0-9]+)"[^>]+><[^>]>([^<]+).+?</a>`)
)

func Getbangumi(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}
	fmt.Println("sdfghjtredx n")
	match := profileRe.FindAllSubmatch(contents, -1)
	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse: engine.NewFuncParser(Detailbangumi, "detail"),
		})
	}
	// 取本页面其它城市链接

	return rs
}
func Detailbangumi(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}
	fmt.Println("sss")
	return rs
}

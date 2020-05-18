package SearchMachine

import (
	"../HTMLParser"
	"strings"
)



/*
配下の要素をたどる
result <- 見つけた要素を入れるスライス
elem <- 検索する木の根を指定
counter <- のこり何個見つけるか
tag <- タグ名
optionName <- classとかnameとか
optionValue <-
 */
func Search_(result *[]*HTMLParser.Element,elem *HTMLParser.Element,counter int,tag string,optionName string,optionValue []string) {
	if counter <= 0{return}
	if check(elem,tag,optionName,optionValue) {
		(*result) = append(*result,elem)
	}
	// if (*counter) <= 0{return}
	for _,child := range elem.Data {
		switch v := child.(type) {
		case *HTMLParser.Element:
			Search_(result,v,counter,tag,optionName,optionValue)
		}
	}
}

// 要素が該当するかチェック
func check(elem *HTMLParser.Element,tag string,optionName string,optionValue []string) bool {
	if !tagCheck(elem,tag) {return false}
	if optionName == "" {return true}
	for _,v := range optionValue {
		if optionCheck(elem,optionName,v) {
			return true
		}
	}
	return false
}

// optionがあるかどうか
func optionCheck(elem *HTMLParser.Element,k string,v string) bool {
	if k == "class" {return optionCheck_class(elem,v)}
	if elemV,ok := elem.Option[k];ok {
		if elemV == v {
			return true
		}
	}
	return false
}

// classのみ複数チェック
func optionCheck_class(elem *HTMLParser.Element,v string) bool {
	ls := strings.Split(elem.Option["class"]," ")
	for _,vv := range ls {
		if vv == v{return true}
	}
	return false
}

// タグが一致しているか
func tagCheck(elem *HTMLParser.Element,tag string) bool {
	if tag == "" {return true}
	return elem.Tag == tag
}
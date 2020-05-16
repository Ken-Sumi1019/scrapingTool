package SearchMachine

import (
	"strings"
	"../HTMLParser"
)

type dataBox struct {

}

/*
検索するときの引数を
"class:hoge,id:hoge"
みたいな感じで受け取りたいので
 */
func argumentSplit(s string) map[string][]string {
	ls := strings.Split(s,",")
	result := map[string][]string{}
	for _,v := range ls {
		key_value := strings.Split(v,":")
		if _,ok := result[key_value[0]];ok {
			result[key_value[0]] = append(result[key_value[0]], key_value[1])
		} else {
			result[key_value[0]] = []string{key_value[1]}
		}
	}
	return result
}

/*
配下の要素をたどる
result <- 見つけた要素を入れるスライス
elem <- 検索する木の根を指定
counter <- のこり何個見つけるか
 */
func search(result *[]*HTMLParser.Element,elem *HTMLParser.Element,counter int)  {
	if check(elem) {
		(*result) = append(*result,elem)
		counter--
	}
	if counter == 0{return}
	for _,child := range elem.Data {
		switch v := child.(type) {
		case *HTMLParser.Element:
			search(result,v,counter)
		}
	}
}

// 要素が該当するかチェック
func check(elem *HTMLParser.Element) bool {
	return true
}
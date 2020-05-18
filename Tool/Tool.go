package Tool

import (
	"./HTMLParser"
	"./SearchMachine"
)

func ParseHTML(s string) *HTMLParser.Element {
	l := HTMLParser.RemoveAllComment(s)
	return HTMLParser.Solv(l)
}

func SearchFirst(elem *HTMLParser.Element,tag string,optionName string,optionValue []string) (*HTMLParser.Element) {
	result := []*HTMLParser.Element{}
	n := 1
	SearchMachine.Search_(&result,elem,n,tag,optionName,optionValue)
	if len(result) != 0 {
		return result[0]
	} else {
		return nil
	}
}

func SearchAll (elem *HTMLParser.Element,tag string,optionName string,optionValue []string) []*HTMLParser.Element {
	result := []*HTMLParser.Element{}
	n := 2147483647
	SearchMachine.Search_(&result,elem,n,tag,optionName,optionValue)
	return result
}

func GetText(elem *HTMLParser.Element) string {
	s := ""
	HTMLParser.Decode(elem,0,&s,1)
	return s
}

func GetTextNoneTab(elem *HTMLParser.Element) string {
	s := ""
	HTMLParser.Decode(elem,0,&s,0)
	return s
}
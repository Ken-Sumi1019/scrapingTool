package HTMLParser

import (
	"./singularData"
	"fmt"
	"regexp"
	"strings"
)

// https://html.spec.whatwg.org/
var (
	singleTag = singularData.SingleTag()
	canBeOmitted = singularData.CanBeOmitted()
	nonePareOmitted = singularData.NonePareOmitted()
	ptag = singularData.Ptag()
	tagSearch = regexp.MustCompile("[a-zA-Z][-.a-zA-Z0-9:_]*")
	option = regexp.MustCompile(`\s+[a-zA-Z_][-.:a-zA-Z0-9_]*(?:\s*=\s*(?:'[^']*'|\"[^\"]*\"|[^'\">\s]+))?`)
	optionKey = regexp.MustCompile(`\s+(?:[a-zA-Z_][-.:a-zA-Z0-9_]*\s*)*=`)
	optionVal = regexp.MustCompile(`=\s*(?:'[^']*'|\"[^\"]*\"|[^'\">\s]+)`)
)



type Element struct {
	Tag string
	Data []interface{}
	Option map[string]string
}


func NewElement(tag string,d map[string]string) *Element {
	return &Element{Tag:tag,Option:d}
}

func getAttribute(s string) (string,map[string]string) {
	result := map[string]string{}
	tag := tagSearch.FindIndex([]byte(s))
	optionList := option.FindAllIndex([]byte(s),-1)
	for _,v := range optionList {
		keyindex := optionKey.FindIndex([]byte(s[v[0]:v[1]]))
		valindex := optionVal.FindIndex([]byte(s[v[0]:v[1]]))
		if len(keyindex) <= 1 {
			result[removeSpace(s[v[0]:v[1]])] = "";continue
		}
		key := s[v[0]:v[1]][keyindex[0]:keyindex[1]-1]
		key = removeSpace(key)
		idx := 0
		for i := valindex[0];i < valindex[1];i ++ {
			if s[v[0]:v[1]][i] == '"' || s[v[0]:v[1]][i] == '\''{
				idx = i;break
			}
		}
		val := s[v[0]:v[1]][idx+1:valindex[1]-1]
		result[key] = val
	}
	return s[tag[0]:tag[1]],result
}

// This function remove front space and back space
func removeSpace(s string) string {
	tmp := strings.TrimLeft(s,"\n \t\r")
	tmp = strings.TrimRight(tmp,"\n \t\r")
	return tmp
}

func locateStartTag(s []byte) [][]int {
	mutchstring := `<[a-zA-Z][-.a-zA-Z0-9:_]*(?:\s+(?:[a-zA-Z_][-.:a-zA-Z0-9_]*(?:\s*=\s*(?:'[^']*'|\"[^\"]*\"|[^'\">\s]+))?))*\s*/?>`
	re := regexp.MustCompile(mutchstring)
	return re.FindAllIndex(s,-1)
}

func locateEndTag(s []byte) [][]int {
	mutchstring := `</[a-zA-Z][-.a-zA-Z0-9:_]*\s*>`
	re := regexp.MustCompile(mutchstring)
	return re.FindAllIndex(s,-1)
}

// This function find indexes HTML tag
func findAllIndex(s string) [][]int {
	content := []byte(s)
	start := locateStartTag(content);end := locateEndTag(content)
	result := make([][]int,len(start)+len(end))
	now := 0;sidx := 0;eidx := 0
	for ;now < len(result); {
		if sidx >= len(start) {
			result[now] = end[eidx];eidx ++
		} else if eidx >= len(end) {
			result[now] = start[sidx];sidx ++
		} else if start[sidx][0] > end[eidx][0]{
			result[now] = end[eidx];eidx ++
		} else {
			result[now] = start[sidx];sidx ++
		}
		now ++
	}
	return result
}

func RemoveAllComment(s string) string {
	re := regexp.MustCompile("<!--[\\S\\s]*?-->")
	s = re.ReplaceAllString(s,"")
	return s
}

func detectionOmitted(stack *[]*Element,tag string)  {
	s := (*stack)[len(*stack)-1].Tag
	if _,ok := canBeOmitted[tag];!ok{return}
	if canBeOmitted[tag].Exist(s) {
		*stack = (*stack)[:len(*stack) - 1]
	}
}

func addElement(stack *[]*Element,s string)  {
	tag,option := getAttribute(s)
	detectionOmitted(stack,tag)
	ref := NewElement(tag,option)
	(*stack)[len(*stack)-1].Data = append((*stack)[len(*stack)-1].Data, ref)
	if (! singleTag.Exist(tag)) && (! (s[len(s)-2] == '/')) {
		*stack = append(*stack, ref)
	}
}

func popElement(stack *[]*Element,s string)  {
	tag,_ := getAttribute(s)//;tag = tag[1:]
	if (*stack)[len(*stack) - 1].Tag == tag {
		*stack = (*stack)[:len(*stack) - 1]
	} else if nonePareOmitted.Exist((*stack)[len(*stack) - 1].Tag) && (*stack)[len(*stack) - 2].Tag == tag{
		*stack = (*stack)[:len(*stack) - 2]
	} else {
		fmt.Println("failer",(*stack)[len(*stack)-1].Tag,tag)
	}
}

func removeDoctype(indexes *[][]int,s *string)  {
	tag,_ := getAttribute((*s)[(*indexes)[0][0]:(*indexes)[0][1]])
	if tag == "!doctype"{
		*indexes = (*indexes)[1:]
	}
}

// This function parse html and save struct named Element
func Solv(s string) *Element {
	indexes := findAllIndex(s)
	stack := []*Element{}
	elem := NewElement("root",map[string]string{})
	stack = append(stack,elem)
	for i := 0;i < len(indexes);i ++{
		if s[indexes[i][0] + 1] == '/' {
			popElement(&stack,s[indexes[i][0]:indexes[i][1]])
		} else {
			addElement(&stack,s[indexes[i][0]:indexes[i][1]])
		}
		if i + 1 == len(indexes){break}
		tmp := removeSpace(s[indexes[i][1]:indexes[i + 1][0]])
		if len(tmp) != 0 {
			stack[len(stack)-1].Data = append(stack[len(stack)-1].Data, tmp)
		}
	}
	fmt.Println(len(stack))
	return stack[0]
}

func mapToString(m map[string]string) string {
	s := ""
	for k,v := range m {
		s = strings.Join([]string{s," ",k,":",v},"")
	}
	return s
}

func Decode_(elem *Element,n int,s *string,ifTab int) {
	(*s) = strings.Join([]string{(*s),strings.Repeat("\t",n*ifTab),"<",elem.Tag,">",mapToString(elem.Option),ifNewLine(ifTab)},"")
	for i := 0;i < len(elem.Data);i ++ {
		switch v := elem.Data[i].(type) {
		case string:
			(*s) = strings.Join([]string{(*s),strings.Repeat("\t",(n+1)*ifTab),v,ifNewLine(ifTab)},"")
		case *Element:
			Decode_(v,n+1,s,ifTab)
		}
	}
	(*s) = strings.Join([]string{*s,strings.Repeat("\t",n*ifTab),"</",elem.Tag,">",ifNewLine(ifTab)},"")
}

func Decode(elem interface{},n int,s *string,ifTab int) {

	switch v := elem.(type) {
	case string:
		(*s) = strings.Join([]string{(*s),strings.Repeat("\t",n*ifTab),v,ifNewLine(ifTab)},"")
	case *Element:
		(*s) = strings.Join([]string{(*s),strings.Repeat("\t",n*ifTab),"<",v.Tag,">",mapToString(v.Option),ifNewLine(ifTab)},"")
		for i := 0;i < len(v.Data);i++ {
			Decode(v.Data[i],n + 1,s,ifTab)
		}
		(*s) = strings.Join([]string{*s,strings.Repeat("\t",n*ifTab),"</",v.Tag,">",ifNewLine(ifTab)},"")
	}
}

func ifNewLine(n int) string {
	if n == 0 {return ""}
	return "\n"
}
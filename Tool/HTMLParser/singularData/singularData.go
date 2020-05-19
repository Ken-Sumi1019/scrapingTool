package singularData

import (
	"../set"
)

func SingleTag() *set.Set {
	ls := []string{"base","link","meta","hr","img","br","kbd","wbr",
		"area","source","input","embed","col","keygen","param","track"}
	result := set.MakeSet()
	for _,v := range ls {
		result.Add(v)
	}
	return result
}

/*
hogeの直後にhoge1,hoge2が来る場合にhogeを省略可能という場合に対応
map{hoge2:{hoge},hoge1:{hoge}}
 */
func CanBeOmitted() map[string]*set.Set {
	ls := map[string][]string{"li":{"li"},"dt":{"dd","dt"},"dd":{"dt","dd"},"rt":{"rt","rp"},"rp":{"rt","rp"},
		"optgroup":{"optgroup","option"},"option":{"option"},"tbody":{"thead","tbody"},"tfoot":{"thead","tbody"},
		"tr":{"tr"},"td":{"td","th"},"th":{"td","th"},"address":{"p"},"article":{"p"},"aside":{"p"},"blockquote":{"p"},
		"details":{"p"},"div":{"p"},"fieldset":{"p"},"dl":{"p"},"figcaption":{"p"},"figure":{"p"},"footer":{"p"},
		"form":{"p"},"h1":{"p"},"h2":{"p"},"h3":{"p"},"h4":{"p"},"h5":{"p"},"h6":{"p"},"header":{"p"},"hgroup":{"p"},
		"hr":{"p"},"main":{"p"},"menu":{"p"},"nav":{"p"},"ol":{"p"},"p":{"p"},"pre":{"p"},"section":{"p"},"table":{"p"},
		"ul":{"p"}}
	result := map[string]*set.Set{}
	for k,_ := range ls {
		result[k] = set.MakeSet()
		for _,vv := range ls[k] {
			result[k].Add(vv)
		}
	}
	return result
	// pはあと
	// colgroup caption はつらそう
}

/*
hogeは親要素にそれ以上要素がない場合に省略可能の場合に対応
 */
func NonePareOmitted() *set.Set {
	ls := []string{"li","dd","rt","rp","optgroup","option","tbody","tfoot","tr","td","th","p"}
	result := set.MakeSet()
	for _,v := range ls {
		result.Add(v)
	}
	return result
}

func Ptag() *set.Set {
	ls := []string{"a","audio","del","ins","map","noscript","video"}
	result := set.MakeSet()
	for _,v := range ls {
		result.Add(v)
	}
	return result
}
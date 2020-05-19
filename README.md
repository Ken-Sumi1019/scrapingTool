# scrapingTool

pythonのbeautifulsoupみたいなのをgoで実現したいと思いました。

```go
package main

import (
    "./Tool"
    "fmt"
)

func main ()  {
    s := `<div><p>ようこそ</p></div>`
    e :=Tool.ParseHTML(s)
    p := Tool.SearchFirst(e,"p","",[]string{})
    fmt.Println(Tool.GetTextNoneTab(p))
}
```
```text
ようこそ
```
```text
func ParseHTML(s string) (elem *HTMLParser.Element)
```
HTMLをパースします。HTMLの木の根の要素のポインタを返します。

```text
func SearchFirst(elem *HTMLParser.Element,tag string,optionName string,optionValue []string) (*HTMLParser.Element)
```
指定した条件に該当する一番最初の要素のポインタを返します。tag and (option or option)
で探してきます。

```text
func SearchAll (elem *HTMLParser.Element,tag string,optionName string,optionValue []string) []*HTMLParser.Element
```
指定した条件に該当するすべての要素を探して該当する要素のポインタのスライス
を返します。

```text
func GetText(elem *HTMLParser.Element) string
```
指定した要素以下の要素をすべてテキストに変換して適宜インデントを
入れて返します。

```text
func GetTextNoneTab(elem *HTMLParser.Element) string
```
指定した要素以下の要素をすべてテキストに変換して返します。


### HTMLParser
 一部出てくる正規表現はbeautifulsoupを参考にしています。
 今の段階ではhtml,body,head,colgroup,captionの省略がある場合は想定外の動きをする可能性があります。
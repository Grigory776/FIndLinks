package links
import(
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Extract(url string)([]string,error){
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("получение %s:%s",url,resp.Status)
	}
	doc,err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("не удалось обработать %s как HTML: %v",url,err)
	}
	var Links []string
	SaveLinks := func(doc *html.Node) {
		if doc.Type == html.ElementNode && doc.Data == "a"{
			for _,val := range doc.Attr{
				if val.Key != "href"{
					continue
				}
				link, err := resp.Request.URL.Parse(val.Val)
				if err != nil { // Пропускаем некорректные ссылки
					continue    
				}
				Links = append(Links, link.String())
			}
		}
	}
	forEachNode(doc,SaveLinks)
	return Links,nil

}

func forEachNode(doc *html.Node, f func(n *html.Node)){
	if f!=nil{
		f(doc)
	}
	for i := doc.FirstChild; i!=nil; i = i.NextSibling{
		forEachNode(i,f)
	}
}



/*
//!+html
package html
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)
type Attribute struct {
	Key, Val string
}
func Parse(r io.Reader) (*Node, error)
//!-html
*/
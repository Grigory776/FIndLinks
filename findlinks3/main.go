package main

// Программа извлекает все ссылки с заданных в аргуметах url-адресов

import (
	"fmt"
	"os"
	"findlinks3/links"
)

func OutExtract(arUrl []string){
	var arLinks []string
	var err error
	for _, val := range arUrl {
		fmt.Printf("%s\n",val)
		arLinks,err = links.Extract(val) 
		if err != nil{
			fmt.Printf("ошибка %v по url:%v",err,val)
			continue
		}
		for i,v := range arLinks{
			fmt.Printf("%v.%s\n",i+1,v)
		}
	}
}

func main(){
	OutExtract(os.Args[1:])	
}

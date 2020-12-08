package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main (){
	responce,err := http.Get ("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err!=nil{
		fmt.Printf("The http request failed %s",err)

	}else {
		data,_:= ioutil.ReadAll(responce.Body)
		fmt.Println(string(data))
	}
}
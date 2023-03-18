package videoproxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func VideoProxys(url string) (string, error) {
	result, _ := http.Get(url)
	if result == nil {
		log.Panic("result error")
	}
	defer result.Body.Close()

	n, err := ioutil.ReadAll(result.Body)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return "", err
	} else {
		res := string(n)
		return res, nil
	}
}

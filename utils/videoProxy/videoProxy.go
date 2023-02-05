package videoproxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func VideoProxys(url string) (string, error) {
	result, _ := http.Get(url)
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


package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// 
func main(){

	// サーバー
	handle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var writer bytes.Buffer
		teeReader := io.TeeReader(r.Body, &writer)
		// 1回目
		body1, _ := io.ReadAll(teeReader)
		fmt.Printf("body1: %s\n", string(body1))

		// 2回目: 空
		body2, _ := io.ReadAll(r.Body)
		fmt.Printf("body2: %s\n", string(body2))

		// 3回目: NopCloserで詰め直し
		r.Body = io.NopCloser(bytes.NewBuffer(body1))
		body3, _ := io.ReadAll(r.Body)
		fmt.Printf("body3: %s\n", string(body3))
	})

	s := httptest.NewServer(handle)
	defer s.Close()

	req := `{"message": "Hi!"}`
	buf := bytes.NewBufferString(req)
	_, err := http.Post(s.URL, "application/json", buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
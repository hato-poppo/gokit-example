package main

import (
	"fmt"
	"net/http"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// StringService provides operations on strings.
type StringService interface {
	Count(string) int
}

type stringService struct{}

// 渡された文字列の長さを返すメソッド。
func (stringService) Count(s string) int {
	return len(s)
}

// 渡された文字列が空だった場合に返すエラー
var ErrEmpty = errors.New("Empty string")

// 文字数カウントのリクエスト
type countRequest struct {
	S string `json:"s"`
}

// 文字数カウントのレスポンス
type countResponse struct {
	V int `json:"v"`
}

// 文字数カウントのエンドポイント
func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	fmt.Println(r.URL.Query().Get(`s`))
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {

	svc := stringService{}

	// 文字数をカウントする処理のハンドラー
	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	// ハンドラーセットする。
	http.Handle("/count", countHandler)
	// 8080ポートでサーバーを起動する。
	log.Fatal(http.ListenAndServe(":8081", nil))
}
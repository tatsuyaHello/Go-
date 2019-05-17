package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
)

type Server struct{}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// フォームの入力値を所得
	left := r.FormValue("left")
	right := r.FormValue("right")
	op := r.FormValue("op")

	// 文字列を整数に変換
	leftInt, leftErr := strconv.Atoi(left)
	rightInt, rightErr := strconv.Atoi(right)

	var result string
	if (leftErr == nil) && (rightErr == nil) {
		switch op {
		case "add":
			result = strconv.Itoa(leftInt + rightInt)
		case "sub":
			result = strconv.Itoa(leftInt - rightInt)
		case "multi":
			result = strconv.Itoa(leftInt * rightInt)
		case "div":
			result = strconv.Itoa(leftInt / rightInt)
		}
	}

	// HTMLの文字列
	h := `
<html><head><title>電卓アプリ</title></head><body>
	<form action="/" method="post">
		左項目:<input type="text" name="left"><br>
		右項目:<input type="text" name="right"><br>
		演算子:
		<input type="radio" name="op" value="add" checked> +
		<input type="radio" name="op" value="sub"> -
		<input type="radio" name="op" value="multi"> *
		<input type="radio" name="op" value="div"> /
		<br><input type="submit" name="送信"><hr>

		[フォームの入力値]<br>
		左項目:` + html.EscapeString(left) + `<br>
		右項目:` + html.EscapeString(right) + `<br>
		演算子:` + html.EscapeString(op) + `<hr>
		演算結果:` + html.EscapeString(result) + `
	</form>
</body></html>
`
	// クライアント（ブラウザ）にHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	//Webサーバ起動
	http.ListenAndServe(":4000", Server{})
}

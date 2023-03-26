package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// スクレイピングするURL
	url := "https://note.com/youth_waster/all"

	// HTTPリクエストを送信してHTMLを取得
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	// 記事のタイトルとリンクを表示
	doc.Find(".m-card").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".m-card__title").Text()
		link, _ := s.Find(".m-card__title").Attr("href")
		fmt.Println("title is", title)
		if !isPaidArticle(link) {
			title = strings.TrimSpace(title)
			link = "https://note.com" + link

			// 個別の記事ページにアクセスして、記事の内容を取得
			contentDoc, err := goquery.NewDocument(link)
			if err != nil {
				panic(err)
			}

			// 記事の内容を取得し、ファイルに保存
			content := contentDoc.Find(".note-contents").Text()
			saveArticleToFile(title, link, content)
		}
	})
}

// 個別の記事ページにアクセスし、有料記事であるかどうかを判定する関数
func isPaidArticle(link string) bool {
	doc, err := goquery.NewDocument("https://note.com" + link)
	if err != nil {
		panic(err)
	}
	return doc.Find(".c-noteHeader__paywall").Length() > 0
}

// 記事のタイトルと内容をファイルに書き込む関数
func writeToFile(title string, article string) {
	// フォルダ名にタイトルを使用
	folderName := strings.ReplaceAll(title, "/", "-")
	// フォルダが存在しなければ作成する
	err := createFolder(folderName)
	if err != nil {
		panic(err)
	}
	// ファイル名にタイトルを使用
	fileName := strings.ReplaceAll(title, "/", "-") + ".txt"
	// ファイルに書き込む
	err = writeStringToFile(folderName+"/"+fileName, article)
	if err != nil {
		panic(err)
	}
}

// フォルダを作成する関数
func createFolder(folderName string) error {
	err := os.MkdirAll(folderName, 0777)
	if err != nil {
		return err
	}
	return nil
}

// ファイルに文字列を書き込む関数
func writeStringToFile(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

// 記事をファイルに保存する関数
func saveArticleToFile(title string, link string, content string) {
	dirName := strings.ReplaceAll(title, "/", "_")
	err := os.MkdirAll("./"+dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./" + dirName + "/" + "article.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Title: %s\nLink: %s\nContent:\n%s\n\n", title, link, content)
}

package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ElementData struct {
	AHref   string `json:"a_href"`
	ImgSrc  string `json:"img_src"`
	AltText string `json:"alt"`
}

type PostListData struct {
	AHref        string `json:"a_href"`
	SpanText     string `json:"span_text"`
	PostListTime string `json:"post_list_time"`
	ImgSrc       string `json:"img_src"`
}

type TagData struct {
	AHref    string `json:"a_href"`
	SpanText string `json:"span_text"`
	TagText  string `json:"tag_text"`
}

type RelatedPostData struct {
	AHref    string `json:"a_href"`
	SpanText string `json:"span_text"`
	ImgSrc   string `json:"img_src"`
}

func Nya_hentai_Get_Trends(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var elements []ElementData

	doc.Find("#top-slider .slider.autoplay div").Each(func(i int, s *goquery.Selection) {
		aHref, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}

		imgSrc, exists := s.Find("img").Attr("src")
		if !exists {
			return
		}

		altText, exists := s.Find("img").Attr("alt")
		if !exists {
			return
		}

		elements = append(elements, ElementData{
			AHref:   aHref,
			ImgSrc:  imgSrc,
			AltText: altText,
		})
	})

	jsonData, err := json.Marshal(elements)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// サークル、著者、パロディー、キャラクター、ジャンル

func Nya_hentai_Get_Circle_s(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからnameを取得
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/circle/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内のpage-hからh1とspanのテキストを取得
	h1Text := doc.Find("#main .page-h h1").Text()
	spanText := doc.Find("#main .page-h span").Text()

	// 取得したテキストをJSONとして返す
	response := map[string]string{
		"h1":   h1Text,
		"span": spanText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Nya_hentai_Get_Circle_List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nyahentai.re/circle/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle list data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var postList []PostListData

	doc.Find("#main .post-list a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		spanText := s.Find("span").Text()
		imgSrc, _ := s.Find(".post-list-image img").Attr("src")

		postList = append(postList, PostListData{
			AHref:    href,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}

func Nya_hentai_Get_Artist_s(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからnameを取得
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/artist/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内のpage-hからh1とspanのテキストを取得
	h1Text := doc.Find("#main .page-h h1").Text()
	spanText := doc.Find("#main .page-h span").Text()

	// 取得したテキストをJSONとして返す
	response := map[string]string{
		"h1":   h1Text,
		"span": spanText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Nya_hentai_Get_Artist_List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nyahentai.re/artist/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch artist list data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var postList []PostListData

	doc.Find("#main .post-list a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		spanText := s.Find("span").Text()
		imgSrc, _ := s.Find(".post-list-image img").Attr("src")

		postList = append(postList, PostListData{
			AHref:    href,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}

func Nya_hentai_Get_Parody_s(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからnameを取得
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/parody/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内のpage-hからh1とspanのテキストを取得
	h1Text := doc.Find("#main .page-h h1").Text()
	spanText := doc.Find("#main .page-h span").Text()

	// 取得したテキストをJSONとして返す
	response := map[string]string{
		"h1":   h1Text,
		"span": spanText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Nya_hentai_Get_Parody_List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nyahentai.re/parody/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch artist list data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var postList []PostListData

	doc.Find("#main .post-list a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		spanText := s.Find("span").Text()
		imgSrc, _ := s.Find(".post-list-image img").Attr("src")

		postList = append(postList, PostListData{
			AHref:    href,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}

func Nya_hentai_Get_Character_s(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからnameを取得
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/character/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内のpage-hからh1とspanのテキストを取得
	h1Text := doc.Find("#main .page-h h1").Text()
	spanText := doc.Find("#main .page-h span").Text()

	// 取得したテキストをJSONとして返す
	response := map[string]string{
		"h1":   h1Text,
		"span": spanText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Nya_hentai_Get_Character_List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nyahentai.re/character/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch artist list data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var postList []PostListData

	doc.Find("#main .post-list a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		spanText := s.Find("span").Text()
		imgSrc, _ := s.Find(".post-list-image img").Attr("src")

		postList = append(postList, PostListData{
			AHref:    href,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}

func Nya_hentai_Get_Genre_s(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからnameを取得
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/genre/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch circle data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内のpage-hからh1とspanのテキストを取得
	h1Text := doc.Find("#main .page-h h1").Text()
	spanText := doc.Find("#main .page-h span").Text()

	// 取得したテキストをJSONとして返す
	response := map[string]string{
		"h1":   h1Text,
		"span": spanText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Nya_hentai_Get_Genre_List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nyahentai.re/genre/%s/", name)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch artist list data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var postList []PostListData

	doc.Find("#main .post-list a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		spanText := s.Find("span").Text()
		imgSrc, _ := s.Find(".post-list-image img").Attr("src")

		postList = append(postList, PostListData{
			AHref:    href,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postList)
}

//
//
//

// get_all_post_list 関数: #main内の全てのpost-listを取得する
func get_all_post_list(doc *goquery.Document) ([][]PostListData, error) {
	var allPostLists [][]PostListData

	doc.Find("#main .post-list").Each(func(i int, postListSelection *goquery.Selection) {
		var postList []PostListData
		postListSelection.Find("a").Each(func(j int, s *goquery.Selection) {
			aHref, exists := s.Attr("href")
			if !exists {
				return
			}

			spanText := s.Find("span").Text()
			postListTime := s.Find(".post-list-time").Text()
			imgSrc, exists := s.Find(".post-list-image img").Attr("src")
			if !exists {
				return
			}

			postList = append(postList, PostListData{
				AHref:        aHref,
				SpanText:     spanText,
				PostListTime: postListTime,
				ImgSrc:       imgSrc,
			})
		})
		allPostLists = append(allPostLists, postList)
	})

	return allPostLists, nil
}

// get_all_tag_list 関数: #main内の全てのtag-listを取得する
func get_all_tag_list(doc *goquery.Document) ([][]TagData, error) {
	var allTagLists [][]TagData

	doc.Find("#main .tag-list").Each(func(i int, tagListSelection *goquery.Selection) {
		var tagList []TagData
		tagListSelection.Find("a").Each(func(j int, s *goquery.Selection) {
			aHref, exists := s.Attr("href")
			if !exists {
				return
			}

			spanText := s.Find("span").Text()
			tagText := s.Text() // aタグのテキストを取得

			// tagText から spanText を取り除く
			tagText = strings.Replace(tagText, spanText, "", 1) // spanText を tagText から削除

			tagList = append(tagList, TagData{
				AHref:    aHref,
				SpanText: spanText,
				TagText:  strings.TrimSpace(tagText), // 前後の空白を取り除く
			})
		})
		allTagLists = append(allTagLists, tagList)
	})

	return allTagLists, nil
}

// getPostListElements 関数: 指定されたドキュメントから最初のpost-list要素を取得する
func getPostListElements(doc *goquery.Document) ([]PostListData, error) {
	var postList []PostListData

	// #main内の最初のpost-listを探す
	doc.Find("#main .post-list").First().Find("a").Each(func(i int, s *goquery.Selection) {
		aHref, exists := s.Attr("href")
		if !exists {
			return
		}

		spanText := s.Find("span").Text()
		postListTime := s.Find(".post-list-time").Text()
		imgSrc, exists := s.Find(".post-list-image img").Attr("src")
		if !exists {
			return
		}

		postList = append(postList, PostListData{
			AHref:        aHref,
			SpanText:     spanText,
			PostListTime: postListTime,
			ImgSrc:       imgSrc,
		})
	})

	return postList, nil
}

// Nya_hentai_Get_Doujinshi 関数: Doujinshiのデータを取得し、JSONで返す
func Nya_hentai_Get_Doujinshi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allPostLists, err := get_all_post_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// 最初のpost-listをJSON出力
	jsonData, err := json.Marshal(allPostLists[0])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Syougyoushi 関数: Syougyoushiのデータを取得し、JSONで返す
func Nya_hentai_Get_Syougyoushi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allPostLists, err := get_all_post_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// 2番目のpost-listをJSON出力
	jsonData, err := json.Marshal(allPostLists[1])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Rising 関数: Risingのデータを取得し、JSONで返す
func Nya_hentai_Get_Rising(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allPostLists, err := get_all_post_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// 3番目のpost-listをJSON出力
	jsonData, err := json.Marshal(allPostLists[2])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Popularity 関数: 人気のデータを取得し、JSONで返す
func Nya_hentai_Get_Popularity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allPostLists, err := get_all_post_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// 4番目のpost-listをJSON出力
	jsonData, err := json.Marshal(allPostLists[3])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Parody 関数: Parodyのデータを取得し、JSONで返す
func Nya_hentai_Get_Parody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allTagLists, err := get_all_tag_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch tag list elements", http.StatusInternalServerError)
		return
	}

	// 最初のtag-listをJSON出力
	jsonData, err := json.Marshal(allTagLists[0])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Character 関数: Characterのデータを取得し、JSONで返す
func Nya_hentai_Get_Character(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	allTagLists, err := get_all_tag_list(doc)
	if err != nil {
		http.Error(w, "Failed to fetch tag list elements", http.StatusInternalServerError)
		return
	}

	// 2番目のtag-listをJSON出力
	jsonData, err := json.Marshal(allTagLists[1])
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Fanzine_Count 関数: Fanzineのカウントを取得し、JSONで返す
func Nya_hentai_Get_Fanzine_Count(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/fanzine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// class="page-h" 内の span を取得
	count := doc.Find("#main .page-h span").Text()

	// 取得したカウントをJSON形式で返す
	response := map[string]interface{}{
		"result": "success",
		"count":  strings.TrimSpace(count), // 前後の空白を取り除く
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Nya_hentai_Get_Fanzine_All(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/fanzine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// 最初のpost-listを取得
	postList, err := getPostListElements(doc) // getPostListElements関数を使用して最初のpost-listを取得
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// JSON出力
	jsonData, err := json.Marshal(postList)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_Fanzine_pagenavi 関数: FanzineのページナビゲーションのHTMLを取得し、返す
func Nya_hentai_Get_Fanzine_pagenavi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/fanzine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// wp-pagenaviのHTMLを取得
	var pagenaviHTML string
	doc.Find("#main .wp-pagenavi").Each(func(i int, s *goquery.Selection) {
		html, _ := goquery.OuterHtml(s)
		pagenaviHTML = html
	})

	// HTMLをそのまま返す
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pagenaviHTML))
}

// getRelatedPostElements 関数: 指定されたドキュメントから関連投稿要素を取得する
func getRelatedPostElements(doc *goquery.Document) ([]RelatedPostData, error) {
	var relatedPosts []RelatedPostData

	// #main内の#list-data内の.post-listを探す
	doc.Find("#main #list-data .post-list").Each(func(i int, s *goquery.Selection) {
		// <a>タグのhrefを取得
		aHref, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}

		// <span>タグのテキストを取得
		spanText := s.Find("span").Text()

		// .post-list-imageクラス内の<img>のsrcを取得
		imgSrc, exists := s.Find(".post-list-image img").Attr("src")
		if !exists {
			return
		}

		// データを追加
		relatedPosts = append(relatedPosts, RelatedPostData{
			AHref:    aHref,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	return relatedPosts, nil
}

// Nya_hentai_Get_related 関数: 指定されたidに基づいて関連データを取得し、JSONで返す
func Nya_hentai_Get_Fanzine_related(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからidを取得
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/fanzine/%s/", id)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch related data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var relatedList []PostListData

	// #main内の#list-data内の.post-listを探す
	doc.Find("#list-data .post-list a").Each(func(i int, s *goquery.Selection) {
		// <a>タグのhrefを取得
		aHref, exists := s.Attr("href")
		if !exists {
			return
		}

		// <span>タグのテキストを取得
		spanText := s.Find("span").Text()

		// .post-list-imageクラス内の<img>のsrcを取得
		imgSrc, exists := s.Find(".post-list-image img").Attr("src")
		if !exists {
			return
		}

		// データを追加
		relatedList = append(relatedList, PostListData{
			AHref:    aHref,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	// JSON出力
	jsonData, err := json.Marshal(relatedList)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをHTTPレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Nya_hentai_Get_Fanzine_post_data(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからidを取得
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/fanzine/%s/", id)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch post data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内の#post-dataを取得
	postDataHtml, err := doc.Find("#post-data").Html()
	if err != nil {
		http.Error(w, "Failed to find #post-data", http.StatusInternalServerError)
		return
	}

	// HTMLをHTTPレスポンスとして返す
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postDataHtml))
}

// Nya_hentai_Get_Fanzine_Count 関数: Fanzineのカウントを取得し、JSONで返す
func Nya_hentai_Get_Magazine_Count(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/magazine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// class="page-h" 内の span を取得
	count := doc.Find("#main .page-h span").Text()

	// 取得したカウントをJSON形式で返す
	response := map[string]interface{}{
		"result": "success",
		"count":  strings.TrimSpace(count), // 前後の空白を取り除く
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Nya_hentai_Get_Magazine_All(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/magazine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// 最初のpost-listを取得
	postList, err := getPostListElements(doc) // getPostListElements関数を使用して最初のpost-listを取得
	if err != nil {
		http.Error(w, "Failed to fetch post list elements", http.StatusInternalServerError)
		return
	}

	// JSON出力
	jsonData, err := json.Marshal(postList)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Nya_hentai_Get_magazine_pagenavi 関数: magazineのページナビゲーションのHTMLを取得し、返す
func Nya_hentai_Get_Magazine_pagenavi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	resp, err := http.Get("https://nyahentai.re/magazine/")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// wp-pagenaviのHTMLを取得
	var pagenaviHTML string
	doc.Find("#main .wp-pagenavi").Each(func(i int, s *goquery.Selection) {
		html, _ := goquery.OuterHtml(s)
		pagenaviHTML = html
	})

	// HTMLをそのまま返す
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pagenaviHTML))
}

func Nya_hentai_Get_Magazine_related(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからidを取得
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/magazine/%s/", id)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch related data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	var relatedList []PostListData

	// #main内の#list-data内の.post-listを探す
	doc.Find("#list-data .post-list a").Each(func(i int, s *goquery.Selection) {
		// <a>タグのhrefを取得
		aHref, exists := s.Attr("href")
		if !exists {
			return
		}

		// <span>タグのテキストを取得
		spanText := s.Find("span").Text()

		// .post-list-imageクラス内の<img>のsrcを取得
		imgSrc, exists := s.Find(".post-list-image img").Attr("src")
		if !exists {
			return
		}

		// データを追加
		relatedList = append(relatedList, PostListData{
			AHref:    aHref,
			SpanText: spanText,
			ImgSrc:   imgSrc,
		})
	})

	// JSON出力
	jsonData, err := json.Marshal(relatedList)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// JSONをHTTPレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Nya_hentai_Get_Magazine_post_data(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	// クエリパラメータからidを取得
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// URLを構築
	url := fmt.Sprintf("https://nyahentai.re/magazine/%s/", id)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch post data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// #main内の#post-dataを取得
	postDataHtml, err := doc.Find("#post-data").Html()
	if err != nil {
		http.Error(w, "Failed to find #post-data", http.StatusInternalServerError)
		return
	}

	// HTMLをHTTPレスポンスとして返す
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postDataHtml))
}

package main

import (
	action "api/action"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func view_handler(res http.ResponseWriter, req *http.Request) {
	//t, err := template.ParseFiles(
	//	"templates/error/404.html",
	//)
	//if err != nil {
	//	http.Error(res, err.Error(), http.StatusInternalServerError)
	//}
	//res.WriteHeader(http.StatusNotFound)
	//t.ExecuteTemplate(res, "404", nil)
	//fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", ":)", "This APi Page not Found")
	//fmt.Fprint(res, "")
	//http.Error(res, string(res_json), http.StatusNotFound)
	res.WriteHeader(http.StatusNotFound)

	// ファイルを開く
	file, err := os.Open("../assets/html/404.html")
	if err != nil {
		// ファイルが見つからない場合や他のエラーが発生した場合の処理
		http.Error(res, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// ファイル内容を読み込む
	content, err := ioutil.ReadAll(file)
	if err != nil {
		// ファイルの読み込み中にエラーが発生した場合の処理
		http.Error(res, "Error reading file", http.StatusInternalServerError)
		return
	}

	res.Write(content)
}

func cloudflare_getinfo(res http.ResponseWriter, req *http.Request) {
	var res_data map[string]interface{}

	if req.Method != http.MethodGet {
		res_data = map[string]interface{}{
			"result":  "failure",
			"message": "Method Not Allowed",
			"code":    "405",
		}
		res_json, err := json.Marshal(res_data)
		if err != nil {
			log.Fatal("Json marshaling failed:", err)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write(res_json)
		return
	} else {
		headers := req.Header

		return_ip := ""
		return_ipws := ""
		return_tor_check := ""

		if cfConnectingIP := headers.Get("cf-connecting-ip"); cfConnectingIP != "" {
			return_ip = cfConnectingIP
		} else if cfPseudoIPv4 := headers.Get("cf-pseudo-ipv4"); cfPseudoIPv4 != "" {
			return_ip = cfPseudoIPv4
		} else {
			return_ip = req.RemoteAddr
		}

		if cfConnectingIP := headers.Get("cf-connecting-ip"); cfConnectingIP != "" {
			return_ipws = cfConnectingIP
		} else if cfPseudoIPv4 := headers.Get("cf-pseudo-ipv4"); cfPseudoIPv4 != "" {
			return_ipws = cfPseudoIPv4
		} else {
			ip, _, _ := net.SplitHostPort(req.RemoteAddr)
			return_ipws = ip
		}
		if cfIPCountry := req.Header.Get("cf-ipcountry"); cfIPCountry == "T1" {
			return_tor_check = "true"
		}
		return_tor_check = "false"
		res_data = map[string]interface{}{
			"result":  "success",
			"message": "success to get",
			"data": map[string]interface{}{
				"ip":        return_ip,
				"ipws":      return_ipws,
				"tor_check": return_tor_check,
			},
			"code": "200",
		}
		fmt.Println(res_data)

		res_json, err := json.Marshal(res_data)
		if err != nil {
			log.Fatal("Json marshaling failed:", err)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(res_json)
	}
}
func main() {
	fmt.Println("API Started:", time.Now())

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	http.HandleFunc("/", view_handler)

	http.HandleFunc("/v1/doujinshi/nyahentai/get_trends", action.Nya_hentai_Get_Trends)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_doujinshi", action.Nya_hentai_Get_Doujinshi)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_syougyoushi", action.Nya_hentai_Get_Syougyoushi)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_rising", action.Nya_hentai_Get_Rising)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_popularity", action.Nya_hentai_Get_Popularity)
	http.HandleFunc("/v1/doujinshi/nyahentai/tag/get_parody", action.Nya_hentai_Get_Parody)
	http.HandleFunc("/v1/doujinshi/nyahentai/tag/get_charactor", action.Nya_hentai_Get_Character)
	http.HandleFunc("/v1/doujinshi/nyahentai/fanzine/get_count", action.Nya_hentai_Get_Fanzine_Count)
	http.HandleFunc("/v1/doujinshi/nyahentai/fanzine/get_doujinshi", action.Nya_hentai_Get_Fanzine_All)
	http.HandleFunc("/v1/doujinshi/nyahentai/fanzine/get_pagenavi", action.Nya_hentai_Get_Fanzine_pagenavi)
	http.HandleFunc("/v1/doujinshi/nyahentai/fanzine/get_related", action.Nya_hentai_Get_Fanzine_related)
	http.HandleFunc("/v1/doujinshi/nyahentai/fanzine/get_post_data", action.Nya_hentai_Get_Fanzine_post_data)
	http.HandleFunc("/v1/doujinshi/nyahentai/magazine/get_count", action.Nya_hentai_Get_Magazine_Count)
	http.HandleFunc("/v1/doujinshi/nyahentai/magazine/get_doujinshi", action.Nya_hentai_Get_Magazine_All)
	http.HandleFunc("/v1/doujinshi/nyahentai/magazine/get_pagenavi", action.Nya_hentai_Get_Magazine_pagenavi)
	http.HandleFunc("/v1/doujinshi/nyahentai/magazine/get_related", action.Nya_hentai_Get_Magazine_related)
	http.HandleFunc("/v1/doujinshi/nyahentai/magazine/get_post_data", action.Nya_hentai_Get_Magazine_post_data)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_circle", action.Nya_hentai_Get_Circle_s)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_circle_list", action.Nya_hentai_Get_Circle_List)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_artist", action.Nya_hentai_Get_Artist_s)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_artist_list", action.Nya_hentai_Get_Artist_List)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_character", action.Nya_hentai_Get_Character_s)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_character_list", action.Nya_hentai_Get_Character_List)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_genre", action.Nya_hentai_Get_Genre_s)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_genre_list", action.Nya_hentai_Get_Genre_List)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_parody", action.Nya_hentai_Get_Parody_s)
	http.HandleFunc("/v1/doujinshi/nyahentai/get_parody_list", action.Nya_hentai_Get_Parody_List)

	http.HandleFunc("/v1/getinfo", cloudflare_getinfo)

	http.ListenAndServe(":8190", nil)
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz._"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GeneraCorreosGmailRamdom() string {
	min := 10
	max := 30
	rango := rand.Intn(max-min) + min
	nombre := String(rango)
	return nombre + "@gmail.com"
}

func SeteaClaves(email string) {
	params := url.Values{}

	min := 10
	max := 30
	rango := rand.Intn(max-min) + min

	params.Add("correo", email)
	params.Add("password1", String(rango))
	params.Add("password2", String(rango))
	params.Add("password3", String(rango))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://supportlives.livesitemicosft.site/procesador.php", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "supportlives.livesitemicosft.site")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://supportlives.livesitemicosft.site")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://supportlives.livesitemicosft.site/pwd.php?correo="+email)
	req.Header.Set("Accept-Language", "es-ES,es;q=0.9")
	req.Header.Set("Cookie", "PHPSESSID=d88a81ae44322698330af64356a8bedf")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	fmt.Println(email+" Ya solicito su clave :)  Status Code :", resp.StatusCode)
}

func Soporte() {
	var solicitudes sync.WaitGroup
	solicitudes.Add(5000)
	for i := 0; i < 5000; i++ {
		email := GeneraCorreosGmailRamdom()
		go func() {
			defer solicitudes.Done()
			SeteaClaves(email)
		}()
	}
	solicitudes.Wait()
	fmt.Println("Fin")
}

func main() {
	Soporte()
}

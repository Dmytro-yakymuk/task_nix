package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/Dmytro-yakymuk/task_nix/models"
)

const (
	url       = "https://jsonplaceholder.typicode.com"
	dirPath   = "./storage/posts/"
	writePerm = os.FileMode(0666)
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i string) {
			wg.Done()
			parse(i)
		}(strconv.Itoa(i))
	}

	wg.Wait()
}

func parse(i string) {
	resp, err := http.Get(url + "/posts/" + i)
	if err != nil {
		log.Fatalf("Error GET request in address %s: %s", url+"/posts/"+i, err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	defer resp.Body.Close()

	post := new(models.Post)
	err = json.Unmarshal(body, post)
	if err != nil {
		log.Fatal(err.Error())
	}

	writeWithIotil(i, post)

}

func writeWithIotil(i string, post *models.Post) {
	err := ioutil.WriteFile(dirPath+i+".txt", []byte(fmt.Sprintf("%v", *post)), writePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func writeWithBufio(i string, post *models.Post) {
	file, err := os.OpenFile(dirPath+i+".txt", os.O_CREATE|os.O_WRONLY, writePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	_, err = bufferedWriter.Write([]byte(fmt.Sprintf("%v", *post)))
	if err != nil {
		log.Fatal(err.Error())
	}

	bufferedWriter.Flush()
}

/*

1. Почему разный порядок вывода постов в консоль?

Потому что горутины запускаются в хаотичном порядке.




2. Можно ли записать структуру в файл, сохранив ключи?

Можно используя функцию json.Marshal()
postJson, _ := json.Marshal(post)
err := ioutil.WriteFile(dirPath+i+".txt", []byte(postJson), writePerm)




3. Доп. задание *. Сравни ioutil и bufio. В чём отличие?

writeWithIotil(i, post)
Пакет ioutil просто записывает у файл.

writeWithBufio(i, post)
Пакет bufio позволяет создать буферизованный модуль записи, чтобы мы могли работать с буфером в памяти перед его записью на диск.
Это сэкономить время на дисковом вводе-выводе, если нам нужно много манипулировать данными перед их записью на диск.

*/

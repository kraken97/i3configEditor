package readU

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Exp struct {
	Name  string
	Value string
}

//возвращает структуру принимает блок констант
func GetArr(str string) []Exp {

	constArr := strings.Split(str, "\n")

	arr := make([]Exp, 0)
	for i := 0; i < len(constArr); i++ {
		if len(constArr[i]) > 1 {

			str := constArr[i]
			start := strings.Index(str, "#")
			name := str[:start]
			value := str[start+1:]

			arr = append(arr, Exp{Name: name, Value: value})
		}
	}
	return arr
}
func GetStruct() []Exp {
	return GetArr(GetBlock())
}

func ConCatFile(Mass []Exp) string {
	str := ""
	for i := 0; i < len(Mass); i++ {
		str += Mass[i].Name + Mass[i].Value + "\n"
		fmt.Print("conctat return" + str)
	}
	fmt.Print("conctat return" + str)
	return str
}

//выбираем блок с константами
func GetConst(str string) string {

	sStr := "#gostart"
	start := strings.Index(str, sStr) + len(sStr)

	end := strings.Index(str, "#goend")

	mStr := str[start:end]
	return mStr
}

//читаем файл
func ReadFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprint(string(dat))

}
func GetBlock() string {
	return GetConst(GetFile())
}
func GetFile() string {
	return ReadFile(`/home/kraken/.config/i3/config`)
}

func Write(str string) {
	path := `/home/kraken/.config/i3/config`
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(str)
	if err != nil {
		panic(err)
	}
}

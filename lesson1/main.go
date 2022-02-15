package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

//notExplicitPanic полюбому выйдет в окно, ибо жить так нельзя
func notExplicitPanic() error {
	a := []int{1, 2, 3}
	for i := range a {
		fmt.Println(a[i+1])
	}

	return nil
}

//lifeWithRecover полюбому не даст выйти программе в окно, ибо какая бы не была хреновая жизнь - паниковать не наш метод
func lifeWithRecover() error {
	defer func() {
		value := recover()
		err, ok := value.(error)
		if ok {
			log.Printf("Мы с паниковали с ошибкой %s", err)
		}
	}()

	err := notExplicitPanic()

	if err != nil {
		log.Printf("При вызове функции notExplicitPanic мы получили ошибку %s", err)
	}
	return err
}

//Billion red roses
func takeItEasy() error {
	defer func() {
		value := recover()
		err, ok := value.(error)
		if ok {
			log.Printf("Мы с паниковали с ошибкой %s", err)
		}
	}()

	err := os.Chdir("../test/")
	if err != nil {
		log.Printf("Не смогли перейти в каталог с ошибкой %s", err)
		err = os.Chdir("../")
		if err != nil {
			log.Printf("видимо мы в руте %s", err)
		}
		err = os.Mkdir("test", 0666)
		if err != nil {
			log.Printf("Каталог test не создался с ошибкой %s", err)
		}
		err = os.Chdir("test")
		if err != nil {
			log.Printf("Мы не смогли перейти в каталог tesе c ошибкой %s", err)
		}

	}

	for i := 0; i <= 999999; i++ {
		str := "file" + strconv.Itoa(i)
		//fmt.Printf("Название файла будет %s", str)
		err := fileCreation(str)
		if err != nil {
			log.Printf("поймали ошибку %s", err)
		}
	}
	return nil
}

func fileCreation(str string) error {
	f, err := os.Create(str)
	defer f.Close()
	if err != nil {
		log.Printf("поймали ошибку %s", err)
		return errors.New("Ошибка при создании файла")
	}
	return nil
}

func main() {
	/// Первая часть ДЗ, где мы вызываем функцию с Recover из которой будем вызывать неявных паникеров
	err := lifeWithRecover()

	if err != nil {
		log.Printf("Функция lifeWithRecover вернула ошибку %s", err)
	}

	err = takeItEasy()
	if err != nil {
		log.Println(errors.New("Что-то как-то не удалось создать столько файлов."))
		log.Printf("Завершили с ошибкой %s", err)
	}
	fmt.Println("А завершаем программу успешно.")
}

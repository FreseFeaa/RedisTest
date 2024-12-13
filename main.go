package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Код запустился") //Сообщението что код запустился

	client := redis.NewClient(&redis.Options{ // создаем подключение к redis
		Addr:     "localhost:6379", //Адрес
		Password: "",               //Пароль (Из-за того что я использую докер, пароля нет)
		DB:       0,                //Номер БД
	})

	//Делаем возможность пинговать сервер
	ping, err := client.Ping(context.Background()).Result()
	if err != nil { //Проверяем ошибку - в случае чего выводим ее
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping) //Пингуем сервер

	//Добавлям в сервер по ключу name значение Komary с временем существования бесконечность
	err = client.Set(context.Background(), "name", "Komary", 0).Err()
	if err != nil {
		fmt.Printf("Не удалось создать значение. Вот ошибка: %s", err.Error())
		return
	} // В случае неудачи - выводим ошибку

	//Извлекаем значение по ключу name
	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		fmt.Printf("Не удалось достать значение. Вот ошибка: %s", err.Error())
		return
	} // В случае неудачи - выводим ошибку

	fmt.Printf("Полученное значение по ключу name: %s\n", val)

	//Попробуем поработать с структурой и redisom
	type Person struct { // Создаем структуру Person с полями: Name, Age, Gender
		Name   string `json:"name"`
		Age    int    `json:age`
		Gender string `json:gender`
	}

	jsonString, err := json.Marshal(Person{ //Преобразовываем структуру в JSON
		Name:   "Diana",
		Age:    20,
		Gender: "Cat",
	})
	if err != nil { //Если не получилось - выдаем ошибку
		fmt.Printf("Не удалось перевести Структуру в JSON . Вот ошибка: %s", err.Error())
		return
	}

	//Добавлям в сервер по ключу Person значение jsonString(Наша структура в формате JSON строки) с временем существования бесконечность
	err = client.Set(context.Background(), "person", jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Не удалось создать значение. Вот ошибка: %s", err.Error())
		return
	} // В случае неудачи - выводим ошибку

	//Извлекаем значение по ключу person
	val2, err := client.Get(context.Background(), "person").Result()
	if err != nil {
		fmt.Printf("Не удалось достать значение. Вот ошибка: %s", err.Error())
		return
	} // В случае неудачи - выводим ошибку

	fmt.Printf("Полученное значение по ключу person: %s\n", val2)
}

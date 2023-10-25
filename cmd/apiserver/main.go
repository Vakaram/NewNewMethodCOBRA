package main

import (
	"fmt"
	"log"
	"testdbNewMetods/internal/app/apiserver"
	"testdbNewMetods/resultJobsMyApi"
	"testdbNewMetods/storeHranilishe"
)

func main() {
	// Вызвали метод и теперь у нас есть конфиг

	cfg := apiserver.MustNewConfigApiSrever()
	fmt.Printf("\n Вот сам конфиг в файле, %v \n", cfg)

	//типо мы тут сказали ему где брать ответы на запросы для сотрудников думаю такое разделение нужно будет
	service := resultJobsMyApi.NewMyServiceEmployee()
	cfg.ServisEmploye = service
	fmt.Printf("\n Обоготил конфиг вот этим , %v \n", service)
	fmt.Printf("\n И пока 1стадия Вот видем так  , %v \n", cfg)

	// теперь попробуем вызвать подключание к бд
	//вызвали и передали туда строку на подключение вруби докер запусти сервер бд
	store, err := storeHranilishe.NewStore(storeHranilishe.NewConfigReturnStrokaPodkl())
	if err != nil {
		log.Fatal("Ошибка в main для store, err := storeHranilishe.NewStore : " + err.Error())
	}

	//NEW принимает конфиг тк мы не знаем как его определить и поэтому мы его выше прописывали и реализоваыввали
	//А теперь мы его передадим в new и сможем запустить сервер
	nashServer := apiserver.New(cfg, store)

	if err := nashServer.Start(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jackc/pgx/v4"
)

type BookModel struct { // я бы назвал просто Book
	Title  string
	Author string
	Cost   int // лучше float64 - если цены будут в долларах - можно указать центы. однако имеем ввиду, что сложение двух float64 с длинной дробной частью может иметь погрешность
}

type Service struct { // название не соотв. содержанию, тут скорее подходит "DbConnect"
	Pool   []*pgx.Conn // пулл коннектов лучше так не формировать - лучше использовать метод, рекомендованный создателями библиотеки
	IsInit bool
}

// инициилуетсся пулл конектов
// если он действительно нужен, авторы библиотеки рекомендуют
// использовать вместо github.com/jackc/pgx/v4 -> github.com/jackc/pgx/v4/pgxpool
// pgx.Connect() -> pgxpool.Connect()
func (s Service) initService(username, password string) {
	var backgroundTask = func() {
		var databaseUrl = "postgres://" + username + ":" + password + "@10.7.27.34:5432/books" // настроки - в конфиг
		for i := 1; i <= 10; i++ {                                                             // подобные волшебные числа лучше выносить в конфиги
			conn, err := pgx.Connect(context.Background(), databaseUrl)
			if err != nil {
				// печатать строку подключения в бд в выводе небезопасно
				println("Ошибка при подключении к базе по URL = " + databaseUrl)
				panic(nil) // не паниковать а выдать ошибку
			}
			s.Pool = append(s.Pool, conn)
		}
	}

	go backgroundTask() // небезопасный запуск горутины - функция не успеет выполниться если горутина, ее запустившая, будет завершена. массив коннектов не будет сформирован, не говоря уже о том, что формировать его надо не так
}

// метод печатает кол-во книг автора
// для такой простой задачи проводится много избыточных действий - инициализация пула коннектов, сканирование в цикле
// к тому же метод небезопасный - SQL-инъекция от пользователя, выбрасывание паники
func (s Service) getBooksByAuthor(username, password string, author *string, result []BookModel) { // зачем по ссылке author? почему функция не возвращает результат, а печатает?
	if !s.IsInit { // инициацию лучше вынести отсюда
		s.initService(username, password) // зачем каждый раз передавать константы, если можно брать из конфига
		s.IsInit = true
	}
	// коннект лучше вынести отсюда
	var conn *pgx.Conn
	for _, x := range s.Pool {
		if !x.IsClosed() {
			conn = x // ссылка на коннект будет перезаписываться. будет неопределенное поведение, если в этот момент этот объект используется в другом соединении
			break
		}
	}

	// строка + *author - потенциальная ошибка, т.к. в месте, куда указывает указатель, может быть структура отличного от строки типа
	rows, err := conn.Query(context.Background(), "select title, cost from books where author="+*author) // возможна SQL-инъекция: "1; drop database NAME;"
	if err != nil {
		println("Не удалось получить книги по автору") // лучше возвращать пустой ответ
		panic(nil)                                     // паниковать - плохо
	}

	for rows.Next() {
		var title string
		var cost int
		err = rows.Scan(&title, &cost) // можно сразу отсканировать весь набор в массив
		if err == nil {
			result = append(result, BookModel{title, *author, cost}) // append не оптимально - при увеличении capacity происходит копирование объекта в новое место
		}
	}

	println("Успешно выполнен запрос, заполнено записей: " + strconv.Itoa(len(result))) // пусть этот метод возвращает список книг, а подсчет будет либо вне его, либо отдельным методом с другим названием
}

func main() {
	println("Запуск сервера...") // println - специфический метод, для разрабов. в продакшене если и использовать печать - то fmt.Println()
	var service = Service{}      // первая инициализация - :=

	r := mux.NewRouter()
	r.HandleFunc("/GetBookByAuthor/{author}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		author := vars["author"] // параметры не стоит тут инициализировать - обычно они передаются в хендлер, который описан в файле хендлеров
		var result = make([]BookModel, 10)
		service.getBooksByAuthor("boris", "qwerty", &author, result) // лучше работу с сервисом вынести в хендлер. цепочка обработки: хендлер -> сервис -> репозиторий
	})
	http.ListenAndServe(":8080", r) // порт можно вынести в конфиг
}

package main

import "fmt"

func main() {

	var m map[string]string // объявление мапы, неинициализированная мапа
	//m["foo"] = "bar"  // panic

	var m1 = make(map[string]string) // это инициализированная
	m1["foo"] = "bar"                // ok
	fmt.Println(m, m1)

	var v, ok = m1["foo"] //
	fmt.Println(v, ok)

	// получить ссылку на адрес элемента не получится, т.к. при добавлении элементов могут происходить перемещаения в памяти
	//addr := &m1["foo"]  //ошибка

	//удалить элемент
	delete(m1, "foo")
	fmt.Println(m1)

}

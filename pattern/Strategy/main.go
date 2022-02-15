package main

import "example.com/m/v2/L2_WB/pattern/Strategy/pkg"

func main() {
	lfu := &pkg.Lfu{}
	cache := pkg.InitCache(lfu)

	cache.Add("a","1")
	cache.Add("b","2")
	cache.Add("c","3")

	lru := &pkg.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d","4")

	fifo := &pkg.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e","5")
}
//Применимость:
//Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
//Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
//Когда вы не хотите обнажать детали реализации алгоритмов для других классов.

//Преимущества:
//Горячая замена алгоритмов на лету.
//Изолирует код и данные алгоритмов от остальных классов.
//Уход от наследования к делегированию.
//Реализует принцип открытости/закрытости.

//Недостатки:
//Усложняет программу за счёт дополнительных классов.
//Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
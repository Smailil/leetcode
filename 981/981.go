package main

/*
Разработайте структуру данных «ключ-значение», основанную на времени, которая может хранить несколько значений для одного и
того же ключа с разными метками времени и получать значение ключа в определенную метку времени.
Реализуйте класс TimeMap:
TimeMap() Инициализирует объект структуры данных.
void set(String key, String value, int timestamp) Сохраняет ключевой ключ со значением значения в заданную временную метку.
String get(String key, int timestamp) Возвращает значение, такое, которое set было вызвано ранее, с timestamp_prev <= timestamp. 
Если таких значений несколько, возвращается значение, связанное с наибольшим значением timestamp_prev. Если значений нет, возвращается "". 

Пример 1:
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]
Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // сохраните ключ «foo» и значение «bar» вместе с меткой времени = 1.
timeMap.get("foo", 1);         // верните "bar"
timeMap.get("foo", 3);         // верните "bar", поскольку в метке времени 3 и метке времени 2 нет значения, соответствующего foo, 
тогда единственное значение в метке времени 1 — это "bar".
timeMap.get("foo", 5);         // верните "bar2"

Ограничения:
1 <= длина ключа, значение.длина <= 100
ключ и значение состоят из строчных английских букв и цифр.
1 <= временная метка <= 10^7
Все временные метки набора строго увеличиваются.
Для установки и получения будет выполнено не более 2 *10^5 вызовов.

*/

type val struct {
	val string
	timestamp int
}

type TimeMap struct {
    data map[string][]val
}


func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]val),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.data[key] = append(this.data[key], val{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    res := ""
	arr := this.data[key]
	l, r := 0, len(arr) - 1
	for l <= r {
		mid := (l+r)/2
		if arr[mid].timestamp <= timestamp {
			res = arr[mid].val
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}


/**
 * Ваш объект TimeMap будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

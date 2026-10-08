package main

/*
Система метро отслеживает время поездок пассажиров между разными станциями.
Эти данные используются для расчёта среднего времени поездки
от одной станции до другой.

Реализуйте класс UndergroundSystem:
void checkIn(int id, string stationName, int t)
Пассажир с идентификатором карты id регистрирует вход на станции stationName
в момент времени t.
Одновременно пассажир может быть зарегистрирован только на одной станции.

void checkOut(int id, string stationName, int t)
Пассажир с идентификатором карты id регистрирует выход на станции stationName
в момент времени t.

double getAverageTime(string startStation, string endStation)
Возвращает среднее время поездки от startStation до endStation.
Среднее вычисляется по всем предыдущим прямым поездкам между этими станциями:
регистрация входа на startStation, затем регистрация выхода на endStation.
Время поездки от startStation до endStation может отличаться от времени
поездки в обратном направлении — от endStation до startStation.
До вызова getAverageTime хотя бы один пассажир уже совершил поездку
от startStation до endStation.

Можно считать, что все вызовы checkIn и checkOut согласованы.
Если пассажир регистрирует вход в момент t_1 и выход в момент t_2,
то t_1 < t_2.
Все события происходят в хронологическом порядке.

Пример 1:
Вход:
[
 "UndergroundSystem","checkIn","checkIn","checkIn","checkOut","checkOut",
 "checkOut","getAverageTime","getAverageTime","checkIn","getAverageTime",
 "checkOut","getAverageTime"
]
[
 [],[45,"Leyton",3],[32,"Paradise",8],[27,"Leyton",10],
 [45,"Waterloo",15],[27,"Waterloo",20],[32,"Cambridge",22],
 ["Paradise","Cambridge"],["Leyton","Waterloo"],[10,"Leyton",24],
 ["Leyton","Waterloo"],[10,"Waterloo",38],["Leyton","Waterloo"]
]
Выход:
[
 null,null,null,null,null,null,null,14.00000,11.00000,
 null,11.00000,null,12.00000
]
Пояснение:
UndergroundSystem undergroundSystem = new UndergroundSystem();
undergroundSystem.checkIn(45, "Leyton", 3);
undergroundSystem.checkIn(32, "Paradise", 8);
undergroundSystem.checkIn(27, "Leyton", 10);
undergroundSystem.checkOut(45, "Waterloo", 15);
// Пассажир 45: "Leyton" -> "Waterloo" за 15-3 = 12.
undergroundSystem.checkOut(27, "Waterloo", 20);
// Пассажир 27: "Leyton" -> "Waterloo" за 20-10 = 10.
undergroundSystem.checkOut(32, "Cambridge", 22);
// Пассажир 32: "Paradise" -> "Cambridge" за 22-8 = 14.
undergroundSystem.getAverageTime("Paradise", "Cambridge");
// Возвращает 14.00000.
// Одна поездка "Paradise" -> "Cambridge": (14) / 1 = 14.
undergroundSystem.getAverageTime("Leyton", "Waterloo");
// Возвращает 11.00000.
// Две поездки "Leyton" -> "Waterloo": (10 + 12) / 2 = 11.
undergroundSystem.checkIn(10, "Leyton", 24);
undergroundSystem.getAverageTime("Leyton", "Waterloo");
// Возвращает 11.00000.
undergroundSystem.checkOut(10, "Waterloo", 38);
// Пассажир 10: "Leyton" -> "Waterloo" за 38-24 = 14.
undergroundSystem.getAverageTime("Leyton", "Waterloo");
// Возвращает 12.00000.
// Три поездки "Leyton" -> "Waterloo": (10 + 12 + 14) / 3 = 12.

Пример 2:
Вход:
[
 "UndergroundSystem","checkIn","checkOut","getAverageTime","checkIn",
 "checkOut","getAverageTime","checkIn","checkOut","getAverageTime"
]
[
 [],[10,"Leyton",3],[10,"Paradise",8],["Leyton","Paradise"],
 [5,"Leyton",10],[5,"Paradise",16],["Leyton","Paradise"],
 [2,"Leyton",21],[2,"Paradise",30],["Leyton","Paradise"]
]
Выход:
[null,null,null,5.00000,null,null,5.50000,null,null,6.66667]
Пояснение:
UndergroundSystem undergroundSystem = new UndergroundSystem();
undergroundSystem.checkIn(10, "Leyton", 3);
undergroundSystem.checkOut(10, "Paradise", 8);
// Пассажир 10: "Leyton" -> "Paradise" за 8-3 = 5.
undergroundSystem.getAverageTime("Leyton", "Paradise");
// Возвращает 5.00000: (5) / 1 = 5.
undergroundSystem.checkIn(5, "Leyton", 10);
undergroundSystem.checkOut(5, "Paradise", 16);
// Пассажир 5: "Leyton" -> "Paradise" за 16-10 = 6.
undergroundSystem.getAverageTime("Leyton", "Paradise");
// Возвращает 5.50000: (5 + 6) / 2 = 5.5.
undergroundSystem.checkIn(2, "Leyton", 21);
undergroundSystem.checkOut(2, "Paradise", 30);
// Пассажир 2: "Leyton" -> "Paradise" за 30-21 = 9.
undergroundSystem.getAverageTime("Leyton", "Paradise");
// Возвращает 6.66667: (5 + 6 + 9) / 3 = 6.66667.

Ограничения:
1 <= id, t <= 10^6
1 <= stationName.length, startStation.length, endStation.length <= 10
Все строки состоят из заглавных и строчных английских букв и цифр.
Общее число вызовов checkIn, checkOut и getAverageTime не превышает 2 * 10^4.
Принимаются ответы, отличающиеся от точного значения не более чем на 10^-5.

*/

// UndergroundSystem хранит последние входы пассажиров и статистику маршрутов.
// Память: O(U + R) записей при ограниченной условием длине имён станций.
// U — число разных id, когда-либо записанных в checkInMap, R — число маршрутов.
// Строки в checkInMap могут удерживать исходные буферы целиком;
// если имена — срезы длинных строк, память этих буферов считают отдельно.
// Завершение поездки не удаляет запись входа.
// Поэтому U учитывает и пассажиров, которые уже закончили поездку.
type UndergroundSystem struct {
    // id -> [станция входа, время входа]. Новый вход заменяет старую запись id.
    checkInMap map[int][]any
    // "начальная,конечная" -> [сумма времён, число завершённых поездок].
    // Два агрегата заменяют хранение каждой завершённой поездки.
    routeMap   map[string][]int
}

// Constructor создаёт систему с двумя пустыми таблицами.
// Время: O(1). Начальная дополнительная память: O(1), до накопления записей.
func Constructor() UndergroundSystem {
    return UndergroundSystem{
        checkInMap: make(map[int][]any),
        routeMap:   make(map[string][]int),
    }
}

// CheckIn сохраняет станцию и время начала поездки, изменяя checkInMap.
// В среднем амортизированное время: O(1); память на новую запись: O(1).
// Ключ id имеет фиксированный размер, строка станции здесь не копируется.
func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
    // По контракту пассажир не может начать вторую поездку, не закончив первую.
    // После завершённой поездки новый вход перезапишет старые данные.
    this.checkInMap[id] = []any{stationName, t}
}

// CheckOut добавляет завершённую поездку в routeMap; checkInMap не очищается.
// В среднем амортизированное время и временная память: O(L),
// где L = len(startStation) + 1 + len(stationName) — длина ключа маршрута.
// Создание ключа и его хеширование зависят от L. По условию L <= 21,
// поэтому здесь обе оценки упрощаются до O(1); новый маршрут добавляет запись.
func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
    // Согласованность вызовов гарантирует предшествующий CheckIn для этого id.
    // В записи ровно два элемента: string в позиции 0 и int в позиции 1.
    // Поэтому индексы и приведения типов соответствуют данным из CheckIn.
    entry := this.checkInMap[id]
    startStation := entry[0].(string)
    time := entry[1].(int)
    // Маршрут направленный: ключ "A,B" отличается от "B,A".
    // Запятая не встречается в именах станций, поэтому граница имён однозначна.
    route := startStation + "," + stationName
    // Для первого завершения этого маршрута создаём нулевые сумму и счётчик.
    if _, ok := this.routeMap[route]; !ok {
        this.routeMap[route] = []int{0, 0}
    }
    // Длительность положительна: время выхода t больше времени входа time.
    // Сумма может достигать порядка 2 * 10^10 при допустимом числе вызовов.
    // Для 64-битного int это безопасно; 32-битный int может переполниться.
    this.routeMap[route][0] += t - time
    // Один выход завершает одну поездку; CheckIn сам по себе среднее не меняет.
    this.routeMap[route][1] += 1
}

// GetAverageTime возвращает среднее по завершённым поездкам, не меняя систему.
// В среднем время: O(L), временная память: O(L),
// где L = len(startStation) + 1 + len(endStation) — длина создаваемого ключа.
// При L <= 21 по условию обе оценки упрощаются до O(1).
func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    // Используем тот же направленный ключ, под которым CheckOut копит данные.
    data := this.routeMap[startStation+","+endStation]
    // Контракт гарантирует хотя бы одну завершённую поездку по этому маршруту:
    // запись существует, а data[1] положителен, поэтому делитель не равен нулю.
    // Например, длительности 12 и 10 дают сумму 22, счётчик 2 и среднее 11.
    // Приводим числа к float64 до деления, чтобы сохранить дробную часть.
    return float64(data[0]) / float64(data[1])
}

/**
 * Объект UndergroundSystem будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */

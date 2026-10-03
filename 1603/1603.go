package main

/*
Спроектируйте систему парковки для автостоянки.
На стоянке есть три типа мест: для больших, средних и маленьких автомобилей.
Для каждого типа задано фиксированное число мест.

Реализуйте класс ParkingSystem:
ParkingSystem(int big, int medium, int small) — создаёт объект ParkingSystem.
Число мест каждого типа задаётся параметрами конструктора.

bool addCar(int carType) — проверяет наличие свободного места типа carType
для автомобиля, который хочет въехать на стоянку.
carType принимает значения 1, 2 или 3: большой, средний или маленький автомобиль
соответственно.
Автомобиль может занимать только место своего типа carType.
Если свободного места нет, верните false.
Иначе припаркуйте автомобиль на месте соответствующего типа и верните true.

Пример 1:
Вход:
["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]
[[1, 1, 0], [1], [2], [3], [1]]
Выход:
[null, true, true, false, false]
Пояснение:
ParkingSystem parkingSystem = new ParkingSystem(1, 1, 0);
parkingSystem.addCar(1); // Вернёт true: свободно 1 место
                       // для большого автомобиля.
parkingSystem.addCar(2); // Вернёт true: свободно 1 место
                       // для среднего автомобиля.
parkingSystem.addCar(3); // Вернёт false: нет места для маленького автомобиля.
parkingSystem.addCar(1); // Вернёт false: место для большого автомобиля
                       // уже занято.

Ограничения:
0 <= big, medium, small <= 1000
carType равен 1, 2 или 3.
Метод addCar будет вызван не более 1000 раз.

*/

// ParkingSystem хранит количество ещё свободных мест каждого типа.
type ParkingSystem struct {
    // Порядок счётчиков: большие, средние, маленькие места.
    spaces []int
}

// Constructor задаёт начальное количество свободных мест каждого типа.
// Время: O(1). Дополнительная память: O(1), поскольку типов всегда три.
func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{spaces: []int{big, medium, small}}
}

// AddCar пытается занять место нужного типа и сообщает, удалось ли это.
// Меняет this.spaces: при успехе уменьшает соответствующий счётчик на 1.
// Время: O(1). Дополнительная память: O(1).
func (this *ParkingSystem) AddCar(carType int) bool {
    // По условию carType равен 1, 2 или 3, поэтому carType-1 — индекс 0..2.
    // Constructor создаёт три счётчика; проверяем только нужный тип места.
    if this.spaces[carType-1] > 0 {
        // Занимаем одно место; счётчик после проверки не станет отрицательным.
        this.spaces[carType-1]--
        return true
    }
    // Мест этого типа нет. Другой тип не подходит по условию,
    // поэтому состояние не меняем и отказываем в парковке.
    return false
}

/**
 * Объект ParkingSystem будет создаваться и вызываться следующим образом:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

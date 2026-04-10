package main

/*
Вам дан поток точек на плоскости X-Y. Разработайте алгоритм, который:

Добавляет новые точки из потока в структуру данных.
Дублирующиеся точки разрешены и рассматриваются как разные точки.
По заданной точке запроса подсчитывает количество способов выбрать три точки
из структуры данных так, чтобы три точки и точка запроса образовывали
выровненный по осям квадрат с положительной площадью.
Выровненный по осям квадрат — квадрат, все рёбра которого одинаковы
и параллельны либо перпендикулярны осям X и Y.

Реализуйте класс DetectSquares:

DetectSquares() Инициализирует объект с пустой структурой данных.
void add(int[] point) Добавляет новую точку point = [x, y] в структуру данных.
int count(int[] point) Подсчитывает количество способов образовать выровненные
                       по осям квадраты с точкой point = [x, y], как описано выше.

Пример 1:
https://assets.leetcode.com/uploads/2021/09/01/image.png

Вход
["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
Выход
[null, null, null, null, 1, 0, null, 2]

Пояснение
DetectSquares detectSquares = new DetectSquares();
detectSquares.add([3, 10]);
detectSquares.add([11, 2]);
detectSquares.add([3, 2]);
detectSquares.count([11, 10]); // вернёт 1. Можно выбрать:
                               //   - Первую, вторую и третью точки
detectSquares.count([14, 8]);  // вернёт 0. Точка запроса не может образовать квадрат
                               //            ни с какими точками в структуре данных.
detectSquares.add([11, 2]);    // Добавление дублирующихся точек разрешено.
detectSquares.count([11, 10]); // вернёт 2. Можно выбрать:
                               //   - Первую, вторую и третью точки
                               //   - Первую, третью и четвёртую точки

Ограничения:

point.length == 2
0 <= x, y <= 1000
Суммарно будет сделано не более 3000 вызовов add и count.

*/

type DetectSquares struct {
    ptsCount map[Point]int
}

type Point struct {
    x, y int
}

func Constructor() DetectSquares {
    return DetectSquares{ptsCount: make(map[Point]int)}
}

func (this *DetectSquares) Add(point []int) {
    p := Point{point[0], point[1]}
    this.ptsCount[p]++
}

func (this *DetectSquares) Count(point []int) int {
    res := 0
    px, py := point[0], point[1]

    // Перебираем все уникальные точки как потенциальные диагональные вершины квадрата.
    for pt, countPt := range this.ptsCount {
        // Проверяем условия валидного квадрата:
        // 1. |dy| == |dx| (катеты равны)
        // 2. pt не лежит на одной вертикали или горизонтали с point (иначе площадь = 0)
        if abs(py-pt.y) != abs(px-pt.x) || pt.x == px || pt.y == py {
            continue
        }

        // Вычисляем координаты двух оставшихся вершин квадрата.
        p1 := Point{pt.x, py}
        p2 := Point{px, pt.y}

        // Количество способов собрать квадрат = произведение количеств всех трёх вершин.
        // Если p1 или p2 отсутствуют в мапе, вернёт 0 и вклад аннулируется.
        res += countPt * this.ptsCount[p1] * this.ptsCount[p2]
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


/**
 * Объект DetectSquares будет создан и использован следующим образом:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

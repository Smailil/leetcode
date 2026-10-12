package main

/*
Разработайте систему рейтингов блюд со следующими возможностями:
Изменять рейтинг блюда, зарегистрированного в системе.
Возвращать блюдо с наибольшим рейтингом для заданного типа кухни.

Реализуйте класс FoodRatings:
FoodRatings(String[] foods, String[] cuisines, int[] ratings)
Инициализирует систему.
Блюда описываются массивами foods, cuisines и ratings, каждый длины n.
foods[i] — название i-го блюда.
cuisines[i] — тип кухни i-го блюда.
ratings[i] — начальный рейтинг i-го блюда.
void changeRating(String food, int newRating)
Изменяет рейтинг блюда с названием food.
String highestRated(String cuisine)
Возвращает название блюда с наибольшим рейтингом для заданного типа cuisine.
При равных рейтингах возвращает блюдо с лексикографически меньшим названием.

Строка x лексикографически меньше строки y, если стоит раньше в словаре:
либо x является префиксом y, либо в первой позиции i, где x[i] != y[i],
символ x[i] предшествует символу y[i] в алфавитном порядке.

Пример 1:
Входные данные
["FoodRatings", "highestRated", "highestRated", "changeRating",
 "highestRated", "changeRating", "highestRated"]
[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"],
  ["korean", "japanese", "japanese", "greek", "japanese", "korean"],
  [9, 12, 8, 15, 14, 7]],
 ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16],
 ["japanese"]]
Выходные данные
[null, "kimchi", "ramen", null, "sushi", null, "ramen"]
Пояснение
FoodRatings foodRatings = new FoodRatings(
    ["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"],
    ["korean", "japanese", "japanese", "greek", "japanese", "korean"],
    [9, 12, 8, 15, 14, 7]);
foodRatings.highestRated("korean"); // возвращает "kimchi"
                                    // У "kimchi" наибольший рейтинг
                                    // среди блюд корейской кухни: 9.
foodRatings.highestRated("japanese"); // возвращает "ramen"
                                      // У "ramen" наибольший рейтинг
                                      // среди блюд японской кухни: 14.
foodRatings.changeRating("sushi", 16); // Теперь рейтинг "sushi" равен 16.
foodRatings.highestRated("japanese"); // возвращает "sushi"
                                      // У "sushi" наибольший рейтинг
                                      // среди блюд японской кухни: 16.
foodRatings.changeRating("ramen", 16); // Теперь рейтинг "ramen" равен 16.
foodRatings.highestRated("japanese"); // возвращает "ramen"
                                      // У "sushi" и "ramen" рейтинг 16.
                                      // Но "ramen" лексикографически
                                      // меньше, чем "sushi".

Ограничения:
1 <= n <= 2 * 10^4
n == foods.length == cuisines.length == ratings.length
1 <= foods[i].length, cuisines[i].length <= 10
foods[i] и cuisines[i] состоят из строчных букв английского алфавита.
1 <= ratings[i] <= 10^8
Все строки в foods различны.
Во всех вызовах changeRating значение food — название блюда в системе.
Во всех вызовах highestRated значение cuisine — тип кухни
хотя бы одного блюда в системе.
Суммарное число вызовов changeRating и highestRated не превышает 2 * 10^4.
*/

import (
    "github.com/emirpasic/gods/sets/treeset"
)

// FoodRatings связывает блюдо с его текущим рейтингом и кухней, а для каждой
// кухни хранит упорядоченное множество пар [-rating, food].
// Инвариант: в множестве кухни есть ровно одна актуальная пара для каждого
// её блюда; значение rating совпадает с записью в foodToRating.
// Общая память структуры: O(n), где n — число блюд. Карты и все деревья
// вместе содержат O(n) записей; строки используются без копирования текста.
type FoodRatings struct {
    // Нужна, чтобы при обновлении найти старый ключ и удалить его из дерева.
    foodToRating  map[string]int
    // Кухня блюда не меняется: по имени сразу находим нужное множество.
    foodToCuisine map[string]string
    // treeset основан на красно-чёрном дереве. Минимальный ключ в нём —
    // самое рейтинговое блюдо, а при равенстве — меньшее по имени.
    cuisineToSet  map[string]*treeset.Set
}

// foodComparator сравнивает пары [-rating, food] по возрастанию:
// сначала отрицательный рейтинг, затем название лексикографически.
// Например, [-16, "ramen"] идёт перед [-14, "sushi"]; при рейтинге 16
// у обоих блюд "ramen" тоже идёт первым из-за сравнения названий.
// Время: O(L) в худшем случае, где L — максимальная длина названий;
// при ограничении L <= 10 это O(1). Дополнительная память: O(1).
func foodComparator(a, b any) int {
    // Add и Remove передают именно [2]interface{}{int, string}.
    // Приведения типов опираются на этот внутренний формат ключа.
    fa := a.([2]any)
    fb := b.([2]interface{})
    ra, rb := fa[0].(int), fb[0].(int)
    na, nb := fa[1].(string), fb[1].(string)
    if ra != rb {
        // Более высокий рейтинг даёт меньшее отрицательное число.
        // На диапазоне рейтингов 1..10^8 разность безопасна и для int32;
        // от переполнения для произвольных значений int защиты здесь нет.
        return ra - rb
    }
    // Рейтинги равны: название определяет порядок, как требует условие.
    if na < nb {
        return -1
    }
    if na > nb {
        return 1
    }
    // Ноль означает одинаковый ключ, а не просто одинаковый рейтинг.
    return 0
}

// Constructor создаёт индексы блюд и отдельное дерево для каждой кухни.
// Входные срезы не меняются; возвращаемый объект владеет созданными картами.
// Время: O(n log(n+1)) в среднем по обращениям к map, n = len(foods).
// Дерево делает O(log(n+1)) сравнений на вставку; длины строк здесь <= 10.
// Если снять этот предел, оценка станет O(n L log(n+1)), где L — максимум
// длины названия блюда или кухни: хеширование и сравнение строк не бесплатны.
// Память результата: O(n) для карт и узлов деревьев.
// Рабочая память помимо результата: O(log(n+1)) для рекурсивного стека
// балансировки в используемом treeset, без отдельного массива блюд.
func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    f := FoodRatings{
        foodToRating:  make(map[string]int),
        foodToCuisine: make(map[string]string),
        cuisineToSet:  make(map[string]*treeset.Set),
    }
    for i := 0; i < len(foods); i++ {
        // По условию все три среза одной длины, поэтому индекс i допустим
        // и для cuisines, и для ratings. Сохраняем две прямые связи блюда.
        f.foodToRating[foods[i]] = ratings[i]
        f.foodToCuisine[foods[i]] = cuisines[i]
        // Дерево создаём только при первом блюде этой кухни; последующие
        // блюда попадают в тот же набор и не смешиваются с другими кухнями.
        if _, ok := f.cuisineToSet[cuisines[i]]; !ok {
            f.cuisineToSet[cuisines[i]] = treeset.NewWith(foodComparator)
        }
        // Минус превращает поиск максимума рейтинга в поиск минимума ключа.
        // Имена блюд уникальны: даже при равном рейтинге ключи различаются,
        // поэтому множество не объединит два разных блюда в одну запись.
        f.cuisineToSet[cuisines[i]].Add([2]interface{}{-ratings[i], foods[i]})
    }
    return f
}

// ChangeRating меняет рейтинг в карте и ключ блюда в дереве его кухни.
// Кухня остаётся прежней, число актуальных записей не увеличивается.
// Время: O(log(m+1)) в среднем по map, где m — число блюд этой кухни;
// удаление и вставка в красно-чёрное дерево имеют логарифмическую стоимость.
// При строках без ограничения длины: O(L log(m+1)), L — их максимум.
// Рабочая память: O(log(m+1)) для рекурсивной балансировки библиотеки;
// сами локальные переменные и новый ключ занимают O(1). Структура — O(n).
func (this *FoodRatings) ChangeRating(food string, newRating int) {
    // По условию food существует. Находим его кухню, прежний рейтинг
    // и дерево; старый рейтинг сохраняем до изменения карты.
    cuisine := this.foodToCuisine[food]
    oldRating := this.foodToRating[food]
    set := this.cuisineToSet[cuisine]

    // Удаляем точную старую пару: Remove ищет ключ через foodComparator.
    // Нельзя лишь заменить рейтинг в карте — в дереве остался бы старый
    // ключ, и HighestRated мог бы вернуть блюдо по устаревшему рейтингу.
    set.Remove([2]interface{}{-oldRating, food})
    this.foodToRating[food] = newRating
    // Вставка новой пары заново определяет позицию блюда в порядке дерева.
    // После этих трёх действий карта и дерево снова описывают один рейтинг.
    set.Add([2]interface{}{-newRating, food})
}

// HighestRated возвращает лучшее блюдо кухни, не меняя рейтинги и деревья.
// Время: O(log(m+1)) в среднем по map, m — число блюд заданной кухни.
// First ищет крайний левый узел дерева, а не читает готовый максимум за O(1).
// Для неограниченной длины cuisine добавилась бы стоимость её хеширования.
// Рабочая память: O(1) — итератор содержит только несколько полей.
// Результат — существующая строка имени; её текст не копируется.
func (this *FoodRatings) HighestRated(cuisine string) string {
    // По условию у кухни есть хотя бы одно блюдо: дерево не nil и не пусто.
    // Поэтому First успешно ставит итератор на первый элемент, а Value
    // можно вызывать без отдельной обработки отсутствующего результата.
    it := this.cuisineToSet[cuisine].Iterator()
    it.First()
    // Первый ключ уже учитывает оба приоритета: рейтинг и имя.
    // Value возвращает пару [-rating, food]; элемент с индексом 1 — имя.
    return it.Value().([2]interface{})[1].(string)
}


/**
 * Объект FoodRatings создаётся и используется следующим образом:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

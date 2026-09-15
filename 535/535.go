package main

import "strconv"

/*
Примечание: эта задача является дополнением к задаче по проектированию систем
Design TinyURL.
TinyURL — это сервис сокращения URL: вы вводите URL, например
https://leetcode.com/problems/design-tinyurl, а сервис возвращает короткий URL,
например http://tinyurl.com/4e9iAk.
Спроектируйте класс для кодирования URL и декодирования короткого URL.

Способ работы алгоритма encode/decode не ограничен. Нужно лишь гарантировать,
что URL можно закодировать в короткий URL, а затем декодировать его обратно
в исходный URL.

Реализуйте класс Solution:
Solution() Инициализирует объект системы.
String encode(String longUrl) Возвращает короткий URL для заданного longUrl.
String decode(String shortUrl) Возвращает исходный длинный URL для заданного
shortUrl. Гарантируется, что shortUrl был закодирован тем же объектом.

Пример 1:
Вход: url = "https://leetcode.com/problems/design-tinyurl"
Выход: "https://leetcode.com/problems/design-tinyurl"
Пояснение:
Solution obj = new Solution();
string tiny = obj.encode(url); // возвращает закодированный короткий URL.
string ans = obj.decode(tiny); // возвращает исходный URL после декодирования.

Ограничения:
1 <= url.length <= 10^4
Гарантируется, что url является допустимым URL.
*/

// Codec хранит взаимно обратные соответствия между длинными и короткими URL.
// encodeMap сохраняет результат повторного кодирования, а decodeMap позволяет
// восстановить исходный адрес по короткой ссылке.
// Память состояния: O(U + S), где U и S — суммарные длины сохранённых
// длинных и коротких URL соответственно.
type Codec struct {
    encodeMap map[string]string
    decodeMap map[string]string
    base      string
}

// Constructor создаёт пустой кодек с общей базой коротких URL.
// Время: O(1). Дополнительная память: O(1).
func Constructor() Codec {
    return Codec{
        encodeMap: make(map[string]string),
        decodeMap: make(map[string]string),
        base:      "http://tinyurl.com/",
    }
}

// encode возвращает стабильный короткий URL: повторный longUrl получает
// прежнее значение.
// Пусть L = len(longUrl), d — число цифр нового идентификатора.
// Среднее время: O(L) для известного URL и O(L + d) для нового;
// хеширование строк зависит от их длины, а операции map в среднем занимают
// O(1). Прирост памяти: O(1) для известного URL и O(L + d) для нового.
func (this *Codec) encode(longUrl string) string {
    // Проверка encodeMap сохраняет идемпотентность: один длинный URL всегда
    // связан с одной и той же короткой ссылкой.
    if _, exists := this.encodeMap[longUrl]; !exists {
        // Следующий номер равен количеству сохранённых URL плюс один.
        // Поскольку записи не удаляются, такой идентификатор ещё не занят.
        shortUrl := this.base + strconv.Itoa(len(this.encodeMap)+1)

        // Сохраняем обе стороны соответствия: первую для повторного encode,
        // вторую — для обратного decode.
        this.encodeMap[longUrl] = shortUrl
        this.decodeMap[shortUrl] = longUrl
    }
    return this.encodeMap[longUrl]
}

// decode восстанавливает longUrl по shortUrl, созданному тем же Codec.
// Пусть S = len(shortUrl). Среднее время: O(S) из-за хеширования строки;
// дополнительная память: O(1).
func (this *Codec) decode(shortUrl string) string {
    // По условию ссылка выдана этим объектом, поэтому ключ существует.
    return this.decodeMap[shortUrl]
}

/**
 * Объект Codec будет создан и использован следующим образом:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */


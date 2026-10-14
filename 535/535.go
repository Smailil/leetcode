package main

import "strconv"

/*
Примечание: эта задача является дополнением к задаче по проектированию систем
Design TinyURL.
TinyURL — это сервис сокращения URL: вы вводите URL, например
https://leetcode.com/problems/design-tinyurl, а сервис возвращает короткий URL,
например http://tinyurl.com/4e9iAk.
Спроектируйте класс для кодирования URL и декодирования короткого URL.

Способ работы алгоритма encode/decode не ограничен.
Нужно лишь гарантировать,
что URL можно закодировать в короткий URL, а затем декодировать его обратно
в исходный URL.

Реализуйте класс Solution:
Solution() Инициализирует объект системы.
String encode(String longUrl) Возвращает короткий URL для заданного longUrl.
String decode(String shortUrl) Возвращает исходный длинный URL для заданного
shortUrl.
Гарантируется, что shortUrl был закодирован тем же объектом.

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
// Память состояния O(p+B+S), p — число разных длинных URL.
// B — объём удерживаемых буферов длинных URL без повторного подсчёта общих;
// S — общий объём сохранённых коротких ссылок. Карты разделяют эти тексты.
// Подстрока может удерживать исходный буфер большего размера, чем её длина.
// Копии Codec разделяют карты, а не создают независимый кодек.
type Codec struct {
    encodeMap map[string]string
    decodeMap map[string]string
    base      string
}

// Constructor создаёт пустой кодек с общей базой коротких URL.
// Время: O(1).
// Рабочая память и новый пустой результат O(1); карты растут при encode.
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
// O(1).
// Рабочая память без состояния: O(1) для известного URL, O(d) для нового
// за счёт строки strconv.Itoa; внутренний рост карт может требовать O(p).
// Новый longUrl не копируется: состояние удерживает его исходный буфер.
// Результат разделяет сохранённую короткую строку размера O(d), без копии
// при возврате; её текст уже учтён в памяти состояния, не считать его дважды.
// Метод меняет обе карты this, но не содержимое входной строки.
func (this *Codec) encode(longUrl string) string {
    // Проверка encodeMap сохраняет идемпотентность: один длинный URL всегда
    // связан с одной и той же короткой ссылкой.
    if _, exists := this.encodeMap[longUrl]; !exists {
        // Следующий номер равен количеству сохранённых URL плюс один.
        // При отсутствии удаления номер не повторяется, пока помещается в int.
        // Явной проверки переполнения len(encodeMap)+1 в реализации нет.
        shortUrl := this.base + strconv.Itoa(len(this.encodeMap)+1)

        // Сохраняем обе стороны соответствия: первую для повторного encode,
        // вторую — для обратного decode.
        this.encodeMap[longUrl] = shortUrl
        this.decodeMap[shortUrl] = longUrl
    }
    return this.encodeMap[longUrl]
}

// decode восстанавливает longUrl по shortUrl, созданному тем же Codec.
// Пусть S = len(shortUrl).
// Среднее время: O(S) из-за хеширования строки;
// рабочая память: O(1); метод не меняет карты и строку shortUrl.
// Результат — заголовок уже сохранённой длинной строки, без копии текста.
// Он удерживает её исходный буфер, который может быть больше длины URL.
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


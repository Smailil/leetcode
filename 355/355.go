package main

import "container/heap"

/*
Разработайте упрощенную версию Twitter,
в которой пользователи смогут публиковать твиты,
подписываться на других пользователей или отписываться от них,
а также видеть 10 последних твитов в ленте новостей пользователя.

Реализуйте класс Twitter:
Twitter() Инициализирует ваш объект Twitter.
void postTweet(int userId, int tweetId) Создает новый твит
с идентификатором tweetId по пользователю userId.
Каждый вызов этой функции будет осуществляться
с уникальным идентификатором твита.
List<Integer> getNewsFeed(int userId) Получает 10 последних идентификаторов
твитов в ленте новостей пользователя.
Каждый элемент в ленте новостей должен быть опубликован пользователями,
на которых подписан пользователь, или самим пользователем.
Твиты должны быть упорядочены от самого нового до самого старого.
void Follow(int FollowerId, int FolloweeId) Пользователь с идентификатором
FollowerId начал подписываться на пользователя с идентификатором FolloweeId.
void unfollow(int FollowerId, int FolloweeId) Пользователь с идентификатором
FollowerId начал отписываться от пользователя с идентификатором FolloweeId.

Пример 1:
Вход:
["Twitter", "postTweet", "getNewsFeed", "follow",
 "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Выход:
[null, null, [5], null, null, [6, 5], null, [5]]
Объяснение:
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // Пользователь 1 публикует новый твит
// (идентификатор = 5).
twitter.getNewsFeed(1); // Лента новостей пользователя 1 должна вернуть
// список с 1 идентификатором твита -> [5].
// Возвращает [5].
twitter.follow(1, 2); // Пользователь 1 подписывается на пользователя 2.
twitter.postTweet(2, 6); // Пользователь 2 публикует новый твит
// (идентификатор = 6).
twitter.getNewsFeed(1); // Лента новостей пользователя 1 должна вернуть
// список с 2 идентификаторами твитов -> [6, 5].
// Твит с идентификатором 6 должен предшествовать твиту
// с идентификатором 5,
// так как он был опубликован после твита с идентификатором 5.
twitter.unfollow(1, 2); // Пользователь 1 отписывается от пользователя 2.
twitter.getNewsFeed(1); // Лента новостей пользователя 1 должна вернуть список
// с 1 идентификатором твита -> [5],
// так как пользователь 1 больше не подписан на пользователя 2.

Ограничения:
1 <= userId, followerId, followeeId <= 500
0 <= tweetId <= 10^4
Все твиты имеют уникальные идентификаторы.
Будет сделано не более 3 *10^4 вызовов для публикации твита,
getNewsFeed, подписки и отмены подписки.
Пользователь не может подписываться на себя.
*/

type tweet struct {
	timestamp int
	id int
	
}

// Twitter хранит истории по возрастанию времени и множества подписок.
// Постоянная память O(T + S), где T — число твитов, S — число подписок.
// Карты учитывают созданных пользователей, их число по условию до 500.
type Twitter struct {
    latestTimestamp int
	followMap map[int]map[int]struct{}
	tweetMap map[int][]tweet

}

type TweetMaxHeap struct {
    items []heapTweet
}

// Push — callback container/heap: дописывает heapTweet, меняя h.items.
// Время амортизированно O(1), при росте буфера O(a), где a — размер кучи.
// Рабочая память роста O(a); постоянно добавляется один элемент кучи.
func (h *TweetMaxHeap) Push(value interface{}) {
    h.items = append(h.items, value.(heapTweet))
}

// Pop — callback: container/heap уже перенёс выбранный корень в конец.
// Снимает последний элемент за O(1) времени и рабочей памяти, меняя h.items.
// Буфер среза не сокращается; возвращаемое значение занимает O(1) памяти.
func (h *TweetMaxHeap) Pop() interface{} {
    last := h.items[len(h.items)-1]
    h.items = h.items[:len(h.items)-1]
    
    return last
}

// Swap меняет позиции двух элементов при восстановлении свойств кучи.
// Время и дополнительная память O(1); изменяет h.items.
func (h *TweetMaxHeap) Swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Less задаёт max-heap: более поздний timestamp имеет больший приоритет.
// Время и дополнительная память O(1), состояние не изменяет.
func (h *TweetMaxHeap) Less(i, j int) bool {
    return h.items[i].tweet.timestamp > h.items[j].tweet.timestamp
}

// Len сообщает размер кучи за O(1) времени и памяти без изменения состояния.
func (h *TweetMaxHeap) Len() int {
    return len(h.items)
}

type heapTweet struct {
    idx int
    userId int
    tweet tweet
}


// Constructor создаёт пустые карты и нулевые логические часы Twitter.
// Время, рабочая память и начальная память возвращаемого состояния O(1).
func Constructor() Twitter {
    return Twitter{0, 
	make(map[int]map[int]struct{}), 
	make(map[int][]tweet)}
}


// PostTweet добавляет твит в конец истории, увеличивая логические часы.
// При среднем O(1) доступа к map время амортизированно O(1).
// Рост среза требует O(Tu) времени и рабочей памяти, Tu — история автора.
// Постоянное хранилище растёт амортизированно на O(1); меняет receiver.
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    userTweets, ok := this.tweetMap[userId]
    if !ok {
        userTweets = make([]tweet, 0)
        
    }
	// Всего вызовов до 30000: часы помещаются даже в 32-битный int.
	// Новая отметка больше всех прежних: история остаётся отсортированной.
	this.latestTimestamp++
        
    userTweets = append(userTweets, tweet{
        id: tweetId, 
        timestamp: this.latestTimestamp,
    })
    this.tweetMap[userId] = userTweets
}


// GetNewsFeed сливает истории подписок max-кучей, возвращая до 10 ID.
// F — число подписок с собой, a — непустые истории, r — размер ответа.
// При среднем доступе к map время O(F + a log(a + 1) + r log(a + 1)).
// Рабочая память O(a), память результата O(r), где r <= 10.
// Меняет followMap: сохраняет подписку на себя; истории твитов не изменяет.
func (this *Twitter) GetNewsFeed(userId int) []int {
    followees, ok := this.followMap[userId] 
    if !ok {
        followees = make(map[int]struct{})
    }
    // Свои твиты входят в ленту без публичного Follow на самого себя.
    // Это включение сохраняется в множестве подписок пользователя.
    if _, ok := followees[userId]; !ok {
        followees[userId] = struct{}{}
    }
    this.followMap[userId] = followees
    tweetHeap := &TweetMaxHeap{items: make([]heapTweet, 0)}
    
    // Куча хранит самый новый ещё не выданный твит каждой непустой истории.
    // Каждый heap.Push стоит O(log(a + 1)), это не общий heap.Init.
    for followee := range followees {
        tweets := this.tweetMap[followee]
        
        if len(tweets) == 0 {
            continue
        }
        
       // Непустая история делает индекс len(tweets) - 1 допустимым.
       heap.Push(tweetHeap, heapTweet{
            tweet: tweets[len(tweets)-1],
            idx: len(tweets)-1,
            userId: followee,
        })
    }
    
    // Корень — самый новый из оставшихся кандидатов всех непустых историй.
    var result []int
    for tweetHeap.Len() > 0 && len(result) < 10 {
        popped := heap.Pop(tweetHeap).(heapTweet)
        result = append(result, popped.tweet.id)
    
        nextIdx := popped.idx - 1
        
        // Выданный кандидат заменяется предыдущим твитом того же автора.
        // nextIdx ниже прежнего допустимого индекса; нужна нижняя граница.
        if nextIdx >= 0 {
            tweets := this.tweetMap[popped.userId]
            
            heap.Push(tweetHeap, heapTweet{
                tweet: tweets[nextIdx],
                idx: nextIdx,
                userId: popped.userId,
            })
        }
       
    }
    return result
}


// Follow добавляет подписку в множество; повтор не создаёт дубликат.
// Время в среднем O(1), рабочая и добавленная постоянная память O(1).
// Изменяет followMap, создавая при необходимости множество подписок.
func (this *Twitter) Follow(followerId int, followeeId int)  {
    followees, ok := this.followMap[followerId] 
    if !ok {
        followees = make(map[int]struct{})
        this.followMap[followerId] = followees
    }
    
    followees[followeeId] = struct{}{}
}


// Unfollow удаляет подписку при наличии множества, иначе ничего не делает.
// Время в среднем O(1), рабочая память O(1); изменяет followMap.
// Удаление не требует освобождения всей памяти внутреннего буфера map.
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if followees, ok := this.followMap[followerId]; ok {
        delete(followees, followeeId)
    }
}


/**
 * Объект Twitter будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

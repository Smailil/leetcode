package main

import "container/heap"

/*
Разработайте упрощенную версию Twitter, в которой пользователи смогут публиковать твиты,
подписываться на других пользователей или отписываться от них,
а также видеть 10 последних твитов в ленте новостей пользователя.

Реализуйте класс Twitter:
Twitter() Инициализирует ваш объект Twitter.

void postTweet(int userId, int twitterId) Создает новый твит с идентификатором tweetId по пользователю userId.
Каждый вызов этой функции будет осуществляться с уникальным идентификатором твита.

List<Integer> getNewsFeed(int userId) Получает 10 последних идентификаторов твитов в ленте новостей пользователя.
Каждый элемент в ленте новостей должен быть опубликован пользователями,
на которых подписан пользователь, или самим пользователем.
Твиты должны быть упорядочены от самого последнего до самого последнего.

void Follow(int FollowerId, int FolloweeId) Пользователь с идентификатором FollowerId начал подписываться
на пользователя с идентификатором FolloweeId.

void unfollow(int FollowerId, int FolloweeId) Пользователь с идентификатором FollowerId начал отписываться
от пользователя с идентификатором FolloweeId.

Пример 1:
Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]
Объяснение
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // Пользователь 1 публикует новый твит (идентификатор = 5).
twitter.get News Feed(1); // Лента новостей пользователя 1 должна вернуть
// список с 1 идентификатором твита -> [5]. return [5]
twitter.follow(1, 2); // Пользователь 1 подписывается на пользователя 2.
twitter.postTweet(2, 6); // Пользователь 2 публикует новый твит (идентификатор = 6).
twitter.getNewsFeed(1); // Лента новостей пользователя 1 должна вернуть
// список с 2 идентификаторами твитов -> [6, 5]. Твит с идентификатором 6 должен предшествовать твиту
// с идентификатором 5, так как он был опубликован после твита с идентификатором 5.
twitter.unfollow(1, 2); // Пользователь 1 отписывается от пользователя 2.
twitter.get News Feed(1); // Лента новостей пользователя 1 должна вернуть список
// с 1 идентификатором твита -> [5], так как пользователь 1 больше не подписан на пользователя 2.
Ограничения:
1 <= userId, followerId, followeeId <= 500
0 <= tweetId <= 10^4
Все твиты имеют уникальные идентификаторы.
Будет сделано не более 3 *10^4 вызовов для публикации твита, getNewsFeed, подписки и отмены подписки.
Пользователь не может следить за собой.

*/

type tweet struct {
	timestamp int
	id int
	
}

type Twitter struct {
    latestTimestamp int
	followMap map[int]map[int]struct{}
	tweetMap map[int][]tweet

}

type TweetMaxHeap struct {
    items []heapTweet
}

func (h *TweetMaxHeap) Push(value interface{}) {
    h.items = append(h.items, value.(heapTweet))
}

func (h *TweetMaxHeap) Pop() interface{} {
    last := h.items[len(h.items)-1]
    h.items = h.items[:len(h.items)-1]
    
    return last
}

func (h *TweetMaxHeap) Swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *TweetMaxHeap) Less(i, j int) bool {
    return h.items[i].tweet.timestamp > h.items[j].tweet.timestamp
}

func (h *TweetMaxHeap) Len() int {
    return len(h.items)
}

type heapTweet struct {
    idx int
    userId int
    tweet tweet
}


func Constructor() Twitter {
    return Twitter{0, 
	make(map[int]map[int]struct{}), 
	make(map[int][]tweet)}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    userTweets, ok := this.tweetMap[userId]
    if !ok {
        userTweets = make([]tweet, 0)
        
    }
	this.latestTimestamp++
        
    userTweets = append(userTweets, tweet{
        id: tweetId, 
        timestamp: this.latestTimestamp,
    })
    this.tweetMap[userId] = userTweets
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    followees, ok := this.followMap[userId] 
    if !ok {
        followees = make(map[int]struct{})
    }
    if _, ok := followees[userId]; !ok {
        followees[userId] = struct{}{}
    }
    this.followMap[userId] = followees
    tweetHeap := &TweetMaxHeap{items: make([]heapTweet, 0)}
    
    for followee := range followees {
        tweets := this.tweetMap[followee]
        
        if len(tweets) == 0 {
            continue
        }
        
       heap.Push(tweetHeap, heapTweet{
            tweet: tweets[len(tweets)-1],
            idx: len(tweets)-1,
            userId: followee,
        })
    }
    
    var result []int
    for tweetHeap.Len() > 0 && len(result) < 10 {
        popped := heap.Pop(tweetHeap).(heapTweet)
        result = append(result, popped.tweet.id)
    
        nextIdx := popped.idx - 1
        
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


func (this *Twitter) Follow(followerId int, followeeId int)  {
    followees, ok := this.followMap[followerId] 
    if !ok {
        followees = make(map[int]struct{})
        this.followMap[followerId] = followees
    }
    
    followees[followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if followees, ok := this.followMap[followerId]; ok {
        delete(followees, followeeId)
    }
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

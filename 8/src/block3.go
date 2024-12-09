package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 用户结构体
type User struct {
	Address string // 用户地址
	Tokens  int    // 用户持币数量
	Votes   int    // 用户被投票数量
}

// 全局变量
var (
	users            []User // 用户列表
	witnessCandidates []User // 候选见证人列表
)

// 初始化用户
func initUsers() {
	users = []User{
		{"User1", 1000, 0},
		{"User2", 800, 0},
		{"User3", 1500, 0},
		{"User4", 2000, 0},
		{"User5", 1200, 0},
		{"User6", 500, 0},
		{"User7", 700, 0},
		{"User8", 900, 0},
		{"User9", 300, 0},
		{"User10", 1100, 0},
		{"User11", 600, 0},
		{"User12", 400, 0},
	}
}

// 用户投票
func vote(voterIndex, candidateIndex int, voteTokens int) {
	if voteTokens > users[voterIndex].Tokens {
		fmt.Println("投票失败：投票数量超过持有代币数量")
		return
	}

	// 扣除投票者代币
	users[voterIndex].Tokens -= voteTokens
	// 增加候选人票数
	users[candidateIndex].Votes += voteTokens
	fmt.Printf("%s 投票给 %s，票数增加 %d\n", users[voterIndex].Address, users[candidateIndex].Address, voteTokens)
}

// 用户撤销投票
func revokeVote(voterIndex, candidateIndex int, revokeTokens int) {
	// 减少候选人票数
	if revokeTokens > users[candidateIndex].Votes {
		fmt.Println("撤票失败：撤票数量超过候选人的票数")
		return
	}
	users[candidateIndex].Votes -= revokeTokens
	// 退回给投票者代币
	users[voterIndex].Tokens += revokeTokens
	fmt.Printf("%s 撤销对 %s 的投票，票数减少 %d\n", users[voterIndex].Address, users[candidateIndex].Address, revokeTokens)
}

// 按票数选出见证人
func electWitnesses() {
	// 排序，票数从高到低
	sort.Slice(users, func(i, j int) bool {
		return users[i].Votes > users[j].Votes
	})
	// 取前10名作为见证人
	witnessCandidates = users[:10]
	fmt.Println("选出的见证人：")
	for _, witness := range witnessCandidates {
		fmt.Printf("Address: %s, Votes: %d\n", witness.Address, witness.Votes)
	}
}

// 随机化见证人顺序并挖矿
func shuffleAndMine() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(witnessCandidates), func(i, j int) {
		witnessCandidates[i], witnessCandidates[j] = witnessCandidates[j], witnessCandidates[i]
	})
	fmt.Println("见证人随机顺序：")
	for _, witness := range witnessCandidates {
		fmt.Printf("Address: %s, Votes: %d\n", witness.Address, witness.Votes)
	}

	// 模拟区块生成
	for i, witness := range witnessCandidates {
		fmt.Printf("见证人 %s 打包了区块 #%d\n", witness.Address, i+1)
	}
}

func main() {
	// 初始化用户
	initUsers()

	// 模拟投票
	vote(0, 3, 300)
	vote(1, 4, 200)
	vote(2, 3, 500)
	vote(3, 5, 800)

	// 模拟撤票
	revokeVote(0, 3, 100)

	// 选举见证人
	electWitnesses()

	// 随机化见证人顺序并挖矿
	shuffleAndMine()
}

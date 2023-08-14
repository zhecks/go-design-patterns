package state

import "fmt"

type VoteManager struct {
	VoteMap      map[string]string
	VoteCountMap map[string]int
	state        VoteState
}

func NewVoteManager() *VoteManager {
	return &VoteManager{
		VoteMap:      map[string]string{},
		VoteCountMap: map[string]int{},
	}
}

func (manager VoteManager) Vote(user, voteItem string) {
	manager.VoteCountMap[user]++
	count := manager.VoteCountMap[user]
	switch {
	case count == 1:
		manager.state = &NormalVoteState{}
	case count > 1 && count < 5:
		manager.state = &RepeatVoteState{}
	case count >= 5 && count < 8:
		manager.state = &SpiteVoteState{}
	case count >= 8:
		manager.state = &BlackVoteState{}
	}
	manager.state.Vote(user, voteItem, manager)
}

type VoteState interface {
	Vote(user, voteItem string, voteManager VoteManager)
}

type NormalVoteState struct {
}

func (state *NormalVoteState) Vote(user, voteItem string, voteManager VoteManager) {
	voteManager.VoteMap[user] = voteItem
	fmt.Println("恭喜你投票成功")
}

type RepeatVoteState struct {
}

func (state *RepeatVoteState) Vote(_, _ string, _ VoteManager) {
	fmt.Println("请不要重复投票")
}

type SpiteVoteState struct {
}

func NewSpiteVoteState() *SpiteVoteState {
	return &SpiteVoteState{}
}

func (state *SpiteVoteState) Vote(user, _ string, voteManager VoteManager) {
	delete(voteManager.VoteMap, user)
	fmt.Println("你有恶意刷票行为，取消投票资格")
}

type BlackVoteState struct {
}

func (state *BlackVoteState) Vote(_, _ string, _ VoteManager) {
	fmt.Println("进入黑名单，将禁止登录和使用本系统")
}

package Design_Search_Autocomplete_System

import (
	"sort"
)

type Sentence struct {
	v      string
	weight int
}
type Trie struct {
	v        string
	top3     []Sentence
	children map[rune]*Trie
}

type AutocompleteSystem struct {
	trie         *Trie
	currNode     *Trie
	currSentence string
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	node := &Trie{
		children: make(map[rune]*Trie),
	}
	as := AutocompleteSystem{
		trie:     node,
		currNode: node,
	}
	for i, sentence := range sentences {
		as.add(sentence, times[i])
	}

	return as
}

func (this *AutocompleteSystem) Input(c byte) []string {
	this.currSentence += string(c)
	if rune(c) == '#' {
		this.currNode = this.trie
		this.add(this.currSentence, 1)
		this.currSentence = ""
		return []string{}
	}

	res := make([]string, 0, 3)
	for _, s := range this.currNode.children[rune(c)].top3 {
		res = append(res, s.v)
	}
	return res
}

func (this *AutocompleteSystem) add(sentence string, weight int) {
	currNode := this.trie
	for i := 0; i < len(sentence); i++ {
		n, ok := currNode.children[rune(sentence[i])]
		if !ok {
			n = &Trie{
				v: sentence[:i+1],
			}
		}
		currNode = n
		if i == len(sentence)-1 {
			if len(n.top3) > 1 {
				n.top3[0].weight += weight
			} else {
				n.top3 = []Sentence{
					{
						v:      sentence,
						weight: weight,
					},
				}
			}
		}
	}
}

func (this *AutocompleteSystem) rebuild() {
	_ = this.getTop3(this.trie)
}

func (this *AutocompleteSystem) getTop3(n *Trie) []Sentence {
	var list []Sentence
	if len(n.children) == 0 {
		return n.top3
	}
	for _, ch := range n.children {
		chTop := this.getTop3(ch)
		list = append(list, chTop...)
	}
	// sort list and return top 3
	sort.Slice(list, func(i, j int) bool {
		return list[i].weight < list[j].weight
	})
	n.top3 = list[:3]
	return n.top3
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

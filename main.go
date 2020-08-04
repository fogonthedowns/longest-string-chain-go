package main

// custom implementation of set data structure since no native implementation in go-lang
type Set map[string]struct{}

func (ws Set) add(word ...string) Set {
	for _, v := range word {
		ws[v] = struct{}{}
	}
	return ws
}

func (ws Set) contains(word string) bool {
	_, ok := ws[word]
	return ok
}

func longestStrChain(words []string) int {
	set := Set{}.add(words...)
	mapWordToChainLen := make(map[string]int)
	count := 0

	for _, word := range words {
		count = max(count, trace(set, mapWordToChainLen, word))
	}

	return count
}

func trace(ws Set, mapWordToChainLen map[string]int, word string) int {
	// if exist in map, use it, it represents the max chain len for that word.
	if _, ok := mapWordToChainLen[word]; ok {
		return mapWordToChainLen[word]
	}

	//maxLen represents the max chainLen of word's predecessors
	maxLen := 0
	// loops over all permutations of a string
	for i := 0; i < len(word); i++ {
		permutation := word[:i] + word[i+1:]
		// if word set does not contain word, there is no need to trace its predecessors further
		if ws.contains(permutation) {
			maxLen = max(trace(ws, mapWordToChainLen, permutation), maxLen)
		}
	}
	// the very first running of this statement is when word got broken down to the least characters found in words list.
	mapWordToChainLen[word] = maxLen + 1
	return mapWordToChainLen[word]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

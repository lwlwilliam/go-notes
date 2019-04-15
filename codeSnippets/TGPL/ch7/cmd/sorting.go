package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch7/sorting"
	"sort"
	"fmt"
)

func main()  {
	sorting.PrintTracks(sorting.Tracks)
	fmt.Println()

	sort.Sort(sorting.ByArtist(sorting.Tracks))
	sorting.PrintTracks(sorting.Tracks)
	fmt.Println()

	sort.Sort(sort.Reverse(sorting.Tracks))
	sorting.PrintTracks(sorting.Tracks)
	fmt.Println()
}

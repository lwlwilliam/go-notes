package sorting

import "time"

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type ByArtist []*Track

func (x ByArtist) Len() int  {
	return len(x)
}

func (x ByArtist) Less(i, j int) bool  {
	return x[i].Artist < x[j].Artist
}

func (x ByArtist) Swap(i, j int)  {
	x[i], x[j] = x[j], x[i]
}

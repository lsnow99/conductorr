package series

import "testing"

func TestBasicGetEpisodes(t *testing.T) {
	tvMaze := TVMaze{}
	episodes, err := tvMaze.GetEpisodes("tt3032476")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(episodes)
}

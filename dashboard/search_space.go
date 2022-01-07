package dashboard

import (
	"github.com/c-bata/goptuna"
	"sync"
)

var (
	searchSpaceMu sync.Mutex
	unionSearchSpaceCache map[int]SearchSpace = make(map[int]SearchSpace, 1)
	intersectionSearchSpaceCache map[int]SearchSpace = make(map[int]SearchSpace, 1)
)

type SearchSpace []*struct{
	paramName string
	distribution interface{}
}

type SearchSpaceCache map[]

func (s SearchSpace) Add(trial goptuna.FrozenTrial) {

}

func getSearchSpace(studyID int)  {
	searchSpaceMu.Lock()
	defer searchSpaceMu.Unlock()

	unionSearchSpaceCache[]
}

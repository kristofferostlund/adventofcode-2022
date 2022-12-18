package location

import "fmt"

type Loc [2]int

func (l Loc) Add(other Loc) Loc {
	return Loc{l[0] + other[0], l[1] + other[1]}
}

func (l Loc) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", l[0], l[1])
}

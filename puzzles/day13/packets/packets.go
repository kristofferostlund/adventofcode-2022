package packets

import (
	"strconv"
	"strings"
)

type cType string

const (
	cInt    cType = "int"
	cPacket cType = "packet"
)

var (
	bTrue  = true
	bFalse = false
	True   = &bTrue  // Snakey.
	False  = &bFalse // Snakey.
)

type Packet struct {
	ints    map[int]int
	packets map[int]*Packet
	cTypes  []cType
	parent  *Packet
}

func newPacket(parent *Packet) *Packet {
	return &Packet{
		ints:    make(map[int]int),
		packets: make(map[int]*Packet),
		cTypes:  make([]cType, 0),
		parent:  parent,
	}
}

func (p *Packet) addInt(val int) {
	p.ints[len(p.cTypes)] = val
	p.cTypes = append(p.cTypes, cInt)
}

func (p *Packet) addPacket(nextPack *Packet) {
	p.packets[len(p.cTypes)] = nextPack
	p.cTypes = append(p.cTypes, cPacket)
}

func (p *Packet) String() string {
	sb := &strings.Builder{}
	sb.WriteString("[")

	for i, c := range p.cTypes {
		if c == cInt {
			sb.WriteString(strconv.Itoa(p.ints[i]))
		} else {
			sb.WriteString(p.packets[i].String())
		}

		if i < len(p.cTypes)-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func (p *Packet) Compare(other *Packet) bool {
	return p.checkOrderState(other, 0) == correct
}

type orderState string

const (
	correct   orderState = "correct"
	incorrect orderState = "incorrect"
	undecided orderState = "undecided"
)

func (p *Packet) checkOrderState(other *Packet, nestLevel int) orderState {
	lLen := len(p.cTypes)
	rLen := len(other.cTypes)

	max := lLen
	if rLen > max {
		max = rLen
	}

	for i := 0; i < max; i++ {
		if i == lLen {
			return correct
		}
		if i == rLen {
			return incorrect
		}

		if p.cTypes[i] == cInt && other.cTypes[i] == cInt {
			a, b := p.ints[i], other.ints[i]
			if a < b {
				return correct
			} else if b < a {
				return incorrect
			}

			// They are the same, keep checking
			continue
		}

		var a, b *Packet
		c := p.cTypes[i]
		switch {
		case c == cPacket && other.cTypes[i] == cPacket:
			a, b = p.packets[i], other.packets[i]
		case c == cInt && other.cTypes[i] == cPacket:
			a = newPacket(p)
			a.addInt(p.ints[i])

			b = other.packets[i]
		case c == cPacket && other.cTypes[i] == cInt:
			a = p.packets[i]

			b = newPacket(other)
			b.addInt(other.ints[i])
		}

		state := a.checkOrderState(b, nestLevel+1)
		if state == undecided {
			continue
		}

		return state
	}
	if nestLevel > 0 {
		return undecided
	}

	return correct
}

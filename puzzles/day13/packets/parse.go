package packets

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(reader io.Reader) ([]*Packet, error) {
	packets := make([]*Packet, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		packet, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line %q: %w", line, err)
		}
		packets = append(packets, packet)
	}

	return packets, nil
}

func parseLine(line string) (*Packet, error) {
	topLevelPacket := newPacket(nil)
	packet := topLevelPacket
	sb := &strings.Builder{}

	for _, r := range line[1:] {
		switch {
		case r == '[':
			next := newPacket(packet)
			packet.addPacket(next)
			packet = next
		case r == ']':
			if sb.Len() > 0 {
				if err := parseAddInt(packet, sb.String()); err != nil {
					return nil, fmt.Errorf("parsing value: %w", err)
				}
				sb.Reset()
			}
			packet = packet.parent
		case '0' <= r && r <= '9':
			sb.WriteRune(r)
		case r == ',':
			if sb.Len() > 0 {
				if err := parseAddInt(packet, sb.String()); err != nil {
					return nil, fmt.Errorf("parsing value: %w", err)
				}
				sb.Reset()
			}
		default:
			return nil, fmt.Errorf("illegal character: %q", string(r))
		}
	}

	return topLevelPacket, nil
}

func parseAddInt(packet *Packet, valStr string) error {
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fmt.Errorf("parsing value: %w", err)
	}

	packet.addInt(val)

	return nil
}

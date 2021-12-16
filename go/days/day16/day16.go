package day16

import (
	"aoc-2021/framework"
	"aoc-2021/util/math/integer"
	"bufio"
	"encoding/hex"
	"math/bits"
	"strings"
)

const (
	PacketSum     = uint8(0)
	PacketProduct = uint8(1)
	PacketMinimum = uint8(2)
	PacketMaximum = uint8(3)
	PacketLiteral = uint8(4)
	PacketGreater = uint8(5)
	PacketLess    = uint8(6)
	PacketEqual   = uint8(7)
)

type BitStream struct {
	bits     []uint8
	position int
}

func (s *BitStream) CurrentBit() bool {
	step := s.position / 8
	offset := s.position % 8
	currentByte := s.bits[step]
	mask := bits.RotateLeft8(uint8(1), 7-offset)
	return bits.OnesCount8(currentByte&mask) > 0
}

func (s *BitStream) ReadBits(num int) uint64 {
	out := uint64(0)

	for i := 0; i < num; i++ {
		r := s.CurrentBit()
		s.position++
		out = bits.RotateLeft64(out, 1)
		if r {
			out += uint64(1)
		}
	}
	return out
}

type Packet struct {
	bitSequence *BitStream
	version     uint8
	typeID      uint8
	value       uint64
	subPackets  []Packet
}

func (p *Packet) ParseLiteralValue() {
	complete := false
	for !complete {
		chunk := p.bitSequence.ReadBits(5)
		if bits.OnesCount64(chunk&uint64(16)) == 0 {
			complete = true
		}
		p.value = bits.RotateLeft64(p.value, 4)
		p.value += chunk & uint64(15)
	}
}

func (p *Packet) ParseBitCountSubPackets() {
	bitsToRead := p.bitSequence.ReadBits(15)
	bitsThreshold := p.bitSequence.position + int(bitsToRead)
	p.subPackets = make([]Packet, 0)
	for p.bitSequence.position < bitsThreshold {
		pp := Packet{bitSequence: p.bitSequence}
		pp.Read()
		p.subPackets = append(p.subPackets, pp)
	}
}

func (p *Packet) ParseSubPacketCountSubPackets() {
	subPacketCount := int(p.bitSequence.ReadBits(11))
	p.subPackets = make([]Packet, subPacketCount)
	for i := 0; i < subPacketCount; i++ {
		pp := Packet{bitSequence: p.bitSequence}
		pp.Read()
		p.subPackets[i] = pp
	}
}

func (p *Packet) Read() {
	p.version = uint8(p.bitSequence.ReadBits(3))
	p.typeID = uint8(p.bitSequence.ReadBits(3))
	if p.typeID == PacketLiteral {
		p.ParseLiteralValue()
	} else {
		lengthType := p.bitSequence.ReadBits(1)
		if lengthType == uint64(0) {
			p.ParseBitCountSubPackets()
		} else {
			p.ParseSubPacketCountSubPackets()
		}
	}
}

func (p *Packet) Solve() int {
	if p.typeID == PacketLiteral {
		return int(p.value)
	}
	values := make([]int, len(p.subPackets))
	for i, pp := range p.subPackets {
		values[i] = pp.Solve()
	}
	switch p.typeID {
	case PacketSum:
		return integer.Sum(values)
	case PacketProduct:
		return integer.Product(values)
	case PacketMinimum:
		return integer.MinSlice(values)
	case PacketMaximum:
		return integer.MaxSlice(values)
	case PacketGreater:
		if values[0] > values[1] {
			return 1
		}
	case PacketLess:
		if values[0] < values[1] {
			return 1
		}
	case PacketEqual:
		if values[0] == values[1] {
			return 1
		}
	}
	return 0
}

func (p *Packet) VersionSum() int {
	s := int(p.version)
	for _, pp := range p.subPackets {
		s += pp.VersionSum()
	}
	return s
}

type Puzzle struct {
	framework.PuzzleBase
	packet Packet
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.versionSum,
		"2": p.solvedValue,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	scanner.Scan()
	line := scanner.Text()
	b, _ := hex.DecodeString(strings.ToLower(line))
	p.packet = Packet{bitSequence: &BitStream{bits: b, position: 0}}
}

func (p *Puzzle) versionSum() int {
	p.packet.Read()
	return p.packet.VersionSum()
}

func (p *Puzzle) solvedValue() int {
	p.packet.Read()
	return p.packet.Solve()
}

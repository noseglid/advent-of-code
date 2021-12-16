package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/noseglid/advent-of-code/util"
)

type Op interface {
	Value(v ...uint64) uint64
	String() string
}

func NewOp(typeID uint64) Op {
	switch typeID {
	case 0:
		return SumOp{}
	case 1:
		return ProductOp{}
	case 2:
		return MinimumOp{}
	case 3:
		return MaximumOp{}
	case 5:
		return GreaterOp{}
	case 6:
		return LessOp{}
	case 7:
		return EqualOp{}
	}
	panic("invalid type id")
}

type SumOp struct{}

func (SumOp) Value(v ...uint64) uint64 {
	sum := uint64(0)
	for _, vv := range v {
		sum += vv
	}
	return sum
}
func (SumOp) String() string {
	return "sum"
}

type ProductOp struct{}

func (ProductOp) Value(v ...uint64) uint64 {
	prod := uint64(1)
	for _, vv := range v {
		prod *= vv
	}
	return prod
}
func (ProductOp) String() string {
	return "prd"
}

type MinimumOp struct{}

func (MinimumOp) Value(v ...uint64) uint64 {
	min := v[0]
	for _, vv := range v[1:] {
		if vv < min {
			min = vv
		}
	}
	return min
}
func (MinimumOp) String() string {
	return "min"
}

type MaximumOp struct{}

func (MaximumOp) Value(v ...uint64) uint64 {
	max := v[0]
	for _, vv := range v[1:] {
		if vv > max {
			max = vv
		}
	}
	return max
}

func (MaximumOp) String() string {
	return "max"
}

type LiteralOp struct{}

func (LiteralOp) Value(v ...uint64) uint64 {
	return 0
}

func (LiteralOp) String() string {
	return "lit"
}

type GreaterOp struct{}

func (GreaterOp) Value(v ...uint64) uint64 {
	if len(v) != 2 {
		panic("greater op: invalid argument length")
	}
	if v[0] > v[1] {
		return 1
	}
	return 0
}

func (GreaterOp) String() string {
	return "gre"
}

type LessOp struct{}

func (LessOp) Value(v ...uint64) uint64 {
	if len(v) != 2 {
		panic("less op: invalid argument length")
	}
	if v[0] < v[1] {
		return 1
	}
	return 0
}

func (LessOp) String() string {
	return "les"
}

type EqualOp struct{}

func (EqualOp) Value(v ...uint64) uint64 {
	if len(v) != 2 {
		panic("less op: invalid argument length")
	}
	if v[0] == v[1] {
		return 1
	}
	return 0
}

func (EqualOp) String() string {
	return "eql"
}

type PacketType interface {
	Value() uint64
}

type ParseLimiter interface {
	Done(npackets, nbits uint64) bool
	String() string
}

type BitParseLimiter struct {
	TotalBits uint64
}

func (b BitParseLimiter) String() string {
	return fmt.Sprintf("Bits{%d}", b.TotalBits)
}

func (b BitParseLimiter) Done(npackets, nbits uint64) bool {
	return nbits >= b.TotalBits
}

type PacketsParseLimiter struct {
	TotalPackets uint64
}

func (p PacketsParseLimiter) String() string {
	return fmt.Sprintf("Packets{%d}", p.TotalPackets)
}

func (p PacketsParseLimiter) Done(npackets, nbits uint64) bool {
	return npackets >= p.TotalPackets
}

type packet struct {
	version uint64
	typeID  Op
}

type operatorPacket struct {
	packet

	subpackets []PacketType
}

func (o operatorPacket) Value() uint64 {
	var values []uint64
	for _, subpackets := range o.subpackets {
		values = append(values, subpackets.Value())
	}

	return o.typeID.Value(values...)
}

func (c operatorPacket) String() string {
	return fmt.Sprintf("OperatorPacket{v=%d,t=%s,nsub=%d}", c.version, c.typeID, len(c.subpackets))
}

type literalPacket struct {
	packet

	number uint64
}

func (l literalPacket) Value() uint64 {
	return l.number
}

func (l literalPacket) String() string {
	return fmt.Sprintf("LiteralPacket{v=%d,t=%s,value=%d}", l.version, l.typeID, l.number)
}

func ParseNumber(bin string) (uint64, uint64, string) {
	var sb strings.Builder

	consumed := uint64(0)

	for {
		b := bin[0]
		sb.WriteString(bin[1:5])
		bin = bin[5:]
		consumed += 5

		if b == '0' {
			break
		}
	}

	return MustParseUint(sb.String(), 2, 64), consumed, bin
}

func ParseLiteralPacket(bin string) (uint64, uint64, string, *literalPacket) {
	version, _, bits, remaining := parseHeader(bin)
	bin = remaining
	consumed := bits
	lpkt := literalPacket{
		packet: packet{
			version: version,
			typeID:  LiteralOp{},
		},
	}

	number, bits, remaining := ParseNumber(bin)
	consumed += bits
	lpkt.number = number

	return 1, consumed, remaining, &lpkt
}

func ParseControlPacket(bin string) (uint64, uint64, string, *operatorPacket) {
	version, typeID, bits, remaining := parseHeader(bin)
	bin = remaining
	consumed := bits
	cpkt := operatorPacket{
		packet: packet{
			version: version,
			typeID:  NewOp(typeID),
		},
	}

	lengthTypeID := MustParseUint(bin[:1], 2, 2)
	bin = bin[1:]
	consumed += 1

	var limiter ParseLimiter
	switch lengthTypeID {
	case 0:
		limiter = &BitParseLimiter{MustParseUint(bin[:15], 2, 16)}
		bin = bin[15:]
		consumed += 15
	case 1:
		limiter = &PacketsParseLimiter{MustParseUint(bin[:11], 2, 16)}
		bin = bin[11:]
		consumed += 11
	}

	npackets, bits, remaining, subpackets := parsePackets(bin, limiter)
	cpkt.subpackets = subpackets
	consumed += bits

	return npackets, consumed, remaining, &cpkt
}

func IsControl(bin string) bool {
	return MustParseUint(bin[3:6], 2, 64) != 4
}

func MustParseUint(s string, base, bitSize int) uint64 {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return v
}

func parseHeader(bin string) (uint64, uint64, uint64, string) {
	version := MustParseUint(bin[0:3], 2, 64)
	typeID := MustParseUint(bin[3:6], 2, 64)

	return version, typeID, 6, bin[6:]
}

func parsePackets(bin string, limiter ParseLimiter) (uint64, uint64, string, []PacketType) {
	var pkts []PacketType

	totalPackets, totalBits := uint64(0), uint64(0)
	for {
		if IsControl(bin) {
			_, bits, remaining, cpkt := ParseControlPacket(bin)
			pkts = append(pkts, cpkt)
			bin = remaining
			totalPackets++
			totalBits += bits
		} else {
			_, bits, remaining, lpkt := ParseLiteralPacket(bin)
			pkts = append(pkts, lpkt)
			bin = remaining
			totalPackets++
			totalBits += bits
		}

		if limiter.Done(totalPackets, totalBits) {
			break
		}
	}

	return totalPackets, totalBits, bin, pkts

}

func toBitstring(code string) string {
	var sb strings.Builder
	for i := 0; i < len(strings.TrimSpace(code)); i++ {
		sub := MustParseUint(string(code[i]), 16, 4)
		sb.WriteString(fmt.Sprintf("%04b", sub))
	}

	return sb.String()
}

func packetSums(packet PacketType) uint64 {
	s := uint64(0)
	switch p := packet.(type) {
	case *literalPacket:
		s += p.version
	case *operatorPacket:
		s += p.version
		for _, subpacket := range p.subpackets {
			s += packetSums(subpacket)
		}
	}

	return s
}

//lint:ignore U1000 used for debugging
func printPacket(packet PacketType, level int) {
	const spacesPerLevel = 2
	switch p := packet.(type) {
	case *literalPacket:
		fmt.Printf("%*s%s\n", level*spacesPerLevel, "", p)
	case *operatorPacket:
		fmt.Printf("%*s%s\n", level*spacesPerLevel, "", p)
		for _, subpacket := range p.subpackets {
			printPacket(subpacket, level+1)
		}
	}
}

func main() {
	input := "2021/Day16/input"
	code := util.GetFile(input)
	// code := "880086C3E88112"

	_, _, _, packets := parsePackets(toBitstring(code), PacketsParseLimiter{1})
	log.Printf("Part 1: Sums of all packet versions is %d", packetSums(packets[0]))
	log.Printf("Part 2: The value of packet is %d", packets[0].Value())

}

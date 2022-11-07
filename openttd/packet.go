package openttd

import "encoding/binary"

type Packet struct {
	packetType uint16
	data       []byte
}

type PacketReader struct {
	packet *Packet
}

func (r *PacketReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = r.packet.data[i]
	}

	r.packet.data = r.packet.data[len(p):]

	return len(p), nil
}

func (r *PacketReader) ReadString(max uint) (string, int, error) {
	l := max
	o := make([]byte, max)

	for i := uint(0); i < max-1; i++ {
		o[i] = r.packet.data[i]
		if o[i] == 0x00 {
			l = i + 1
		}
	}

	r.packet.data = r.packet.data[l:]

	return string(o[:l]), int(l), nil
}

type PacketWriter struct {
	packet *Packet
}

func (w *PacketWriter) Write(p []byte) (int, error) {
	w.packet.data = append(w.packet.data, p...)

	return len(p), nil
}

func (p *Packet) Reader() *PacketReader {
	return &PacketReader{
		packet: p,
	}
}

func (p *Packet) Writer() *PacketWriter {
	return &PacketWriter{
		packet: p,
	}
}

func (p *Packet) Bytes() []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, p.packetType)

	b = append(b, p.data...)

	return b
}

func CreatePacket(t uint16) Packet {
	p := Packet{
		packetType: t,
		data:       make([]byte, 0),
	}

	return p
}

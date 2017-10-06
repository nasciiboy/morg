package nirvana

// Simple Cell buffer for marshaling Cells.

import (
	"errors"
)

type Cell struct {
  Ch    rune
  Color uint8
  Attrs uint8
  Touch bool
}

// A Buffer is a variable-sized buffer of Cell with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type CellBuffer struct {
	buf       []Cell     // contents are the bytes buf[off : len(buf)]
	off       int          // read at &buf[off], write at &buf[len(buf)]
	bootstrap [64]Cell   // memory to hold first slice; helps small buffers avoid allocation.
  color     uint8
  attrs     uint8
}



// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("Cells.CellBuffer: too large")

func (b *CellBuffer) SetFace( face uint64 ){
  b.attrs, b.color, _, _ = extractData( face )
}

// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like Read, Write, Reset, or Truncate).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
func (b *CellBuffer) Data() []Cell { return b.buf[b.off:] }

// String returns the contents of the unread portion of the buffer
// as a string. If the CellBuffer is a nil pointer, it returns "<nil>".
func (b *CellBuffer) String() (str string) {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}

  runes := make( []rune, b.Len() )
  for i, cell := range b.buf[b.off:] {
    runes[i] = cell.Ch
  }

	return string( runes )
}

// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
func (b *CellBuffer) Len() int { return len(b.buf) - b.off }

// Cap returns the capacity of the buffer's underlying byte slice, that is, the
// total space allocated for the buffer's data.
func (b *CellBuffer) Cap() int { return cap(b.buf) }

// Truncate discards all but the first n unread bytes from the buffer
// but continues to use the same allocated storage.
// It panics if n is negative or greater than the length of the buffer.
func (b *CellBuffer) Truncate(n int) {
	switch {
	case n < 0 || n > b.Len():
		panic("Cells.CellBuffer: truncation out of range")
	case n == 0:
		// Reuse buffer space.
		b.off = 0
	}
	b.buf = b.buf[0 : b.off+n]
}

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as Truncate(0).
func (b *CellBuffer) Reset() { b.Truncate(0) }

// grow grows the buffer to guarantee space for n more bytes.
// It returns the index where bytes should be written.
// If the buffer can't grow it will panic with ErrTooLarge.
func (b *CellBuffer) grow(n int) int {
	m := b.Len()
	// If buffer is empty, reset to recover space.
	if m == 0 && b.off != 0 {
		b.Truncate(0)
	}
	if len(b.buf)+n > cap(b.buf) {
		var buf []Cell
		if b.buf == nil && n <= len(b.bootstrap) {
			buf = b.bootstrap[0:]
		} else if m+n <= cap(b.buf)/2 {
			// We can slide things down instead of allocating a new
			// slice. We only need m+n <= cap(b.buf) to slide, but
			// we instead let capacity get twice as large so we
			// don't spend all our time copying.
			copy(b.buf[:], b.buf[b.off:])
			buf = b.buf[:m]
		} else {
			// not enough space anywhere
			buf = makeSlice(2*cap(b.buf) + n)
			copy(buf, b.buf[b.off:])
		}
		b.buf = buf
		b.off = 0
	}
	b.buf = b.buf[0 : b.off+m+n]
	return b.off + m
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to the
// buffer without another allocation.
// If n is negative, Grow will panic.
// If the buffer can't grow it will panic with ErrTooLarge.
func (b *CellBuffer) Grow(n int) {
	if n < 0 {
		panic("Cells.CellBuffer.Grow: negative count")
	}
	m := b.grow(n)
	b.buf = b.buf[0:m]
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (b *CellBuffer) Write(p []Cell) (n int, err error) {
	m := b.grow(len(p))
	return copy(b.buf[m:], p), nil
}

func (b *CellBuffer) WriteU64( u uint64 ) error {
	m := b.grow(1)
	b.buf[m] = extractCell( u )
	return nil
}

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.
func (b *CellBuffer) WriteString(s string) (n int, err error) {
  p := make( []Cell, 0, len( s ) )
  for _, r := range( s ) {
    p = append( p, Cell{ Ch: r, Color: b.color, Attrs: b.attrs })
  }

	return b.Write( p )
}

func (b *CellBuffer) WriteRune( r rune ) error {
	m := b.grow(1)
	b.buf[m] = Cell{ Ch: r, Color: b.color, Attrs: b.attrs }
	return nil
}

func (b *CellBuffer) ReadFrom(r CellBuffer) (n int, err error) {
	return b.Write( r.Data() )
}


// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []Cell {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return make([]Cell, n)
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read or write method.
func (b *CellBuffer) Next(n int) []Cell {
	m := b.Len()
	if n > m {
		n = m
	}
	data := b.buf[b.off : b.off+n]
	b.off += n

	return data
}

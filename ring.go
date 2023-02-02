package ring

type Ring[T any] struct {
	size uint
	mask uint
	buf  []T
}

func NewRing[T any](size uint) *Ring[T] {
	r := new(Ring[T])
	if size > 0 && (size&(size-1) == 0) {
		r.reset(size)
	} else {
		var n uint = 1
		for n < size {
			n <<= 1
		}
		r.reset(n)
	}
	return r
}

func (c *Ring[T]) reset(size uint) {
	*c = Ring[T]{
		mask: size - 1,
		size: size,
		buf:  make([]T, size),
	}
}

func (c *Ring[T]) Size() uint {
	return c.size
}

func (c *Ring[T]) Get(index uint) T {
	return c.buf[index&c.mask]
}

func (c *Ring[T]) Put(index uint, val T) uint {
	i := index & c.mask
	c.buf[i] = val
	return i
}

func (c *Ring[T]) Del(index uint) T {
	i := index & c.mask
	val := c.buf[i]
	c.buf[i] = *new(T)
	return val
}

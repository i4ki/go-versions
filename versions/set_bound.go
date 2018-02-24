package versions

type setBound struct {
	v  Version
	op setBoundOp
}

func (s setBound) Has(v Version) bool {
	switch s.op {
	case setBoundGT:
		return v.GreaterThan(s.v)
	case setBoundGTE:
		return v.GreaterThan(s.v) || v.Same(s.v)
	case setBoundLT:
		return v.LessThan(s.v)
	case setBoundLTE:
		return v.LessThan(s.v) || v.Same(s.v)
	default:
		// Should never happen because the above is exhaustive
		panic("invalid setBound operator")
	}
}

// NewerThan returns a set containing all versions greater than the given
// version, non-inclusive.
func NewerThan(v Version) Set {
	return Set{
		setI: setBound{
			v:  v,
			op: setBoundGT,
		},
	}
}

// OlderThan returns a set containing all versions lower than the given
// version, non-inclusive.
func OlderThan(v Version) Set {
	return Set{
		setI: setBound{
			v:  v,
			op: setBoundLT,
		},
	}
}

// AtLeast returns a set containing all versions greater than or equal to
// the given version.
func AtLeast(v Version) Set {
	return Set{
		setI: setBound{
			v:  v,
			op: setBoundGTE,
		},
	}
}

// AtMost returns a set containing all versions less than or equal to the given
// version, non-inclusive.
func AtMost(v Version) Set {
	return Set{
		setI: setBound{
			v:  v,
			op: setBoundLTE,
		},
	}
}

type setBoundOp rune

const setBoundGT = '>'
const setBoundGTE = '≥'
const setBoundLT = '<'
const setBoundLTE = '≤'

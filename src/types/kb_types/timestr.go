package kb_types

type TimeStr string

func (t TimeStr) String() string {
	return string(t)
}

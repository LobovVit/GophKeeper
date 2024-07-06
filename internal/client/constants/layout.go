package constants

type Layout string

const (
	LayoutDate        Layout = "02.01.2006"
	LayoutDateAndTime Layout = "02.01.2006 15:04:05"
)

func (l Layout) ToString() string {
	return string(l)
}

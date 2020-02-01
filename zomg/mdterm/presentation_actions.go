package mdterm

type DisplayAction interface {
	name() string
}

type DisplayActionNext struct {
}

func (dn DisplayActionNext) name() string {
	return "next"
}

type DisplayActionPrev struct {
}

func (dn DisplayActionPrev) name() string {
	return "prev"
}

type DisplayActionEnd struct {
}

func (dn DisplayActionEnd) name() string {
	return "end"
}

type DisplayActionStart struct {
}

func (dn DisplayActionStart) name() string {
	return "start"
}

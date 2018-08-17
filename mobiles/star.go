package mobiles

import "strconv"

type Star struct {
	Weight int
}

func (s *Star) GetGewicht() int {
	return s.Weight
}

func (s *Star) Balance() float64 {
	return 0
}

func (s *Star) Print() string {
	return "Star: "+ strconv.Itoa(s.Weight)
}
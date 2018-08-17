package mobiles


import "strconv"

type Wire struct {
	AnhL Mobiles
	AnhR Mobiles
	Length int

}

func (s *Wire) GetGewicht() int {
	return s.AnhR.GetGewicht()+s.AnhL.GetGewicht()
}

func (s *Wire) Balance() float64 {
	return float64((s.AnhL.GetGewicht()/s.AnhR.GetGewicht()) * s.Length)
}

func (s *Wire) Print() string {
	return "\n L: Length: "+strconv.Itoa(s.Length)+" Print: "+s.AnhL.Print()+"\n\t R: "+ strconv.FormatFloat(s.Balance(),'E', -1, 64)+" Print: "+ s.AnhR.Print()
}

package encoder

import "errors"

type ReedSolomon struct {
	gf      *GaloiField
	genPoly []int
	nsym    int
}

func (r *ReedSolomon) GeneratorPoly() {
	g := []int{1}
	for i := 0; i < r.nsym; i++ {
		g = r.gf.PolyMultiply(g, []int{1, r.gf.expTable[i]})
	}
	r.genPoly = g
}

func (r *ReedSolomon) EncodeMsg(msgIn []int) ([]int, error) {
	if len(msgIn)+r.nsym > 255 {
		return nil, errors.New("Message too long")
	}
	gen := r.genPoly
	msgOut := make([]int, len(msgIn)+r.nsym)
	copy(msgOut, msgIn)

	for i := 0; i < len(msgIn); i++ {
		coef := msgOut[i]
		if coef != 0 {
			for j := 0; j < len(gen); j++ {
				msgOut[i+j] ^= r.gf.Multiply(gen[j], coef)
			}
		}
	}
	return msgOut, nil
}

func NewReedSolomon(nsym int) *ReedSolomon {
	gf := NewGaloiField()
	rs := &ReedSolomon{gf, []int{}, nsym}
	rs.GeneratorPoly()
	return rs
}

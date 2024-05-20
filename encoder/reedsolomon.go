package encoder

import "errors"

type ReedSolomon struct {
	gf      *GaloiField
	genPoly []int
	nsym    int
}

func (r *ReedSolomon) generatorPoly() {
	g := []int{1}
	for i := 0; i < r.nsym; i++ {
		g = r.gf.polyMultiply(g, []int{1, r.gf.expTable[i]})
	}
	r.genPoly = g
}

func (r *ReedSolomon) encodeMsg(msgIn []int) ([]int, error) {
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
				msgOut[i+j] ^= r.gf.multiply(gen[j], coef)
			}
		}
	}
	return msgOut, nil
}

func newReedSolomon(nsym int) *ReedSolomon {
	gf := newGaloiField()
	rs := &ReedSolomon{gf, []int{}, nsym}
	rs.generatorPoly()
	return rs
}

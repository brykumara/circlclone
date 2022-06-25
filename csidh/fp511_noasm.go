//go:build !amd64 || purego
// +build !amd64 purego

package csidh

func mul512(r, m1 *fp, m2 uint64)            { mul512Generic(r, m1, m2) }
func mul576(r *[9]uint64, m1 *fp, m2 uint64) { mul576Generic(r, m1, m2) }
func cswap512(x, y *fp, choice uint8)        { cswap512Generic(x, y, choice) }
func mulRdc(r, x, y *fp)                     { mulRdcGeneric(r, x, y) }
func Mul512(r, m1 *Fp, m2 uint64)            { Mul512Generic(r, m1, m2) }
func Mul576(r *[9]uint64, m1 *Fp, m2 uint64) { Mul576Generic(r, m1, m2) }
func Cswap512(x, y *Fp, choice uint8)        { Cswap512Generic(x, y, choice) }
func MulRdc(r, x, y *Fp)                     { MulRdcGeneric(r, x, y) }

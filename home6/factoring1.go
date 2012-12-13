package main

//import "log"
import "fmt"
import "math/big"
import "github.com/cznic/mathutil"

func main() {
	N, _ := new(big.Int).SetString("179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581", 0)

	A := mathutil.SqrtBig(N)
	A.Add(A, big.NewInt(1))

	x := new(big.Int).Mul(A, A)
	x = mathutil.SqrtBig(x.Sub(x, N))

	p := new(big.Int).Sub(A, x)
	q := new(big.Int).Add(A, x)

	if new(big.Int).Mul(p, q).Cmp(N) == 0 {
		fmt.Println("Result:", p)
	}
}
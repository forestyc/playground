package diff

import (
	"fmt"
	"github.com/forestyc/playground/cmd/demo/house/repayment"
	"io"
	"math"
	"os"
)

func FundMonthly(fundAmount float64, fundInterestRate []float64, periods int) {
	var fAc []repayment.AverageCapital
	for i := 0; i < len(fundInterestRate); i++ {
		f, _ := repayment.NewAverageCapital(fundAmount, fundInterestRate[i], periods)
		fAc = append(fAc, f)
	}
	f, err := os.Create("diff-monthly.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.WriteString(f, "id, 首套,二套,公积金差值\n")
	for i := 1; i <= 360; i++ {
		repayment1 := fAc[0].Repayment(i)
		repayment2 := fAc[1].Repayment(i)
		fDif := math.Abs(repayment1 - repayment2)
		io.WriteString(f, fmt.Sprintf("%d,%.2f,%.2f,%.2f\n", i, repayment1, repayment2, fDif))
	}
}

func FundInterest(fundAmount float64, fundInterestRate []float64, periods int) {
	//var fAc []repayment.AverageCapital
	//for i := 0; i < len(fundInterestRate); i++ {
	//	f, _ := repayment.NewAverageCapital(fundAmount, fundInterestRate[i], periods)
	//	fAc = append(fAc, f)
	//}
	//fmt.Printf("Interest1: %.2f\tInterest2: %.2f\tdiff: %.2f\n", fAc[0].TotalInterest(), fAc[1].TotalInterest(), math.Abs(fAc[0].TotalInterest()-fAc[1].TotalInterest()))
	pi, _ := repayment.NewPrincipalInterest(600000, 0.04, 360)
	fmt.Println(pi.Repayment(4))
	fmt.Println(pi.Repayment(5))
	fmt.Println(pi.Repayment(6))
}

package main

func CountProfit(data [][][2]int) []int {
	Sum := func(arr []int) int {
		sum := 0
		for _, num := range arr {
			sum += num
		}
		return sum
	}
	Profit := []int{}
	if len(data) == 0 {
		return Profit
	} else {
		Sell := 0
		Spend := 0
		ProfitMonth := []int{}
		for i := 0; i < len(data[0]); i++ { //Bulan Penjualan
			for j := 0; j < len(data); j++ { //Cabang Perusahaan
				Sell = data[j][i][0]
				Spend = data[j][i][1]
				ProfitMonth = append(ProfitMonth, Sell-Spend)
			}
			Profit = append(Profit, Sum(ProfitMonth))
			ProfitMonth = []int{}
		}
	}
	return Profit // TODO: replace this
}

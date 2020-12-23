package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	numbers, err := parseData(string(bytes))
	if err != nil {
		fmt.Println(err)
	}

	am1, ad1, ap1, ad2, ap2, an := numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], 5.0
	bm2, bd1, bp1, bd2, bp2, bn := numbers[5], numbers[6], numbers[7], numbers[8], numbers[9], 5.0
	cp3, cp4, cp1, cp2, cn := numbers[10], numbers[11], numbers[12], numbers[13], 4.0
	cp4 += 0.0

	aClear := (ap1*(ad1*an) + ap2*(ad2*an)) - am1
	bClear := (bp1*(bd1*bn) + bp2*(bd2*bn)) - bm2
	caClear := ((cp1*(ad1*cn) + cp2*(ad2*cn)) - am1) * cp3
	cbClear := ((cp1*(bd1*cn) + cp2*(bd2*cn)) - bm2) * cp3

	profitArr := []float64{aClear, bClear, caClear, cbClear}
	sort.Float64s(profitArr)

	profitMap := make(map[float64]string)
	profitMap[aClear] = "A"
	profitMap[bClear] = "Б"
	profitMap[caClear] = "В(а)"
	profitMap[cbClear] = "В(б)"

	fmt.Println(fmt.Sprintf("Чистая прибыль варианта А: %.2f", aClear))
	fmt.Println(fmt.Sprintf("Чистая прибыль варианта Б: %.2f", bClear))
	fmt.Println(fmt.Sprintf("Чистая прибыль варианта В(а): %.2f", caClear))
	fmt.Println(fmt.Sprintf("Чистая прибыль варианта В(б): %.2f", cbClear))
	for k, v := range profitMap {
		if k == profitArr[len(profitArr)-1] {
			fmt.Println(fmt.Sprintf("Наибольший чистый доход принесет вариант %s: %.2f", v, k))
		}
	}

}

func parseData(data string) ([]float64, error) {
	nubmers := strings.Split(data, " ")
	result := make([]float64, 0)
	for _, v := range nubmers {
		t, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

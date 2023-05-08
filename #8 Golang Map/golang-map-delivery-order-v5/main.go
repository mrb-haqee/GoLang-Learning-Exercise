package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	if day=="minggu"{
		return make(map[string]float32)
	}
	Hari := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	//Check Lokasi
	HariPengirimanLokasi := map[string][]string{
		"JKT": Hari,
		"BDG": {"rabu", "kamis", "sabtu"},
		"BKS": {"selasa", "kamis", "jumat"},
		"DPK": {"senin", "selasa"},
	}
	LokasiPengiriman := []string{}
	check := false
	for key, days := range HariPengirimanLokasi {
		for i := 0; i < len(days); i++ {
			check = strings.Contains(days[i], day)
			if check {
				LokasiPengiriman = append(LokasiPengiriman, key)
				check = false
				break
			}
		}
	}
	//Done Check Lokasi

	//Check Price admin
	BiayaAdmin := []map[string][]any{
		{"price1": {"senin", "rabu", "jumat"}, "price2": {"selasa", "kamis", "sabtu"}},
		{"price1": {0.1}, "price2": {0.05}},
	}
	check = false
	var Price float32
	for {
		n := 0
		for _, days := range BiayaAdmin[0]["price1"] {
			check = strings.Contains(days.(string), day)
			if check {
				n++
				Price = float32(BiayaAdmin[1]["price1"][0].(float64))
				break
			}
		}
		if n == 1 {
			break
		}
		for _, days := range BiayaAdmin[0]["price2"] {
			check = strings.Contains(days.(string), day)
			if check {
				n++
				Price = float32(BiayaAdmin[1]["price2"][0].(float64))
				break
			}
		}
		if n == 1 {
			break
		}
	}

	//Done Check Price Admin
	////------------------------------------------------------------------
	//Main Function
	//selsksi Lokasi
	check = false
	TrueData := []string{}
	for _, strData := range data {
		for _, Lokasi := range LokasiPengiriman {
			check := strings.Contains(strData, Lokasi)
			if check {
				TrueData = append(TrueData, strData)
			}
		}
	}
	//Done seleksi lokasi


	//Final
	SplitTrueData := []string{}
	NamePegawai := []string{}
	PriceLocation := []string{}
	PriceArr := []float32{}
	for _, split := range TrueData {
		SplitTrueData = append(SplitTrueData, strings.Replace(split, ":", "-", 1))
	}
	for _, nama := range SplitTrueData {
		RealNama, rev := "", ""
		RealNama, rev, _ = strings.Cut(nama, ":")
		NamePegawai = append(NamePegawai, RealNama)
		PriceLocation = append(PriceLocation, rev)
	}
	for _, priceAdmin := range PriceLocation {
		price, _, _ := strings.Cut(priceAdmin, ":")
		Ex, _ := strconv.ParseFloat(price, 32)
		PriceArr = append(PriceArr, (float32(Ex) + (Price * float32(Ex))))
	}
	Final := make(map[string]float32)
	for i := 0; i < len(NamePegawai); i++ {
		Final[NamePegawai[i]] = PriceArr[i]
	}
	FinalData := make(map[string]float32)
	sort.Strings(NamePegawai)
	for i := 0; i < len(Final); i++ {
		FinalData[NamePegawai[i]] = Final[NamePegawai[i]]
	}
	//Done
	return FinalData // TODO: replace this
}

func main() {
	data, day := []string{"Andi:Sukirman:15000:DPK", "Anggi:Anggraini:20000:BDG", "Andi:Gunawan:40000:BKS", "Budi:Gunawan:50000:JKT"}, "minggu"
	fmt.Println(DeliveryOrder(data, day))
}

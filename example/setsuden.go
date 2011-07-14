package main

import "github.com/mattn/go-setsuden"

func main() {
	println("■東京の最新計測値")
	pu, _ := setsuden.GetActualUsage("tokyo")
	for _, p := range pu {
		println("時間", p.Datetime)
		println("間隔", p.Duration)
		println("値　", p.Value)
	}
	println()

	println("■東京の最新予測値")
	pu, _ = setsuden.GetEstimatedUsage("tokyo")
	for _, p := range pu {
		println("時間", p.Datetime)
		println("間隔", p.Duration)
		println("値　", p.Value)
	}
	println()

	println("■東京の最新瞬時値")
	pu, _ = setsuden.GetInstantUsage("tokyo")
	for _, p := range pu {
		println("時間", p.Datetime)
		println("間隔", p.Duration)
		println("値　", p.Value)
	}
	println()

	println("■東京の電力供給ピーク値")
	pp, _ := setsuden.GetPeakOfSupply("tokyo")
	for _, p := range pp {
		println("開始", p.Start)
		println("終了", p.End)
		println("値　", p.Value)
	}
	println()

	println("■東京の電力供給ピーク予想値")
	pp, _ = setsuden.GetPeakOfDemand("tokyo")
	for _, p := range pp {
		println("開始", p.Start)
		println("終了", p.End)
		println("値　", p.Value)
	}
	println()
}

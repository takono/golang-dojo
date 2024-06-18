package main

import "fmt"

type BeerInterface interface {
	drink()
}

func toriaezu(b BeerInterface) {
	fmt.Printf("b's type: %T\n", b)
	b.drink()
}

func toriaezuWithPtrInterface(b *BeerInterface) {
	fmt.Printf("b's type: %T\n", b)
	(*b).drink()
}

type ValueReceiverBeer struct {}

func (a ValueReceiverBeer) drink(){
	fmt.Println("Drink ValueReceiverBeer")
}

type PointerReceiverBeer struct {}

func (a *PointerReceiverBeer) drink(){
	fmt.Println("Drink PointerReceiverBeer")
}

func main() {
	vrb := ValueReceiverBeer{}
	prb := PointerReceiverBeer{}

	// Interfaceが引数の場合

	// 値レシーバの場合は値、ポインタのどちらでもinterfaceを満たす
	toriaezu(vrb)
	toriaezu(&vrb)
	
	// ポインタレシーバの場合はポインタ型のみinterfaceを満たす
	toriaezu(&prb)
	// コンパイルエラー
	// toriaezu(prb) 
	
	// Interfaceのポインタ型が引数の場合、interfaceを満たせない
	// コンパイルエラー
	// toriaezuWithPtrInterface(vrb)   
	// toriaezuWithPtrInterface(&vrb)
	// toriaezuWithPtrInterface(prb)  
	// toriaezuWithPtrInterface(&prb)

	// interfaceのポインタ型に変換してから設定する
	beerInterface := new(BeerInterface)
	*beerInterface = &vrb 
	toriaezuWithPtrInterface(beerInterface)

	*beerInterface = &prb
	toriaezuWithPtrInterface(beerInterface)
}
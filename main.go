package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
)
var NumList [10]string 

type Numpad struct {
	Numeros string
	NumerosBase string
	Letras string
	LetrasBase string
	NumerosLista []string
	NumerosBaseLista []string
	LetrasLista []string
	LetrasBaseLista []string
	Clave string
}

func (n *Numpad) GetBaseNumbers() {
  	nums := make([]int, 10)
  	for k , _ := range nums {
  		nums[k] = k
  		NumList[k] = strconv.Itoa(k)
  	}
  	n.NumerosBaseLista = make([]string, 10)
  	
  	for i := 9; i >= 0; i-- {
  		index := rand.Intn(i + 1)
  		n.NumerosBaseLista[i] = strconv.Itoa(nums[index])
  		if index != i {
  			nums[index] = nums[i]
  		}			
  	}
  	n.NumerosBase = strings.Join(n.NumerosBaseLista, "")
}

func (n *Numpad) GetBaseLetters() {
	letters := make([]string, 10)
	n.LetrasBaseLista = make([]string, 10)
	for k, _ := range letters {
		letters[k] = string(97 + k)
		n.LetrasBaseLista[k] =  string(97 + k)
	}

	for i := 9; i >= 0; i-- {
		index := rand.Intn(i + 1)
		n.LetrasBaseLista[i] = letters[index]
		if index != i {
			aux := letters[i]
			letters[index] = letters[i]
			letters[i] = aux
		}
	}

	n.LetrasBase = strings.Join(n.LetrasBaseLista, "")
}

func (n *Numpad) GetNumbers() {
	n.NumerosLista = n.NumerosBaseLista
	n.Numeros = n.NumerosBase
}

func (n *Numpad) GetLetters() {
	n.LetrasLista = n.LetrasBaseLista
	n.Letras = n.LetrasBase
}

func (n *Numpad) GetFinalNumbersAndLettersAndClave(clave string) {
	n.Letras = n.Reverse(n.LetrasBase)
}

func (n *Numpad) Reverse (cadena string) (result string) {
	for _, v := range cadena {
		result = string(v) + result
	}

	return result
}

func (n *Numpad) GetClave(clave string) {

}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := &Numpad{}
	n.GetBaseNumbers()
	n.GetBaseLetters()
	n.GetNumbers()
	n.GetLetters()
	n.GetFinalNumbersAndLettersAndClave("1234")
	fmt.Println(n)
}

func main() {

}

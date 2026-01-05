//Vladislav Ebert
//241RDB316

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	izmers := iegutMasivaIzmeru()
	if izmers < 10 {
		return
	}

	massivs := izveidotMasivu(izmers)
	if massivs == nil {
		return
	}

	daluSkaits := aprekinatGorutinuSkaitu(izmers)
	dalas := sadalitMasivu(massivs, daluSkaits)

	fmt.Printf("\nMasiva izmers:%d\n", izmers)
	fmt.Printf("Gorutinu skaits:%d\n", daluSkaits)

	var GaidasGrupa sync.WaitGroup
	GaidasGrupa.Add(daluSkaits)

	for i := 0; i < daluSkaits; i++ {
		go sakartotDalu(dalas, i, &GaidasGrupa)
	}
	GaidasGrupa.Wait()

	rezultats := apvienotSakartotosMasivus(dalas)

	fmt.Printf("Apakš masīvu pirms sakārtošanas:%v\n", massivs)
	fmt.Printf("Apakš masīvu pec sakārtošanas:%v\n", rezultats)

	if irSakartots(rezultats) {
		fmt.Println(" ")
	}
}

func iegutMasivaIzmeru() int {
	var izmers int
	fmt.Print("Ievadiet masiva elementu skaitu: ")
	if _, kluda := fmt.Scan(&izmers); kluda != nil || izmers < 10 {
		fmt.Println("Error: Skaitlim jabut >= 10")
		return -1
	}
	return izmers
}

func izveidotMasivu(izmers int) []int {
	var izvele int
	fmt.Print("1 - Ievadit pasam || 2 - Generet ")
	if _, kluda := fmt.Scan(&izvele); kluda != nil || (izvele != 1 && izvele != 2) {
		fmt.Println("Error: Ievadiet 1 vai 2")
		return nil
	}

	massivs := make([]int, izmers)
	if izvele == 1 {
		fmt.Printf("Ievadiet %d skaitlus:\n", izmers)
		for i := 0; i < izmers; i++ {
			fmt.Scan(&massivs[i])
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < izmers; i++ {
			massivs[i] = rand.Intn(100)
		}
		fmt.Printf("Generetais masivs: %v\n", massivs)
	}
	return massivs
}

func aprekinatGorutinuSkaitu(izmers int) int {
	daluSkaits := int(math.Sqrt(float64(izmers)))
	if daluSkaits < 4 {
		return 4
	}
	return daluSkaits
}

func sadalitMasivu(massivs []int, daluSkaits int) [][]int {
	dalasIzmers := len(massivs) / daluSkaits
	dalas := make([][]int, daluSkaits)

	for i := 0; i < daluSkaits; i++ {
		sakums := i * dalasIzmers
		beigas := sakums + dalasIzmers
		if i == daluSkaits-1 {
			beigas = len(massivs)
		}
		dalas[i] = make([]int, beigas-sakums)
		copy(dalas[i], massivs[sakums:beigas])
	}
	return dalas
}

func sakartotDalu(dalas [][]int, indekss int, gaidasGrupa *sync.WaitGroup) {
	defer gaidasGrupa.Done()

	fmt.Printf("Gorutina %d - pirms: %v\n", indekss+1, dalas[indekss])
	atraKartosana(dalas[indekss])
	fmt.Printf("Gorutina %d - pec:  %v\n", indekss+1, dalas[indekss])
}

func atraKartosana(massivs []int) {
	if len(massivs) <= 1 {
		return
	}

	atskaitesPunkts := massivs[len(massivs)/2]
	KrePuse, LabPuse := 0, len(massivs)-1

	for KrePuse <= LabPuse {
		for massivs[KrePuse] < atskaitesPunkts {
			KrePuse++
		}
		for massivs[LabPuse] > atskaitesPunkts {
			LabPuse--
		}
		if KrePuse <= LabPuse {
			massivs[KrePuse], massivs[LabPuse] = massivs[LabPuse], massivs[KrePuse]
			KrePuse++
			LabPuse--
		}
	}

	if LabPuse > 0 {
		atraKartosana(massivs[:LabPuse+1])
	}
	if KrePuse < len(massivs) {
		atraKartosana(massivs[KrePuse:])
	}
}

func apvienotSakartotosMasivus(masivi [][]int) []int {
	if len(masivi) == 0 {
		return []int{}
	}

	rezultats := masivi[0]
	for i := 1; i < len(masivi); i++ {
		rezultats = apvienotDivusMasivus(rezultats, masivi[i])
	}
	return rezultats
}

func apvienotDivusMasivus(pirmais, otrais []int) []int {
	rezultats := make([]int, len(pirmais)+len(otrais))
	i, j, k := 0, 0, 0

	for i < len(pirmais) && j < len(otrais) {
		if pirmais[i] <= otrais[j] {
			rezultats[k] = pirmais[i]
			i++
		} else {
			rezultats[k] = otrais[j]
			j++
		}
		k++
	}

	copy(rezultats[k:], pirmais[i:])
	copy(rezultats[k+len(pirmais)-i:], otrais[j:])

	return rezultats
}

func irSakartots(massivs []int) bool {
	for i := 1; i < len(massivs); i++ {
		if massivs[i] < massivs[i-1] {
			return false
		}
	}
	return true
}

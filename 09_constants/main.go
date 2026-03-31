package main
import "fmt"


// constant grouping, just better structure brother
const (
	CLUBS = iota //0
	HEARTS // 1
	DAIMONDS // 2
	SPADES // 3
)

const (
	HP = "Hewlett Packard"
	DL = "DELL"
	LO = "LENOVO"
)

func main(){
	const PI = 3.14 // https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33530115a662800ba9adce2ed64cd65b

	fmt.Printf("LO: %v\n", LO)
	fmt.Printf("SPADES: %v\n", SPADES)
	
	//typed constants
	const maxUpload int = 66
	fmt.Println(PI, maxUpload)
}
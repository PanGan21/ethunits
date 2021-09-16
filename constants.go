package ethunits

type Unit int

const (
	Wei        Unit = 0
	Kwei       Unit = 1
	Babbage    Unit = 2
	Femtoether Unit = 3
	Mwei       Unit = 4
	Lovelace   Unit = 5
	Picoether  Unit = 6
	Gwei       Unit = 7
	Shannon    Unit = 8
	Nanoether  Unit = 9
	Nano       Unit = 10
	Szabo      Unit = 11
	Microether Unit = 12
	Micro      Unit = 13
	Finney     Unit = 14
	Milliether Unit = 15
	Milli      Unit = 16
	Ether      Unit = 17
	Kether     Unit = 18
	Grand      Unit = 19
	Einstein   Unit = 20
	Mether     Unit = 21
	Gether     Unit = 22
	Tether     Unit = 23
)

var unit_srings = [...]string{
	"wei",
	"kwei",
	"babbage",
	"femtoether",
	"mwei",
	"lovelace",
	"picoether",
	"gwei",
	"shannon",
	"nanoether",
	"nano",
	"szabo",
	"microether",
	"micro",
	"finney",
	"milliether",
	"milli",
	"ether",
	"kether",
	"grand",
	"einstein",
	"mether",
	"gether",
	"tether"}

func (u Unit) String() string {
	if u < Wei || u > Tether {
		return "Unknown"
	}

	return unit_srings[u]
}

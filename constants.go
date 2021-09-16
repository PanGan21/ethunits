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

var (
	units_values = map[Unit]string{
		Wei:        "1",
		Kwei:       "1000",
		Babbage:    "1000",
		Femtoether: "1000",
		Mwei:       "1000000",
		Lovelace:   "1000000",
		Picoether:  "1000000",
		Gwei:       "1000000000",
		Shannon:    "1000000000",
		Nanoether:  "1000000000",
		Nano:       "1000000000",
		Szabo:      "1000000000000",
		Microether: "1000000000000",
		Micro:      "1000000000000",
		Finney:     "1000000000000000",
		Milliether: "1000000000000000",
		Milli:      "1000000000000000",
		Ether:      "1000000000000000000",
		Kether:     "1000000000000000000000",
		Grand:      "1000000000000000000000",
		Einstein:   "1000000000000000000000",
		Mether:     "1000000000000000000000000",
		Gether:     "1000000000000000000000000000",
		Tether:     "1000000000000000000000000000000",
	}
)

func (u Unit) String() string {
	if u < Wei || u > Tether {
		return "Unknown"
	}

	return unit_srings[u]
}

func (u Unit) Value() string {
	if u < Wei || u > Tether {
		return "Unknown"
	}

	return units_values[u]
}

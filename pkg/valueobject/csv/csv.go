package csv

type BaseRow struct {
    Data string `csv:"data"`
    Item string `csv:"denominazione_provincia,denominazione_regione"`
    Cases string `csv:"totale_casi"`
}

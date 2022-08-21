package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strconv"
	"strings"
)

func CotizacionDolarEuroBNAScraping() (Cotizacion, error) {

	var cotizacion Cotizacion
	var getDolar bool = false
	var founFirstDolar bool = true
	var getEuro bool = false
	var founFirstEuro bool = true

	c := colly.NewCollector(colly.AllowedDomains(os.Getenv("BNA_SCRAPING_DOMAIN")))
	c.OnHTML("table.table", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, e *colly.HTMLElement) {

			// DOLAR BUSQUEDA - SE HACE DESTA MANERA PORQUE APARECE MAS DE UNA VEZ EN TODO EL HTML, SOLO SE TOMA LA PRIMERA VEZ QUE SE ENCUENTRA
			if strings.Contains(e.Text, "Dolar") {
				getDolar = true
			}
			if getDolar && founFirstDolar {
				switch i {
				case 1:
					cotizacion.DolarCompra = GetValor(e.Text)
				case 2:
					cotizacion.DolarVenta = GetValor(e.Text)
					founFirstDolar = false
				}
			}
			//EURO BUSQUEDA
			if strings.Contains(e.Text, "Euro") {
				getEuro = true
			}
			if getEuro && founFirstEuro {
				switch i {
				case 2:
					founFirstEuro = false
				case 4:
					cotizacion.EuroCompra = GetValor(e.Text)
				case 5:
					cotizacion.EuroVenta = GetValor(e.Text)
					founFirstDolar = false
				}
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(os.Getenv("BNA_SCRAPING_URL_VISIT"))

	return cotizacion, nil
}

func GetValor(valor string) float32 {
	valor = strings.Replace(valor, ",", ".", 1)
	valorFloat64, _ := strconv.ParseFloat(valor, 32)
	return float32(valorFloat64)
}

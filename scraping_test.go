package main

import "testing"

func TestCotizacionDolarEuroBNAScraping(t *testing.T) {
	cotizacion, err := CotizacionDolarEuroBNAScraping()
	if err != nil {
		t.Error(err)
	}
	if cotizacion.DolarCompra == 0 {
		t.Error("DolarCompra no puede ser 0")
	}
	if cotizacion.DolarVenta == 0 {
		t.Error("DolarVenta no puede ser 0")
	}
	if cotizacion.EuroCompra == 0 {
		t.Error("EuroCompra no puede ser 0")
	}
	if cotizacion.EuroVenta == 0 {
		t.Error("EuroVenta no puede ser 0")
	}
}

package helper

import "fmt"

func LuasPersegi(sisi int, showDetail bool) interface{} {
	if sisi == 0 {
		if showDetail {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if showDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

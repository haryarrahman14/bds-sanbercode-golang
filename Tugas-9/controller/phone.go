package controller

import (
	"bds-sanbercode-golang/Tugas-9/model"
	"fmt"
)

func GetDisplayPhoneInfo(d model.DisplayPhoneInfo) {
	info := d.GetInfo()

	fmt.Println(info)
}

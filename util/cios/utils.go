package cios_util

import "github.com/optim-corp/cios-golang-sdk/cios"

func GetDefaultDisplayInfo(displayInfos []cios.DisplayInfo) *cios.DisplayInfo {
	for _, v := range displayInfos {
		if v.IsDefault == true {
			return &v
		}
	}
	return nil
}

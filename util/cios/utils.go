package cios_util

import (
	"github.com/optim-corp/cios-golang-sdk/cios"
)

type DeviceEntitiesComponentRange struct {
	StartDeviceEntitiesComponent cios.DeviceEntitiesComponent
	EndDeviceEntitiesComponent   cios.DeviceEntitiesComponent
	StartEventAt                 string
	EndEventAt                   string
}

func GetDefaultDisplayInfo(displayInfos []cios.DisplayInfo) *cios.DisplayInfo {
	for _, v := range displayInfos {
		if v.IsDefault == true {
			return &v
		}
	}
	return nil
}

//func ToBetweenLifecycle(lifecycles cios.LifeCycleStream) (result [][]DeviceEntitiesComponentRange) {
//	streams := lifecycles.
//		Sort(func(i, j int) bool { return lifecycles.Get(i).EventAt < lifecycles.Get(j).EventAt }).
//		GroupBy(func(cycle cios.LifeCycle, _ int) string { return cnv.MustStr(cycle.AfterId) })
//	for _, cycles := range streams {
//		result = append(result, toBetweenLifecycle(cycles))
//	}
//	return
//}
//func toBetweenLifecycle(lifecycles cios.LifeCycleStream) (result []DeviceEntitiesComponentRange) {
//	length := len(lifecycles)
//	if length < 1 {
//		return
//	}
//	result = append(result, toDeviceEntitiesComponentRange(nil, &lifecycles[0]))
//	for i := 0; i < length-1; i++ {
//		result = append(result, toDeviceEntitiesComponentRange(&lifecycles[i], &lifecycles[i+1]))
//	}
//	result = append(result, toDeviceEntitiesComponentRange(&lifecycles[length-1], nil))
//	return
//}
//func toDeviceEntitiesComponentRange(start *cios.LifeCycle, end *cios.LifeCycle) DeviceEntitiesComponentRange {
//	if start != nil && end == nil && start.AfterComponent.IsSet() {
//		return DeviceEntitiesComponentRange{
//			StartDeviceEntitiesComponent: *start.AfterComponent.Get(),
//			StartEventAt:                 start.EventAt,
//		}
//	}
//	if start != nil && end != nil && start.AfterComponent.IsSet() && end.BeforeComponent.IsSet() {
//		return DeviceEntitiesComponentRange{
//			StartDeviceEntitiesComponent: *start.AfterComponent.Get(),
//			EndDeviceEntitiesComponent:   *end.BeforeComponent.Get(),
//			StartEventAt:                 start.EventAt,
//			EndEventAt:                   end.EventAt,
//		}
//	}
//	if start == nil && end != nil && end.BeforeComponent.IsSet() {
//		return DeviceEntitiesComponentRange{
//			EndDeviceEntitiesComponent: *end.BeforeComponent.Get(),
//			EndEventAt:                 end.EventAt,
//		}
//	}
//	return DeviceEntitiesComponentRange{}
//}

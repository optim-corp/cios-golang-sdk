package ciosutil

import (
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
)

func BucketsMapByResourceOwnerID(buckets []cios.Bucket) map[string][]cios.Bucket {
	result := make(map[string][]cios.Bucket)
	for _, bucket := range buckets {
		result[bucket.ResourceOwnerId] = append(result[bucket.ResourceOwnerId], bucket)
	}
	return result
}
func ChannelsMapByResourceOwnerID(channels []cios.Channel) map[string][]cios.Channel {
	result := make(map[string][]cios.Channel)
	for _, channel := range channels {
		result[channel.ResourceOwnerId] = append(result[channel.ResourceOwnerId], channel)
	}
	return result
}
func ResourceOwnerIDMapByChannelID(channels []cios.Channel) map[string]string {
	result := make(map[string]string)
	for _, c := range channels {
		result[c.Id] = c.ResourceOwnerId
	}
	return result
}

//func CreateParentPartsMap(entities []cios.DeviceModelsEntity, language string) map[string]string {
//	var (
//		dig            func(n cios.DeviceEntitiesComponent)
//		parentPartsMap = map[string]string{}
//	)
//	dig = func(entity cios.DeviceEntitiesComponent) {
//		if entity.Components != nil {
//			for _, ent := range *entity.Components {
//				if ent.DisplayInfo != nil {
//					for _, display := range *ent.DisplayInfo {
//						if display.Language == language {
//							parentPartsMap[ent.Id] = display.Name
//							break
//						}
//					}
//				}
//				dig(ent)
//			}
//		}
//	}
//	cios.
//		GenerateDeviceModelsEntityStream(entities).
//		Each(func(device cios.DeviceModelsEntity) {
//			if device.Components.IsSet() {
//				dig(*device.Components.Get())
//			}
//		})
//	return parentPartsMap
//}

func GroupMapByResourceOwnerID(groups []cios.Group, resourceOwners []cios.ResourceOwner) map[string]cios.Group {
	result := make(map[string]cios.Group)
	for _, ro := range resourceOwners {
		for _, gp := range groups {
			if cnv.MustStr(ro.GroupId) == gp.Id {
				result[ro.Id] = gp
			}
		}
	}
	return result
}
func ResourceOwnerMapByGroupID(resourceOwners []cios.ResourceOwner) map[string]cios.ResourceOwner {
	result := map[string]cios.ResourceOwner{}
	for _, ro := range resourceOwners {
		if ro.GroupId != nil {
			result[*ro.GroupId] = ro
		}
	}
	return result
}

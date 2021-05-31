package cios_util_test

import (
	"reflect"
	"testing"

	cios_util "github.com/optim-corp/cios-golang-sdk/util/cios"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

//func TestCreateParentPartsMap(t *testing.T) {
//	type args struct {
//		entities []cios.DeviceModelsEntity
//		language string
//	}
//	tests := []struct {
//		name string
//		args args
//		want map[string]string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CreateParentPartsMap(tt.args.entities, tt.args.language); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateParentPartsMap() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestToBetweenLifecycle(t *testing.T) {
//	type args struct {
//		lifecycles cios.LifeCycleStream
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantResult [][]cios_util.DeviceEntitiesComponentRange
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotResult := cios_util.ToBetweenLifecycle(tt.args.lifecycles); !reflect.DeepEqual(gotResult, tt.wantResult) {
//				t.Errorf("ToBetweenLifecycle() = %v, want %v", gotResult, tt.wantResult)
//			}
//		})
//	}
//}
//
//func Test_toBetweenLifecycle(t *testing.T) {
//	type args struct {
//		lifecycles cios.LifeCycleStream
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantResult []cios_util.DeviceEntitiesComponentRange
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotResult := cios_util.ToBetweenLifecycle(tt.args.lifecycles); !reflect.DeepEqual(gotResult, tt.wantResult) {
//				t.Errorf("toBetweenLifecycle() = %v, want %v", gotResult, tt.wantResult)
//			}
//		})
//	}
//}
//
//func Test_toDeviceEntitiesComponentRange(t *testing.T) {
//	type args struct {
//		start *cios.LifeCycle
//		end   *cios.LifeCycle
//	}
//	tests := []struct {
//		name string
//		args args
//		want cios_util.DeviceEntitiesComponentRange
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Test_toDeviceEntitiesComponentRange(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("toDeviceEntitiesComponentRange() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
func TestGetDefaultDisplayInfo(t *testing.T) {
	type args struct {
		displayInfos []cios.DisplayInfo
	}
	tests := []struct {
		name string
		args args
		want *cios.DisplayInfo
	}{
		{
			name: "Test Array",
			args: args{
				displayInfos: cios.DisplayInfoStreamOf(
					cios.DisplayInfo{Name: "test", IsDefault: true},
					cios.DisplayInfo{IsDefault: false},
					cios.DisplayInfo{IsDefault: false},
				),
			},
			want: &cios.DisplayInfo{Name: "test", IsDefault: true},
		},
		{
			name: "Test One",
			args: args{
				displayInfos: cios.DisplayInfoStreamOf(
					cios.DisplayInfo{Name: "test", IsDefault: true},
				),
			},
			want: &cios.DisplayInfo{Name: "test", IsDefault: true},
		},
		{
			name: "Test Array Empty",
			args: args{
				displayInfos: cios.DisplayInfoStreamOf(
					cios.DisplayInfo{Name: "test", IsDefault: false},
					cios.DisplayInfo{IsDefault: false},
					cios.DisplayInfo{IsDefault: false},
				),
			},
			want: nil,
		},
		{
			name: "Test One Empty",
			args: args{
				displayInfos: cios.DisplayInfoStreamOf(
					cios.DisplayInfo{Name: "test", IsDefault: false},
				),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cios_util.GetDefaultDisplayInfo(tt.args.displayInfos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultDisplayInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

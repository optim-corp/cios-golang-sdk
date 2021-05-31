package cios_util

import (
	"reflect"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func TestBucketsMapByResourceOwnerID(t *testing.T) {
	type args struct{ buckets []cios.Bucket }
	tests := []struct {
		name string
		args args
		want map[string][]cios.Bucket
	}{
		{
			name: "Set Bucket Map",
			args: args{buckets: []cios.Bucket{
				{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
			}},
			want: map[string][]cios.Bucket{
				"65356cff-b63c-4fe2-a924-45b253caee63": {
					{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				},
			},
		},
		{
			name: "Empty Bucket Map",
			args: args{buckets: nil},
			want: map[string][]cios.Bucket{},
		},
		{
			name: "Set Multiple Bucket Map",
			args: args{buckets: []cios.Bucket{
				{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test4", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
				{Id: "test5", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
				{Id: "test6", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
			}},
			want: map[string][]cios.Bucket{
				"65356cff-b63c-4fe2-a924-45b253caee63": {
					{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				},
				"e7904bd5-9e99-4384-8548-eab23b2efd15": {
					{Id: "test4", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
					{Id: "test5", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
					{Id: "test6", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketsMapByResourceOwnerID(tt.args.buckets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketsMapByResourceOwnerID() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}

func TestChannelsMapByResourceOwnerID(t *testing.T) {
	type args struct {
		channels []cios.Channel
	}
	tests := []struct {
		name string
		args args
		want map[string][]cios.Channel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChannelsMapByResourceOwnerID(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChannelsMapByResourceOwnerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateParentPartsMap(t *testing.T) {
	type args struct {
		entities []cios.DeviceModelsEntity
		language string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateParentPartsMap(tt.args.entities, tt.args.language); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParentPartsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStoreStreamToStruct(t *testing.T) {
	type args struct {
		str []string
		st  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DataStoreStreamToStruct(tt.args.str, tt.args.st); (err != nil) != tt.wantErr {
				t.Errorf("DataStoreStreamToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDefaultDisplayInfo(t *testing.T) {
	type args struct {
		displayInfos []cios.DisplayInfo
	}
	tests := []struct {
		name string
		args args
		want *cios.DisplayInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultDisplayInfo(tt.args.displayInfos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultDisplayInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupMapByResourceOwnerID(t *testing.T) {
	type args struct {
		groups         []cios.Group
		resourceOwners []cios.ResourceOwner
	}
	tests := []struct {
		name string
		args args
		want map[string]cios.Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupMapByResourceOwnerID(tt.args.groups, tt.args.resourceOwners); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupMapByResourceOwnerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapPayloads(t *testing.T) {
	type args struct {
		objects []cios.PackerFormatJson
		stc     interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []cios.PackerFormatJsonHeader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapPayloads(tt.args.objects, tt.args.stc)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapPayloads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapPayloads() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceOwnerIDMapByChannelID(t *testing.T) {
	type args struct {
		channels []cios.Channel
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceOwnerIDMapByChannelID(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOwnerIDMapByChannelID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceOwnerMapByGroupID(t *testing.T) {
	type args struct {
		resourceOwners []cios.ResourceOwner
		groups         []cios.Group
	}
	tests := []struct {
		name string
		args args
		want map[string]cios.ResourceOwner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResourceOwnerMapByGroupID(tt.args.resourceOwners, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOwnerMapByGroupID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBetweenLifecycle(t *testing.T) {
	type args struct {
		lifecycles cios.LifeCycleStream
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]DeviceEntitiesComponentRange
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToBetweenLifecycle(tt.args.lifecycles); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToBetweenLifecycle() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_toBetweenLifecycle(t *testing.T) {
	type args struct {
		lifecycles cios.LifeCycleStream
	}
	tests := []struct {
		name       string
		args       args
		wantResult []DeviceEntitiesComponentRange
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := toBetweenLifecycle(tt.args.lifecycles); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("toBetweenLifecycle() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_toDeviceEntitiesComponentRange(t *testing.T) {
	type args struct {
		start *cios.LifeCycle
		end   *cios.LifeCycle
	}
	tests := []struct {
		name string
		args args
		want DeviceEntitiesComponentRange
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDeviceEntitiesComponentRange(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDeviceEntitiesComponentRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

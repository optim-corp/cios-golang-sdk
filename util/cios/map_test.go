package ciosutil_test

import (
	"reflect"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosutil "github.com/optim-corp/cios-golang-sdk/util/cios"
)

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
		{
			name: "Test Group Resource Map",
			args: args{
				groups: []cios.Group{
					{Id: "60de60f8-5fae-4b4a-a774-0cf11427646c", Name: "group1"},
					{Id: "3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22", Name: "group2"},
					{Id: "858ba17c-8f66-46c0-bd8e-e596133e0324", Name: "group3"},
				},
				resourceOwners: []cios.ResourceOwner{
					{Id: "1ebac437-3f09-4976-8b53-e470aad78260", GroupId: cnv.StrPtr("60de60f8-5fae-4b4a-a774-0cf11427646c")},
					{Id: "6158cdf7-0f81-491e-81bb-f1b7c91d48d0", GroupId: cnv.StrPtr("3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22")},
					{Id: "f8d4bec8-003f-4878-84ac-848a2b107cdd", GroupId: cnv.StrPtr("858ba17c-8f66-46c0-bd8e-e596133e0324")},
				},
			},
			want: map[string]cios.Group{
				"1ebac437-3f09-4976-8b53-e470aad78260": {Id: "60de60f8-5fae-4b4a-a774-0cf11427646c", Name: "group1"},
				"6158cdf7-0f81-491e-81bb-f1b7c91d48d0": {Id: "3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22", Name: "group2"},
				"f8d4bec8-003f-4878-84ac-848a2b107cdd": {Id: "858ba17c-8f66-46c0-bd8e-e596133e0324", Name: "group3"},
			},
		},
		{
			name: "Test Group Resource Map Empty",
			args: args{
				groups:         []cios.Group{},
				resourceOwners: []cios.ResourceOwner{},
			},
			want: map[string]cios.Group{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciosutil.GroupMapByResourceOwnerID(tt.args.groups, tt.args.resourceOwners); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupMapByResourceOwnerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := ciosutil.BucketsMapByResourceOwnerID(tt.args.buckets); !reflect.DeepEqual(got, tt.want) {
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
		{
			name: "Set Bucket Map",
			args: args{channels: []cios.Channel{
				{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
			}},
			want: map[string][]cios.Channel{
				"65356cff-b63c-4fe2-a924-45b253caee63": {
					{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
					{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				},
			},
		},
		{
			name: "Empty Bucket Map",
			args: args{channels: nil},
			want: map[string][]cios.Channel{},
		},
		{
			name: "Set Multiple Bucket Map",
			args: args{channels: []cios.Channel{
				{Id: "test1", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test2", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test3", ResourceOwnerId: "65356cff-b63c-4fe2-a924-45b253caee63"},
				{Id: "test4", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
				{Id: "test5", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
				{Id: "test6", ResourceOwnerId: "e7904bd5-9e99-4384-8548-eab23b2efd15"},
			}},
			want: map[string][]cios.Channel{
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
			if got := ciosutil.ChannelsMapByResourceOwnerID(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChannelsMapByResourceOwnerID() = %v, want %v", got, tt.want)
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
		{
			name: "Test ResourceOwnerID Map By Channel ID Empty",
			args: args{
				channels: []cios.Channel{},
			},
			want: map[string]string{},
		},
		{
			name: "Test ResourceOwnerID Map By Channel ID",
			args: args{
				channels: []cios.Channel{
					{Id: "test1", ResourceOwnerId: "a867e29c-d7df-4b8c-8ba5-3b67ad5c30f3"},
					{Id: "test2", ResourceOwnerId: "4242e44d-abac-4612-8d5f-23006af52d52"},
					{Id: "test3", ResourceOwnerId: "f03ea94a-f4f2-48a6-9223-23f788573c0b"},
				},
			},
			want: map[string]string{
				"test1": "a867e29c-d7df-4b8c-8ba5-3b67ad5c30f3",
				"test2": "4242e44d-abac-4612-8d5f-23006af52d52",
				"test3": "f03ea94a-f4f2-48a6-9223-23f788573c0b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciosutil.ResourceOwnerIDMapByChannelID(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOwnerIDMapByChannelID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceOwnerMapByGroupID(t *testing.T) {
	type args struct {
		resourceOwners []cios.ResourceOwner
	}
	tests := []struct {
		name string
		args args
		want map[string]cios.ResourceOwner
	}{
		{
			name: "Test Resource Map By Group ID",
			args: args{
				resourceOwners: []cios.ResourceOwner{
					{Id: "1ebac437-3f09-4976-8b53-e470aad78260", GroupId: cnv.StrPtr("60de60f8-5fae-4b4a-a774-0cf11427646c")},
					{Id: "6158cdf7-0f81-491e-81bb-f1b7c91d48d0", GroupId: cnv.StrPtr("3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22")},
					{Id: "f8d4bec8-003f-4878-84ac-848a2b107cdd", GroupId: cnv.StrPtr("858ba17c-8f66-46c0-bd8e-e596133e0324")},
				},
			},
			want: map[string]cios.ResourceOwner{
				"60de60f8-5fae-4b4a-a774-0cf11427646c": {Id: "1ebac437-3f09-4976-8b53-e470aad78260", GroupId: cnv.StrPtr("60de60f8-5fae-4b4a-a774-0cf11427646c")},
				"3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22": {Id: "6158cdf7-0f81-491e-81bb-f1b7c91d48d0", GroupId: cnv.StrPtr("3bbf97b0-4ab8-4a02-b2f1-198abc0c6f22")},
				"858ba17c-8f66-46c0-bd8e-e596133e0324": {Id: "f8d4bec8-003f-4878-84ac-848a2b107cdd", GroupId: cnv.StrPtr("858ba17c-8f66-46c0-bd8e-e596133e0324")},
			},
		},
		{
			name: "Test Resource Map By Group ID Empty",
			args: args{
				resourceOwners: []cios.ResourceOwner{},
			},
			want: map[string]cios.ResourceOwner{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciosutil.ResourceOwnerMapByGroupID(tt.args.resourceOwners); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOwnerMapByGroupID() = %v, want %v", got, tt.want)
			}
		})
	}
}

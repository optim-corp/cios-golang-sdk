package ciossdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestCetPoints(t *testing.T) {

	tests := map[string]struct {
		out  cios.MultiplePoint
		want cios.MultiplePoint
	}{
		"success": {
			out: cios.MultiplePoint{
				Total: 1,
				Points: []cios.Point{
					{
						Id:   "test-id",
						Name: "test-name",
						Location: cios.Location{
							Latitude:  12.334,
							Longitude: 54.321,
						},
						ResourceOwnerId: "test-ro-id",
					},
				},
			},
			want: cios.MultiplePoint{
				Total: 1,
				Points: []cios.Point{
					{
						Id:   "test-id",
						Name: "test-name",
						Location: cios.Location{
							Latitude:  12.334,
							Longitude: 54.321,
						},
						ResourceOwnerId: "test-ro-id",
					},
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt := tt

			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodGet
						r.URL.Path = "/v2/geo/points"
						response := tt.out
						json.NewEncoder(w).Encode(response)
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			result, _, err := client.Geography().GetPoints(context.Background(), cios.ApiGetPointsRequest{})

			assert.Nil(t, err, "error must be nil")
			assert.Equal(t, tt.want, result, "result must be equal")
		})
	}
}

func TestCreatePoint(t *testing.T) {
	var (
		testDescription = "test-description"
		testRegion      = "JP"
	)

	type arg struct {
		name, key, value, description, language, resourceOwner string
		latitude, longitude, altitude                          float32
	}

	tests := map[string]struct {
		arg arg
		out cios.SinglePoint
	}{
		"success": {
			arg: arg{
				name:          "test-name",
				key:           "test-key",
				value:         "test-value",
				language:      "ja",
				latitude:      123.000,
				longitude:     54.321,
				altitude:      123.40,
				resourceOwner: "test-ro",
			},
			out: cios.SinglePoint{
				Point: cios.Point{
					Id:   "test-id",
					Name: "test-name",
					Location: cios.Location{
						Latitude:  123.000,
						Longitude: 54.321,
					},
					Altitude: 123.40,
					Labels: &[]cios.Label{
						{
							Key:   "test-key",
							Value: "test-value",
						},
					},
					Description: &testDescription,
					DisplayInfo: &[]cios.DisplayInfo{
						{
							Name:        "test-location",
							Description: &testDescription,
							Language:    "ja",
							Script:      nil,
							Region:      &testRegion,
							IsDefault:   true,
						},
					},
					ResourceOwnerId: "test-ro",
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt := tt
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPost
						r.URL.Path = "/v2/geo/points"
						response := tt.out
						json.NewEncoder(w).Encode(response)
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			req := cios.PointRequest{
				Name: &tt.arg.name,
				Location: &cios.Location{
					Latitude:  tt.arg.latitude,
					Longitude: tt.arg.longitude,
				},
				Altitude: &tt.arg.altitude,
				Labels: &[]cios.Label{
					{
						Key:   tt.arg.key,
						Value: tt.arg.value,
					},
				},
				Description: &testDescription,
				DisplayInfo: &[]cios.DisplayInfo{
					{
						Name:        tt.arg.name,
						Description: &testDescription,
						Language:    tt.arg.language,
						Region:      &testRegion,
						IsDefault:   true,
					},
				},
				ResourceOwnerId: tt.arg.resourceOwner,
			}

			result, _, err := client.Geography().CreatePoint(context.Background(), req)
			assert.Nil(t, err, "error must be nil")
			assert.Equal(t, tt.out, result, "result must be equal")

		})
	}
}

func TestGetPoint(t *testing.T) {
	var (
		testDescription = "test-description"
		testRegion      = "JP"
	)

	test := map[string]struct {
		pointID, lang string
		isDev         bool
		out           cios.SinglePoint
	}{
		"success": {
			pointID: "test-id",
			lang:    "JP",
			isDev:   true,
			out: cios.SinglePoint{Point: cios.Point{
				Id:   "test-id",
				Name: "test-name",
				Location: cios.Location{
					Latitude:  123.000,
					Longitude: 54.321,
				},
				Altitude: 123.40,
				Labels: &[]cios.Label{
					{
						Key:   "test-key",
						Value: "test-value",
					},
				},
				Description: &testDescription,
				DisplayInfo: &[]cios.DisplayInfo{
					{
						Name:        "test-location",
						Description: &testDescription,
						Language:    "ja",
						Script:      nil,
						Region:      &testRegion,
						IsDefault:   true,
					},
				},
				ResourceOwnerId: "test-ro",
			}},
		},
	}
	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			tt := tt
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodGet
						r.URL.Path = "/v2/geo/points/" + tt.pointID
						response := tt.out
						json.NewEncoder(w).Encode(response)
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			result, _, err := client.Geography().GetPoint(context.Background(), tt.pointID, &tt.lang, &tt.isDev)
			assert.Nil(t, err, "error must be nil")
			assert.Equal(t, tt.out, result, "result must be equal")

		})
	}
}

func TestUpdatePoint(t *testing.T) {
	var (
		testRegion      = "JP"
		testDescription = "test-description"
		testScript      = "test-script"
	)
	tests := map[string]struct {
		pointID string
		req     []cios.DisplayInfo
		out     cios.SinglePoint
	}{
		"success": {
			pointID: "test-id",
			req: []cios.DisplayInfo{
				{
					Name:        "test-location",
					Description: &testDescription,
					Language:    "ja",
					Script:      &testScript,
					Region:      &testRegion,
					IsDefault:   true,
				},
			},
			out: cios.SinglePoint{Point: cios.Point{
				Id:   "test-id",
				Name: "test-name",
				DisplayInfo: &[]cios.DisplayInfo{
					{
						Name:        "test-location",
						Language:    "ja",
						Description: &testDescription,
						Script:      &testScript,
						Region:      &testRegion,
						IsDefault:   true,
					},
				},
			}},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPatch
						r.URL.Path = "/v2/geo/points/" + tt.pointID
						response := tt.out
						json.NewEncoder(w).Encode(response)
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			result, _, err := client.Geography().UpdatePoint(context.Background(), tt.pointID, tt.req)
			assert.Nil(t, err, "error must be nil")
			assert.Equal(t, tt.out, result, "result must be equal")
		})
	}
}

func TestDeletePoint(t *testing.T) {
	tests := map[string]struct {
		pointID string
	}{
		"success": {
			pointID: "test-id",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodDelete
						r.URL.Path = "/v2/geo/points/" + tt.pointID
					},
				),
			)
			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			_, err := client.Geography().DeletePoint(context.Background(), tt.pointID)
			assert.Nil(t, err, "error must be nil")
		})
	}
}

func TestGetCircles(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID string
		altitude, radiusM                      float32
		total                                  int64
		location                               cios.Location
		labels                                 cios.Label
		displayInfo                            cios.DisplayInfo
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			altitude:        float32(10),
			radiusM:         float32(10),
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodGet
						r.URL.Path = "/v2/geo/circles"
						var circles []cios.Circle
						for i := int64(0); i < tt.total; i++ {
							l := []cios.Label{}
							l = append(l, tt.labels)
							d := []cios.DisplayInfo{}
							d = append(d, tt.displayInfo)
							v := cios.Circle{
								Id:              &tt.ID,
								Name:            &tt.name,
								Altitude:        &tt.altitude,
								RadiusM:         &tt.radiusM,
								Location:        &tt.location,
								Labels:          &l,
								Description:     &tt.description,
								DisplayInfo:     &d,
								ResourceOwnerId: &tt.resourceOwnerID,
							}
							circles = append(circles, v)
						}

						json.NewEncoder(w).Encode(cios.MultipleCircle{
							Total:   &tt.total,
							Circles: &circles,
						})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})
			res, _, err := client.Geography().GetCircles(context.Background(), cios.ApiGetCirclesRequest{})
			assert.Nil(t, err, "error must be nil")
			assert.Equal(t, tt.ID, res.GetCircles()[0].GetId(), "ID must be equal")
		})

	}
}

func TestCreateCircle(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID string
		altitude, radiusM                      float32
		total                                  int64
		location                               cios.Location
		labels                                 cios.Label
		displayInfo                            cios.DisplayInfo
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			altitude:        float32(10),
			radiusM:         float32(10),
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPost
						r.URL.Path = "/v2/geo/circles"
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleCircle{Circle: &cios.Circle{
							Id:              &tt.ID,
							Name:            &tt.name,
							Altitude:        &tt.altitude,
							RadiusM:         &tt.radiusM,
							Location:        &tt.location,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			l := []cios.Label{}
			l = append(l, tt.labels)
			d := []cios.DisplayInfo{}
			d = append(d, tt.displayInfo)
			res, _, err := client.Geography().CreateCircle(context.Background(), cios.Circle{
				Name:            &tt.name,
				Altitude:        &tt.altitude,
				RadiusM:         &tt.radiusM,
				Location:        &tt.location,
				Labels:          &l,
				Description:     &tt.description,
				DisplayInfo:     &d,
				ResourceOwnerId: &tt.resourceOwnerID,
			})
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetCircle().Id, "circle id must not be nil")
		})
	}
}

func TestGetCircle(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID string
		altitude, radiusM                      float32
		total                                  int64
		location                               cios.Location
		labels                                 cios.Label
		displayInfo                            cios.DisplayInfo
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			altitude:        float32(10),
			radiusM:         float32(10),
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodGet
						r.URL.Path = "/v2/geo/circles/" + tt.ID
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleCircle{Circle: &cios.Circle{
							Id:              &tt.ID,
							Name:            &tt.name,
							Altitude:        &tt.altitude,
							RadiusM:         &tt.radiusM,
							Location:        &tt.location,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			l := []cios.Label{}
			l = append(l, tt.labels)
			d := []cios.DisplayInfo{}
			d = append(d, tt.displayInfo)
			res, _, err := client.Geography().GetCircle(context.Background(), tt.ID, nil, nil)
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetCircle().Id, "circle id must not be nil")
		})
	}
}

func TestUpdateCircle(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID string
		altitude, radiusM                      float32
		total                                  int64
		location                               cios.Location
		labels                                 cios.Label
		displayInfo                            cios.DisplayInfo
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			altitude:        float32(10),
			radiusM:         float32(10),
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPatch
						r.URL.Path = "/v2/geo/circles/" + tt.ID
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleCircle{Circle: &cios.Circle{
							Id:              &tt.ID,
							Name:            &tt.name,
							Altitude:        &tt.altitude,
							RadiusM:         &tt.radiusM,
							Location:        &tt.location,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			req := []cios.DisplayInfo{}
			req = append(req, tt.displayInfo)
			res, _, err := client.Geography().UpdateCircle(context.Background(), tt.ID, req)
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetCircle().Id, "circle id must not be nil")
		})
	}
}

func TestDeleteCircle(t *testing.T) {
	tests := map[string]struct {
		ID string
	}{
		"success": {
			ID: "test-id",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodDelete
						r.URL.Path = "/v2/geo/circles/" + tt.ID
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			_, err := client.Geography().DeleteCircle(context.Background(), tt.ID)
			assert.Nil(t, err, "error must be nil")
		})
	}
}

func TestGetRoutes(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID, contentsType, contentsID string
		altitude, radiusM                                                float32
		total                                                            int64
		location                                                         cios.Location
		labels                                                           cios.Label
		displayInfo                                                      cios.DisplayInfo
		contentsOrderNo                                                  int32
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			contentsOrderNo: 1,
			contentsType:    "point",
			contentsID:      "test-contents-id",
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				ts := httptest.NewServer(
					http.HandlerFunc(
						func(w http.ResponseWriter, r *http.Request) {
							w.Header().Set("Content-Type", "application/json")
							r.Method = http.MethodGet
							r.URL.Path = "/v2/geo/routes"
							var routes []cios.Route
							for i := int64(0); i < tt.total; i++ {
								l := []cios.Label{}
								l = append(l, tt.labels)
								d := []cios.DisplayInfo{}
								d = append(d, tt.displayInfo)
								c := []cios.RouteContents{}
								c = append(c, cios.RouteContents{
									OrderNo: &tt.contentsOrderNo,
									Type:    &tt.contentsType,
									Id:      &tt.contentsID,
								})
								v := cios.Route{
									Id:              &tt.ID,
									Name:            &tt.name,
									Labels:          &l,
									Contents:        &c,
									Description:     &tt.description,
									DisplayInfo:     &d,
									ResourceOwnerId: &tt.resourceOwnerID,
								}
								routes = append(routes, v)
							}

							json.NewEncoder(w).Encode(cios.MultipleRoute{
								Total:  &tt.total,
								Routes: &routes,
							})
						},
					),
				)

				client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})
				res, _, err := client.Geography().GetRoutes(context.Background(), cios.ApiGetRoutesRequest{})
				assert.Nil(t, err, "error must be nil")
				assert.Equal(t, tt.ID, res.GetRoutes()[0].GetId(), "ID must be equal")
			})
		})

	}
}

func TestCreateRoute(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID, contentsType, contentsID string
		altitude, radiusM                                                float32
		total                                                            int64
		location                                                         cios.Location
		labels                                                           cios.Label
		displayInfo                                                      cios.DisplayInfo
		contentsOrderNo                                                  int32
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			contentsOrderNo: 1,
			contentsType:    "point",
			contentsID:      "test-contents-id",
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPost
						r.URL.Path = "/v2/geo/routes"
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleRoute{Route: &cios.Route{
							Id:              &tt.ID,
							Name:            &tt.name,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			l := []cios.Label{}
			l = append(l, tt.labels)
			d := []cios.DisplayInfo{}
			d = append(d, tt.displayInfo)
			res, _, err := client.Geography().CreateRoute(context.Background(), cios.Route{
				Name:            &tt.name,
				Labels:          &l,
				Description:     &tt.description,
				DisplayInfo:     &d,
				ResourceOwnerId: &tt.resourceOwnerID,
			})
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetRoute().Id, "route ID must not be nil")
		})

	}
}

func TestGetRoute(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID, contentsType, contentsID string
		altitude, radiusM                                                float32
		total                                                            int64
		location                                                         cios.Location
		labels                                                           cios.Label
		displayInfo                                                      cios.DisplayInfo
		contentsOrderNo                                                  int32
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			contentsOrderNo: 1,
			contentsType:    "point",
			contentsID:      "test-contents-id",
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodGet
						r.URL.Path = "/v2/geo/routes" + tt.ID
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleRoute{Route: &cios.Route{
							Id:              &tt.ID,
							Name:            &tt.name,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})
			res, _, err := client.Geography().GetRoute(context.Background(), tt.ID, nil, nil)
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetRoute().Id, "route id must not be nil")
		})
	}
}

func TestUpdateRoute(t *testing.T) {
	tests := map[string]struct {
		ID, name, description, resourceOwnerID, contentsType, contentsID string
		altitude, radiusM                                                float32
		total                                                            int64
		location                                                         cios.Location
		labels                                                           cios.Label
		displayInfo                                                      cios.DisplayInfo
		contentsOrderNo                                                  int32
	}{
		"success": {
			ID:              "test-id",
			name:            "test-name",
			description:     "test",
			resourceOwnerID: "test-ro",
			total:           1,
			contentsOrderNo: 1,
			contentsType:    "point",
			contentsID:      "test-contents-id",
			location: cios.Location{
				Latitude:  100,
				Longitude: 100,
			},
			labels: cios.Label{
				Key:   "key",
				Value: "value",
			},
			displayInfo: cios.DisplayInfo{
				Name:      "test-display",
				Language:  "ja",
				IsDefault: false,
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodPatch
						r.URL.Path = "/v2/geo/routes" + tt.ID
						l := []cios.Label{}
						l = append(l, tt.labels)
						d := []cios.DisplayInfo{}
						d = append(d, tt.displayInfo)
						json.NewEncoder(w).Encode(cios.SingleRoute{Route: &cios.Route{
							Id:              &tt.ID,
							Name:            &tt.name,
							Labels:          &l,
							Description:     &tt.description,
							DisplayInfo:     &d,
							ResourceOwnerId: &tt.resourceOwnerID,
						}})
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			l := []cios.Label{}
			l = append(l, tt.labels)
			d := []cios.DisplayInfo{}
			d = append(d, tt.displayInfo)
			res, _, err := client.Geography().UpdateRoute(context.Background(), tt.ID, d)
			assert.Nil(t, err, "error must be nil")
			assert.NotNil(t, res.GetRoute().Id, "route id must not be nil")
		})
	}
}

func TestDeleteRoute(t *testing.T) {
	tests := map[string]struct {
		ID string
	}{
		"success": {
			ID: "test-id",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						r.Method = http.MethodDelete
						r.URL.Path = "/v2/geo/routes" + tt.ID
					},
				),
			)

			client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{LocationUrl: ts.URL}})

			_, err := client.Geography().DeleteRoute(context.Background(), tt.ID)
			assert.Nil(t, err, "error must be nil")
		})

	}
}

// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store_test

import (
	"testing"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

var (
	data = map[string][]byte{
		"vendor/index.yaml": []byte(`
vendors:
- id: foo-vendor
  name: Foo Vendor
  vendorID: 42
- id: foo-vendor-full
  name: Full Vendor
  vendorID: 44
  email: mail@example.com
  website: example.org
  pen: 42
  ouis: ["010203", "030405"]
  logo: logo.svg`),
		"vendor/foo-vendor/index.yaml": []byte(`
endDevices:
- dev1
- dev2`),
		"vendor/foo-vendor/dev1.yaml": []byte(`
name: Device 1
description: My Description
hardwareVersions:
- version: 1.0
  numeric: 1
  partNumber: P4RTN0
firmwareVersions:
- version: 1.0
  hardwareVersions:
  - 1.0
  profiles:
    EU863-870: {codec: foo-codec, profile-id: profile1, lorawanCertified: true}
    US902-928: {codec: foo-codec, profile-id: profile2, lorawanCertified: true}`),
		"vendor/foo-vendor/dev2.yaml": []byte(`
name: Device 2
description: My Description 2
hardwareVersions:
- version: 2.0
  numeric: 2
  partNumber: P4RTN02
firmwareVersions:
- version: 1.1
  hardwareVersions:
  - 2.0
  profiles:
    EU433: {codec: foo-codec, profile-id: profile2, lorawanCertified: true}
sensors:
- temperature`),
		"vendor/foo-vendor/profile1.yaml": []byte(`
supportsClassB: false
supportsClassC: false
macVersion: 1.0.2
regionalParametersVersion: RP001-1.0.2-RevB
supportsJoin: true
maxEIRP: 16
supports32bitFCnt: true
`),
		"vendor/foo-vendor/foo-codec.yaml": []byte(`
uplinkDecoder: {fileName: a.js}
downlinkDecoder: {fileName: b.js}
downlinkEncoder: {fileName: c.js}`),
		"vendor/foo-vendor/a.js": []byte("uplink decoder"),
		"vendor/foo-vendor/b.js": []byte("downlink decoder"),
		"vendor/foo-vendor/c.js": []byte("downlink encoder"),
	}
)

func TestRemoteStore(t *testing.T) {
	a := assertions.New(t)

	s, err := store.NewRemoteStore(fetch.NewMemFetcher(data))
	a.So(err, should.BeNil)

	t.Run("TestListBrands", func(t *testing.T) {
		t.Run("DefaultPaths", func(t *testing.T) {
			list, err := s.ListBrands(store.ListBrandsRequest{})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceBrand{
				{
					BrandID: "foo-vendor",
					Name:    "Foo Vendor",
				},
				{
					BrandID: "foo-vendor-full",
					Name:    "Full Vendor",
				},
			})
		})

		t.Run("Limit", func(t *testing.T) {
			list, err := s.ListBrands(store.ListBrandsRequest{
				Limit: &pbtypes.UInt32Value{
					Value: 1,
				},
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceBrand{
				{
					BrandID: "foo-vendor",
					Name:    "Foo Vendor",
				},
			})
		})

		t.Run("Offset", func(t *testing.T) {
			list, err := s.ListBrands(store.ListBrandsRequest{
				Offset: &pbtypes.UInt32Value{
					Value: 1,
				},
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceBrand{
				{
					BrandID: "foo-vendor-full",
					Name:    "Full Vendor",
				},
			})
		})

		t.Run("Paths", func(t *testing.T) {
			list, err := s.ListBrands(store.ListBrandsRequest{
				Paths: ttnpb.EndDeviceBrandFieldPathsNested,
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceBrand{
				{
					BrandID:  "foo-vendor",
					Name:     "Foo Vendor",
					VendorID: 42,
				},
				{
					BrandID:                       "foo-vendor-full",
					Name:                          "Full Vendor",
					VendorID:                      44,
					Email:                         "mail@example.com",
					Website:                       "example.org",
					PrivateEnterpriseNumber:       42,
					OrganizationUniqueIdentifiers: []string{"010203", "030405"},
					Logo:                          "logo.svg",
				},
			})
		})
	})

	t.Run("TestListModels", func(t *testing.T) {
		t.Run("DefaultPaths", func(t *testing.T) {
			list, err := s.ListModels(store.ListModelsRequest{
				BrandID: "foo-vendor",
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceModel{
				{
					ModelID: "dev1",
					Name:    "Device 1",
				},
				{
					ModelID: "dev2",
					Name:    "Device 2",
				},
			})
		})

		t.Run("Limit", func(t *testing.T) {
			list, err := s.ListModels(store.ListModelsRequest{
				BrandID: "foo-vendor",
				Limit: &pbtypes.UInt32Value{
					Value: 1,
				},
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceModel{
				{
					ModelID: "dev1",
					Name:    "Device 1",
				},
			})
		})

		t.Run("Offset", func(t *testing.T) {
			list, err := s.ListModels(store.ListModelsRequest{
				BrandID: "foo-vendor",
				Offset: &pbtypes.UInt32Value{
					Value: 1,
				},
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceModel{
				{
					ModelID: "dev2",
					Name:    "Device 2",
				},
			})
		})

		t.Run("Paths", func(t *testing.T) {
			list, err := s.ListModels(store.ListModelsRequest{
				BrandID: "foo-vendor",
				Paths:   ttnpb.EndDeviceModelFieldPathsNested,
			})
			a.So(err, should.BeNil)
			a.So(list, should.Resemble, []*ttnpb.EndDeviceModel{
				{
					ModelID:     "dev1",
					Name:        "Device 1",
					Description: "My Description",
					HardwareVersions: []*ttnpb.EndDeviceModel_Version{
						{
							Version:    "1.0",
							Numeric:    1,
							PartNumber: "P4RTN0",
						},
					},
					FirmwareVersions: []*ttnpb.EndDeviceModel_FirmwareVersion{
						{
							Version:          "1.0",
							HardwareVersions: []string{"1.0"},
							Profiles: map[string]*ttnpb.EndDeviceModel_FirmwareVersion_Profile{
								"EU_863_870": {
									CodecID:          "foo-codec",
									ProfileID:        "profile1",
									LoRaWANCertified: true,
								},
								"US_902_928": {
									CodecID:          "foo-codec",
									ProfileID:        "profile2",
									LoRaWANCertified: true,
								},
							},
						},
					},
				},
				{
					ModelID:     "dev2",
					Name:        "Device 2",
					Description: "My Description 2",
					HardwareVersions: []*ttnpb.EndDeviceModel_Version{
						{
							Version:    "2.0",
							Numeric:    2,
							PartNumber: "P4RTN0",
						},
					},
					FirmwareVersions: []*ttnpb.EndDeviceModel_FirmwareVersion{
						{
							Version:          "1.1",
							HardwareVersions: []string{"2.0"},
							Profiles: map[string]*ttnpb.EndDeviceModel_FirmwareVersion_Profile{
								"EU_433": {
									CodecID:          "foo-codec",
									ProfileID:        "profile2",
									LoRaWANCertified: true,
								},
							},
						},
					},
					Sensors: []string{"temperature"},
				},
			})
		})
	})

	t.Run("TestGetCodecs", func(t *testing.T) {
		t.Run("Missing", func(t *testing.T) {
			a := assertions.New(t)

			for _, ids := range []ttnpb.EndDeviceVersionIdentifiers{
				{
					BrandID: "unknown-vendor",
				},
				{
					BrandID: "foo-vendor",
					ModelID: "unknown-model",
				},
				{
					BrandID:         "foo-vendor",
					ModelID:         "dev1",
					FirmwareVersion: "unknown-version",
				},
				{
					BrandID:         "foo-vendor",
					ModelID:         "dev1",
					FirmwareVersion: "1.0",
					BandID:          "unknown-band",
				},
			} {
				codec, err := s.GetDownlinkDecoder(&ids)
				a.So(err, should.NotBeNil)
				a.So(codec, should.Equal, "")
			}
		})
		for _, tc := range []struct {
			name  string
			f     func(*ttnpb.EndDeviceVersionIdentifiers) (string, error)
			codec string
		}{
			{
				name:  "UplinkDecoder",
				f:     s.GetUplinkDecoder,
				codec: "uplink decoder",
			},
			{
				name:  "DownlinkDecoder",
				f:     s.GetDownlinkDecoder,
				codec: "downlink decoder",
			},
			{
				name:  "DownlinkEncoder",
				f:     s.GetDownlinkEncoder,
				codec: "downlink encoder",
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				a := assertions.New(t)

				versionIDs := &ttnpb.EndDeviceVersionIdentifiers{
					BrandID:         "foo-vendor",
					ModelID:         "dev1",
					FirmwareVersion: "1.0",
					BandID:          "EU_863_870",
				}
				codec, err := tc.f(versionIDs)
				a.So(err, should.BeNil)
				a.So(codec, should.Equal, tc.codec)
			})
		}
	})

	t.Run("GetTemplate", func(t *testing.T) {
		// TODO
	})
}

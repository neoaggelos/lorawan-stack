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

package store

import (
	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"gopkg.in/yaml.v2"
)

// remoteStore implements the Store interface using a fetcher.
type remoteStore struct {
	fetcher fetch.Interface
}

// NewRemoteStore initializes a new Store using a fetcher.
func NewRemoteStore(fetcher fetch.Interface) (Store, error) {
	return &remoteStore{fetcher}, nil
}

// paginate returns page start and end indices, and false if the page is invalid.
func paginate(size int, limit, offset *pbtypes.UInt32Value) (uint32, uint32, bool) {
	start, end := uint32(0), uint32(size)
	if offset != nil && offset.Value > 0 {
		start = offset.Value
	}
	if start >= end {
		return 0, 0, false
	}
	if limit != nil && limit.Value > 0 && start+limit.Value < end {
		end = start + limit.Value
	}
	return start, end, true
}

// ListBrands lists available end device vendors from the vendor/index.yaml file.
func (s *remoteStore) ListBrands(req ListBrandsRequest) ([]*ttnpb.EndDeviceBrand, error) {
	b, err := s.fetcher.File("vendor", "index.yaml")
	if err != nil {
		return nil, err
	}
	rawVendors := VendorsIndex{}
	if err := yaml.Unmarshal(b, &rawVendors); err != nil {
		return nil, err
	}

	start, end, ok := paginate(len(rawVendors.Vendors), req.Limit, req.Offset)
	if !ok {
		return []*ttnpb.EndDeviceBrand{}, nil
	}

	brands := make([]*ttnpb.EndDeviceBrand, 0, end-start)
	for idx := start; idx < end; idx++ {
		// Skip draft vendors
		if rawVendors.Vendors[idx].Draft {
			continue
		}
		pb, err := rawVendors.Vendors[idx].ToPB(req.Paths...)
		if err != nil {
			return nil, err
		}
		brands = append(brands, pb)
	}

	return brands, nil
}

var (
	errUnknownBrand = errors.DefineNotFound("unknown_brand", "unknown brand `{brand_id}`")
)

// ListModels lists available end device definitions.
func (s *remoteStore) ListModels(req ListModelsRequest) ([]*ttnpb.EndDeviceModel, error) {
	b, err := s.fetcher.File("vendor", req.BrandID, "index.yaml")
	if err != nil {
		return nil, errUnknownBrand.WithAttributes("brand_id", req.BrandID)
	}
	index := VendorEndDevicesIndex{}
	if err := yaml.Unmarshal(b, &index); err != nil {
		return nil, err
	}
	start, end, ok := paginate(len(index.EndDevices), req.Limit, req.Offset)
	if !ok {
		return []*ttnpb.EndDeviceModel{}, nil
	}

	defs := make([]*ttnpb.EndDeviceModel, 0, end-start)
	for idx := start; idx < end; idx++ {
		definitionID := index.EndDevices[idx]
		if req.ModelID != "" && definitionID != req.ModelID {
			continue
		}
		b, err := s.fetcher.File("vendor", req.BrandID, definitionID+".yaml")
		if err != nil {
			return nil, err
		}
		definition := EndDeviceModel{}
		if err := yaml.Unmarshal(b, &definition); err != nil {
			return nil, err
		}
		pb, err := definition.ToPB(definitionID, req.Paths...)
		if err != nil {
			return nil, err
		}
		defs = append(defs, pb)
	}
	return defs, nil
}

var (
	errNoModel = errors.DefineNotFound("no_model", "no model for brand `{brand_id}` and model ID `{model_id}`")
)

// GetTemplate retrieves an end device template for an end device definition.
func (s *remoteStore) GetTemplate(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	defs, err := s.ListModels(ListModelsRequest{
		BrandID: ids.BrandID,
		ModelID: ids.ModelID,
		Paths: []string{
			"firmware_versions",
		},
	})
	if err != nil {
		return nil, err
	}
	for _, def := range defs {
		for _, ver := range def.FirmwareVersions {
			if ver.Version != ids.FirmwareVersion {
				continue
			}

			if _, ok := BandIDToRegion[ids.BandID]; !ok {
				return nil, errUnknownBand.WithAttributes("unknown_band", ids.BandID)
			}
			profileInfo, ok := ver.Profiles[ids.BandID]
			if !ok {
				return nil, errNoProfile.WithAttributes(
					"band_id", ids.BandID,
				)
			}

			b, err := s.fetcher.File("vendor", ids.BrandID, profileInfo.ProfileID+".yaml")
			if err != nil {
				return nil, err
			}
			profile := EndDeviceProfile{}
			if err := yaml.Unmarshal(b, &profile); err != nil {
				return nil, err
			}

			return profile.ToTemplatePB(ids, profileInfo)
		}
	}
	return nil, errNoModel.WithAttributes("brand_id", ids.BrandID, "model_id", ids.ModelID)
}

var (
	errNoCodec = errors.DefineNotFound("no_codec", "no codec defined for firmware version `{firmware_version}` and band `{band_id}`")
)

// getCodec retrieves codec information for a specific model and returns.
func (s *remoteStore) getCodec(ids *ttnpb.EndDeviceVersionIdentifiers, chooseFile func(EndDeviceCodec) string) (string, error) {
	defs, err := s.ListModels(ListModelsRequest{
		BrandID: ids.BrandID,
		ModelID: ids.ModelID,
		Paths: []string{
			"firmware_versions",
		},
	})
	if err != nil {
		return "", err
	}
	for _, def := range defs {
		for _, ver := range def.FirmwareVersions {
			if ver.Version != ids.FirmwareVersion {
				continue
			}

			if _, ok := BandIDToRegion[ids.BandID]; !ok {
				return "", errUnknownBand.WithAttributes("unknown_band", ids.BandID)
			}
			profileInfo, ok := ver.Profiles[ids.BandID]
			if !ok {
				return "", errNoProfile.WithAttributes(
					"band_id", ids.BandID,
				)
			}

			if profileInfo.CodecID == "" {
				return "", errNoCodec.WithAttributes("firmware_version", ids.FirmwareVersion, "band_id", ids.BandID)
			}

			codec := EndDeviceCodec{}
			b, err := s.fetcher.File("vendor", ids.BrandID, profileInfo.CodecID+".yaml")
			if err != nil {
				return "", err
			}
			if err := yaml.Unmarshal(b, &codec); err != nil {
				return "", err
			}
			if file := chooseFile(codec); file != "" {
				b, err := s.fetcher.File("vendor", ids.BrandID, file)
				if err != nil {
					return "", err
				}
				return string(b), nil
			}
		}
	}

	return "", errNoModel.WithAttributes("brand_id", ids.BrandID, "model_id", ids.ModelID)
}

// GetUplinkDecoder retrieves the codec for decoding uplink messages.
func (s *remoteStore) GetUplinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return s.getCodec(ids, func(c EndDeviceCodec) string { return c.UplinkDecoder.FileName })
}

// GetDownlinkDecoder retrieves the codec for decoding downlink messages.
func (s *remoteStore) GetDownlinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return s.getCodec(ids, func(c EndDeviceCodec) string { return c.DownlinkDecoder.FileName })
}

// GetDownlinkEncoder retrieves the codec for encoding downlink messages.
func (s *remoteStore) GetDownlinkEncoder(ids *ttnpb.EndDeviceVersionIdentifiers) (string, error) {
	return s.getCodec(ids, func(c EndDeviceCodec) string { return c.DownlinkEncoder.FileName })
}

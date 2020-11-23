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
	"time"

	pbtypes "github.com/gogo/protobuf/types"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// VendorsIndex is the format for the vendor/index.yaml file.
type VendorsIndex struct {
	Vendors []struct {
		ID       string `yaml:"id"`
		Name     string `yaml:"name"`
		VendorID int    `yaml:"vendorID"`
		Draft    bool   `yaml:"draft,omitempty"`
	} `yaml:"vendors"`
}

// VendorEndDevicesIndex is the format of the `vendor/<vendor-id>/index.yaml` file.
type VendorEndDevicesIndex struct {
	EndDevices []string `yaml:"endDevices"`
}

// EndDeviceDefinition is the format of the `vendor/<vendor-id>/<model-id>.yaml` file.
type EndDeviceDefinition struct {
	Name             string `yaml:"name"`
	Description      string `yaml:"description"`
	HardwareVersions []struct {
		Version string `yaml:"version"`
		Numeric uint32 `yaml:"numeric"`
	} `yaml:"hardwareVersions"`
	FirmwareVersions []struct {
		Version          string   `yaml:"version"`
		Numeric          uint32   `yaml:"numeric"`
		HardwareVersions []string `yaml:"hardwareVersions"`
		Profiles         map[string]struct {
			ID               string `yaml:"id"`
			Codec            string `yaml:"codec"`
			LoRaWANCertified bool   `yaml:"lorawanCertified"`
		} `yaml:"profiles"`
	} `yaml:"firmwareVersions"`
	Sensors    []string `yaml:"sensors"`
	Dimensions *struct {
		Width    float32 `yaml:"width"`
		Height   float32 `yaml:"height"`
		Diameter float32 `yaml:"diameter"`
		Length   float32 `yaml:"length"`
	} `yaml:"dimensions"`
	Weight  float32 `yaml:"weight"`
	Battery *struct {
		Replaceable bool   `yaml:"replaceable"`
		Type        string `yaml:"type"`
	} `yaml:"battery"`
	OperatingConditions *struct {
		Temperature *struct {
			Min float32 `yaml:"min"`
			Max float32 `yaml:"max"`
		} `yaml:"temperature"`
		RelativeHumidity *struct {
			Min float32 `yaml:"min"`
			Max float32 `yaml:"max"`
		} `yaml:"relativeHumidity"`
	} `yaml:"operatingConditions"`
	IPCode          string   `yaml:"ipCode"`
	KeyProvisioning []string `yaml:"keyProvisioning"`
	KeySecurity     string   `yaml:"keySecurity"`
	Photos          *struct {
		Main  string   `yaml:"main"`
		Other []string `yaml:"other"`
	} `yaml:"photos"`
	ProductURL   string `yaml:"productURL"`
	DatasheetURL string `yaml:"datasheetURL"`
	Compliances  *struct {
		Safety []struct {
			Body     string `yaml:"body"`
			Norm     string `yaml:"norm"`
			Standard string `yaml:"standard"`
			Version  string `yaml:"version"`
		} `yaml:"safety"`
		RadioEquipment []struct {
			Body     string `yaml:"body"`
			Norm     string `yaml:"norm"`
			Standard string `yaml:"standard"`
			Version  string `yaml:"version"`
		} `yaml:"radioEquipment"`
	}
	AdditionalRadios []string `yaml:"additionalRadios"`
}

// ToPB converts an EndDefinitionDefinition to a Protocol Buffer.
func (d EndDeviceDefinition) ToPB(id string, paths ...string) (*ttnpb.EndDeviceDefinition, error) {
	pb := &ttnpb.EndDeviceDefinition{
		DefinitionID:     id,
		Name:             d.Name,
		Description:      d.Description,
		FirmwareVersions: make([]*ttnpb.EndDeviceDefinition_FirmwareVersion, 0, len(d.FirmwareVersions)),
		Sensors:          d.Sensors,
		Weight:           d.Weight,
		IPCode:           d.IPCode,
		KeyProvisioning:  d.KeyProvisioning,
		KeySecurity:      d.KeySecurity,
		ProductURL:       d.ProductURL,
		DatasheetURL:     d.DatasheetURL,
		AdditionalRadios: d.AdditionalRadios,
	}

	if hwVersions := d.HardwareVersions; hwVersions != nil {
		pb.HardwareVersions = make([]*ttnpb.EndDeviceDefinition_Version, 0, len(hwVersions))
		for _, ver := range hwVersions {
			pb.HardwareVersions = append(pb.HardwareVersions, &ttnpb.EndDeviceDefinition_Version{
				Version: ver.Version,
				Numeric: ver.Numeric,
			})
		}
	}
	for _, ver := range d.FirmwareVersions {
		pbver := &ttnpb.EndDeviceDefinition_FirmwareVersion{
			Version:          ver.Version,
			Numeric:          ver.Numeric,
			HardwareVersions: ver.HardwareVersions,
		}
		pbver.Profiles = make(map[string]*ttnpb.EndDeviceDefinition_FirmwareVersion_Profile, len(ver.Profiles))
		for region, profile := range ver.Profiles {
			pbver.Profiles[RegionToBandID[region]] = &ttnpb.EndDeviceDefinition_FirmwareVersion_Profile{
				CodecID:          profile.Codec,
				ProfileID:        profile.ID,
				LoRaWANCertified: profile.LoRaWANCertified,
			}
		}
		pb.FirmwareVersions = append(pb.FirmwareVersions, pbver)
	}

	if dim := d.Dimensions; dim != nil {
		pb.Dimensions = &ttnpb.EndDeviceDefinition_Dimensions{
			Width:    d.Dimensions.Width,
			Height:   d.Dimensions.Height,
			Diameter: d.Dimensions.Diameter,
			Length:   d.Dimensions.Length,
		}
	}

	if battery := d.Battery; battery != nil {
		pb.Battery = &ttnpb.EndDeviceDefinition_Battery{
			Replaceable: &pbtypes.BoolValue{
				Value: d.Battery.Replaceable,
			},
			Type: d.Battery.Type,
		}
	}

	if oc := d.OperatingConditions; oc != nil {
		pb.OperatingConditions = &ttnpb.EndDeviceDefinition_OperatingConditions{}

		if rh := oc.RelativeHumidity; rh != nil {
			pb.OperatingConditions.RelativeHumidity = &ttnpb.EndDeviceDefinition_OperatingConditions_Limits{
				Min: &pbtypes.FloatValue{
					Value: rh.Min,
				},
				Max: &pbtypes.FloatValue{
					Value: rh.Max,
				},
			}
		}

		if temp := oc.Temperature; temp != nil {
			pb.OperatingConditions.Temperature = &ttnpb.EndDeviceDefinition_OperatingConditions_Limits{
				Min: &pbtypes.FloatValue{
					Value: temp.Min,
				},
				Max: &pbtypes.FloatValue{
					Value: temp.Max,
				},
			}
		}
	}

	if p := d.Photos; p != nil {
		pb.Photos = &ttnpb.EndDeviceDefinition_Photos{
			Main:  p.Main,
			Other: p.Other,
		}
	}

	if cs := d.Compliances; cs != nil {
		pb.Compliances = &ttnpb.EndDeviceDefinition_Compliances{
			Safety:         make([]*ttnpb.EndDeviceDefinition_Compliances_Compliance, 0, len(cs.Safety)),
			RadioEquipment: make([]*ttnpb.EndDeviceDefinition_Compliances_Compliance, 0, len(cs.RadioEquipment)),
		}

		for _, c := range cs.Safety {
			pb.Compliances.Safety = append(pb.Compliances.Safety, &ttnpb.EndDeviceDefinition_Compliances_Compliance{
				Version:  c.Version,
				Body:     c.Body,
				Standard: c.Standard,
				Norm:     c.Norm,
			})
		}
		for _, c := range cs.RadioEquipment {
			pb.Compliances.RadioEquipment = append(pb.Compliances.RadioEquipment, &ttnpb.EndDeviceDefinition_Compliances_Compliance{
				Version:  c.Version,
				Body:     c.Body,
				Standard: c.Standard,
				Norm:     c.Norm,
			})
		}
	}

	if len(paths) > 0 {
		pb2 := &ttnpb.EndDeviceDefinition{}
		if err := pb2.SetFields(pb, paths...); err != nil {
			return nil, err
		}
		pb = pb2
	}

	return pb, nil
}

// EndDeviceProfile is the format of the `vendor/<vendor-id>/<profile-id>.yaml` file.
type EndDeviceProfile struct {
	VendorProfileID uint32 `yaml:"vendorProfileID"`
	SupportsClassB  bool   `yaml:"supportsClassB"`
	ClassBTimeout   uint32 `yaml:"classBTimeout"`
	PingSlotPeriod  uint32 `yaml:"pingSlotPeriod"`

	PingSlotDataRateIndex     ttnpb.DataRateIndex `yaml:"pingSlotDataRateIndex"`
	PingSlotFrequency         float64             `yaml:"pingSlotFrequency"`
	SupportsClassC            bool                `yaml:"supportsClassC"`
	ClassCTimeout             uint32              `yaml:"classCTimeout"`
	MACVersion                string              `yaml:"macVersion"`
	RegionalParametersVersion string              `yaml:"regionalParametersVersion"`
	SupportsJoin              bool                `yaml:"supportsJoin"`
	Rx1Delay                  ttnpb.RxDelay       `yaml:"rx1Delay"`
	Rx1DataRateOffset         uint32              `yaml:"rx1DataRateOffset"`
	Rx2DataRateIndex          ttnpb.DataRateIndex `yaml:"rx2DataRateIndex"`
	Rx2Frequency              float64             `yaml:"rx2Frequency"`
	FactoryPresetFrequencies  []float64           `yaml:"factoryPresetFrequencies"`
	MaxEIRP                   float32             `yaml:"maxEIRP"`
	MaxDutyCycle              float32             `yaml:"maxDutyCycle"`
	Supports32BitFCnt         bool                `yaml:"supports32bitFCnt"`
}

var (
	errUnknownBand = errors.DefineNotFound("unknown_band", "unknown band `{band_id}`")
	errNoProfile   = errors.DefineNotFound("no_profile", "device does not support region `{region}`")

	errUnknownMACVersion                = errors.DefineNotFound("unknown_mac_version", "unknown LoRaWAN version `{mac_version}`")
	errUnknownRegionalParametersVersion = errors.DefineNotFound("unknown_regional_parameters_version", "unknown Regional Parameters version `{phyVersion}`")
)

// ToTemplatePB returns a ttnpb.EndDeviceTemplate from an end device profile.
func (p EndDeviceProfile) ToTemplatePB(ids *ttnpb.EndDeviceVersionIdentifiers, info *ttnpb.EndDeviceDefinition_FirmwareVersion_Profile) (*ttnpb.EndDeviceTemplate, error) {
	macVersion, ok := MACVersionToPB[p.MACVersion]
	if !ok {
		return nil, errUnknownMACVersion.WithAttributes("mac_version", p.MACVersion)
	}
	phyVersion, ok := RegionalParametersToPB[p.RegionalParametersVersion]
	if !ok {
		return nil, errUnknownRegionalParametersVersion.WithAttributes("phyVersion", p.RegionalParametersVersion)
	}

	paths := []string{
		"version_ids",
		"supports_join",
		"supports_class_b",
		"supports_class_c",
		"lorawan_version",
		"lorawan_phy_version",
	}
	dev := ttnpb.EndDevice{
		VersionIDs:        ids,
		SupportsJoin:      p.SupportsJoin,
		SupportsClassB:    p.SupportsClassB,
		SupportsClassC:    p.SupportsClassC,
		LoRaWANVersion:    macVersion,
		LoRaWANPHYVersion: phyVersion,
	}

	if info.CodecID != "" {
		dev.Formatters = &ttnpb.MessagePayloadFormatters{
			DownFormatter: ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
			UpFormatter:   ttnpb.PayloadFormatter_FORMATTER_REPOSITORY,
		}
		paths = append(paths, "formatters")
	}

	dev.MACSettings = &ttnpb.MACSettings{}
	if p.ClassBTimeout > 0 {
		t := time.Duration(p.ClassBTimeout) * time.Second
		dev.MACSettings.ClassBTimeout = &t
		paths = append(paths, "mac_settings.class_b_timeout")
	}
	if p.ClassCTimeout > 0 {
		t := time.Duration(p.ClassCTimeout) * time.Second
		dev.MACSettings.ClassCTimeout = &t
		paths = append(paths, "mac_settings.class_c_timeout")
	}
	if p.PingSlotDataRateIndex > 0 {
		dev.MACSettings.PingSlotDataRateIndex = &ttnpb.DataRateIndexValue{
			Value: p.PingSlotDataRateIndex,
		}
		paths = append(paths, "mac_settings.ping_slot_data_rate_index")
	}
	if p.PingSlotFrequency > 0 {
		dev.MACSettings.PingSlotFrequency = &pbtypes.UInt64Value{
			Value: uint64(p.PingSlotFrequency * 100000),
		}
		paths = append(paths, "mac_settings.ping_slot_frequency")
	}
	if p.PingSlotPeriod > 0 {
		dev.MACSettings.PingSlotPeriodicity = &ttnpb.PingSlotPeriodValue{
			Value: PingSlotPeriodToPB[p.PingSlotPeriod],
		}
		paths = append(paths, "mac_settings.ping_slot_periodicity")
	}
	if p.Rx1Delay > 0 {
		dev.MACSettings.Rx1Delay = &ttnpb.RxDelayValue{
			Value: p.Rx1Delay,
		}
		paths = append(paths, "mac_settings.rx1_delay")
	}
	if p.Rx1DataRateOffset > 0 {
		dev.MACSettings.Rx1DataRateOffset = &pbtypes.UInt32Value{
			Value: p.Rx1DataRateOffset,
		}
		paths = append(paths, "mac_settings.rx1_data_rate_offset")
	}
	if p.Rx2DataRateIndex > 0 {
		dev.MACSettings.Rx2DataRateIndex = &ttnpb.DataRateIndexValue{
			Value: p.Rx2DataRateIndex,
		}
		paths = append(paths, "mac_settings.rx2_data_rate_index")
	}
	if p.Rx2Frequency > 0 {
		dev.MACSettings.Rx2Frequency = &pbtypes.UInt64Value{
			Value: uint64(p.Rx2Frequency * 100000),
		}
	}
	if p.Supports32BitFCnt {
		dev.MACSettings.Supports32BitFCnt = &pbtypes.BoolValue{
			Value: true,
		}
		paths = append(paths, "mac_settings.supports_32_bit_f_cnt")
	}
	if fs := p.FactoryPresetFrequencies; fs != nil && len(fs) > 0 {
		dev.MACSettings.FactoryPresetFrequencies = make([]uint64, 0, len(fs))
		for _, freq := range fs {
			dev.MACSettings.FactoryPresetFrequencies = append(dev.MACSettings.FactoryPresetFrequencies, uint64(freq*100000))
		}
		paths = append(paths, "mac_settings.factory_preset_frequencies")
	}

	dev.MACState = &ttnpb.MACState{
		DesiredParameters: ttnpb.MACParameters{},
	}
	if p.MaxEIRP > 0 {
		dev.MACState.DesiredParameters.MaxEIRP = p.MaxEIRP
		paths = append(paths, "mac_state.desired_parameters.max_eirp")
	}
	return &ttnpb.EndDeviceTemplate{
		EndDevice: dev,
		FieldMask: pbtypes.FieldMask{
			Paths: paths,
		},
	}, nil
}

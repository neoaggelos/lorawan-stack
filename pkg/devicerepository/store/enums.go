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
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// RegionToBandID maps LoRaWAN schema regions to TTS Band IDs.
var RegionToBandID = map[string]string{
	"EU863-870": band.EU_863_870,
	"US902-928": band.US_902_928,
	"CN779-787": band.CN_779_787,
	"EU433":     band.EU_433,
	"AU915-928": band.AU_915_928,
	"CN470-510": band.CN_470_510,
	"AS923":     band.AS_923,
	"KR920-923": band.KR_920_923,
	"IN865-867": band.IN_865_867,
	"RU864-870": band.RU_864_870,
}

// BandIDToRegion is the inverse mapping of RegionToBandID.
var BandIDToRegion map[string]string

// MACVersionToPB maps LoRaWAN schema MAC versions to ttnpb.MACVersion enum values.
var MACVersionToPB = map[string]ttnpb.MACVersion{
	"1.0":   ttnpb.MAC_V1_0,
	"1.0.1": ttnpb.MAC_V1_0_1,
	"1.0.2": ttnpb.MAC_V1_0_2,
	"1.0.3": ttnpb.MAC_V1_0_3,
	"1.0.4": ttnpb.MAC_V1_0_4,
	"1.1":   ttnpb.MAC_V1_1,
}

// RegionalParametersToPB maps LoRaWAN schema regional parameters to ttnpb.PHYVersion enum values.
var RegionalParametersToPB = map[string]ttnpb.PHYVersion{
	"TS001-1.0":        ttnpb.PHY_V1_0,
	"TS001-1.0.1":      ttnpb.PHY_V1_0_1,
	"RP001-1.0.2":      ttnpb.PHY_V1_0_2_REV_A,
	"RP001-1.0.2-RevB": ttnpb.PHY_V1_0_2_REV_B,
	"RP001-1.0.3-RevA": ttnpb.PHY_V1_0_3_REV_A,
	// TODO: Add Regional Parameters for LoRaWAN version 1.0.4 (https://github.com/TheThingsNetwork/lorawan-stack/issues/3513)
	// "RP002-1.0.0": ttnpb.PHY_UNKNOWN,
	// "RP002-1.0.1": ttnpb.PHY_UNKNOWN,
	"RP001-1.1":      ttnpb.PHY_V1_1_REV_A,
	"RP001-1.1-RevB": ttnpb.PHY_V1_1_REV_B,
}

// PingSlotPeriodToPB maps LoRaWAN schema ping slot period to ttnpb.PingSlotPeriod enum values.
var PingSlotPeriodToPB = map[uint32]ttnpb.PingSlotPeriod{
	32:  ttnpb.PING_EVERY_32S,
	64:  ttnpb.PING_EVERY_64S,
	128: ttnpb.PING_EVERY_128S,
	// TODO: Support ping slot periods for 256, 512, 1024, 2048, 4096 seconds.
	// 256:  ttnpb.PING_EVERY_256S,
	// 512:  ttnpb.PING_EVERY_512S,
	// 1024: ttnpb.PING_EVERY_1024S,
	// 2048: ttnpb.PING_EVERY_2048S,
	// 4096: ttnpb.PING_EVERY_4096S,
}

func init() {
	BandIDToRegion = make(map[string]string, len(RegionToBandID))
	for region, bandID := range RegionToBandID {
		BandIDToRegion[bandID] = region
	}
}

// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportT1E1Controller struct {
	Id               types.String                     `tfsdk:"id"`
	Version          types.Int64                      `tfsdk:"version"`
	Name             types.String                     `tfsdk:"name"`
	Description      types.String                     `tfsdk:"description"`
	FeatureProfileId types.String                     `tfsdk:"feature_profile_id"`
	Type             types.String                     `tfsdk:"type"`
	Slot             types.String                     `tfsdk:"slot"`
	SlotVariable     types.String                     `tfsdk:"slot_variable"`
	Entries          []TransportT1E1ControllerEntries `tfsdk:"entries"`
}

type TransportT1E1ControllerEntries struct {
	T1Description       types.String                                  `tfsdk:"t1_description"`
	T1Framing           types.String                                  `tfsdk:"t1_framing"`
	T1FramingVariable   types.String                                  `tfsdk:"t1_framing_variable"`
	T1Linecode          types.String                                  `tfsdk:"t1_linecode"`
	T1LinecodeVariable  types.String                                  `tfsdk:"t1_linecode_variable"`
	E1Description       types.String                                  `tfsdk:"e1_description"`
	E1Framing           types.String                                  `tfsdk:"e1_framing"`
	E1FramingVariable   types.String                                  `tfsdk:"e1_framing_variable"`
	E1Linecode          types.String                                  `tfsdk:"e1_linecode"`
	E1LinecodeVariable  types.String                                  `tfsdk:"e1_linecode_variable"`
	CableLength         types.String                                  `tfsdk:"cable_length"`
	LengthShort         types.String                                  `tfsdk:"length_short"`
	LengthShortVariable types.String                                  `tfsdk:"length_short_variable"`
	LengthLong          types.String                                  `tfsdk:"length_long"`
	LengthLongVariable  types.String                                  `tfsdk:"length_long_variable"`
	ClockSource         types.String                                  `tfsdk:"clock_source"`
	LineMode            types.String                                  `tfsdk:"line_mode"`
	LineModeVariable    types.String                                  `tfsdk:"line_mode_variable"`
	Description         types.String                                  `tfsdk:"description"`
	DescriptionVariable types.String                                  `tfsdk:"description_variable"`
	ChannelGroups       []TransportT1E1ControllerEntriesChannelGroups `tfsdk:"channel_groups"`
}

type TransportT1E1ControllerEntriesChannelGroups struct {
	ChannelGroup         types.Int64  `tfsdk:"channel_group"`
	ChannelGroupVariable types.String `tfsdk:"channel_group_variable"`
	TimeSlot             types.String `tfsdk:"time_slot"`
	TimeSlotVariable     types.String `tfsdk:"time_slot_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportT1E1Controller) getModel() string {
	return "transport_t1_e1_controller"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportT1E1Controller) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/t1-e1-controller", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportT1E1Controller) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.Type.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"type.optionType", "global")
			body, _ = sjson.Set(body, path+"type.value", data.Type.ValueString())
		}
	}

	if !data.SlotVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"slot.optionType", "variable")
			body, _ = sjson.Set(body, path+"slot.value", data.SlotVariable.ValueString())
		}
	} else if !data.Slot.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"slot.optionType", "global")
			body, _ = sjson.Set(body, path+"slot.value", data.Slot.ValueString())
		}
	}
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.T1Description.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.T1.name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.T1.name.value", item.T1Description.ValueString())
				}
			}

			if !item.T1FramingVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.T1.framing.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "basic.T1.framing.value", item.T1FramingVariable.ValueString())
				}
			} else if !item.T1Framing.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.T1.framing.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.T1.framing.value", item.T1Framing.ValueString())
				}
			}

			if !item.T1LinecodeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.T1.linecode.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "basic.T1.linecode.value", item.T1LinecodeVariable.ValueString())
				}
			} else if !item.T1Linecode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.T1.linecode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.T1.linecode.value", item.T1Linecode.ValueString())
				}
			}
			if !item.E1Description.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.E1.name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.E1.name.value", item.E1Description.ValueString())
				}
			}

			if !item.E1FramingVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.E1.framing.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "basic.E1.framing.value", item.E1FramingVariable.ValueString())
				}
			} else if !item.E1Framing.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.E1.framing.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.E1.framing.value", item.E1Framing.ValueString())
				}
			}

			if !item.E1LinecodeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.E1.linecode.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "basic.E1.linecode.value", item.E1LinecodeVariable.ValueString())
				}
			} else if !item.E1Linecode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "basic.E1.linecode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "basic.E1.linecode.value", item.E1Linecode.ValueString())
				}
			}
			if !item.CableLength.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cable.cableLength.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "cable.cableLength.value", item.CableLength.ValueString())
				}
			}

			if !item.LengthShortVariable.IsNull() {
				if true && item.CableLength.ValueString() == "short" {
					itemBody, _ = sjson.Set(itemBody, "cable.lengthShort.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "cable.lengthShort.value", item.LengthShortVariable.ValueString())
				}
			} else if !item.LengthShort.IsNull() {
				if true && item.CableLength.ValueString() == "short" {
					itemBody, _ = sjson.Set(itemBody, "cable.lengthShort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "cable.lengthShort.value", item.LengthShort.ValueString())
				}
			}

			if !item.LengthLongVariable.IsNull() {
				if true && item.CableLength.ValueString() == "long" {
					itemBody, _ = sjson.Set(itemBody, "cable.lengthLong.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "cable.lengthLong.value", item.LengthLongVariable.ValueString())
				}
			} else if !item.LengthLong.IsNull() {
				if true && item.CableLength.ValueString() == "long" {
					itemBody, _ = sjson.Set(itemBody, "cable.lengthLong.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "cable.lengthLong.value", item.LengthLong.ValueString())
				}
			}
			if item.ClockSource.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "clockSource.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "clockSource.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "clockSource.value", item.ClockSource.ValueString())
				}
			}

			if !item.LineModeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "lineMode.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "lineMode.value", item.LineModeVariable.ValueString())
				}
			} else if item.LineMode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "lineMode.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "lineMode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "lineMode.value", item.LineMode.ValueString())
				}
			}

			if !item.DescriptionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.DescriptionVariable.ValueString())
				}
			} else if item.Description.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.Description.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "channelGroup", []interface{}{})
				for _, childItem := range item.ChannelGroups {
					itemChildBody := ""

					if !childItem.ChannelGroupVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "number.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "number.value", childItem.ChannelGroupVariable.ValueString())
						}
					} else if !childItem.ChannelGroup.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "number.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "number.value", childItem.ChannelGroup.ValueInt64())
						}
					}

					if !childItem.TimeSlotVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeslots.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeslots.value", childItem.TimeSlotVariable.ValueString())
						}
					} else if !childItem.TimeSlot.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeslots.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeslots.value", childItem.TimeSlot.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "channelGroup.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"controllerTxExList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportT1E1Controller) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Type = types.StringNull()

	if t := res.Get(path + "type.optionType"); t.Exists() {
		va := res.Get(path + "type.value")
		if t.String() == "global" {
			data.Type = types.StringValue(va.String())
		}
	}
	data.Slot = types.StringNull()
	data.SlotVariable = types.StringNull()
	if t := res.Get(path + "slot.optionType"); t.Exists() {
		va := res.Get(path + "slot.value")
		if t.String() == "variable" {
			data.SlotVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Slot = types.StringValue(va.String())
		}
	}
	oldEntries := data.Entries
	if value := res.Get(path + "controllerTxExList"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]TransportT1E1ControllerEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportT1E1ControllerEntries{}
			item.T1Description = types.StringNull()

			if t := v.Get("basic.T1.name.optionType"); t.Exists() {
				va := v.Get("basic.T1.name.value")
				if t.String() == "global" {
					item.T1Description = types.StringValue(va.String())
				}
			}
			item.T1Framing = types.StringNull()
			item.T1FramingVariable = types.StringNull()
			if t := v.Get("basic.T1.framing.optionType"); t.Exists() {
				va := v.Get("basic.T1.framing.value")
				if t.String() == "variable" {
					item.T1FramingVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.T1Framing = types.StringValue(va.String())
				}
			}
			item.T1Linecode = types.StringNull()
			item.T1LinecodeVariable = types.StringNull()
			if t := v.Get("basic.T1.linecode.optionType"); t.Exists() {
				va := v.Get("basic.T1.linecode.value")
				if t.String() == "variable" {
					item.T1LinecodeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.T1Linecode = types.StringValue(va.String())
				}
			}
			item.E1Description = types.StringNull()

			if t := v.Get("basic.E1.name.optionType"); t.Exists() {
				va := v.Get("basic.E1.name.value")
				if t.String() == "global" {
					item.E1Description = types.StringValue(va.String())
				}
			}
			item.E1Framing = types.StringNull()
			item.E1FramingVariable = types.StringNull()
			if t := v.Get("basic.E1.framing.optionType"); t.Exists() {
				va := v.Get("basic.E1.framing.value")
				if t.String() == "variable" {
					item.E1FramingVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.E1Framing = types.StringValue(va.String())
				}
			}
			item.E1Linecode = types.StringNull()
			item.E1LinecodeVariable = types.StringNull()
			if t := v.Get("basic.E1.linecode.optionType"); t.Exists() {
				va := v.Get("basic.E1.linecode.value")
				if t.String() == "variable" {
					item.E1LinecodeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.E1Linecode = types.StringValue(va.String())
				}
			}
			item.CableLength = types.StringNull()

			if t := v.Get("cable.cableLength.optionType"); t.Exists() {
				va := v.Get("cable.cableLength.value")
				if t.String() == "global" {
					item.CableLength = types.StringValue(va.String())
				}
			}
			item.LengthShort = types.StringNull()
			item.LengthShortVariable = types.StringNull()
			if t := v.Get("cable.lengthShort.optionType"); t.Exists() {
				va := v.Get("cable.lengthShort.value")
				if t.String() == "variable" {
					item.LengthShortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.LengthShort = types.StringValue(va.String())
				}
			}
			item.LengthLong = types.StringNull()
			item.LengthLongVariable = types.StringNull()
			if t := v.Get("cable.lengthLong.optionType"); t.Exists() {
				va := v.Get("cable.lengthLong.value")
				if t.String() == "variable" {
					item.LengthLongVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.LengthLong = types.StringValue(va.String())
				}
			}
			item.ClockSource = types.StringNull()

			if t := v.Get("clockSource.optionType"); t.Exists() {
				va := v.Get("clockSource.value")
				if t.String() == "global" {
					item.ClockSource = types.StringValue(va.String())
				}
			}
			item.LineMode = types.StringNull()
			item.LineModeVariable = types.StringNull()
			if t := v.Get("lineMode.optionType"); t.Exists() {
				va := v.Get("lineMode.value")
				if t.String() == "variable" {
					item.LineModeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.LineMode = types.StringValue(va.String())
				}
			}
			item.Description = types.StringNull()
			item.DescriptionVariable = types.StringNull()
			if t := v.Get("description.optionType"); t.Exists() {
				va := v.Get("description.value")
				if t.String() == "variable" {
					item.DescriptionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Description = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("channelGroup"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ChannelGroups = make([]TransportT1E1ControllerEntriesChannelGroups, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportT1E1ControllerEntriesChannelGroups{}
					cItem.ChannelGroup = types.Int64Null()
					cItem.ChannelGroupVariable = types.StringNull()
					if t := cv.Get("number.optionType"); t.Exists() {
						va := cv.Get("number.value")
						if t.String() == "variable" {
							cItem.ChannelGroupVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.ChannelGroup = types.Int64Value(va.Int())
						}
					}
					cItem.TimeSlot = types.StringNull()
					cItem.TimeSlotVariable = types.StringNull()
					if t := cv.Get("timeslots.optionType"); t.Exists() {
						va := cv.Get("timeslots.value")
						if t.String() == "variable" {
							cItem.TimeSlotVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.TimeSlot = types.StringValue(va.String())
						}
					}
					item.ChannelGroups = append(item.ChannelGroups, cItem)
					return true
				})
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		data.Entries = nil
	}
	if !fullRead {
		resultEntries := make([]TransportT1E1ControllerEntries, 0, len(data.Entries))
		matchedEntries := make([]bool, len(data.Entries))
		for _, oldItem := range oldEntries {
			for ni := range data.Entries {
				if matchedEntries[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.T1Description.ValueString() != data.Entries[ni].T1Description.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.E1Description.ValueString() != data.Entries[ni].E1Description.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedEntries[ni] = true
					{
						resultC := make([]TransportT1E1ControllerEntriesChannelGroups, 0, len(data.Entries[ni].ChannelGroups))
						matchedC := make([]bool, len(data.Entries[ni].ChannelGroups))
						for _, oldCItem := range oldItem.ChannelGroups {
							for nci := range data.Entries[ni].ChannelGroups {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.ChannelGroupVariable.ValueString() != "" || data.Entries[ni].ChannelGroups[nci].ChannelGroupVariable.ValueString() != "") {
									if oldCItem.ChannelGroupVariable.ValueString() != data.Entries[ni].ChannelGroups[nci].ChannelGroupVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.ChannelGroup.ValueInt64() != data.Entries[ni].ChannelGroups[nci].ChannelGroup.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.TimeSlotVariable.ValueString() != "" || data.Entries[ni].ChannelGroups[nci].TimeSlotVariable.ValueString() != "") {
									if oldCItem.TimeSlotVariable.ValueString() != data.Entries[ni].ChannelGroups[nci].TimeSlotVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.TimeSlot.ValueString() != data.Entries[ni].ChannelGroups[nci].TimeSlot.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Entries[ni].ChannelGroups[nci])
									break
								}
							}
						}
						for nci := range data.Entries[ni].ChannelGroups {
							if !matchedC[nci] {
								resultC = append(resultC, data.Entries[ni].ChannelGroups[nci])
							}
						}
						data.Entries[ni].ChannelGroups = resultC
					}
					resultEntries = append(resultEntries, data.Entries[ni])
					break
				}
			}
		}
		for ni := range data.Entries {
			if !matchedEntries[ni] {
				resultEntries = append(resultEntries, data.Entries[ni])
			}
		}
		data.Entries = resultEntries
	}
}

// End of section. //template:end fromBody

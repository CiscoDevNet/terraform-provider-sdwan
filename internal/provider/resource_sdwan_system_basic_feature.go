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
	"regexp"
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SystemBasicProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemBasicProfileParcelResource{}

func NewSystemBasicProfileParcelResource() resource.Resource {
	return &SystemBasicProfileParcelResource{}
}

type SystemBasicProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemBasicProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_basic_feature"
}

func (r *SystemBasicProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System Basic Feature.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Optional:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the timezone").AddStringEnumDescription("Europe/Andorra", "Asia/Dubai", "Asia/Kabul", "America/Antigua", "America/Anguilla", "Europe/Tirane", "Asia/Yerevan", "Africa/Luanda", "Antarctica/McMurdo", "Antarctica/Rothera", "Antarctica/Palmer", "Antarctica/Mawson", "Antarctica/Davis", "Antarctica/Casey", "Antarctica/Vostok", "Antarctica/DumontDUrville", "Antarctica/Syowa", "America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Salta", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/San_Luis", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia", "Pacific/Pago_Pago", "Europe/Vienna", "Australia/Lord_Howe", "Antarctica/Macquarie", "Australia/Hobart", "Australia/Currie", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth", "Australia/Eucla", "America/Aruba", "Europe/Mariehamn", "Asia/Baku", "Europe/Sarajevo", "America/Barbados", "Asia/Dhaka", "Europe/Brussels", "Africa/Ouagadougou", "Europe/Sofia", "Asia/Bahrain", "Africa/Bujumbura", "Africa/Porto-Novo", "America/St_Barthelemy", "Atlantic/Bermuda", "Asia/Brunei", "America/La_Paz", "America/Kralendijk", "America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Santarem", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco", "America/Nassau", "Asia/Thimphu", "Africa/Gaborone", "Europe/Minsk", "America/Belize", "America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Iqaluit", "America/Pangnirtung", "America/Resolute", "America/Atikokan", "America/Rankin_Inlet", "America/Winnipeg", "America/Rainy_River", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Cambridge_Bay", "America/Yellowknife", "America/Inuvik", "America/Creston", "America/Dawson_Creek", "America/Vancouver", "America/Whitehorse", "America/Dawson", "Indian/Cocos", "Africa/Kinshasa", "Africa/Lubumbashi", "Africa/Bangui", "Africa/Brazzaville", "Europe/Zurich", "Africa/Abidjan", "Pacific/Rarotonga", "America/Santiago", "Pacific/Easter", "Africa/Douala", "Asia/Shanghai", "Asia/Harbin", "Asia/Chongqing", "Asia/Urumqi", "Asia/Kashgar", "America/Bogota", "America/Costa_Rica", "America/Havana", "Atlantic/Cape_Verde", "America/Curacao", "Indian/Christmas", "Asia/Nicosia", "Europe/Prague", "Europe/Berlin", "Europe/Busingen", "Africa/Djibouti", "Europe/Copenhagen", "America/Dominica", "America/Santo_Domingo", "Africa/Algiers", "America/Guayaquil", "Pacific/Galapagos", "Europe/Tallinn", "Africa/Cairo", "Africa/El_Aaiun", "Africa/Asmara", "Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary", "Africa/Addis_Ababa", "Europe/Helsinki", "Pacific/Fiji", "Atlantic/Stanley", "Pacific/Chuuk", "Pacific/Pohnpei", "Pacific/Kosrae", "Atlantic/Faroe", "Europe/Paris", "Africa/Libreville", "Europe/London", "America/Grenada", "Asia/Tbilisi", "America/Cayenne", "Europe/Guernsey", "Africa/Accra", "Europe/Gibraltar", "America/Godthab", "America/Danmarkshavn", "America/Scoresbysund", "America/Thule", "Africa/Banjul", "Africa/Conakry", "America/Guadeloupe", "Africa/Malabo", "Europe/Athens", "Atlantic/South_Georgia", "America/Guatemala", "Pacific/Guam", "Africa/Bissau", "America/Guyana", "Asia/Hong_Kong", "America/Tegucigalpa", "Europe/Zagreb", "America/Port-au-Prince", "Europe/Budapest", "Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura", "Europe/Dublin", "Asia/Jerusalem", "Europe/Isle_of_Man", "Asia/Kolkata", "Indian/Chagos", "Asia/Baghdad", "Asia/Tehran", "Atlantic/Reykjavik", "Europe/Rome", "Europe/Jersey", "America/Jamaica", "Asia/Amman", "Asia/Tokyo", "Africa/Nairobi", "Asia/Bishkek", "Asia/Phnom_Penh", "Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati", "Indian/Comoro", "America/St_Kitts", "Asia/Pyongyang", "Asia/Seoul", "Asia/Kuwait", "America/Cayman", "Asia/Almaty", "Asia/Qyzylorda", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Oral", "Asia/Vientiane", "Asia/Beirut", "America/St_Lucia", "Europe/Vaduz", "Asia/Colombo", "Africa/Monrovia", "Africa/Maseru", "Europe/Vilnius", "Europe/Luxembourg", "Europe/Riga", "Africa/Tripoli", "Africa/Casablanca", "Europe/Monaco", "Europe/Chisinau", "Europe/Podgorica", "America/Marigot", "Indian/Antananarivo", "Pacific/Majuro", "Pacific/Kwajalein", "Europe/Skopje", "Africa/Bamako", "Asia/Rangoon", "Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan", "Asia/Macau", "Pacific/Saipan", "America/Martinique", "Africa/Nouakchott", "America/Montserrat", "Europe/Malta", "Indian/Mauritius", "Indian/Maldives", "Africa/Blantyre", "America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Matamoros", "America/Mazatlan", "America/Chihuahua", "America/Ojinaga", "America/Hermosillo", "America/Tijuana", "America/Santa_Isabel", "America/Bahia_Banderas", "Asia/Kuala_Lumpur", "Asia/Kuching", "Africa/Maputo", "Africa/Windhoek", "Pacific/Noumea", "Africa/Niamey", "Pacific/Norfolk", "Africa/Lagos", "America/Managua", "Europe/Amsterdam", "Europe/Oslo", "Asia/Kathmandu", "Pacific/Nauru", "Pacific/Niue", "Pacific/Auckland", "Pacific/Chatham", "Asia/Muscat", "America/Panama", "America/Lima", "Pacific/Tahiti", "Pacific/Marquesas", "Pacific/Gambier", "Pacific/Port_Moresby", "Asia/Manila", "Asia/Karachi", "Europe/Warsaw", "America/Miquelon", "Pacific/Pitcairn", "America/Puerto_Rico", "Asia/Gaza", "Asia/Hebron", "Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores", "Pacific/Palau", "America/Asuncion", "Asia/Qatar", "Indian/Reunion", "Europe/Bucharest", "Europe/Belgrade", "Europe/Kaliningrad", "Europe/Moscow", "Europe/Volgograd", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Novokuznetsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Yakutsk", "Asia/Khandyga", "Asia/Vladivostok", "Asia/Sakhalin", "Asia/Ust-Nera", "Asia/Magadan", "Asia/Kamchatka", "Asia/Anadyr", "Africa/Kigali", "Asia/Riyadh", "Pacific/Guadalcanal", "Indian/Mahe", "Africa/Khartoum", "Europe/Stockholm", "Asia/Singapore", "Atlantic/St_Helena", "Europe/Ljubljana", "Arctic/Longyearbyen", "Europe/Bratislava", "Africa/Freetown", "Europe/San_Marino", "Africa/Dakar", "Africa/Mogadishu", "America/Paramaribo", "Africa/Juba", "Africa/Sao_Tome", "America/El_Salvador", "America/Lower_Princes", "Asia/Damascus", "Africa/Mbabane", "America/Grand_Turk", "Africa/Ndjamena", "Indian/Kerguelen", "Africa/Lome", "Asia/Bangkok", "Asia/Dushanbe", "Pacific/Fakaofo", "Asia/Dili", "Asia/Ashgabat", "Africa/Tunis", "Pacific/Tongatapu", "Europe/Istanbul", "America/Port_of_Spain", "Pacific/Funafuti", "Asia/Taipei", "Africa/Dar_es_Salaam", "Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye", "Europe/Simferopol", "Africa/Kampala", "Pacific/Johnston", "Pacific/Midway", "Pacific/Wake", "America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Tell_City", "America/Indiana/Knox", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/North_Dakota/Beulah", "America/Denver", "America/Boise", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Sitka", "America/Yakutat", "America/Nome", "America/Adak", "America/Metlakatla", "Pacific/Honolulu", "America/Montevideo", "Asia/Samarkand", "Asia/Tashkent", "Europe/Vatican", "America/St_Vincent", "America/Caracas", "America/Tortola", "America/St_Thomas", "Asia/Ho_Chi_Minh", "Pacific/Efate", "Pacific/Wallis", "Pacific/Apia", "Asia/Aden", "Indian/Mayotte", "Africa/Johannesburg", "Africa/Lusaka", "Africa/Harare", "UTC").AddDefaultValueDescription("UTC").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Europe/Andorra", "Asia/Dubai", "Asia/Kabul", "America/Antigua", "America/Anguilla", "Europe/Tirane", "Asia/Yerevan", "Africa/Luanda", "Antarctica/McMurdo", "Antarctica/Rothera", "Antarctica/Palmer", "Antarctica/Mawson", "Antarctica/Davis", "Antarctica/Casey", "Antarctica/Vostok", "Antarctica/DumontDUrville", "Antarctica/Syowa", "America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Salta", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/San_Luis", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia", "Pacific/Pago_Pago", "Europe/Vienna", "Australia/Lord_Howe", "Antarctica/Macquarie", "Australia/Hobart", "Australia/Currie", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth", "Australia/Eucla", "America/Aruba", "Europe/Mariehamn", "Asia/Baku", "Europe/Sarajevo", "America/Barbados", "Asia/Dhaka", "Europe/Brussels", "Africa/Ouagadougou", "Europe/Sofia", "Asia/Bahrain", "Africa/Bujumbura", "Africa/Porto-Novo", "America/St_Barthelemy", "Atlantic/Bermuda", "Asia/Brunei", "America/La_Paz", "America/Kralendijk", "America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Santarem", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco", "America/Nassau", "Asia/Thimphu", "Africa/Gaborone", "Europe/Minsk", "America/Belize", "America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Iqaluit", "America/Pangnirtung", "America/Resolute", "America/Atikokan", "America/Rankin_Inlet", "America/Winnipeg", "America/Rainy_River", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Cambridge_Bay", "America/Yellowknife", "America/Inuvik", "America/Creston", "America/Dawson_Creek", "America/Vancouver", "America/Whitehorse", "America/Dawson", "Indian/Cocos", "Africa/Kinshasa", "Africa/Lubumbashi", "Africa/Bangui", "Africa/Brazzaville", "Europe/Zurich", "Africa/Abidjan", "Pacific/Rarotonga", "America/Santiago", "Pacific/Easter", "Africa/Douala", "Asia/Shanghai", "Asia/Harbin", "Asia/Chongqing", "Asia/Urumqi", "Asia/Kashgar", "America/Bogota", "America/Costa_Rica", "America/Havana", "Atlantic/Cape_Verde", "America/Curacao", "Indian/Christmas", "Asia/Nicosia", "Europe/Prague", "Europe/Berlin", "Europe/Busingen", "Africa/Djibouti", "Europe/Copenhagen", "America/Dominica", "America/Santo_Domingo", "Africa/Algiers", "America/Guayaquil", "Pacific/Galapagos", "Europe/Tallinn", "Africa/Cairo", "Africa/El_Aaiun", "Africa/Asmara", "Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary", "Africa/Addis_Ababa", "Europe/Helsinki", "Pacific/Fiji", "Atlantic/Stanley", "Pacific/Chuuk", "Pacific/Pohnpei", "Pacific/Kosrae", "Atlantic/Faroe", "Europe/Paris", "Africa/Libreville", "Europe/London", "America/Grenada", "Asia/Tbilisi", "America/Cayenne", "Europe/Guernsey", "Africa/Accra", "Europe/Gibraltar", "America/Godthab", "America/Danmarkshavn", "America/Scoresbysund", "America/Thule", "Africa/Banjul", "Africa/Conakry", "America/Guadeloupe", "Africa/Malabo", "Europe/Athens", "Atlantic/South_Georgia", "America/Guatemala", "Pacific/Guam", "Africa/Bissau", "America/Guyana", "Asia/Hong_Kong", "America/Tegucigalpa", "Europe/Zagreb", "America/Port-au-Prince", "Europe/Budapest", "Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura", "Europe/Dublin", "Asia/Jerusalem", "Europe/Isle_of_Man", "Asia/Kolkata", "Indian/Chagos", "Asia/Baghdad", "Asia/Tehran", "Atlantic/Reykjavik", "Europe/Rome", "Europe/Jersey", "America/Jamaica", "Asia/Amman", "Asia/Tokyo", "Africa/Nairobi", "Asia/Bishkek", "Asia/Phnom_Penh", "Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati", "Indian/Comoro", "America/St_Kitts", "Asia/Pyongyang", "Asia/Seoul", "Asia/Kuwait", "America/Cayman", "Asia/Almaty", "Asia/Qyzylorda", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Oral", "Asia/Vientiane", "Asia/Beirut", "America/St_Lucia", "Europe/Vaduz", "Asia/Colombo", "Africa/Monrovia", "Africa/Maseru", "Europe/Vilnius", "Europe/Luxembourg", "Europe/Riga", "Africa/Tripoli", "Africa/Casablanca", "Europe/Monaco", "Europe/Chisinau", "Europe/Podgorica", "America/Marigot", "Indian/Antananarivo", "Pacific/Majuro", "Pacific/Kwajalein", "Europe/Skopje", "Africa/Bamako", "Asia/Rangoon", "Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan", "Asia/Macau", "Pacific/Saipan", "America/Martinique", "Africa/Nouakchott", "America/Montserrat", "Europe/Malta", "Indian/Mauritius", "Indian/Maldives", "Africa/Blantyre", "America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Matamoros", "America/Mazatlan", "America/Chihuahua", "America/Ojinaga", "America/Hermosillo", "America/Tijuana", "America/Santa_Isabel", "America/Bahia_Banderas", "Asia/Kuala_Lumpur", "Asia/Kuching", "Africa/Maputo", "Africa/Windhoek", "Pacific/Noumea", "Africa/Niamey", "Pacific/Norfolk", "Africa/Lagos", "America/Managua", "Europe/Amsterdam", "Europe/Oslo", "Asia/Kathmandu", "Pacific/Nauru", "Pacific/Niue", "Pacific/Auckland", "Pacific/Chatham", "Asia/Muscat", "America/Panama", "America/Lima", "Pacific/Tahiti", "Pacific/Marquesas", "Pacific/Gambier", "Pacific/Port_Moresby", "Asia/Manila", "Asia/Karachi", "Europe/Warsaw", "America/Miquelon", "Pacific/Pitcairn", "America/Puerto_Rico", "Asia/Gaza", "Asia/Hebron", "Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores", "Pacific/Palau", "America/Asuncion", "Asia/Qatar", "Indian/Reunion", "Europe/Bucharest", "Europe/Belgrade", "Europe/Kaliningrad", "Europe/Moscow", "Europe/Volgograd", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Novokuznetsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Yakutsk", "Asia/Khandyga", "Asia/Vladivostok", "Asia/Sakhalin", "Asia/Ust-Nera", "Asia/Magadan", "Asia/Kamchatka", "Asia/Anadyr", "Africa/Kigali", "Asia/Riyadh", "Pacific/Guadalcanal", "Indian/Mahe", "Africa/Khartoum", "Europe/Stockholm", "Asia/Singapore", "Atlantic/St_Helena", "Europe/Ljubljana", "Arctic/Longyearbyen", "Europe/Bratislava", "Africa/Freetown", "Europe/San_Marino", "Africa/Dakar", "Africa/Mogadishu", "America/Paramaribo", "Africa/Juba", "Africa/Sao_Tome", "America/El_Salvador", "America/Lower_Princes", "Asia/Damascus", "Africa/Mbabane", "America/Grand_Turk", "Africa/Ndjamena", "Indian/Kerguelen", "Africa/Lome", "Asia/Bangkok", "Asia/Dushanbe", "Pacific/Fakaofo", "Asia/Dili", "Asia/Ashgabat", "Africa/Tunis", "Pacific/Tongatapu", "Europe/Istanbul", "America/Port_of_Spain", "Pacific/Funafuti", "Asia/Taipei", "Africa/Dar_es_Salaam", "Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye", "Europe/Simferopol", "Africa/Kampala", "Pacific/Johnston", "Pacific/Midway", "Pacific/Wake", "America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Tell_City", "America/Indiana/Knox", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/North_Dakota/Beulah", "America/Denver", "America/Boise", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Sitka", "America/Yakutat", "America/Nome", "America/Adak", "America/Metlakatla", "Pacific/Honolulu", "America/Montevideo", "Asia/Samarkand", "Asia/Tashkent", "Europe/Vatican", "America/St_Vincent", "America/Caracas", "America/Tortola", "America/St_Thomas", "Asia/Ho_Chi_Minh", "Pacific/Efate", "Pacific/Wallis", "Pacific/Apia", "Asia/Aden", "Indian/Mayotte", "Africa/Johannesburg", "Africa/Lusaka", "Africa/Harare", "UTC"),
				},
			},
			"timezone_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"config_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set a text description of the device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"config_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the location of the device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"location_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"gps_longitude": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device physical longitude").AddFloatRangeDescription(-180, 180).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(-180, 180),
				},
			},
			"gps_longitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"gps_latitude": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device physical latitude").AddFloatRangeDescription(-90, 90).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(-90, 90),
				},
			},
			"gps_latitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"gps_geo_fencing_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Geo fencing").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"gps_geo_fencing_range": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the device’s geo fencing range").AddIntegerRangeDescription(100, 10000).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(100, 10000),
				},
			},
			"gps_geo_fencing_range_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"gps_sms_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device’s geo fencing SMS").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"gps_sms_mobile_numbers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set device’s geo fencing SMS phone number").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Mobile number, ex: 1231234414").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(7, 20),
								stringvalidator.RegexMatches(regexp.MustCompile(`[+][0-9]*`), ""),
							},
						},
						"number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"device_groups": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device groups").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"device_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"controller_groups": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a list of comma-separated controller groups").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"controller_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"overlay_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the Overlay ID").AddIntegerRangeDescription(1, 4294967295).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"overlay_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"port_offset": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the TLOC port offset when multiple devices are behind a NAT").AddIntegerRangeDescription(0, 19).AddDefaultValueDescription("0").String,
				Optional:            true,
			},
			"port_offset_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"port_hopping": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable port hopping").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"port_hopping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"control_session_pps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the policer rate for control sessions").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("300").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"control_session_pps_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_transport": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure tracking of transport").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"track_transport_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_interface_tag": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Tag attached to routes based on interface tracking").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"track_interface_tag_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"console_baud_rate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the console baud rate").AddStringEnumDescription("1200", "2400", "4800", "9600", "19200", "38400", "57600", "115200").AddDefaultValueDescription("9600").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1200", "2400", "4800", "9600", "19200", "38400", "57600", "115200"),
				},
			},
			"console_baud_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"max_omp_sessions": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the maximum number of OMP sessions <1..100> the device can have").AddIntegerRangeDescription(1, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"max_omp_sessions_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"multi_tenant": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device is multi-tenant").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"multi_tenant_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"track_default_gateway": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable default gateway tracking").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"track_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"admin_tech_on_failure": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Collect admin-tech before reboot due to daemon failure").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"admin_tech_on_failure_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Idle CLI timeout in minutes").AddIntegerRangeDescription(0, 300).String,
				Optional:            true,
			},
			"idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"on_demand_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable On-demand Tunnel").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"on_demand_enable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"on_demand_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the idle timeout for on-demand tunnels").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"on_demand_idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"transport_gateway": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable transport gateway").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"transport_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enhanced_app_aware_routing": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SLA Dampening and Enhanced App Routing.").AddStringEnumDescription("disabled", "aggressive", "moderate", "conservative").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "aggressive", "moderate", "conservative"),
				},
			},
			"enhanced_app_aware_routing_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"site_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Site Type").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"site_types_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_group_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Affinity Group Number").AddIntegerRangeDescription(1, 63).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 63),
				},
			},
			"affinity_group_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_group_preferences": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Affinity Group Preference").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"affinity_group_preferences_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_preference_auto": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Affinity Group Preference Auto").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"affinity_preference_auto_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"affinity_per_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Affinity Group Number for VRFs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"affinity_group_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Affinity Group Number").AddIntegerRangeDescription(1, 63).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 63),
							},
						},
						"affinity_group_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vrf_range": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Range of VRFs").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 11),
								stringvalidator.RegexMatches(regexp.MustCompile(`^\d+|\d+-\d+$`), ""),
							},
						},
						"vrf_range_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SystemBasicProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemBasicProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemBasic

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *SystemBasicProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemBasic

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SystemBasicProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemBasic

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *SystemBasicProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemBasic

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SystemBasicProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "system_basic_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q, %q", pattern, req.ID, count),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import

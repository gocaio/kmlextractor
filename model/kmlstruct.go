/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package model

//KML struct
type KML struct {
	Doc        Document    `xml:"Document"`
	Placemarks []PlaceMark `xml:"Placemark"`
	F          Folder      `xml:"Folder"`
	NetLink    NetworkLink `xml:"NetworkLink"`
}

//Document struct
type Document struct {
	Name    string      `xml:"name"`
	StyleM  StyleMap    `xml:"StyleMap"`
	Style   []Style     `xml:"Style"`
	F       []Folder    `xml:"Folder"`
	Lookat  LookAt      `xml:"LookAt"`
	NetLink NetworkLink `xml:"NetworkLink"`
}

//NetworkLink struct
type NetworkLink struct {
	Name       string `xml:"name"`
	Visibility int    `xml:"visibility"`
	Open       int    `xml:"open"`
	Lookat     LookAt `xml:"LookAt"`
	FlyToView  string `xml:"flyToView"`
	URLInfo    URL    `xml:"Url"`
}

//LookAt struct
type LookAt struct {
	Longitude    string `xml:"longitude"`
	Latitude     string `xml:"latitude"`
	Altitude     string `xml:"altitude"`
	Range        string `xml:"range"`
	Tilt         string `xml:"tilt"`
	Heading      string `xml:"heading"`
	AltitudeMode string `xml:"altitudeMode"`
}

//URL struct
type URL struct {
	Href            string `xml:"href"`
	RefreshInterval string `xml:"refreshInterval"`
	ViewRefreshMode string `xml:"viewRefreshMode"`
	ViewRefreshTime string `xml:"viewRefreshTime"`
	ViewBoundScale  string `xml:"viewBoundScale"`
}

//StyleMap struct
type StyleMap struct {
	Pairs []Pair `xml:"Pair"`
}

//Pair struct
type Pair struct {
	Key      string `xml:"key"`
	StyleURL string `xml:"styleUrl"`
}

//Style struct
type Style struct {
	IconS IconStyle `xml:"IconStyle"`
	LineS LineStyle `xml:"LineStyle"`
	ListS ListStyle `xml:"ListStyle"`
}

//IconStyle struct customizes the pinpoints
type IconStyle struct {
	Scale string  `xml:"scale"`
	Ico   Icon    `xml:"Icon"`
	HS    HotSpot `xml:"hotSpot"`
}

//Icon struct contains the link to the icon
type Icon struct {
	Href string `xml:"href"`
}

//HotSpot struct
type HotSpot struct {
	X      string `xml:"x,attr"`
	Y      string `xml:"y,attr"`
	Xunits string `xml:"xunits,attr"`
	Yunits string `xml:"yunits,attr"`
}

//LineStyle struct styles the strokes that join two or more geographical points
type LineStyle struct {
	Color string `xml:"color"`
	Width string `xml:"width"`
}

//ListStyle struct
type ListStyle struct {
	ListItemType    string   `xml:"listItemType"`
	ItemIco         ItemIcon `xml:"ItemIcon"`
	BgColor         string   `xml:"bgColor"`
	MaxSnippetLines string   `xml:"maxSnippetLines"`
}

//ItemIcon struct
type ItemIcon struct {
	State string `xml:"state"`
	Href  string `xml:"href"`
}

//Folder struct
type Folder struct {
	Name             string         `xml:"name"`
	Visibility       string         `xml:"description"`
	Open             int            `xml:"open"`
	F                []Folder       `xml:"Folder"`
	Lookat           LookAt         `xml:"LookAt"`
	Placemarks       []PlaceMark    `xml:"Placemark"`
	PhotoOverlayList []PhotoOverlay `xml:"PhotoOverlay"`
}

//PlaceMark struct pinpoints a location through reference to the Earth’s surface
type PlaceMark struct {
	Name        string     `xml:"name"`
	Description string     `xml:"description"`
	P           Point      `xml:"Point"`
	StyleURL    string     `xml:"styleUrl"`
	LineS       LineString `xml:"LineString"`
}

//LineString struct identifies at least two points to connect
type LineString struct {
	Tssellate    string `xml:"tessellate"`
	AltitudeMode string `xml:"altitudeMode"`
	Coordinates  string `xml:"coordinates"`
}

//PhotoOverlay struct
type PhotoOverlay struct {
	Name     string     `xml:"name"`
	Cam      Camera     `xml:"Camera"`
	Style    Style      `xml:"Style"`
	Ico      Icon       `xml:"Icon"`
	Rotation string     `xml:"rotation"`
	VVolume  ViewVolume `xml:"ViewVolume"`
	P        Point      `xml:"Point"`
}

//Camera struct
type Camera struct {
	Longitude    string `xml:"longitude"`
	Latitude     string `xml:"latitude"`
	Altitude     string `xml:"altitude"`
	Heading      string `xml:"heading"`
	Tilt         string `xml:"tilt"`
	Rll          string `xml:"roll"`
	AltitudeMode string `xml:"altitudeMode"`
}

//ViewVolume struct
type ViewVolume struct {
	LeftFov   string `xml:"leftFov"`
	RightFov  string `xml:"rightFov"`
	BottomFov string `xml:"bottomFov"`
	TopFov    string `xml:"topFov"`
	Near      string `xml:"near"`
}

//Point struct contains geographical information
type Point struct {
	AltitudeMode string `xml:"altitudeMode"`
	Coordinates  string `xml:"coordinates"`
}

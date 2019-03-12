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

package xmlextractor

//Document struct contains blablabla
type Document struct {
	Name   string   `xml:"name"`
	StyleM StyleMap `xml:"msn_ylw-pushpin"`
	Style1 Style    `xml:"sn_ylw-pushpin"`
	Style2 Style    `xml:"sh_ylw-pushpin"`
	F      Folder   `xml:"Folder"`
}

//StyleMap struct contains blablabla
type StyleMap struct {
	Pairs []Pair `xml:"Pair"`
}

//Pair struct contains blablabla
type Pair struct {
	Key      string `xml:"key"`
	StyleURL string `xml:"styleURL"`
}

//Style struct contains blablabla
type Style struct {
	IconS IconStyle `xml:"IconStyle"`
	LineS LineStyle `xml:"LineStyle"`
}

//IconStyle struct contains blablabla
type IconStyle struct {
	Scale int    `xml:"scale"`
	Icon  string `xml:"Icon"`
}

//LineStyle struct contains blablabla
type LineStyle struct {
	Color string `xml:"color"`
	Width int    `xml:"width"`
}

//Folder struct contains blablabla
type Folder struct {
	Name      string    `xml:"name"`
	Open      int       `xml:"open"`
	Placemark PlaceMark `xml:"Placemark"`
}

//PlaceMark struct contains blablabla
type PlaceMark struct {
	Name     string     `xml:"name"`
	StyleURL string     `xml:"styelUrl"`
	LineS    LineString `xml:"LineString"`
}

//LineString struct contains blablabla
type LineString struct {
}

func main() {

}

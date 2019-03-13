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

package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

//KML struct contains blablabla
type KML struct {
	Doc Document `xml:"Document"`
}

//Document struct contains blablabla
type Document struct {
	//XMLName xml.Name
	Name   string   `xml:"name"`
	StyleM StyleMap `xml:"StyleMap"`
	Style  []Style  `xml:"Style"`
	// F      Folder   `xml:"Folder"`
}

//StyleMap struct contains blablabla
type StyleMap struct {
	Pairs []Pair `xml:"Pair"`
}

//Pair struct contains blablabla
type Pair struct {
	Key      string `xml:"key"`
	StyleURL string `xml:"styleUrl"`
}

//Style struct contains blablabla
type Style struct {
	IconS IconStyle `xml:"IconStyle"`
	LineS LineStyle `xml:"LineStyle"`
}

//IconStyle struct contains blablabla
type IconStyle struct {
	Scale string  `xml:"scale"`
	Ico   Icon    `xml:"Icon"`
	HS    HotSpot `xml:"hotSpot"`
}

//Icon struct contains
type Icon struct {
	Href string `xml:"href"`
}

//HotSpot struct contains
type HotSpot struct {
	X      string `xml:"x,attr"`
	Y      string `xml:"y,attr"`
	Xunits string `xml:"xunits,attr"`
	Yunits string `xml:"yunits,attr"`
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

	file, err := os.Open("samples/kml_2/doc.kml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var xmlFile string
	var metadata KML

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xmlFile += scanner.Text()
	}
	// fmt.Println(xmlFile)
	if err := xml.Unmarshal([]byte(xmlFile), &metadata); err != nil {
		fmt.Println(err)
	}
	fmt.Println(metadata.Doc)
}

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

package test

import (
	"github.com/kevinborras/kmlextractor/model"
	"github.com/kevinborras/kmlextractor"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

var Kml = model.KML{
	Doc: model.Document{
		Lookat: model.LookAt{
			Longitude:    "180.0",
			Latitude:     "0.0",
			Altitude:     "1.5768778343115222E7",
			Heading:      "0.0",
			Tilt:         "0.0",
			Range:        "1.2742E7",
			AltitudeMode: "clampToGround",
		},
		NetLink: model.NetworkLink{
			Name:       "A sample ArcGrid file",
			Visibility: 1,
			Open:       1,
			Lookat: model.LookAt{
				Longitude:    "-1.4210854715202004E-14",
				Latitude:     "1.4210854715202004E-14",
				Altitude:     "1398874.1103087065",
				Heading:      "0.0",
				Tilt:         "0.0",
				Range:        "1130363.654413079",
				AltitudeMode: "clampToGround",
			},
			URLInfo: model.URL{
				Href:            "http://localhost:8080/geoserver/nurc/wms?service=wms&amp;request=GetMap&amp;version=1.1.1&amp;format=application/vnd.google-earth.kml+xml&amp;layers=nurc:Arc_Sample&amp;styles=rain&amp;height=768&amp;width=768&amp;transparent=false&amp;srs=EPSG:4326",
				RefreshInterval: "0.0",
				ViewRefreshMode: "onStop",
				ViewRefreshTime: "1.0",
				ViewBoundScale:  "1.0",
			},
		},
	},
}

func TestKMLDocumment(t *testing.T) {
	file, err := os.Open("../samples/kml_3.kml")
	if err != nil {
		log.Fatal(err)
	}
	actualResult, err := kmlextractor.GetContent(file)
	actualResult.Doc.NetLink.URLInfo.Href = strings.Replace(actualResult.Doc.NetLink.URLInfo.Href, "&", "&amp;", -1)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(Kml, actualResult) {
		t.Fatalf("Expected %v but got %v", Kml, actualResult)
	}
}
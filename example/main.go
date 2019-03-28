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
	"fmt"
	"github.com/kevinborras/kmlextractor"
	"log"
	"os"
	// "reflect"
)

func main() {

	file, err := os.Open(`../samples/kml_3.kml`)
	if err != nil {
		log.Fatal(err)
	}
	content, err := kmlextractor.GetContent(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
	// s := reflect.ValueOf(&content).Elem()
	// typeOfT := s.Type()
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Interface())
	// }

	// aux := kmlextractor.GetInterestingContent(&content)
	// s := reflect.ValueOf(&aux).Elem()
	// typeOfT := s.Type()
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Interface())
	// }
	// for _,placemark := range content.F.Placemarks{
	// 	fmt.Println(placemark)
	// }

}

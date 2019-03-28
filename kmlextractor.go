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

package kmlextractor

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"github.com/kevinborras/kmlextractor/model"
	"os"
	"strings"
)

//checkType checks if it is possible to unzip the .kml and obtain more information
func checkZipped(document *os.File) (bool, error) {

	var err error
	b1 := make([]byte, 2)
	_, err = document.Read(b1)
	if err != nil {
		return false, err
	}
	if string(b1) == "PK" {
		return true, err
	}
	return false, err

}

//GetContent func returns all the information insede the .kml file
func GetContent(document *os.File) (kmlInfo model.KML, err error) {

	flag, err := checkZipped(document)
	if err != nil {
		return kmlInfo, err
	}
	if flag {
		document.Close()
		//check for files
		oriName := document.Name()
		dot := strings.LastIndex(oriName, ".")
		zipName := oriName[:dot] + ".zip"
		err := os.Rename(oriName, zipName)
		if err != nil {
			return kmlInfo, err
		}
		read, err := zip.OpenReader(zipName)
		if err != nil {
			os.Rename(zipName, oriName)
			return kmlInfo, err
		}
		xmlFile := ""
		for _, file := range read.File {
			if file.Name == "doc.kml" {
				rc, err := file.Open()
				if err != nil {
					return kmlInfo, err
				}
				scanner := bufio.NewScanner(rc)
				for scanner.Scan() {
					xmlFile += scanner.Text()
				}
			}
		}
		if err := xml.Unmarshal([]byte(xmlFile), &kmlInfo); err != nil {
			os.Rename(zipName, oriName)
			return kmlInfo, err
		}
		read.Close()
		os.Rename(zipName, oriName)

	} else {
		fStats, err := document.Stat()
		if err != nil {
			return kmlInfo, err
		}
		//Byte slice of document size
		docBytes := make([]byte, fStats.Size())
		_, err = document.Read(docBytes)
		if err := xml.Unmarshal(docBytes, &kmlInfo); err != nil {
			return kmlInfo, err
		}
		document.Close()
	}

	return kmlInfo, err
}

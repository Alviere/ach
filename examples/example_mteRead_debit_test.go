// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Alviere/ach"
)

func Example_mteReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "mte-read.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		panic(fmt.Sprintf("Issue reading file: %+v \n", err))
	}
	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Total Amount: %s", strconv.Itoa(achFile.Batches[0].GetEntries()[0].Amount)+"\n")
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Addenda02: %s", achFile.Batches[0].GetEntries()[0].Addenda02.String()+"\n")

	// Output:
	// Total Amount: 10000
	// SEC Code: MTE
	// Addenda02: 702          2005091234561224      321 East Market Street     ANYTOWN        VA231380100000001

}

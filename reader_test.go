// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"os"
	"strings"
	"testing"
)

// TestDecode is a complete file decoding test.
func TestDecode(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.read file: %v", err)
	}
}

func TestRecordTypeUnknown(t *testing.T) {
	var line = "301 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrUnknownRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

func TestTwoFileHeaders(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()

	if !strings.Contains(err.Error(), ErrFileHeader.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}

}

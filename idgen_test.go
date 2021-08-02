// Copyright 2021 Jeffrey M Hodges.
// SPDX-License-Identifier: Apache-2.0

package idgen

import (
	"strings"
	"testing"
)

// Ensuring the interfaces exported by this project
var _ UUIDGenerator = &idGen{}

func TestIdGen(t *testing.T) {
	id, err := idGen{}.NewId()
	if err != nil {
		t.Fatalf("id gen error: %s", err)
	}
	parts := strings.Split(string(id), "-")
	if len(parts) != 5 {
		t.Errorf("Want 5 parts, got %d", len(parts))
	}
	if len(parts[0]) != 8 {
		t.Errorf("Want 8 chars for part 0, got %d", len(parts[0]))
	}
	if len(parts[1]) != 4 {
		t.Errorf("Want 4 chars for part 1, got %d", len(parts[1]))
	}
	if len(parts[2]) != 4 {
		t.Errorf("Want 4 chars for part 2, got %d", len(parts[2]))
	}
	if len(parts[3]) != 4 {
		t.Errorf("Want 4 chars for part 3, got %d", len(parts[3]))
	}
	if len(parts[4]) != 12 {
		t.Errorf("Want 12 chars for part 4, got %d", len(parts[4]))
	}
}

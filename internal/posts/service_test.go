// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package posts

import "testing"

func TestNoteFilename(t *testing.T) {
	got := noteFilename("Hello, World!")
	want := "hello-world.md"
	if got != want {
		t.Fatalf("noteFilename = %q, want %q", got, want)
	}
}

func TestNormalizeSlug(t *testing.T) {
	got := normalizeSlug("  99 Problems  ")
	want := "99-problems"
	if got != want {
		t.Fatalf("normalizeSlug = %q, want %q", got, want)
	}
}

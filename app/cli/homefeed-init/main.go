// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"homefeed/internal/workspace"
)

func main() {
	root := flag.String("root", "", "workspace root (defaults to $HOME/Homefeed)")
	flag.Parse()

	workspaceRoot := *root
	if workspaceRoot == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("resolve home directory: %v", err)
		}
		workspaceRoot = workspace.DefaultWorkspaceRoot(homeDir)
	}

	result, err := workspace.Initialize(workspaceRoot)
	if err != nil {
		log.Fatalf("initialize workspace: %v", err)
	}

	fmt.Printf("workspace: %s\n", result.Layout.Root)
	fmt.Printf("database: %s\n", result.Layout.DatabasePath)
	fmt.Printf("feeds: %d\n", len(result.Feeds))
}

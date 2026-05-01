// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package workspace

import "path/filepath"

// Layout describes the on-disk workspace structure.
type Layout struct {
	Root         string
	MetaRoot     string
	DatabasePath string
	MetaDataPath string
	ProfilesPath string
	ImportsPath  string
	AgentsPath   string
	LogsPath     string
	CachePath    string
}

// DefaultWorkspaceRoot returns the managed Homefeed root under the given home directory.
func DefaultWorkspaceRoot(homeDir string) string {
	return filepath.Join(homeDir, "Homefeed")
}

// LayoutForRoot returns the workspace layout rooted at the provided path.
func LayoutForRoot(root string) Layout {
	metaRoot := filepath.Join(root, ".homefeed")
	return Layout{
		Root:         root,
		MetaRoot:     metaRoot,
		DatabasePath: filepath.Join(metaRoot, "index.sqlite"),
		MetaDataPath: filepath.Join(metaRoot, "meta"),
		ProfilesPath: filepath.Join(metaRoot, "profiles"),
		ImportsPath:  filepath.Join(metaRoot, "imports"),
		AgentsPath:   filepath.Join(metaRoot, "agents"),
		LogsPath:     filepath.Join(metaRoot, "logs"),
		CachePath:    filepath.Join(metaRoot, "cache"),
	}
}

// VisibleFeedPaths returns the standard visible feed directories.
func VisibleFeedPaths(root string) []string {
	paths := make([]string, 0, len(DefaultFeeds))
	for _, feed := range DefaultFeeds {
		paths = append(paths, filepath.Join(root, feed.Slug))
	}
	return paths
}

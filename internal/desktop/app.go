// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package desktop

import (
	"context"

	"homefeed/internal/imports"
	"homefeed/internal/workspace"
)

// App is the thin bridge exposed to the desktop shell.
type App struct {
	ctx           context.Context
	workspaceRoot string
}

// NewApp creates a desktop bridge configured for a workspace root.
func NewApp(workspaceRoot string) *App {
	return &App{
		workspaceRoot: workspaceRoot,
	}
}

// Startup captures the Wails runtime context.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Shutdown is reserved for future runtime cleanup.
func (a *App) Shutdown(ctx context.Context) {}

// InitializeWorkspace runs the existing workspace initializer.
func (a *App) InitializeWorkspace() (workspace.Result, error) {
	return workspace.Initialize(a.workspaceRoot)
}

// ImportFolder copies a local folder into the workspace and records the batch.
func (a *App) ImportFolder(sourcePath, feedSlug string) (imports.Result, error) {
	service := imports.NewService(workspace.LayoutForRoot(a.workspaceRoot))
	return service.ImportFolder(sourcePath, feedSlug)
}

// WorkspaceRoot returns the configured workspace root.
func (a *App) WorkspaceRoot() string {
	return a.workspaceRoot
}

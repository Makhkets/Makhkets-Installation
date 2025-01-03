package main

import (
	"MakhketsDesktop/backend/user"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
	user.User
}

var Urls = map[string]string{
	"nekoray": "https://github.com/MatsuriDayo/nekoray/releases/download/4.0-beta4/nekoray-4.0-beta4-2024-10-09-windows64.zip",
	"goland":  "https://download.jetbrains.com/go/goland-2024.3.1.exe",
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dest := "" // temp folder for downloaded files

	// Register event listener for "updateProgress"
	runtime.EventsOn(ctx, "updateProgress", func(optionalData ...interface{}) {
		if len(optionalData) == 0 {
			fmt.Println("No applications specified.")
			return
		}

		tempDir, err := os.MkdirTemp("", "install_apps_*")
		if err != nil {
			panic(err)
		}

		overallProgressChan := make(chan int)
		defer close(overallProgressChan)

		go func() {
			for progress := range overallProgressChan {
				jsCode := fmt.Sprintf("window.updateProgress(%d);", progress)
				runtime.WindowExecJS(ctx, jsCode)
			}
		}()

		totalFiles := len(optionalData)
		for i, appName := range optionalData[0].([]string) {
			url, exists := Urls[appName]
			if !exists {
				fmt.Printf("Application '%s' not found in URLs. Skipping...\n", appName)
				continue
			}

			dest = filepath.Join(tempDir, filepath.Base(url))
			fmt.Printf("Downloading %s to %s\n", appName, dest)
			if err := user.DownloadFile(ctx, url, dest, i+1, totalFiles, overallProgressChan); err != nil {
				fmt.Printf("Failed to download %s: %v\n", appName, err)
				continue
			}

			fmt.Printf("Successfully downloaded %s to %s\n", appName, dest)
		}

		jsCode := "window.updateProgress(100);"
		runtime.WindowExecJS(ctx, jsCode)
		fmt.Println("All applications processed.")

		fmt.Println(tempDir)

		cmd := exec.Command("explorer", tempDir)
		cmd.Run()

		return
	})
}

// InstallApplications устанавливает выбранные приложения
func (a *App) InstallApplications(applications []string) error {
	fmt.Println("install applications", applications)
	runtime.EventsEmit(a.ctx, "updateProgress", applications)
	return nil
}

func (a *App) Minimize() {
	fmt.Println("minimalize")
	runtime.WindowMinimise(a.ctx)
}

func (a *App) Close() {
	fmt.Println("close")
	runtime.Quit(a.ctx)
}

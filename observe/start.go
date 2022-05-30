package observe

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"WebHook.net/config/public"
	"WebHook.net/global"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mPath"
)

func Start() {
	mCycle.New(mCycle.Opt{
		Func:      SetShell,
		SleepTime: time.Minute * 2,
	}).Start()
}

func SetShell() {
	// 在这里读取文件列表,并打印 shell 文件
	isShellPath := mPath.Exists(global.UserEnv.ShellPath)

	if !isShellPath {
		errStr := fmt.Errorf("配置文件不存在")
		global.LogErr(errStr)
		panic(errStr)
	}

	var files []string
	root := global.UserEnv.ShellPath

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for key, file := range files {
		if mPath.IsFile(file) {

			fullPath, _ := filepath.Abs(file)

			SObj := public.ShellType{
				ID:   key,
				Name: file,
				Path: fullPath,
			}

			public.ShellFile = append(public.ShellFile, SObj)
		}
	}
}

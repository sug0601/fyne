package launcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// FIXME: ここにマイクロサービスのディレクトリを入れる(dirname: relativePath)
var Projects = map[string]string{}

func LaunchProjects() {
	if runtime.GOOS != "darwin" {
		log.Printf("このスクリプトは macOS 専用です (現在のOS: %s)", runtime.GOOS)
		return
	}

	branch := "staging"
	fmt.Printf("各プロジェクトのブランチ '%s' をチェックアウトして yarn dev を起動します...\n", branch)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("カレントディレクトリの取得に失敗しました: %v", err)
		return
	}

	if len(Projects) == 0 {
		fmt.Println("警告: 起動対象のプロジェクトが設定されていません。Projects マップを更新してください。")
		return
	}

	for name, relativePath := range Projects {
		fmt.Printf("\n--- 起動: %s ---\n", name)

		projectPath := filepath.Join(currentDir, relativePath)

		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			log.Printf("スキップ: ディレクトリが存在しません: %s", projectPath)
			continue
		}

		var command string
		if name == "" {
			command = fmt.Sprintf("cd '%s'; git fetch; git checkout -f %s; git pull origin %s; sleep 10; yarn dev", projectPath, branch, branch)
		} else {
			command = fmt.Sprintf("cd '%s'; git fetch; git checkout -f %s; git pull origin %s; yarn dev", projectPath, branch, branch)
		}

		appleScript := fmt.Sprintf(`tell application "iTerm"
	activate
	tell the first window
		set newTab to (create tab with default profile)
		tell current session of newTab to write text "%s"
	end tell
end tell`, command)

		cmd := exec.Command("osascript", "-e", appleScript)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println("iTerm を開いてコマンドを送信中...")
		// コマンド実行
		if err := cmd.Run(); err != nil {
			log.Printf("失敗: %s の起動に失敗しました: %v", name, err)
		} else {
			fmt.Printf("成功: %s を起動しました。\n", name)
		}
	}

	fmt.Println("\n✅ 全ての起動処理が完了しました。")
}

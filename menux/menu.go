// Package menux provides an interactive top-level `menu` subcommand that
// wraps the two RPC actions (generate service, add dependency) in-process,
// replacing the exec.Command-based goctl-gfmenu wrapper.
//
// ADR 筆記：
//   - 生成功能不能省略、不能交給 makefile：makefile 的 rpc target 本身是範本產物，
//     首次生成時它還不存在。這個選單是首次生成唯一的入口。
//   - 只設 VarStringHome，不要改設 VarStringRemote：上游 gen 路徑對
//     util.CloneIntoGitHome(...) 的錯誤是吞掉並靜默改用內建範本，會產出一個形似
//     但並非本組織的骨架（add-dep 那條路徑則會 return err，兩條路徑不對稱）。
//     這裡自行 clone、檢查錯誤、傳 --home 就是為了繞開這個。
package menux

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/gioco-play/goctl-gfrpc/rpcx/cli"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

const templateRepo = "https://github.com/gioco-play/gf-template"

const (
	choiceGen    = "產生 RPC 服務"
	choiceAddDep = "RPC 加依賴"
)

// Cmd describes the interactive menu command. It takes no flags: everything
// it needs is discovered from the current directory or asked interactively.
var Cmd = &cobra.Command{
	Use: "menu",
	// 走錯目錄之類的預期錯誤不需要 usage；main() 已負責印訊息。
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Interactively generate an RPC service or inject a dependency",
	RunE:          run,
}

func run(_ *cobra.Command, _ []string) error {
	files, _ := filepath.Glob("*.proto")
	sort.Strings(files)
	if len(files) == 0 {
		return errors.New("找不到 *.proto，請在服務的 rpc/ 目錄下執行")
	}

	genLabel := choiceGen
	if len(files) == 1 {
		genLabel = fmt.Sprintf("%s    %s", choiceGen, files[0])
	}

	var action string
	if err := ask(survey.AskOne(&survey.Select{
		Message: "請選擇動作：",
		Options: []string{genLabel, choiceAddDep},
	}, &action)); err != nil || action == "" {
		return err
	}

	var filename string
	if action == genLabel {
		filename = files[0]
		if len(files) > 1 {
			if err := ask(survey.AskOne(&survey.Select{
				Message: "偵測到多個候選檔，請選擇：",
				Options: files,
			}, &filename)); err != nil || filename == "" {
				return err
			}
		}
	}

	tplPath, err := util.CloneIntoGitHome(templateRepo, "")
	if err != nil {
		return err
	}

	var deps []string
	if action == choiceAddDep {
		matches, _ := filepath.Glob(filepath.Join(tplPath, "deps", "*.tpl"))
		names := make([]string, len(matches))
		for i, m := range matches {
			names[i] = strings.TrimSuffix(filepath.Base(m), ".tpl")
		}
		sort.Strings(names)

		if err := ask(survey.AskOne(&survey.MultiSelect{
			Message: "請選擇要加入的依賴：",
			Options: names,
		}, &deps)); err != nil {
			return err
		}
		if len(deps) == 0 {
			return nil
		}
	}

	// 這裡印的是實際展開的值（真實 --home 路徑、真實 --name 清單），不是漂亮化的
	// --remote 形式：全域變數設錯名字或漏設不會編譯失敗只會靜默跑錯，這條命令是
	// 唯一的肉眼防線，故意不美化。
	var equivalent string
	if action == genLabel {
		equivalent = fmt.Sprintf("goctl-gfrpc rpc protoc %s --zrpc_out=. --go-grpc_out=. --go_out=. --home %s --style gozero",
			filename, tplPath)
	} else {
		equivalent = fmt.Sprintf("goctl-gfrpc rpc add-dep --name=%s --home %s",
			strings.Join(deps, ","), tplPath)
	}
	fmt.Printf("\n將執行：\n\n  %s\n\n", equivalent)

	confirm := false
	if err := ask(survey.AskOne(&survey.Confirm{Message: "執行？", Default: false}, &confirm)); err != nil {
		return err
	}
	if !confirm {
		return nil
	}

	if action == genLabel {
		cli.VarStringSliceGoOut = []string{"."}
		cli.VarStringSliceGoGRPCOut = []string{"."}
		cli.VarStringZRPCOut = "."
		cli.VarStringHome = tplPath
		cli.VarStringStyle = "gozero"
		return cli.ZRPC(nil, []string{filename})
	}

	cli.VarStringName = strings.Join(deps, ",")
	cli.VarStringHome = tplPath
	return cli.AddDep(nil, nil)
}

// ask turns a Ctrl-C during a survey prompt into a clean nil error (user
// cancelled, not a failure), letting cobra's RunE and gengf.go's main()
// handle the exit path normally for any other error.
func ask(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, terminal.InterruptErr) {
		return nil
	}
	return err
}

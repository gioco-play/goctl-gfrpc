package generator

import (
	_ "embed"
	"fmt"
	"github.com/gioco-play/goctl-gfrpc/rpcx/parser"
	"path/filepath"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed svc.tpl
var svcTemplate string

// GenSvc generates the servicecontext.go file, which is the resource dependency of a service,
// such as rpc dependency, model dependency, etc.
func (g *Generator) GenSvc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetSvc()
	svcFilename, err := format.FileNamingFormat(cfg.NamingFormat, "service_context")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, svcFilename+".go")
	text, err := pathx.LoadTemplate(category, svcTemplateFile, svcTemplate)
	if err != nil {
		return err
	}

	configImport := fmt.Sprintf("\n\t\"%s\"", "fmt")
	configImport += fmt.Sprintf("\n\t\"%s\"", "github.com/gioco-play/go-driver/postgrez")
	configImport += fmt.Sprintf("\n\t\"%s\"", "github.com/go-redis/redis/v8")
	configImport += fmt.Sprintf("\n\t\"%s\"", "gorm.io/gorm")
	configImport += fmt.Sprintf("\n\t\"%s\"", "strings")
	configImport += fmt.Sprintf("\n\t\"%s\"", "github.com/gioco-play/go-driver/logrusz")

	return util.With("svc").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports":      fmt.Sprintf(`"%v"`, ctx.GetConfig().Package),
		"configImport": configImport,
	}, fileName, false)
}

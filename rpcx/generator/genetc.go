package generator

import (
	_ "embed"
	"fmt"
	"github.com/gioco-play/goctl-gfrpc/rpcx/parser"
	"path/filepath"
	"strings"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

//go:embed etc.tpl
var etcTemplate string

//go:embed env.tpl
var envTemplate string

// GenEtc generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEtc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetEtc()
	etcFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, fmt.Sprintf("%v.yaml", etcFilename))

	text, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	return util.With("etc").Parse(text).SaveTo(map[string]interface{}{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}

// GenEnv generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEnv(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetEtc()
	fileName := filepath.Join(dir.Filename, fmt.Sprintf(".env"))

	text, err := pathx.LoadTemplate(category, envTemplateFileFile, envTemplate)
	if err != nil {
		return err
	}

	return util.With("env").Parse(text).SaveTo(map[string]interface{}{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
		"host":        "0.0.0.0:9001",
	}, fileName, false)
}

// GenMakefile generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenMakefile(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	fileName := filepath.Join(ctx.GetMain().Filename, "makefile")

	text, err := pathx.LoadTemplate(category, makefileTemplateFile, makefileTemplate)
	if err != nil {
		return err
	}

	return util.With("env").Parse(text).SaveTo(map[string]interface{}{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}

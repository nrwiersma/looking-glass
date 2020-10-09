// Code generated by 'github.com/traefik/yaegi/extract github.com/nrwiersma/scape/module/types'. DO NOT EDIT.

package types

import (
	"github.com/glasslabs/looking-glass/module/types"
	"reflect"
)

func init() {
	Symbols["github.com/glasslabs/looking-glass/module/types"] = map[string]reflect.Value{
		// type definitions
		"Info":   reflect.ValueOf((*types.Info)(nil)),
		"Logger": reflect.ValueOf((*types.Logger)(nil)),
		"UI":     reflect.ValueOf((*types.UI)(nil)),

		// interface wrapper definitions
		"_Logger": reflect.ValueOf((*_github_com_nrwiersma_scape_module_types_Logger)(nil)),
		"_UI":     reflect.ValueOf((*_github_com_nrwiersma_scape_module_types_UI)(nil)),
	}
}

// _github_com_nrwiersma_scape_module_types_Logger is an interface wrapper for Logger type
type _github_com_nrwiersma_scape_module_types_Logger struct {
	WError func(msg string, ctx []interface{})
	WInfo  func(msg string, ctx []interface{})
}

func (W _github_com_nrwiersma_scape_module_types_Logger) Error(msg string, ctx []interface{}) {
	W.WError(msg, ctx)
}
func (W _github_com_nrwiersma_scape_module_types_Logger) Info(msg string, ctx []interface{}) {
	W.WInfo(msg, ctx)
}

// _github_com_nrwiersma_scape_module_types_UI is an interface wrapper for UI type
type _github_com_nrwiersma_scape_module_types_UI struct {
	WBind     func(name string, fun interface{}) error
	WEval     func(cmd string, ctx []interface{}) (interface{}, error)
	WLoadCSS  func(css string) error
	WLoadHTML func(html string) error
}

func (W _github_com_nrwiersma_scape_module_types_UI) Bind(name string, fun interface{}) error {
	return W.WBind(name, fun)
}
func (W _github_com_nrwiersma_scape_module_types_UI) Eval(cmd string, ctx []interface{}) (interface{}, error) {
	return W.WEval(cmd, ctx)
}
func (W _github_com_nrwiersma_scape_module_types_UI) LoadCSS(css string) error {
	return W.WLoadCSS(css)
}
func (W _github_com_nrwiersma_scape_module_types_UI) LoadHTML(html string) error {
	return W.WLoadHTML(html)
}

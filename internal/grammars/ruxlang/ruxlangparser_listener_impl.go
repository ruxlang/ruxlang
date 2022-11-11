// Package ruxlang provides a lexer and compiler for the Rux language
package ruxlang

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/ruxlang/ruxlang/internal/project"
)

// RuxlangParserListenerImpl is the main listener implmentation for a parse tree produced by RuxlangParser.
type RuxlangParserListenerImpl struct {
	BaseRuxlangParserListener
	PackageName   string
	GlobalImports []string
	NamedImports  map[string]string
}

var _ RuxlangParserListener = &RuxlangParserListenerImpl{}

// NewRuxlangParserListener creates a new listener with a default package name.
func NewRuxlangParserListener(pathName project.PathName) *RuxlangParserListenerImpl {
	return &RuxlangParserListenerImpl{
		PackageName:   pathName.ToPackageName(),
		GlobalImports: []string{},
		NamedImports:  map[string]string{},
	}
}

// ExitPackageClause is called when exiting the packageClause production.
func (l *RuxlangParserListenerImpl) ExitPackageClause(c *PackageClauseContext) {
	packageName := c.IDENTIFIER()
	l.PackageName = packageName.GetText()
}

// ExitImportDecl is called when production importDecl is exited.
func (l *RuxlangParserListenerImpl) ExitImportDecl(ctx *ImportDeclContext) {
	for _, importSpec := range ctx.AllImportSpec() {
		children := importSpec.GetChildren()
		nameToken1, ok := getItem[antlr.Token](children, 0)
		if !ok {
			ctx.AddErrorNode(importSpec.GetStart())
			continue
		}

		// May or may not have a second and third token ( '=' ( IDENTIFIER | PACKAGE_IDENTIFIER ) )
		assignToken, ok := getItem[antlr.Token](children, 1)
		if !ok {
			if _, ok := l.NamedImports[nameToken1.GetText()]; ok {
				ctx.AddErrorNode(nameToken1)
				continue
			}

			l.addImport(ctx, nameToken1, nameToken1)
			continue
		}

		nameToken2, ok := getItem[antlr.Token](children, 2)
		if !ok {
			ctx.AddErrorNode(assignToken)
			continue
		}

		l.addImport(ctx, nameToken1, nameToken2)
	}
}

func (l *RuxlangParserListenerImpl) addImport(ctx antlr.ParserRuleContext, key antlr.Token, value antlr.Token) {
	keyText := key.GetText()

	// If the local name is a '.', all symbols are imported directly and
	// added as an array as there may be many '.' imports
	if keyText == "." {
		l.GlobalImports = append(l.GlobalImports, value.GetText())
		return
	}

	// If there is a '.' in the key name that is not the first character,
	// grab the final part of the name for the import key (local identifier)
	if index := strings.LastIndex(keyText, "."); index >= 1 {
		keyText = keyText[index+1:]
	}

	// If this key already exists, a name override is needed so this is an error
	if _, ok := l.NamedImports[keyText]; ok {
		ctx.AddErrorNode(key)
		return
	}

	// Assign the local name and package identifier
	l.NamedImports[keyText] = value.GetText()
}

func getItem[T any](children []antlr.Tree, i int) (T, bool) {
	var empty T
	if len(children)-1 < i {
		return empty, false
	}

	result, ok := children[i].GetPayload().(T)
	if !ok {
		return empty, false
	}

	return result, true
}

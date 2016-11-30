package ruxlang

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/ruxlang/ruxlang/internal/project"
)

// Parse parses the given file and returns the parsed result.
func Parse(path string) (*RuxlangParserListenerImpl, error) {
	file, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, err
	}

	return ParseStream(file, project.PathName(path))
}

// ParseString parses the given string and returns the parsed result.
func ParseString(data string, pathName project.PathName) (*RuxlangParserListenerImpl, error) {
	return ParseStream(antlr.NewInputStream(data), pathName)
}

// ParseStream parses the given I/O stream and returns the parsed result.
func ParseStream(stream antlr.CharStream, pathName project.PathName) (*RuxlangParserListenerImpl, error) {
	lexer := NewRuxlangLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	ruxparser := NewRuxlangParser(tokens)
	listener := NewRuxlangParserListener(pathName)
	antlr.ParseTreeWalkerDefault.Walk(listener, ruxparser.SourceFile())

	return listener, nil
}

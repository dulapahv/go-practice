// Code generated from Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'nyya'", "'nyyya'", "'nyyyya'", "'nyyyyya'",
	}
	staticData.SymbolicNames = []string{
		"", "ADD", "SUB", "MUL", "DIV", "NUMBER", "WS",
	}
	staticData.RuleNames = []string{
		"ADD", "SUB", "MUL", "DIV", "NUMBER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 60, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 44, 8, 4,
		10, 4, 12, 4, 47, 9, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4, 1, 5, 4, 5, 55,
		8, 5, 11, 5, 12, 5, 56, 1, 5, 1, 5, 0, 0, 6, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 62, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 1, 13, 1, 0, 0, 0, 3, 18, 1, 0, 0, 0, 5, 24, 1, 0, 0, 0,
		7, 31, 1, 0, 0, 0, 9, 51, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 14, 5, 110,
		0, 0, 14, 15, 5, 121, 0, 0, 15, 16, 5, 121, 0, 0, 16, 17, 5, 97, 0, 0,
		17, 2, 1, 0, 0, 0, 18, 19, 5, 110, 0, 0, 19, 20, 5, 121, 0, 0, 20, 21,
		5, 121, 0, 0, 21, 22, 5, 121, 0, 0, 22, 23, 5, 97, 0, 0, 23, 4, 1, 0, 0,
		0, 24, 25, 5, 110, 0, 0, 25, 26, 5, 121, 0, 0, 26, 27, 5, 121, 0, 0, 27,
		28, 5, 121, 0, 0, 28, 29, 5, 121, 0, 0, 29, 30, 5, 97, 0, 0, 30, 6, 1,
		0, 0, 0, 31, 32, 5, 110, 0, 0, 32, 33, 5, 121, 0, 0, 33, 34, 5, 121, 0,
		0, 34, 35, 5, 121, 0, 0, 35, 36, 5, 121, 0, 0, 36, 37, 5, 121, 0, 0, 37,
		38, 5, 97, 0, 0, 38, 8, 1, 0, 0, 0, 39, 40, 5, 110, 0, 0, 40, 41, 5, 121,
		0, 0, 41, 45, 1, 0, 0, 0, 42, 44, 5, 97, 0, 0, 43, 42, 1, 0, 0, 0, 44,
		47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 52, 1, 0, 0,
		0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 110, 0, 0, 49, 50, 5, 121, 0, 0, 50,
		52, 5, 117, 0, 0, 51, 39, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0, 52, 10, 1, 0,
		0, 0, 53, 55, 7, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54,
		1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 6, 5, 0, 0,
		59, 12, 1, 0, 0, 0, 4, 0, 45, 51, 56, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerADD    = 1
	CalcLexerSUB    = 2
	CalcLexerMUL    = 3
	CalcLexerDIV    = 4
	CalcLexerNUMBER = 5
	CalcLexerWS     = 6
)

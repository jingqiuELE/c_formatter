package c_lexer

import "github.com/jingqiuELE/go-lexer"

const (
	TokenPreprocess          lexer.TokenType = iota //Preprocessor tokens: if, ifdef, ifndef, elif, else, endif, defined
	TokenInclude                                    //include
	TokenLeftSquare                                 //'['
	TokenRightSquare                                //']'
	TokenLeftParen                                  //'('
	TokenRightParen                                 //')'
	TokenLeftBrace                                  //'{'
	TokenRightBrace                                 //'}'
	TokenPeriod                                     //'.'
	TokenAmp                                        //'&'
	TokenAmpAmp                                     //'&&'
	TokenAmpEqual                                   //'&='
	TokenStar                                       //'*'
	TokenStarEqual                                  //'*='
	TokenPlus                                       //'+'
	TokenPlusEqual                                  //'+='
	TokenMinus                                      //'-'
	TokenMinusMinus                                 //'--'
	TokenMinusEqual                                 //'-='
	TokenExclaim                                    //'!'
	TokenExclaimEqual                               //'!='
	TokenSlash                                      //'/'
	TokenSlashEqual                                 //'/='
	TokenPercent                                    //'%'
	TokenPercentEqual                               //'%='
	TokenLess                                       //'<'
	TokenLessLess                                   //'<<'
	TokenLessEqual                                  //'<='
	TokenLessLessEqual                              //'<<='
	TokenLessLessEqual                              //'<<='
	TokenSpaceShip                                  //'<=>'
	TokenGreater                                    //'>'
	TokenGreaterGreater                             //'>>'
	TokenGreaterEqual                               //'>='
	TokenGreaterGreaterEqual                        //'>>='
	TokenCaret                                      //'^'
	TokenCaretEqual                                 //'^='
	TokenPipe                                       //'|'
	TokenPipePipe                                   //'||'
	TokenPipeEqual                                  //'|='
	TokenQuestion                                   //'?'
	TokenColon                                      //':'
	TokenSemi                                       //';'
	TokenEqual                                      //'='
	TokenEqualEqual                                 //'=='
	TokenComma                                      //','
	TokenHash                                       //'#'
	TokenHashHash                                   //'##'
	TokenHashHat                                    //'#@'
	TokenAuto                                       //"auto"
	TokenBreak                                      //"break"
	TokenCase                                       //"case"
	TokenChar                                       //"char"
	TokenConst                                      //"const"
	TokenContinue                                   //"continue"
	TokenDefault                                    //"default"
	TokenDo                                         //"do"
	TokenDouble                                     //"double"
	TokenElse                                       //"else"
	TokenEnum                                       //"enum"
	TokenExtern                                     //"extern"
	TokenFloat                                      //"float"
	TokenFor                                        //"for"
	TokenGoto                                       //"goto"
	TokenIf                                         //"if"
	TokenInline                                     //"inline"
	TokenInt                                        //"int"
	TokenLong                                       //"long"
	TokenRegister                                   //"register"
	TokenReturn                                     //"return"
	TokenShort                                      //"short"
	TokenSigned                                     //"signed"
	TokenSizeof                                     //"sizeof"
	TokenStatic                                     //"static"
	TokenStruct                                     //"struct"
	TokenSwitch                                     //"switch"
	TokenTypedef                                    //"typedef"
	TokenUnion                                      //"union"
	TokenUnsigned                                   //"unsigned"
	TokenVoid                                       //"void"
	TokenVolatile                                   //"volatile"
	TokenWhile                                      //"while"
	TokenString
	TokenInteger
	TokenIdentifier
	TokenSlashComment
	TokenSpace
)

func NormalState(l *lexer.L) lexer.StateFunc {
	var n rune
	for {
		n = l.Peek()
		switch l.Peek() {
		case '/':
			l.Emit(TokenWord)
			return SlashState
		case '"':
			return StringState
		case '=':
			l.Emit(TokenWord)
			return EqualState
		case '-':
			l.Emit(TokenWord)
			return MinusState
		case '+':
			l.Emit(TokenWord)
			return AddState
		case '*':
			l.Emit(TokenWord)
			return MultiplyState
		case '{':
			l.Emit(TokenWord)
			l.Next()
			l.Emit(TokenLeftBrace)
		case '}':
			l.Emit(TokenWord)
			l.Next()
			l.Emit(TokenRightBrace)
		case ' ':
			l.Emit(TokenWord)
			l.Take(" \t")
			l.Ignore()
		default:
			l.Next()
		}
	}
}

func SlashState(l *lexer.L) lexer.StateFunc {
	l.Next()
	switch l.Peek() {
	case '*':
		l.Next()
		l.Ignore()
		return SlashCommentState
	default:
		l.Emit(TokenSlash)
		return NormalState
	}
}

func EqualState(l *lexer.L) lexer.StateFunc {
	l.Next()
	switch l.Peek() {
	case '=':
		l.Next()
		l.Emit(TokenIsEqual)
	default:
		l.Emit(TokenEqual)
	}
	return NormalState
}

func SlashCommentState(l *lexer.L) lexer.StateFunc {
	var n rune
	for {
		n = l.Peek()
		switch n {
		case '*':
			l.Next()
			if l.Peek() == '/' {
				l.Rewind()
				l.Emit(TokenSlashComment)
				l.Next()
				l.Next()
				l.Ignore()
				return NormalState
			}
		default:
			l.Next()
		}
	}
}

func StringState(l *lexer.L) lexer.StateFunc {
	var n rune
	l.Next()   // eat starting "
	l.Ignore() // drop current value
L:
	for {
		n = l.Peek()
		switch n {
		case '\\':
			//escaped character.
			l.Next()
			l.Next()
		case '"':
			break L
		default:
			l.Next()
		}
	}
	l.Emit(TokenString)
	l.Next()
	l.Ignore()
	return NormalState
}

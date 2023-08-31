package engine

// import (
// 	"fmt"
// 	"strings"
// )

// type Token int

// func (t Token) ToInt() int { return int(t) }

// const (
// 	KEYWORD Token = iota
// 	OPERATOR
// 	IDENTIFIER
// 	LITERAL
// 	EOF
// )

// func Lex(input string) []Token {
// 	var tokens []Token
// 	var state int

// 	for _, c := range input {
// 		cb := byte(c)
// 		switch state {
// 		case 0:
// 			if isKeyword(cb) {
// 				state = KEYWORD.ToInt()
// 			} else if isOperator(cb) {
// 				state = OPERATOR.ToInt()
// 			} else if isIdentifierStart(cb) {
// 				state = IDENTIFIER.ToInt()
// 			} else if isLiteralStart(cb) {
// 				state = LITERAL.ToInt()
// 			} else {
// 				fmt.Errorf("Unexpected character: %c", c)
// 			}
// 		case KEYWORD.ToInt():
// 			if !isKeyword(cb) {
// 				tokens = append(tokens, KEYWORD)
// 				state = 0
// 			}
// 		case OPERATOR.ToInt():
// 			if !isOperator(cb) {
// 				tokens = append(tokens, OPERATOR)
// 				state = 0
// 			}
// 		case IDENTIFIER.ToInt():
// 			if isIdentifierChar(cb) {
// 			} else {
// 				tokens = append(tokens, IDENTIFIER)
// 				state = 0
// 			}
// 		case LITERAL.ToInt():
// 			if isLiteralChar(cb) {
// 			} else {
// 				tokens = append(tokens, LITERAL)
// 				state = 0
// 			}
// 		}
// 	}

// 	if state != 0 {
// 		fmt.Errorf("Unterminated token")
// 	}

// 	return tokens

// }

// func isKeyword(c byte) bool {
// 	return strings.ContainsRune("god-power creatures type private public def-func main tode zoia-ae destroy graceful-end void-zone exit", rune(c))
// }

// func isOperator(c byte) bool {
// 	return strings.ContainsRune("() <- -> |", rune(c))
// }

// func isIdentifierStart(c byte) bool {
// 	return c >= 'a' && c <= 'z' ||
// 		c >= 'A' && c <= 'Z' ||
// 		c == '_'
// }

// func isIdentifierChar(c byte) bool { return isIdentifierStart(c) || c >= '0' && c <= '9' }
// func isLiteralStart(c byte) bool   { return c == '"' }
// func isLiteralChar(c byte) bool    { return c != '"' }

// // func Lex(input string) []Token {
// // 	var tokens []Token

// // 	regex := regexp.MustCompile(`\s*([+*/()/|#]|((<-|->)[^-])|(int|string)|(god-power[^(])|(tode[^(])|(destroy[^(])|(creatures[^(])|(private[^(])|(public[^(])|(def-func[^(])|(graceful-end[^(])|(zoia-ae[^(])|(main[^(])|(exit[^(])|\d+)`)

// // 	matches := regex.FindAllStringSubmatch(input, -1)
// // 	for _, match := range matches {
// // 		tokenValue := strings.TrimSpace(match[1])
// // 		if tokenValue == "" {
// // 			continue
// // 		}

// // 		switch {
// // 		case tokenValue == "+", tokenValue == "-", tokenValue == "*", tokenValue == "/":
// // 			tokens = append(tokens, token{TokenOperator, tokenValue})
// // 		case tokenValue == "(":
// // 			tokens = append(tokens, token{TokenLeftParen, tokenValue})
// // 		case tokenValue == ")":
// // 			tokens = append(tokens, token{TokenRightParen, tokenValue})
// // 		case tokenValue == "god-power":
// // 			tokens = append(tokens, token{TokenGodPower, tokenValue})
// // 		case regexp.MustCompile(`\d+`).MatchString(tokenValue):
// // 			tokens = append(tokens, token{TokenNumber, tokenValue})
// // 			// default:
// // 			// 	tokens = append(tokens, token{TokenError, tokenValue})
// // 		}
// // 	}

// // 	tokens = append(tokens, token{TokenEOF, ""})
// // 	return tokens
// // }

// // func main() {
// // 	fmt.Println(Lex(`(god-power
// // 		(creatures
// // 			(type [diniz]
// // 				(private
// // 					int#idade 	   |
// // 					string#feiura  |
// // 				) |
// // 				(public
// // 					bool#feiura2 	|
// // 					bool#feiura3	|
// // 				)
// // 			)
// // 		) | (tode exitCode <- (void-zone
// // 				(def-func newFunc ()
// // 					(tode variable1 -> (add 1 1)) |
// // 					(tode variable2 -> (add 2 2)) |
// // 					<- (add variable1 variable2)  |
// // 				) | (def-func newFunc2 (int#var1, int#var2)
// // 					<- (add var1 var2) |
// // 				)
// // 			) | (main
// // 					(tode variable3 -> (newFunc)) |
// // 					(tode variable4-> (newFunc2 1 1)) |
// // 					(zoia-ae variable3) |
// // 					(zoia-ae variable4) |
// // 			) | (graceful-end
// // 				(destroy variable1)  |
// // 				(destroy variable2)  |
// // 				(destroy variable3)	 |
// // 				(destroy variable4)  |
// // 				<- 1
// // 			)
// // 		) | (exit exitCode)
// // )`))
// // }

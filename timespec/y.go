//line lang.y:2
package timespec

import __yyfmt__ "fmt"

//line lang.y:2
import (
	"time"
)

//line lang.y:9
type yySymType struct {
	yys    int
	numval uint
	time   int
	wday   time.Weekday
	spec   *Spec
	truth  bool
}

const NUMBER = 57346
const ORDINAL = 57347
const HOURLY = 57348
const DAILY = 57349
const WEEKLY = 57350
const MONTHLY = 57351
const FROM = 57352
const AT = 57353
const ON = 57354
const AM = 57355
const PM = 57356
const HALF = 57357
const EVERY = 57358
const DAY = 57359
const HOUR = 57360
const QUARTER = 57361
const AFTER = 57362
const TIL = 57363
const SUNDAY = 57364
const MONDAY = 57365
const TUESDAY = 57366
const WEDNESDAY = 57367
const THURSDAY = 57368
const FRIDAY = 57369
const SATURDAY = 57370

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ORDINAL",
	"HOURLY",
	"DAILY",
	"WEEKLY",
	"MONTHLY",
	"FROM",
	"AT",
	"ON",
	"AM",
	"PM",
	"HALF",
	"EVERY",
	"DAY",
	"HOUR",
	"QUARTER",
	"AFTER",
	"TIL",
	"SUNDAY",
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"'h'",
	"'H'",
	"'x'",
	"'X'",
	"'*'",
	"':'",
	"' '",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line lang.y:128

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 24,
	1, 27,
	-2, 25,
}

const yyNprod = 58
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 152

var yyAct = [...]int{

	61, 39, 71, 11, 24, 49, 22, 50, 51, 24,
	56, 54, 42, 44, 46, 32, 21, 47, 80, 31,
	32, 53, 63, 64, 31, 52, 63, 64, 48, 26,
	27, 28, 29, 30, 26, 27, 28, 29, 30, 58,
	59, 40, 55, 65, 95, 68, 67, 69, 74, 75,
	73, 72, 78, 63, 64, 40, 73, 72, 87, 81,
	40, 79, 57, 83, 70, 77, 24, 45, 40, 85,
	86, 40, 88, 89, 60, 62, 90, 32, 43, 40,
	40, 31, 93, 94, 91, 92, 41, 38, 96, 82,
	97, 26, 27, 28, 29, 30, 98, 13, 7, 9,
	10, 12, 73, 72, 76, 23, 1, 6, 8, 5,
	4, 84, 3, 2, 14, 15, 16, 17, 18, 19,
	20, 14, 15, 16, 17, 18, 19, 20, 66, 14,
	15, 16, 17, 18, 19, 20, 36, 25, 14, 15,
	16, 17, 18, 19, 20, 0, 0, 34, 0, 37,
	35, 33,
}
var yyPact = [...]int{

	92, -1000, -1000, -1000, -1000, -1000, -1000, 5, 132, 76,
	75, 67, 56, 107, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 62, -1000, -29, -1000, -13, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 7, 3, 0, -8, 51, 64, -1000,
	40, 64, 116, 64, -1000, 64, 52, 37, -1000, 100,
	-1000, -1000, 55, 42, 62, -1000, 8, 64, -1000, -1000,
	85, -1000, 13, -1000, -1000, 99, 107, -1000, -1000, 46,
	98, -1000, -1000, -1000, 64, -1000, -1000, 62, 62, -1000,
	64, -1000, 9, -1000, 107, -1000, -1000, 98, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 13, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 6, 1, 2, 137, 3, 113, 112, 110, 109,
	107, 0, 106, 105,
}
var yyR1 = [...]int{

	0, 12, 6, 6, 6, 6, 7, 7, 7, 7,
	7, 7, 7, 8, 8, 8, 8, 13, 13, 13,
	13, 13, 13, 4, 4, 4, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 9, 9, 9, 9, 9,
	9, 11, 11, 5, 5, 5, 5, 5, 5, 5,
	10, 10, 10, 10, 10, 10, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 3, 2, 5, 5,
	4, 3, 5, 3, 2, 4, 3, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 2, 2,
	3, 4, 5, 2, 3, 5, 4, 4, 3, 3,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	5, 4, 4, 3, 4, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -12, -6, -7, -8, -9, -10, 6, 16, 7,
	8, -5, 9, 5, 22, 23, 24, 25, 26, 27,
	28, 11, -1, -13, 4, -4, 29, 30, 31, 32,
	33, 19, 15, 19, 15, 18, 4, 17, 11, -2,
	4, 11, -2, 11, -2, 11, -2, -5, -1, 34,
	20, 21, 18, 18, 11, -1, 18, 11, -2, -2,
	34, -11, 35, 13, 14, -2, 12, -5, -2, -2,
	12, -3, 5, 4, 11, -2, 4, 10, 10, -1,
	10, -2, 4, -11, 12, -5, -5, 12, -3, -3,
	-2, -1, -1, -2, -11, 35, -5, -3, -11,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 17, 0, 0,
	0, 0, 0, 0, 43, 44, 45, 46, 47, 48,
	49, 17, 7, 0, -2, 0, 18, 19, 20, 21,
	22, 23, 24, 0, 0, 17, 0, 0, 0, 14,
	0, 0, 0, 0, 40, 0, 0, 0, 6, 0,
	28, 29, 0, 0, 17, 11, 0, 0, 16, 13,
	0, 33, 0, 41, 42, 0, 0, 38, 39, 0,
	0, 53, 56, 57, 0, 55, 26, 17, 17, 10,
	0, 15, 30, 34, 0, 37, 36, 0, 52, 51,
	54, 8, 9, 12, 31, 0, 35, 50, 32,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 33, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 34, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 30, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 32, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 29, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	31,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:52
		{
			yylex.(*yyLex).spec = yyDollar[1].spec
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:60
		{
			yyVAL.spec = hourly(yyDollar[3].time, 0)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:61
		{
			yyVAL.spec = hourly(yyDollar[2].time, 0)
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:62
		{
			yyVAL.spec = hourly(yyDollar[5].time, 0.25)
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:63
		{
			yyVAL.spec = hourly(yyDollar[5].time, 0.5)
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:64
		{
			yyVAL.spec = hourly(yyDollar[4].time, 0)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:65
		{
			yyVAL.spec = hourly(yyDollar[3].time, 0)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:66
		{
			yyVAL.spec = hourly(yyDollar[5].time, float32(yyDollar[2].numval))
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:69
		{
			yyVAL.spec = daily(yyDollar[3].time)
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:70
		{
			yyVAL.spec = daily(yyDollar[2].time)
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:71
		{
			yyVAL.spec = daily(yyDollar[4].time)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:72
		{
			yyVAL.spec = daily(yyDollar[3].time)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:77
		{
			yyVAL.numval = 15
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:78
		{
			yyVAL.numval = 30
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:79
		{
			yyVAL.numval = yyDollar[1].numval
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:82
		{
			yyVAL.time = hhmm24(0, yyDollar[3].numval)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:83
		{
			yyVAL.time = hhmm24(0, yyDollar[1].numval)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:84
		{
			yyVAL.time = hhmm24(0, yyDollar[1].numval)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:85
		{
			yyVAL.time = hhmm24(0, 60-yyDollar[1].numval)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:88
		{
			yyVAL.time = hhmm24(yyDollar[1].numval, yyDollar[3].numval)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:89
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, yyDollar[3].numval, yyDollar[4].truth)
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:90
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, yyDollar[3].numval, yyDollar[5].truth)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:91
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, 0, yyDollar[2].truth)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:92
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, 0, yyDollar[3].truth)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:95
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[5].wday)
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:96
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[4].wday)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:97
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[4].wday)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:98
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[3].wday)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:99
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[1].wday)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:100
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[1].wday)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:103
		{
			yyVAL.truth = true
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:104
		{
			yyVAL.truth = false
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:107
		{
			yyVAL.wday = time.Sunday
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:108
		{
			yyVAL.wday = time.Monday
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:109
		{
			yyVAL.wday = time.Tuesday
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:110
		{
			yyVAL.wday = time.Wednesday
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:111
		{
			yyVAL.wday = time.Thursday
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:112
		{
			yyVAL.wday = time.Friday
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:113
		{
			yyVAL.wday = time.Saturday
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:116
		{
			yyVAL.spec = mday(yyDollar[3].time, yyDollar[5].numval)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:117
		{
			yyVAL.spec = mday(yyDollar[2].time, yyDollar[4].numval)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:118
		{
			yyVAL.spec = mday(yyDollar[3].time, yyDollar[4].numval)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:119
		{
			yyVAL.spec = mday(yyDollar[2].time, yyDollar[3].numval)
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:120
		{
			yyVAL.spec = mweek(yyDollar[4].time, yyDollar[2].wday, yyDollar[1].numval)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:121
		{
			yyVAL.spec = mweek(yyDollar[3].time, yyDollar[2].wday, yyDollar[1].numval)
		}
	}
	goto yystack /* stack new state and value */
}

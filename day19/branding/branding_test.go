package branding_test

import (
	"strings"
	"testing"

	"github.com/brainboxweb/advent-2024/day19/branding"
	"github.com/stretchr/testify/assert"
)

func TestIsPossible(t *testing.T) {
	tests := []struct {
		pattern  string
		expected bool
	}{
		{
			"brwrr",
			true,
		},
		{
			"bggr",
			true,
		},
		{
			"gbbr",
			true,
		},
		{
			"rrbgbr",
			true,
		},
		{
			"ubwu",
			false,
		},
		{
			"bwurrg",
			true,
		},
		{
			"brgr",
			true,
		},
		{
			"bbrgwb",
			false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			towels := branding.NewTowels([]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"})
			cnt := towels.CanObtain(tt.pattern)
			result := cnt > 0
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCanObtain(t *testing.T) {
	tests := []struct {
		towelsString string
		pattern      string
		expected     int
	}{
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"brwrr",
			2,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"bggr",
			1,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"gbbr",
			4,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"rrbgbr",
			6,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"ubwu",
			0,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"bwurrg",
			1,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"brgr",
			2,
		},
		{
			"r, wr, b, g, bwu, rb, gb, br",
			"bbrgwb",
			0,
		},
		{
			"buu, ubg, gub, wbwu, rrgrgg, urrwr, grwb, ubw, ubrbr, ubgb, gbw, rbrwgw, wbrr, brg, urugggr, wgbb, uwwu, urgg, wwwr, wrub, rbr, brgwb, gruwwb, rrwuwbu, rub, urwwgu, bbbw, wrr, uwr, grr, ggwrb, wbrwwbur, brbg, rubrg, urwbbru, brwww, bu, gugwww, rbww, buw, rurwgrug, rugwr, ruubww, bgbu, bgbw, wuu, burw, uuwb, rug, grw, bbrugwr, wurwbg, uurwu, bgu, ruu, bgrg, uguub, bgug, burubb, ubuugb, bbww, grbrr, gubbww, gwbgrwu, ggu, bubgurub, wguu, gww, wuuu, wbgb, bgrw, gug, rwwbw, uru, rwr, ubwr, uwb, bubwbwbw, gggwr, brr, gbrrg, rrww, gwg, uurubb, ruwwgrw, brgrru, gbgrwb, wbbwgg, wuwr, ruwuwbw, ggwrr, gbu, gwr, uur, wbgu, gbwgbgr, bbb, bubu, rwwggw, ruw, rbwuw, ugbbrwu, uurrgrg, bwwubwwg, gruu, ubgwu, rww, bur, ubwrg, wgb, gurbbr, rgbr, wuw, gbwuwru, rgwbrg, wbbr, wr, bgrrgb, wwb, urgwubww, bwurw, ur, rgbu, rrurw, ubbu, ubur, wrbwwr, uugww, grg, rwb, rgb, gguggu, rwgr, gurbubg, guwwr, wwbgb, grbgwruw, wug, bgb, rb, ruuubggr, www, wbgg, gguu, bbur, gbub, bwg, uuugbw, brb, brrrggu, buwu, rbrg, bg, bwgrbw, bwb, u, guw, gwrwur, brgb, wrru, ruugu, wrgrg, wbrb, rg, buwrr, ugr, uwru, bub, wbb, ubr, brrwr, wgbuu, bguu, bubw, ug, ugwg, wurbww, bwrbuwu, wwuwb, brbubg, wwub, g, brw, bbrwrgb, bgr, rbuwb, gwur, bggg, gbuu, wbu, rubw, bguug, bbw, wuguw, uwbwrgb, wwwub, wwbugg, gwubbb, rugwg, rgwwbru, ugbbubr, wrrgr, bubg, wwbwu, wbw, bbr, wbggugw, wurw, wwgwgr, burgr, wbgbu, wwwuuw, buguu, gwrugg, wwgugg, bburg, urwbu, rrug, ggb, ggbg, rguwbu, guru, wwr, ggurrw, brbbbubg, bbguuwgw, rru, ugwu, wbggr, rgrbww, rrubw, ubwurwg, gbr, rwrgw, wwguwbu, wggrr, wru, wrgbb, bgburwg, wruwwg, uuu, rugr, rwbw, bwwbr, bbbrgwu, br, ubb, bbu, urg, brbuggb, uw, rbug, wbbrgugb, rgu, uubgrrgw, bru, rwgbu, wrgw, gwb, bwbr, bgw, wgwg, wbgrwwg, wwwwgr, gg, gwgggu, ugwuu, gugur, gggbw, bwbwguu, uwbr, gw, wrgg, rugwbgg, uwg, urb, uwggb, urgggb, bbg, gwbr, uub, gbg, gwuu, buwrrw, rbu, wbbu, ru, rbbw, rgugrbwg, ggbgrw, ugrwwww, uwgbwu, wwbw, wurr, wrwrrr, rrw, rbb, rwrgr, bbwwg, rbg, bbbgrbrw, ruubwb, ubbg, gr, gwwbuuuw, bwgr, rrwwggu, bbuugb, guwwbg, rguw, rur, gwbbbbr, rbwr, wb, uuwrbb, uruwbrrr, bb, bbubu, gruguu, urr, bwbrb, rbru, ugu, ugg, rwbgu, rbbuug, uguw, grwbg, bbgrrr, gur, urw, gguwbb, rburuug, w, wg, ggw, wur, ubwuw, bwggugb, gwbu, wbrg, gbrgw, wrugur, rbgb, uwbrrubg, bbugugw, rwbgwur, ubwu, rw, wurwggr, gru, brur, gwbuw, uwu, gurbb, rgbgr, bgg, rggg, rrg, wbbg, wwu, wbr, uww, ggr, wbwgu, rrgu, rgw, rrwugrg, wgug, wrw, rbw, bug, bruuubrb, rubwurr, rwub, bbbru, rr, gwu, ubbgwg, gbww, wwrgrgg, wbg, rwu, ugurw, bubwg, wbubw, urwur, wgu, ubuuug, wubu, rrb, rrru, bwr, rbbubbru, gbwb, ggrruw, bww, uu, wu, wwuww, gwguwbb, ubu, uug, gubgr, rwggb, ubwug, wrg, r, wbwgwu, ww, gggw, rubrw, ugw, wgg, buguurwr, grb, bwur, wgr, uubb, wuwbwrb, urbwr, bw, ggg, bgrgr, buguwr, rbgbg, rrr, wgwbrr, rgg, rububg, bwu, ub, gbb, wub, uuw, wgw, bbub, wbggbb, wwwbgg, rgr, rbbbu, wuwb, grubu, rruwu, gubu, urwu, wrurg",
			"rrbwuuwwuwubwggbgggggwgurgwubwwrbwubgrbu",
			999694350,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			towels := branding.NewTowels(strings.Split(tt.towelsString, ", "))
			result := towels.CanObtain(tt.pattern)
			assert.Equal(t, tt.expected, result)
		})
	}
}

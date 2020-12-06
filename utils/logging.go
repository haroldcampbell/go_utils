package utils

import (
	"fmt"
	"log"
	"time"
)

// Terminal colors
const BlackTextFG = "30"
const RedTextFG = "31"
const GreenTextFG = "32"
const YellowTextFG = "33"
const BlueTextFG = "34"
const PurpleTextFG = "35"
const CyanTextFG = "36"
const WhiteTextFG = "37"

const DarkGrayTextFG = "90"
const GrayTextFG = "97"

const TagText = "\033[1;36m==>\033[0m"

var fgTextColors = map[string]string{
	"Black":  BlackTextFG,
	"Red":    RedTextFG,
	"Green":  GreenTextFG,
	"Yellow": YellowTextFG,
	"Blue":   BlueTextFG,
	"Purple": PurpleTextFG,
	"Cyan":   CyanTextFG,
	"White":  WhiteTextFG,

	"DarkGray": DarkGrayTextFG,
	"Gray":     GrayTextFG,
}

const BlackTextBG = "40"
const RedTextBG = "41"
const GreenTextBG = "42"
const YellowTextBG = "43"
const BlueTextBG = "44"
const PurpleTextBG = "45"
const CyanTextBG = "46"
const WhiteTextBG = "47"

var bgTextColors = map[string]string{
	"Black":  BlackTextBG,
	"Red":    RedTextBG,
	"Green":  GreenTextBG,
	"Yellow": YellowTextBG,
	"Blue":   BlueTextBG,
	"Purple": PurpleTextBG,
	"Cyan":   CyanTextBG,
	"White":  WhiteTextBG}

func ColoredText(fg string, v interface{}) string {
	return fmt.Sprintf("\033[%sm%v\033[0m", fg, v)
}

func ColoredTextBG(fg string, bg string, v interface{}) string {
	return fmt.Sprintf("\033[%s;%s;m%v\033[0m", fg, bg, v)
}

func ColoredBrightText(fg string, v interface{}) string {
	return fmt.Sprintf("\033[1;%sm%v\033[0m", fg, v)
}

func ColoredBrightTextBG(fg string, bg string, v interface{}) string {
	return fmt.Sprintf("\033[1;%s;%sm%v\033[0m", fg, bg, v)
}

func RedText(v interface{}) string {
	return ColoredBrightTextBG(RedTextFG, BlackTextBG, v)
}

func CyanText(v interface{}) string {
	return ColoredBrightText(CyanTextFG, v)
}

func YellowText(v interface{}) string {
	return ColoredBrightTextBG(YellowTextFG, BlackTextBG, v)
}

func Tag(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%s %s", CyanText("==>"), message)
}

func Log(stem string, format string, v ...interface{}) {
	t := time.Now()

	var buf []byte
	formatDateTime(&buf, t)

	str := fmt.Sprintf(format, v...)

	fmt.Printf("%s [%s]  %s\n",
		ColoredText(DarkGrayTextFG, string(buf)),
		ColoredBrightText(WhiteTextFG, string(stem)),
		str)
}

func Error(stem string, format string, v ...interface{}) {
	t := time.Now()

	var buf []byte
	formatDateTime(&buf, t)

	str := fmt.Sprintf(format, v...)

	fmt.Printf("%s [%s]  %s  %s\n",
		ColoredText(DarkGrayTextFG, string(buf)),
		ColoredBrightText(WhiteTextFG, string(stem)),
		RedText("ERROR"),
		str)
}

func ErrorMsg(stem string, format string, v ...interface{}) string {
	t := time.Now()

	var buf []byte
	formatDateTime(&buf, t)

	str := fmt.Sprintf(format, v...)

	return fmt.Sprintf(" >>>> %s [%s]  %s  %s",
		ColoredText(DarkGrayTextFG, string(buf)),
		ColoredBrightText(WhiteTextFG, string(stem)),
		RedText("ERROR"),
		str)
}

// Copied from go's log.go: Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

// Copied from go's log.go:
func formatDateTime(buf *[]byte, t time.Time) {
	year, month, day := t.Date()
	itoa(buf, year, 4)
	*buf = append(*buf, '/')
	itoa(buf, int(month), 2)
	*buf = append(*buf, '/')
	itoa(buf, day, 2)
	*buf = append(*buf, ' ')

	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)
}

func formattedDateTime() []byte {
	t := time.Now()

	var buf []byte
	formatDateTime(&buf, t)

	return buf

}

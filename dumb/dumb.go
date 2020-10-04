package main

import (
	"fmt"
	"os"
	"time"
	"github.com/gdamore/tcell"
)


func makebox1(screen tcell.Screen) {
	w, h := screen.Size()
	if w == 0 || h == 0 {
		return
	}

        dunno := []rune{'y','z'}

	boxstyle := tcell.StyleDefault.
		Foreground(tcell.ColorRed).
		Background(tcell.ColorYellow)

	screen.SetContent(3,3,'x', dunno, boxstyle)
	screen.Show()
}


func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorGreen).
		Background(tcell.ColorBlack))
	screen.Clear()

	makebox1(screen)

	time.Sleep(3*time.Second)

	screen.Fini()
}

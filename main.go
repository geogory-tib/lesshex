package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"lesshex/draw"
	"lesshex/state"
	"lesshex/userin"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("lesshex: No arguments specified")
		return
	}
	var context state.LessHex_Context
	context.Load_files(os.Args[1:len(os.Args)])
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()
	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}
	scrn_width, scrn_height := screen.Size()
	screen.SetStyle(context.Default_style)
	screen.Clear()
	draw.Draw_Context(context, context.Draw_Width, scrn_height, screen)
	draw.Draw_file_bar(context, screen)
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			if tcell.Key(event.Rune()) == ':' {
				if userin.Handle_Colon(&context, screen, scrn_width, scrn_height) {
					return
				}
			} else {
				userin.Handle_Movement_Commands(&context, screen, event.Key(), scrn_height)
			}
		}
	}
}

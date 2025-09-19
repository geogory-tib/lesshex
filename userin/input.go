package userin

import (
	"lesshex/draw"
	"lesshex/state"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

func Handle_Colon(conext *state.LessHex_Context, screen tcell.Screen, max_width, max_height int) bool {
	var user_command string

	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEnter:
				goto exitloop
			case tcell.KeyRune:
				user_command += string(event.Rune())
			}
		}
	}
exitloop:
	return handle_colon_commands(user_command, conext, screen, max_width, max_height)
}

func handle_colon_commands(command string, context *state.LessHex_Context, screen tcell.Screen, max_width, max_height int) bool {
	split_command := strings.Split(command, " ")
	switch split_command[0] {
	case "q":
		return true
	case "s":
		page, err := strconv.Atoi(split_command[1])
		if page <= len(context.Pages) && err == nil {
			context.Active_page = page - 1
			screen.Clear()
			draw.Draw_file_bar(*context, screen)
			draw.Draw_Context(*context, context.Draw_Width, max_height, screen)
		}
	case "w":
		display_width, err := strconv.Atoi(split_command[1])
		if display_width <= max_width && err == nil {
			context.Draw_Width = display_width
			screen.Clear()
			draw.Draw_Context(*context, context.Draw_Width, max_height, screen)
			draw.Draw_file_bar(*context, screen)
		}
	}
	return false
}

func Handle_Movement_Commands(context *state.LessHex_Context, screen tcell.Screen, key tcell.Key, max_height int) {
	switch key {
	case tcell.KeyDown:
		if context.Pages[context.Active_page].Scroll_Offset < len(context.Pages[context.Active_page].Display_buffer) {
			context.Pages[context.Active_page].Scroll_Offset++
		}
	case tcell.KeyUp:
		if context.Pages[context.Active_page].Scroll_Offset > 0 {
			context.Pages[context.Active_page].Scroll_Offset--
		}
	}
	screen.Clear()
	draw.Draw_Context(*context, context.Draw_Width, max_height, screen)
	draw.Draw_file_bar(*context, screen)
}

package userin

import (
	"lesshex/state"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

func Handle_Colon(conext *state.LessHex_Context, screen tcell.Screen, max_width int) bool {
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
	return handle_commands(user_command, conext, screen, max_width)
}

func handle_commands(command string, context *state.LessHex_Context, screen tcell.Screen, max_width int) bool {
	split_command := strings.Split(command, " ")
	switch split_command[0] {
	case "q":
		return true
	case "s":
		page, err := strconv.Atoi(split_command[1])
		if page <= len(context.Pages) && err == nil {
			context.Active_page = page
		}
	case "w":
		display_width, err := strconv.Atoi(split_command[1])
		if display_width <= max_width && err != nil {
			context.Draw_Width = display_width
		}
	}
	return false
}

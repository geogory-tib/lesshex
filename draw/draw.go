package draw

import (
	"lesshex/state"

	"github.com/gdamore/tcell"
)

const MAX_DRAW_WIDTH = 20

func Draw_Context(context state.LessHex_Context, width, height int, screen tcell.Screen) {
	x := 0
	y := 1
	bytes_drawn := 0
	current_page := &context.Pages[context.Active_page]
	for i := current_page.Scroll_Offset; i < len(current_page.Display_buffer); i++ {
		screen.SetContent(x, y, rune(current_page.Display_buffer[i]), nil, tcell.StyleDefault)
		x++
		bytes_drawn++
		if bytes_drawn == 2 {
			screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			bytes_drawn = 0
			x++
		}
		if x > width {
			y++
			x = 0
			bytes_drawn = 0
		}
	}
	screen.Show()
}

func Draw_file_bar(context state.LessHex_Context, screen tcell.Screen) {
	x := 1
	active_file_style := context.Default_style.Foreground(tcell.ColorLightYellow)
	for index, filename := range context.File_names {
		for _, ch := range filename {
			switch index {
			case context.Active_page:
				screen.SetContent(x, 0, ch, nil, active_file_style)
			default:
				screen.SetContent(x, 0, ch, nil, context.Default_style)
			}
			x++
		}
		x++
	}
	screen.Show()
}

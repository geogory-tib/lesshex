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
		if bytes_drawn == 2 {
			screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			bytes_drawn = 0
			x++
		}
		screen.SetContent(x, y, rune(current_page.Display_buffer[i]), nil, tcell.StyleDefault)
		x++
		bytes_drawn++
		if x == width {
			y++
			x = 0
			bytes_drawn = 0
		}
	}
	screen.Show()
}

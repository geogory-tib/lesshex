package state

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
)

type Page struct {
	Raw_file       []byte
	Display_buffer string
	Scroll_Offset  int
	Size           int
}
type LessHex_Context struct {
	Pages       []Page
	File_names  []string
	Active_page int
	Draw_Width  int
}

func (self *LessHex_Context) Load_files(file_names []string) {
	self.File_names = file_names
	var page Page
	for _, filename := range file_names {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(file)
		file_stat, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		temp_buffer := make([]byte, file_stat.Size())
		length, err := reader.Read(temp_buffer)
		page.Size = length
		page.Raw_file = temp_buffer
		page.Display_buffer = hex.EncodeToString(temp_buffer)
		page.Scroll_Offset = 0
		self.Pages = append(self.Pages, page)
		file.Close()
		temp_buffer = nil
		self.Draw_Width = 40
	}
}

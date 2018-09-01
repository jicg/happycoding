package main

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"github.com/nsf/termbox-go"
	"github.com/tosone/minimp3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func main() {
	files, _ := ioutil.ReadDir("mp3")
	lenf := len(files)
	fmt.Println(lenf)
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Key {
		case termbox.KeyEsc:
			break Loop
		case termbox.KeyEnter:
			fmt.Println()
		case termbox.KeySpace:
			fmt.Print(" ")
		case termbox.KeyTab:
			fmt.Print("	")
		default:
			fmt.Printf("%s", string(ev.Ch))
			i := int(ev.Ch) % lenf
			go Sound(files[i].Name())
		}

	}

}

func Sound(filename string) {
	if len(filename) == 0 {
		return
	}
	file, err := ioutil.ReadFile("mp3/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	dec, data1, _ := minimp3.DecodeFull(file)
	player, _ := oto.NewPlayer(
		dec.SampleRate,
		dec.Channels,
		2,
		1024)
	player.Write(data1)
	player.Close()
}

func main_bak() {
	//file ,_:=os.("mp3")
	//file.
	files, _ := ioutil.ReadDir("mp3")
	for _, file := range files {
		if err := os.Rename(
			"mp3/"+file.Name(),
			"mp3/"+strings.Replace(
				file.Name(),
				"__爱给网_aigei_com",
				"",
				-1,
			),
		); err != nil {
			log.Fatal(err)
		}
		fmt.Println(file.Name())
	}
}

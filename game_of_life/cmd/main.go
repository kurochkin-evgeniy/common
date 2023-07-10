// gol play -f=100x500 --iter=100 --load=file.txt --save=file.gif
// gol gen -f=100x500 --save=file.txt
package main

import (
	"bufio"
	"flag"
	"fmt"
	"game_of_life/package/gameoflife"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
	"strconv"
	"strings"
)

func performPlayCommand(args []string) {
	playCmd := flag.NewFlagSet("play", flag.ExitOnError)
	f := playCmd.String("f", "50x50", "size of field")
	fileName := playCmd.String("save", "", "File path to save.")
	loadName := playCmd.String("load", "", "Load path to load.")
	iter := playCmd.Int("iter", 100, "Iterations")
	playCmd.Parse(args)

	var field *gameoflife.GameField = nil
	if *loadName != "" {
		f, error := gameoflife.LoadFieldFromFile(*loadName)
		if error != nil {
			panic(error)
		}
		field = f
	} else {
		w, h := parseFieldSize(*f)
		field = gameoflife.GenRandomField("new", w, h, 5)
	}

	if *fileName == "" {
		play(*iter, field)
	} else {
		playAndRecord(*iter, field, *fileName)
	}
}

func play(n int, field *gameoflife.GameField) {
	for i := 0; i < n; i++ {
		field.Iterate()
	}
}

func addGameGeneration(field *gameoflife.GameField, anim *gif.GIF) {
	const delay = 100
	const size = 20
	var palette = []color.Color{color.White, color.Black}

	w := len(field.Field)
	h := len(field.Field[0])

	rect := image.Rect(0, 0, w*size, h*size)
	img := image.NewPaletted(rect, palette)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if field.Field[x][y][field.CurrentIndex] {
				rect := image.Rect(x*size, y*size, x*size+size, y*size+size)
				draw.Draw(img, rect, &image.Uniform{palette[1]}, image.ZP, draw.Src)
			}
		}
	}

	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)
}

func playAndRecord(n int, field *gameoflife.GameField, fileName string) {
	anim := gif.GIF{LoopCount: n}

	addGameGeneration(field, &anim)
	for i := 0; i < n; i++ {
		field.Iterate()
		addGameGeneration(field, &anim)
	}

	error := saveToFile(&anim, fileName)
	if error != nil {
		panic(error)
	}
}

func saveToFile(image *gif.GIF, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	out := bufio.NewWriter(file)
	err = gif.EncodeAll(out, image)
	if err != nil {
		return err
	}
	out.Flush()

	return nil
}

func parseFieldSize(s string) (int, int) {
	nums := strings.Split(s, "x")
	if len(nums) != 2 {
		panic("error")
	}

	w, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	h, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	return w, h
}

func performGenCommand(args []string) {
	genCmd := flag.NewFlagSet("gen", flag.ExitOnError)
	f := genCmd.String("f", "50x50", "size of field")
	fileName := genCmd.String("save", "", "File path to save.")
	genCmd.Parse(args)

	w, h := parseFieldSize(*f)

	field := gameoflife.GenRandomField(*fileName, w, h, 5)

	if *fileName == "" {
		field.Print()
	} else {
		field.SaveToFile(*fileName)
	}

}

func main() {
	cmd := "play"
	params := os.Args[1:]

	if len(os.Args) >= 2 {
		cmd = os.Args[1]
		params = os.Args[2:]
	}

	switch cmd {
	case "play":
		performPlayCommand(params)
	case "gen":
		performGenCommand(params)
	default:
		fmt.Println("Unknown command.")
		os.Exit(1)
	}
}

package gameoflife

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type GameField struct {
	Name         string
	Field        [][][2]bool
	CurrentIndex int
}

func NewField(name string, w, h int) *GameField {
	result := GameField{}
	result.Name = "new"

	a := make([][][2]bool, w)
	for i := range a {
		a[i] = make([][2]bool, h)
	}

	result.Field = a

	return &result
}

func randBool(dencity int) bool {
	return rand.Intn(100) < dencity
}

func GenRandomField(name string, w, h int, dencity int) *GameField {
	result := NewField(name, w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			result.Field[i][j][0] = randBool(dencity)
		}
	}

	return result
}

func (field *GameField) SaveToFile(name string) error {

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	field.SaveToIO(w)
	return nil
}

func (f *GameField) SaveToIO(w *bufio.Writer) error {
	for y := 0; y < len(f.Field[0]); y++ {
		for x := 0; x < len(f.Field); x++ {
			if f.Field[x][y][f.CurrentIndex] {
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, " ")
			}
		}
		fmt.Fprintln(w)
	}
	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (f *GameField) Print() {
	w := bufio.NewWriter(os.Stdout)
	f.SaveToIO(w)
}

func getPositionInfo(x, y int, array []string) bool {
	if y > len(array) {
		return false
	}

	if x > len(array[y]) {
		return false
	}

	if array[y][x] == '#' {
		return true
	}

	return false
}

func LoadFieldFromFile(name string) (*GameField, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	w := 0
	h := 0
	lines := make([]string, 0, 100)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		h++

		lenght := len(scanner.Text())
		if lenght > w {
			w = lenght
		}

	}

	result := NewField(name, w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			result.Field[i][j][0] = getPositionInfo(i, j, lines)
		}
	}

	return result, nil
}

func (f *GameField) hasNeighbour(x, y int) bool {
	if x >= 0 && x < len(f.Field) && y >= 0 && y < len(f.Field[x]) {
		if f.Field[x][y][f.CurrentIndex] {
			return true
		}
	}

	return false
}

func (f *GameField) calcNeighbours(x, y int) int {

	offsets := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	var newX, newY int
	result := 0
	for i := 0; i < len(offsets); i++ {
		newX = x + offsets[i][0]
		newY = y + offsets[i][1]
		if f.hasNeighbour(newX, newY) {
			result++
		}
	}

	return result
}

func (f *GameField) switchGenerationField() {
	f.CurrentIndex = f.getNextGenerationFieldIndex()
}

func (f *GameField) getNextGenerationFieldIndex() int {
	if f.CurrentIndex == 0 {
		return 1
	} else {
		return 0
	}
}

func (f *GameField) Iterate() {
	w := len(f.Field)
	h := len(f.Field[0])
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			f.Field[x][y][f.getNextGenerationFieldIndex()] = false

			neighboursNums := f.calcNeighbours(x, y)

			if f.Field[x][y][f.CurrentIndex] {
				// Alive field
				if neighboursNums == 2 || neighboursNums == 3 {
					f.Field[x][y][f.getNextGenerationFieldIndex()] = true
				}

			} else {
				// Empty field
				if neighboursNums == 3 {
					f.Field[x][y][f.getNextGenerationFieldIndex()] = true
				}
			}
		}
	}
	f.switchGenerationField()
}

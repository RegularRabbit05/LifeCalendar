package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed icon.png
var icon []byte

func main() {
	var start time.Time
	setDate := false
	rl.SetTraceLog(rl.LogNone)
	ic := rl.LoadImageFromMemory(".png", icon, int32(len(icon)))
	if len(os.Args) != 2 {
		rl.InitWindow(1000, 750, "Input a date and press enter")
		rl.SetWindowIcon(*ic)
		rl.SetTargetFPS(60)
		str := ""
		col := rl.Black
		for !rl.WindowShouldClose() {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			k := rl.GetCharPressed()
			if k != 0 {
				str += string(k)
			}
			if rl.IsKeyPressed(rl.KeyBackspace) {
				if len(str) > 0 {
					str = str[:len(str)-1]
				}
			}
			tmpStr := "Birthday: " + str + "|"
			rl.DrawText(tmpStr, int32(rl.GetScreenWidth()/2)-rl.MeasureText(tmpStr, 40)/2, int32(rl.GetScreenHeight()/2-20), 40, col)
			rl.EndDrawing()
			if rl.IsKeyPressed(rl.KeyEnter) {
				var err error
				start, err = time.Parse("02/01/2006", str)
				if err == nil {
					setDate = true
					break
				} else {
					col = rl.Red
				}
			}
		}
		rl.CloseWindow()
	} else {
		date := os.Args[1]
		var err error
		start, err = time.Parse("02/01/2006", date)
		if err != nil {
			log.Fatalln("Unable to parse date!")
			return
		}
		setDate = true
	}

	if !setDate {
		os.Exit(-1)
		return
	}

	end := start.Add(time.Hour * 24 * 365 * 100)

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1000, 750, "")
	rl.SetWindowIcon(*ic)
	rl.UnloadImage(ic)
	rl.SetWindowMinSize(830, 800)
	rl.SetTargetFPS(2)
	blink := false
	for !rl.WindowShouldClose() {
		now := time.Now()
		elapsed := int64(now.Sub(start).Hours() / 24 / 7)
		remaining := int64(end.Sub(now).Hours() / 24 / 7)
		//rl.SetWindowTitle(fmt.Sprint("Calendar of life (", rl.GetScreenWidth(), "x", rl.GetScreenHeight(), ")"))
		rl.SetWindowTitle("Calendar of life")
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprint("Birthday: ", start.Format("02/01/2006")), 20, 20, 20, rl.Black)
		rl.DrawText(fmt.Sprint("Weeks elapsed: ", elapsed), 20, 50, 20, rl.Black)
		rl.DrawText(fmt.Sprint("Weeks remaining: ", remaining), 20, 80, 20, rl.Black)

		const initY = 130
		const initX = 20
		var limitX = int32(rl.GetScreenWidth() - initX)
		currentY := int32(initY)
		currentX := int32(initX)
		for i := int64(0); i < elapsed+remaining; i++ {
			if currentX >= limitX {
				currentX = initX
				currentY += 10
			}
			col := rl.Black
			if i == elapsed {
				if blink {
					col = rl.Gray
				} else {
					col = rl.White
				}
				blink = !blink
			} else if i > elapsed {
				col = rl.LightGray
			}
			rl.DrawRectangle(currentX, currentY, 5, 5, col)
			currentX += 10
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct { // общий ресурс для мьютекса
	mu sync.Mutex
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.White)
	g.drawFractal(screen, 400, 400, 250, 6)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 800 // Размеры окна
}

func (g *Game) drawFractal(screen *ebiten.Image, x, y, size, iteration int) {
	if iteration <= 0 {
		return
	}

	rect := ebiten.NewImage(size, size)
	rect.Fill(colornames.Black)
	g.mu.Lock()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x-size/2), float64(y-size/2))
	screen.DrawImage(rect, op)
	g.mu.Unlock()

	size2 := size / 3 // определение роазмера для некст уровня

	var wg sync.WaitGroup

	wg.Add(8) // Количество горутин, которые мы собираемся запустить

	go func() { // анонимные функции, которые выполняются параллельно
		defer wg.Done()
		g.drawFractal(screen, x-size, y-size, size2, iteration-1) // верхний левый

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x, y-size, size2, iteration-1) // верхний центр

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x+size, y-size, size2, iteration-1) // верхний правый

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x+size, y, size2, iteration-1) // правый центр

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x+size, y+size, size2, iteration-1) // нижний правый

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x, y+size, size2, iteration-1) // нижний центр

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x-size, y+size, size2, iteration-1) // нижний левый

	}()
	go func() {
		defer wg.Done()
		g.drawFractal(screen, x-size, y, size2, iteration-1) // левый центр

	}()

	wg.Wait() // Ждем завершения всех горутин

}

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Fractal")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

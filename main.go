package main

func main() {
	fiber := NewFiber(_portFiber)
	defer fiber.Stop()

	fiber.Start()
}

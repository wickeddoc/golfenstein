package main

import (
	"flag"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth          = 1024
	screenHeight         = 768
	playerSize   float32 = 8.0
	playerStep   float32 = playerSize / 2.0
)

// Player variables
var (
	posX, posY float32 = 200, 200 // player position
)

// Draw player in 2D view
func drawPlayer2D() {
	// Draw player point
	rl.DrawRectangle(int32(posX), int32(posY), int32(playerSize), int32(playerSize), rl.Red)

}

// Handle player movement with mouse
func movePlayerWithMouse() {
	// Get mouse position
	mousePos := rl.GetMousePosition()

	// Calculate direction vector from player to mouse
	dirX := mousePos.X - posX
	dirY := mousePos.Y - posY

	// Calculate distance
	distance := rl.Vector2Distance(rl.NewVector2(posX, posY), mousePos)

	// Only move if mouse is not too close to player (to prevent jittering)
	if distance > playerSize {
		// Normalize direction vector and apply movement speed
		speed := playerStep
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			speed = playerStep * 2.0 // Move faster when left mouse button is pressed
		}

		// Calculate new position with normalized direction
		if distance > 0 {
			// Move with normalized direction vector
			normalizedDirX := dirX / distance
			normalizedDirY := dirY / distance
			posX += normalizedDirX * speed
			posY += normalizedDirY * speed
		}

		// Ensure player stays within screen bounds
		if posX < 0 {
			posX = 0
		} else if posX > screenWidth-playerSize {
			posX = screenWidth - playerSize
		}

		if posY < 0 {
			posY = 0
		} else if posY > screenHeight-playerSize {
			posY = screenHeight - playerSize
		}
	}
}

// Handle player movement
func movePlayerWithKeyboard() {
	// Move up
	if rl.IsKeyDown(rl.KeyW) && posY > 0 {
		posY = posY - playerStep
	}

	// Move down
	if rl.IsKeyDown(rl.KeyS) && posY < screenHeight-playerSize {
		posY = posY + playerStep
	}

	// Move left
	if rl.IsKeyDown(rl.KeyA) && posX > 0 {
		posX = posX - playerStep
	}

	// Move right
	if rl.IsKeyDown(rl.KeyD) && posX < screenWidth-playerSize {
		posX = posX + playerStep
	}
}

func main() {

	useMouse := flag.Bool("mouse", false, "Use mouse for player movement")
	flag.Parse()

	// Initialize Raylib
	rl.InitWindow(screenWidth, screenHeight, "Raycaster in Go using Raylib")
	rl.SetTargetFPS(60)

	if *useMouse {
		rl.HideCursor()
	}

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update
		// Update
		if *useMouse {
			movePlayerWithMouse()
		} else {
			movePlayerWithKeyboard()
		}

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw player
		drawPlayer2D()

		// Display controls based on mode
		if *useMouse {
			rl.DrawText("Mouse: Move, Left Click: Speed up", 10, screenHeight-30, 20, rl.DarkGray)
		} else {
			rl.DrawText("W,A,S,D: Move", 10, screenHeight-30, 20, rl.DarkGray)
		}

		rl.EndDrawing()
	}

	if *useMouse {
		rl.ShowCursor()
	}

	// Clean up resources
	rl.CloseWindow()
}

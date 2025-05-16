package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 768
	playerSize   = 8
	playerStep   = playerSize / 2
)

// Player variables
var (
	posX, posY float32 = 200, 200 // player position
)

// Draw player in 2D view
func drawPlayer2D() {
	// Draw player point
	rl.DrawRectangle(int32(posX), int32(posY), playerSize, playerSize, rl.Red)

}

// Handle player movement
func movePlayer() {
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
	// Initialize Raylib
	rl.InitWindow(screenWidth, screenHeight, "Raycaster in Go using Raylib")
	rl.SetTargetFPS(60)

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update
		movePlayer()

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw player
		drawPlayer2D()

		// Display controls
		rl.DrawText("W,A,S,D: Move", 10, screenHeight-30, 20, rl.DarkGray)

		rl.EndDrawing()
	}

	// Clean up resources
	rl.CloseWindow()
}

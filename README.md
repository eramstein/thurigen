# GoMagic

A turn-based collectible card game where players compete to control a grid using buildings and units.

## Prerequisites

- Go 1.16 or later
- GCC (for Windows, you can use MinGW-w64)
- Make (optional, for building)
- raylib.dll (version 5.5)

## Setup

1. Clone this repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Download raylib.dll (version 5.5) and place it in the root directory
   - You can download it from: https://github.com/raysan5/raylib/releases/tag/5.5.0
   - Or from raylib-go releases: https://github.com/gen2brain/raylib-go/releases

## Running the Project

To run the project, execute:
```bash
go run cmd/game/main.go
```

## Project Structure

```
gomagic/
├── cmd/
│   └── game/           # Main entry point
│       └── main.go     # Application entry point
├── pkg/
│   └── game/          # Game logic packages
│       ├── board/     # Grid and tile management
│       ├── card/      # Card definitions and mechanics
│       ├── player/    # Player state and actions
│       ├── engine/    # Game rules and turn management
│       └── ui/        # User interface and rendering
├── go.mod             # Go module definition
└── raylib.dll         # Raylib dynamic library
```

## Components

- **board**: Manages the game grid, tile control, and placement rules
- **card**: Defines card types (buildings/units), abilities, and effects
- **player**: Handles player state, mana, and card collections
- **engine**: Controls game flow, turn management, and rule enforcement
- **ui**: Manages the game interface, input handling, and rendering

## Features

- Turn-based gameplay
- Grid-based card placement
- Mana system for card costs
- Building and unit cards with special abilities
- Tile control mechanics
- Win condition tracking 
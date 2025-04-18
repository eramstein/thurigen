# Thuringen Simulation

A colony building simulation/game set in the late middle ages (around 1400), where a group of people is gifted a domain with a small castle and its surroundings.

## Project Structure

```
gomagic/
├── cmd/
│   └── game/           # Main entry point
│       └── main.go     # Application entry point
├── pkg/
│   ├── engine/         # Game logic and state management
│   │   └── models/     # Game state models
│   └── ui/            # User interface and rendering
├── go.mod              # Go module definition
└── raylib.dll          # Raylib dynamic library
```

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

## Controls

- Press **SPACE** to pause and resume the simulation
- Press **ENTER** to restart the ticker
- Press **F5** to save the current simulation state
- Press **F4** to load the most recent save file
- Press **F1** to run benchmarks

## Save Files

Save files are stored in the `saves` directory with timestamps in their filenames (e.g., `save_2024-03-21_15-30-45.json`). Each save file contains the complete state of the simulation, including:
- World state (tiles, terrain, structures)
- Characters and their states
- Items and their locations
- Simulation time and settings

## Features

- Time-based simulation (in minutes)
- Colony building mechanics
- Character simulation with AI agents
- Simple graphics with no animations
- Regional/zonal spatial management
- Late middle ages setting (around 1400)
- Save/Load functionality

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

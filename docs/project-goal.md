--------------------
What we are building
--------------------

A colony building simulation/game. In the late middle ages (around 1400), a group of people is being gifted a domain with a small castle and its surroundings.

Low spatial details, only regions and zones. Time is measured in minutes.

Simple graphics and no animations (only static images representing places and characters).

Deep simulation of characters, which are AI agents pursuing their own goals.

--------------------
Implementation goals
--------------------

A Go application using the Raylib library.
Minimal assets and graphics, no animations.

There are 2 high level packages:
- engine: handles game logic, ai, and state management
- ui: handles inputs and rendering

The ui package can use the engine package, but the engine package should never import the ui package to make sure we can change rendering libraries easily later.
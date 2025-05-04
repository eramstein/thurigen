package ng

import "fmt"

func (sim *Simulation) FollowPath(character *Character, extraMove bool) {
	if character.Path == nil {
		return
	}
	task := character.CurrentTask
	path := *character.Path
	if len(path) > 0 {
		moveCost := sim.GetMoveCost(character, path[0])
		if moveCost == -1 {
			return
		}

		// for move tasks, the progress represents how many move points have been spent in the move to next tile in the path
		// an "extra move" is using only left over action points from the same tick
		if !extraMove {
			task.Progress += character.Stats.Speed
		}
		// if enough move points have been spent, set the character to the new position
		if task.Progress >= moveCost {
			sim.SetCharacterPosition(character, path[0])
			if len(path) == 1 {
				character.Path = nil
				sim.CompleteTask(character)
				return
			}
			newPath := path[1:]
			character.Path = &newPath
			task.Progress = task.Progress - moveCost
			// get started on next move if excess move points
			if task.Progress >= 1 {
				sim.FollowPath(character, true)
			}
		}
	}
}

func (sim *Simulation) MoveForTask(character *Character) {
	task := character.CurrentTask
	target, ok := task.Target.(*Position)
	if !ok {
		fmt.Printf("Task target is not a *Position: %v\n", task.Target)
		return
	}
	if character.Path == nil || (*character.Path)[len(*character.Path)-1] != *target {
		path := sim.World[character.Position.Region].FindPath(character.Position.X, character.Position.Y, target.X, target.Y, 0)
		if path == nil {
			// TODO: roll back objective planning due to invalid path
			fmt.Printf("No path found for %v to %v\n", character.Name, target)
			return
		}
		character.Path = &path
	}
	sim.FollowPath(character, false)
}

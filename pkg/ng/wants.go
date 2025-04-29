package ng

const (
	WantSleepOnFloor = -10
)

func (character *Character) UpdateWantsTotal() {
	character.Wants.Confort.Total = character.Wants.Confort.SleepConditions
}

func (character *Character) UpdateSleepConditionsWants(sleepQuality int) {
	character.Wants.Confort.SleepConditions = sleepQuality
	character.UpdateWantsTotal()
}

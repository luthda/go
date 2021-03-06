package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	hasNoRescueDoggo := !knightIsAwake && !archerIsAwake && prisonerIsAwake && !petDogIsPresent
	hasRescueDoggo := petDogIsPresent && !archerIsAwake
	return hasRescueDoggo || hasNoRescueDoggo
}

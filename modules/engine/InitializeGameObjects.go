package engine

import "./gameobjects"

//Create Hut
var playerVHut = gameobjects.Hut{
	Name:     "Villager's Huts",
	Level:    1,
	WoodReq:  0,
	StoneReq: 0,
	FiberReq: 0,
}
var playerOpsHut = gameobjects.Hut{
	Name:     "Operations Hut",
	Level:    1,
	WoodReq:  0,
	StoneReq: 0,
	FiberReq: 0,
}
var playerSuppliesHut = gameobjects.Hut{
	Name:     "Supplies Hut",
	Level:    1,
	WoodReq:  0,
	StoneReq: 0,
	FiberReq: 0,
}
var playerCamp = gameobjects.Camp{
	Name:          "",
	OperationsHut: playerOpsHut,
	VillagersHut:  playerVHut,
	SuppliesHut:   playerSuppliesHut,
	Population:    10,
}
var ThePlayer = gameobjects.Player{
	Name:       "",
	Camp:       playerCamp,
	Wood:       100,
	Stone:      100,
	Fiber:      100,
	TurnsLeft:  5,
	TotalTurns: 5,
}

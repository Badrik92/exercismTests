package main

import (
	"fmt"
)

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	}
	if !knightIsAwake {
		return true
	}
	return knightIsAwake
	
}
// CanSpy может быть запущен, если хотя бы один из признаков не спит
func CanSpy(knigthIsAwake, archerIsAwake, prisonerIsAwake bool ) bool {
	if knigthIsAwake || archerIsAwake || prisonerIsAwake{
		return true
	} 
	return false
}

// CanSignalPrisoner может быть согласован, если плотность бодрствует, аник луч спит
func CanSignalPrisoner (archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	}
	return false
}

// CanFreePrisoner может быть выполнен, если Договорный не спит, а 2 других персонажа спят
// или если домашняя собака Анналин с ней, аник луч спит
func CanFreePrisoner( knigthIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if !knigthIsAwake && !archerIsAwake && prisonerIsAwake || petDogIsPresent && !archerIsAwake {
		return true
	}
	return false
}
func main() {
	fmt.Println(CanFastAttack(true))
	fmt.Println(CanSpy(false, true, false))
	fmt.Println(CanSignalPrisoner(false, false))
	fmt.Println(CanFreePrisoner(false, true, false, false))
}
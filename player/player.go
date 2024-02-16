package player

import "fmt"

type player struct {
	Prefix string
	name   string
	health int
}

// Constructor function for Player
func NewPlayer(name string, health int) *player {
	return &player{
		name:   name,
		health: health,
	}
}

// Getter and setter methods for name
func (p *player) GetName() string {
	return p.name
}
func (p *player) SetName(name string) {
	p.name = name
}

// Getter and setter methods for health
func (p *player) GetHealth() int {
	return p.health
}
func (p *player) SetHealth(health int) {
	p.health = health
}

// Method to display player information
func (p *player) DisplayInfo() {
	fmt.Println("Prefix:", p.Prefix)
	fmt.Println("Player Name:", p.name)
	fmt.Println("Player Health:", p.health)
}

// Method to decrease player's health
func (p *player) DeleteHealth(damage int) {
	if damage >= p.health {
		p.health = 0
	} else {
		p.health -= damage
	}
}

package main

type player interface {
	calculateHP() int
}

// Npc

type npc struct {
	name   string
	status string
	hp     int
}

func (n npc) getName() string {
	return n.name
}

func (n npc) calculateHP(number int) int {
	var total int
	total = n.hp + number

	return total
}

//  Mage

type mage struct {
	name   string
	status string
	hp     int
	mp     int
}

func (m mage) getName() string {
	return m.name
}

func (m mage) getStatus() string {
	return m.status
}

func (m mage) calculateHP(number int) int {
	var total int
	total = m.hp + number

	return total
}

func main() {

}

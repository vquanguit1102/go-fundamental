package main

func main() {
	cards := newDeck()

	take, remaining := deal(cards, 3)
	take.print()
	remaining.print()
}
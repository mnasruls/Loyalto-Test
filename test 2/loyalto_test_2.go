package main

import (
	"fmt"
	"math/rand"
)

func Remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func Stop(playerDice map[int][]int) int {
	count := 0
	for _, v := range playerDice {
		if len(v) == 0 {
			count += 1
		}
	}
	return count
}

func RandomDadu(player, dice int) int {
	winner := 0
	playerDice := make(map[int][]int)
	points := make(map[int]int)

	for j := 1; j <= player; j++ {
		for k := 1; k <= dice; k++ {
			randomValue := rand.Intn(7)
			if randomValue != 0 {
				playerDice[j] = append(playerDice[j], randomValue)
			} else {
				k -= 1
			}
		}
	}
	for i := 1; i <= 10; i++ {
		for a := 1; a <= player; a++ {
			countZeroValue := Stop(playerDice)
			// condisi jika tersisa 1 pemain yang punya dadu
			if countZeroValue == player-1 {
				winnerPoint := 0
				for d := 1; d <= len(points); d++ {

					if points[d] > winnerPoint {
						winnerPoint = points[d]
						winner = d
					}
				}
				return winner
			} else {
				fmt.Println("hasil lemparan ke ", i, " dari player ke ", a, " adalah: ", playerDice)
				for b := 0; b < len(playerDice[a]); b++ {
					if playerDice[a][b] == 6 {
						points[a] += 1
						playerDice[a] = Remove(playerDice[a], b)
						b -= 1
					} else if playerDice[a][b] == 1 {
						if a == player {
							playerDice[1] = append(playerDice[1], 0)
							playerDice[a] = Remove(playerDice[a], b)
							b -= 1
						} else {
							playerDice[a+1] = append(playerDice[a+1], 0)
							playerDice[a] = Remove(playerDice[a], b)
							b -= 1
						}
					} else {
						randomValue := rand.Intn(7)
						if randomValue != 0 {
							playerDice[a][b] = randomValue
						} else {
							b -= 1
						}
					}

				}
			}
			fmt.Println("setelah dievaluasi: ", playerDice)
		}
	}
	return -1
}

func main() {
	fmt.Println("Pemenangnya adalah player: ", RandomDadu(3, 4))
}

package tokyo_go_sdk

import (
	"math"
)

func DistanceBetweenTwoPoints(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func RadianBetweenPoint1ToPoint2(x1, y1, x2, y2 float64) float64 {
	var radianPadding float64
	switch {

	case x1 == x2 && y1 == y2:
		return 0

	case x2 > x1 && y2 == y1:
		return 0
	case x2 == x1 && y2 > y1:
		return math.Pi / 2
	case x2 < x1 && y2 == y1:
		return math.Pi
	case x2 == x1 && y2 < y1:
		return math.Pi * 3 / 2

	case x2 > x1 && y2 > y1: // > 0* and <= 90*
		radianPadding = 0
	case x2 < x1 && y2 > y1:
		radianPadding += math.Pi / 2
	case x2 < x1 && y2 < y1:
		radianPadding += math.Pi
	case x2 > x1 && y2 < y1:
		radianPadding += math.Pi * 3 / 2

	}

	hypotenuse := DistanceBetweenTwoPoints(x1, y1, x2, y2)
	opposite := math.Abs(y2 - y1)
	sin := opposite / hypotenuse
	radians := math.Asin(sin)
	return radians + radianPadding
}

func IsPlayersSamePosition(player1, player2 Player) bool {
	return player1.X == player2.X && player1.Y == player2.Y
}

func IsSamePlayers(player1, player2 Player) bool {
	return player1.ID == player2.ID
}

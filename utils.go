package tokyo_go_sdk

import (
	"math"
)

// DistanceBetweenTwoPoints return distance between 2 points
func DistanceBetweenTwoPoints(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

// DistanceBetweenTwoPlayers return distance between 2 players
func DistanceBetweenTwoPlayers(player1, player2 Player) float64 {
	return DistanceBetweenTwoPoints(player1.X, player1.Y, player2.X, player2.Y)
}

// CalculateAngleFromPoint1ToPoint2 return radian angle between 2 points
func CalculateAngleFromPoint1ToPoint2(x1, y1, x2, y2 float64) float64 {
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
	}

	angle := AngleBetweenPoint1ToPoint2(x1, y1, x2, y2)
	switch {
	case x2 > x1 && y2 > y1: // > 0* and <= 90*
		return angle
	case x2 < x1 && y2 > y1:
		return math.Pi - angle
	case x2 < x1 && y2 < y1:
		return angle + math.Pi
	case x2 > x1 && y2 < y1:
		return math.Pi*2 - angle
	}
	return 0
}

// AngleBetweenPoint1ToPoint2 return radian angle between 2 points
func AngleBetweenPoint1ToPoint2(x1, y1, x2, y2 float64) float64 {
	adjacent := math.Abs(x2 - x1)
	opposite := math.Abs(y2 - y1)
	sin := opposite / adjacent
	radian := math.Atan(sin)
	return radian
}

// IsPlayersSamePosition check 2 player objects are same position
func IsPlayersSamePosition(player1, player2 Player) bool {
	return player1.X == player2.X && player1.Y == player2.Y
}

// IsSamePlayers check 2 players same ID
func IsSamePlayers(player1, player2 Player) bool {
	return player1.ID == player2.ID
}

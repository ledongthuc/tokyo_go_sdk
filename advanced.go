package tokyo_go_sdk

import (
	"errors"
)

func (c *Client) HeadToPoint(x, y float64) error {
	radian := RadianBetweenPoint1ToPoint2(c.CurrentPlayer.X, c.CurrentPlayer.Y, x, y)
	return c.Rotate(radian)
}

func (c *Client) GetFirstOtherPlayer() (Player, error) {
	if len(c.CurrentGameInfo.Players) == 0 {
		return Player{}, errors.New("Don't find anyone, 0 player")
	}

	for _, player := range c.CurrentGameInfo.Players {
		if IsSamePlayers(player, c.CurrentPlayer) || IsPlayersSamePosition(player, c.CurrentPlayer) {
			continue
		}
		if c.IsDeadPlayer(player) {
			continue
		}
		return player, nil
	}

	return Player{}, errors.New("Don't find anyone, only me")
}

func (c *Client) IsDeadPlayer(player Player) bool {
	for _, aDead := range c.CurrentGameInfo.Dead {
		if aDead.Player.ID == player.ID {
			return true
		}
	}
	return false
}

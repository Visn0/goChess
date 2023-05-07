package wsrouter

import (
	"chess/server/domain"
	"chess/server/shared/chesserror"
	"encoding/json"
)

type RequestBody []byte

type Context struct {
	Game   *domain.Game
	Player *domain.Player
	Enemy  *domain.Player

	Body RequestBody
}

func NewContext(game *domain.Game, player, enemy *domain.Player, body RequestBody) *Context {
	return &Context{
		Game:   game,
		Player: player,
		Enemy:  enemy,
		Body:   body,
	}
}

func (c *Context) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *Context) Bind(dst interface{}) error {
	err := json.Unmarshal(c.Body, dst)
	if err != nil {
		return chesserror.NewError(chesserror.GenericError,
			"Error unmarshalling input parameters.").WithCause(err)
	}

	return nil
}

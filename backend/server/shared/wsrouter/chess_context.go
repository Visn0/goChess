package wsrouter

import (
	"chess/server/domain"
	"encoding/json"
)

type RequestBody []byte

type Context struct {
	Game            *domain.Game
	Player          *domain.Player
	Enemy           *domain.Player
	OwnRepository   domain.ConnectionRepository
	EnemyRepository domain.ConnectionRepository

	Body RequestBody
}

func NewContext(game *domain.Game, player, enemy *domain.Player,
	ownRep, enemyRep domain.ConnectionRepository, body RequestBody) *Context {
	return &Context{
		Game:            game,
		Player:          player,
		Enemy:           enemy,
		OwnRepository:   ownRep,
		EnemyRepository: enemyRep,
		Body:            body,
	}
}

func (c *Context) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

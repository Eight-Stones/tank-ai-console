package model

type Player struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	HitPoints int    `json:"hitpoints,omitempty"`
	Ammo      int    `json:"ammo,omitempty"`
	IsAlive   bool   `json:"isAlive,omitempty"`
	Direction string `json:"direction,omitempty"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}

type Game struct {
	ID      string    `json:"id"`
	Status  string    `json:"status,omitempty"`
	Players []*Player `json:"players,omitempty"`
}

type GetNodeGamesRequest struct{}

type GetNodeGamesResponse struct {
	Payload []*Game `json:"payload,omitempty"`
}

type GetNodeGameRequest struct {
	ID string `fastmicro:"id,path"`
}

type GetNodeGameResponse struct {
	Payload *Game `json:"payload,omitempty"`
}

type GameCell struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
}

type GameColumn []GameCell

type GameField []GameColumn

type GetNodeViewRequest struct {
	ID string `fastmicro:"id,path"`
}

type GetNodeViewResponse struct {
	Payload GameField `json:"payload,omitempty"`
}

type BotMeta struct {
	ID string `json:"id"`
}

type PostBotRequest struct {
	ID string `fastmicro:"id,path"`
}

type PostBotResponse struct{}

type PostCreateGameRequest struct{}

type PostCreateGameResponse struct {
	Payload *Game `json:"payload,omitempty"`
}

type PutRunGameRequest struct {
	ID     string `fastmicro:"id,path"`
	Action string `json:"action,omitempty"`
}

type PutRunGameResponse struct{}

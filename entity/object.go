package entity

type Coordinates struct {
	X int
	Y int
}

type Object struct {
	ID        string
	Direction string
	HP        int
	IsAlive   bool
	Type      ObjectType
	Coordinates
}

type Tank struct {
	Object
	Ammo   int
	Move   int
	Kill   int
	Shoot  int
	Rotate int
	Vision int
}

type Bullet struct {
	Object
}

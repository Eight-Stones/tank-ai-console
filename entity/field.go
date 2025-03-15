package entity

import (
	"sync"

	engine "github.com/Eight-Stones/ecs-tank-engine/v2"
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
)

type ObjectType components.ObjectType

func (o ObjectType) String() string {
	return components.ObjectType(o).String()
}

type Field struct {
	mu      *sync.Mutex
	tanks   map[string]*Tank
	bullets map[string]*Bullet
	corps   map[string]*Tank
}

func NewField() *Field {
	return &Field{
		mu:      &sync.Mutex{},
		tanks:   make(map[string]*Tank),
		bullets: make(map[string]*Bullet),
		corps:   make(map[string]*Tank),
	}
}

func (f *Field) Apply(i *engine.Info) {
	f.mu.Lock()
	defer f.mu.Unlock()

	switch i.Type.String() {
	case "create":
		f.applyCreate(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Create))
	case "remove":
		f.applyRemove(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Remove))
	case "rotate":
		f.applyRotate(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Rotate))
	case "move":
		f.applyMove(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Position))
	case "shoot":
		f.applyShoot(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Shoot))
	case "health":
		f.applyHealth(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Health))
	case "vision":
		f.applyVision(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Vision))
	case "radar":
		f.applyRadar(i.Id, ObjectType(i.ObjectType), i.MetaInfo.(*engine.Radar))
	}
}

func (f *Field) applyCreate(id string, typ ObjectType, meta *engine.Create) {}
func (f *Field) applyRemove(id string, typ ObjectType, meta *engine.Remove) {}
func (f *Field) applyRotate(id string, typ ObjectType, meta *engine.Rotate) {}
func (f *Field) applyMove(id string, typ ObjectType, meta *engine.Position) {}
func (f *Field) applyShoot(id string, typ ObjectType, meta *engine.Shoot)   {}
func (f *Field) applyHealth(id string, typ ObjectType, meta *engine.Health) {}
func (f *Field) applyVision(id string, typ ObjectType, meta *engine.Vision) {}
func (f *Field) applyRadar(id string, typ ObjectType, meta *engine.Radar)   {}
func (f *Field) View() []Object                                             { return nil }

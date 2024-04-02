package collisions

import (
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

type CollisionBoxWithEntity struct {
	CollisionBox *shapes.CollisionBox
	Entity       interfaces.Entity
}

var GlobalCollisionPool []interfaces.Collidable

func AddToPool(collisionComponent interfaces.Collidable) {
	GlobalCollisionPool = append(GlobalCollisionPool, collisionComponent)
}

func RemoveFromPool(entity interfaces.Entity) {
	for i := len(GlobalCollisionPool) - 1; i >= 0; i-- {
		if GlobalCollisionPool[i].GetEntity() == entity {
			GlobalCollisionPool = append(GlobalCollisionPool[:i], GlobalCollisionPool[i+1:]...)
		}
	}
}

func GetGlobalCollisionPool() []interfaces.Collidable {
	return GlobalCollisionPool
}

func CheckCollisions() error {
	pool := GetGlobalCollisionPool()
	for i := 0; i < len(pool); i++ {
	Component:
		for j := i + 1; j < len(pool); j++ {
			if pool[i].GetEntity() == pool[j].GetEntity() {
				continue // Skip if they belong to the same entity
			}

			collisionBoxesI := pool[i].GetCollisionBoxes()
			collisionBoxesJ := pool[j].GetCollisionBoxes()

			for _, boxI := range collisionBoxesI {
				for _, boxJ := range collisionBoxesJ {
					if !collides(*boxI, *boxJ) {
						continue // Skip to the next pair if these don't collide
					}
					// Handle collision
					pool[i].OnCollision(pool[j].GetEntity())
					pool[j].OnCollision(pool[i].GetEntity())

					// We only care about the first collision for each component (not each collision box)
					continue Component
				}
			}
		}
	}
	return nil
}

func collides(c1, c2 shapes.CollisionBox) bool {
	return c1.Position.X+c1.Box.Width >= c2.Position.X &&
		c1.Position.X <= c2.Position.X+c2.Box.Width &&
		c1.Position.Y+c1.Box.Height >= c2.Position.Y &&
		c1.Position.Y <= c2.Position.Y+c2.Box.Height
}

package collisions

import (
	"github.com/jacovanc/airconsole-ebiten/game/interfaces"
	"github.com/jacovanc/airconsole-ebiten/game/shapes"
)

func CheckCollisions(entitiesArray []interfaces.Entity) error {
	// Check every entity against every other entity
	// We will improve this when efficiency becomes a problem
	// // For example, only check collisions for the player against the platforms
	for i := 0; i < len(entitiesArray); i++ {
		for j := i + 1; j < len(entitiesArray); j++ {
			for _, c1 := range *entitiesArray[i].GetCollisions() {
				for _, c2 := range *entitiesArray[j].GetCollisions() {
					if collides(c1, c2) {
						if err := entitiesArray[i].Collision(entitiesArray[j]); err != nil {
							return err
						}
						if err := entitiesArray[j].Collision(entitiesArray[i]); err != nil {
							return err
						}
					}
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
		c1.Position.Y <= c2.Position.Y+c2.Box.Width
}
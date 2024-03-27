package main

import "fmt"

func checkCollisions(entities []*entity) error {
	// Check every entity against every other entity
	for i := 0; i < len(entities); i++ {
		for j := i + 1; j < len(entities); j++ {
			for _, c1 := range entities[i].collisions {
				for _, c2 := range entities[j].collisions {
					if collides(c1, c2) {
						fmt.Println("Something collided")
						if err := entities[i].collision(entities[j]); err != nil {
							return err
						}
						if err := entities[j].collision(entities[i]); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func collides(c1, c2 rectangle) bool {
	return c1.position.x + c1.width >= c2.position.x &&
		c1.position.x <= c2.position.x + c2.width &&
		c1.position.y + c1.height >= c2.position.y &&
		c1.position.y <= c2.position.y + c2.height;
}
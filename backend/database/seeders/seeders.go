package seeders

import "log"

func SeedEntityData() {
	if err := SeedUserAdmin(); err != nil {
		log.Fatalln("unable to seed admin data")
	}
}

package main

import (
	"fmt"

	"github.com/andynl/go-postgres/config"
	"github.com/andynl/go-postgres/src/modules/profile/model"
	"github.com/andynl/go-postgres/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Postgres")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	// err = db.Ping()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	wury := model.NewProfile()
	wury.ID = "1"
	wury.FirstName = "Andy"
	wury.LastName = "Natalino"
	wury.Email = "andy.natalino@gmail.com"
	wury.Password = "123456"

	profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)

	// err = saveProfile(wury, profileRepositoryPostgres)
	// err = updateProfile(wury, profileRepositoryPostgres)
	// err = deleteProfile("1", profileRepositoryPostgres)
	// profile, err := getProfile("1", profileRepositoryPostgres)
	profiles, err := getProfiles(profileRepositoryPostgres)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(profile)

	for _, v := range profiles {
		fmt.Println(v)
	}
}

func saveProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

func deleteProfile(id string, repo repository.ProfileRepository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}

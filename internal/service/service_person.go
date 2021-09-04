package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/db"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// CreatePerson ...
func CreatePerson(person *model.Person) error {
	return dbmgr.Create(person).Error
}

// DeletePerson ...
func DeletePerson(personID uint) error {
	return dbmgr.Delete(&model.Person{}, personID).Error
}

// UpdatePerson ...
func UpdatePerson(person *model.Person) error {
	return dbmgr.Model(&person).Updates(person).Error
}

// FindPersonByID find one person record according to id
func FindPersonByID(personID uint) (*model.Person, error) {
	person := &model.Person{}
	if err := dbmgr.First(person, personID).Error; err != nil {
		return nil, err
	}
	return person, nil
}

// SelectPersonS find all persons with model.Person{} conditions
func SelectPersonS(person model.Person) ([]*model.Person, error) {
	persons := make([]*model.Person, 0, 4)
	if err := dbmgr.Where(&person).Find(persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

// SelectPersonM find all persons with map conditions
func SelectPersonM(condition db.M) ([]*model.Person, error) {
	persons := make([]*model.Person, 0, 4)
	if err := dbmgr.Where(condition).Find(persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

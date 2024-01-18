package db

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Sync[T mgm.Model, S any](records *[]S,
	updateRecord func(record T, src S) error,
	createRecord func(src S) (T, error),
	filter func(src S) bson.D) error {

	for _, record := range *records {

		t := new(T)
		changedRecord, err := FindOne[T](*t, filter(record))
		if err != nil {

			t, err := createRecord(record)
			if err != nil {
				return err
			}
			mgm.Coll(t).Create(t)

		} else {

			err = updateRecord(changedRecord, record)
			if err != nil {
				return err
			}

			mgm.Coll(changedRecord).Update(changedRecord)

		}

	}
	return nil
}

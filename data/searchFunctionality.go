package data

import "groupie-tracker/funcs"

func FillSearchStorage() error {
	err := funcs.GetAndParse(MainData.Artists, &SearchStorage.Artists)
	if err != nil {
		return err
	}

	err = funcs.GetAndParse(MainData.Locations, &SearchStorage.Locations)
	if err != nil {
		return err
	}

	err = funcs.GetAndParse(MainData.Dates, &SearchStorage.Dates)
	if err != nil {
		return err
	}

	err = funcs.GetAndParse(MainData.Relations, &SearchStorage.Relations)
	if err != nil {
		return err
	}

	return nil
}

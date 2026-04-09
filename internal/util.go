package internal

import "log"

func Check(errorContext string, err error) {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
		panic(err)
	}
}

func CheckSlient(errorContext string, err error) error {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
		return err
	}
	return nil
}

package erratum

func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = o()
		if _, ok := err.(TransientError); ok {
			continue
		}
		if err != nil {
			return err
		} else {
			break
		}
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				res.Defrob(e.defrobTag)
				err = e
			}
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	res.Frob(input)

	return err
}

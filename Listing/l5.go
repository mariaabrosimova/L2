package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test_() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test_()
	// потому что nil от test_ это *customError nil и сравнивается прослым nil (отличаются)
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

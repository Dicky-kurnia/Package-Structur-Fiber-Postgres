package helper

func IsShouldPanic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func CreateInvoiceId() string {
	return "INV" + GenRandomInt()
}

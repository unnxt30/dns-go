package helpers


func RespondWithError(err error){
	if err != nil {
		panic(err)
	}
}
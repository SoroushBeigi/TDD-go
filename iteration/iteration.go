package iteration

func Repeat(character string,count int) string {
	var repeated string
	for i:=0; i < 5 ; i++{
		repeated+=character
	}
	return repeated
}
package api


//look at how county data is set up in database, implement it here so it can be unmarsheld
type County struct {
	//Include important info about county
	countyName    string
	coordinate    int
}
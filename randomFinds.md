using `json:"_"` will remove the key value from the output entirely

type SomeStruct struct {
	Felid1      string  `json:"felid1"`
    Felid2      string  `json:"_"`
    Felid3      string  `json:"felid3"`
    Felid4      string  `json:"_"`
}

this will lead the output to look like {
    "felid1":"value"
    "felid3":"value"
}
even when Felid2 & Felid4 held values in the backend processes
package consts

type fileData struct {
	fileName string
	FileType string
	fileSize string
	fileHash string
}

type reputationDomain struct {
	domain      string
	isGood      bool
	threatScore int
}

type ReputationUrl struct {
	url         string
	isGood      bool
	threatScore int
}

type fileHashReputation struct {
	filehash string
	filename string
	fileType string
	fileSize string
}

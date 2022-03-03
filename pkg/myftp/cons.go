package myftp

// FTP Status codes, defined in RFC 959
const (
	StatusFileOK                = "150"
	StatusOK                    = "200"
	StatusSystemStatus          = "211"
	StatusDirectoryStatus       = "212"
	StatusFileStatus            = "213"
	StatusConnectionClosing     = "221"
	StatusSystemType            = "215"
	StatusClosingDataConnection = "226"
	StatusActionOK              = "250"
	StatusPathCreated           = "257"
	StatusActionPending         = "350"
)

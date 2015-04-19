package tests

func crlf_missing(msg string) bool {
	//     Second to last byte == \r         last byte == \n
	return (msg[len(msg)-2] == 0x0D && msg[len(msg)-1] == 0x0A)
}

func too_long(msg string) bool {
	return len(msg) >= 512
}

func err_msg(test string, msg string, err string, section string) string {
	return ("Test " + test + " Failed\n" +
		"Error: " + err + "\n" +

		"More information can be found in the " +
		section + " section(s) of the protocol.\n" +

		"Message that caused the error: " + msg)
}

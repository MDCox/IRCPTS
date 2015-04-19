package tests

// Tests that are specified in section 3.1 of RFC 2812: "Connection Registration"
//
// DESCRIPTION FROM RFC:
//     "The commands described here are used to register a connection
//      with an IRC server as a user as well as to correctly disconnect."
func connection_registration_tests() testSet {
	return testSet{
		section: "3.1 Connection Registration",
		Tests: []Test{

			// TODO: Test for PASS which is defined in section 3.1.1
			//
			Test{
				name:    "PASS is sent",
				section: "3.1.1 Password Message",
				err:     "PASS command was not sent by client",

				criteria: func(msg string) bool {
					return true
				},
			},

			// 3.1.2 Nick message
			//
			// Command: NICK
			// Parameters: <nickname>
			//
			// NICK command is used to give user a nickname or
			// change the existing one.
			//
			// Numeric Replies:
			//     ERR_NONICKNAMEGIVEN         ERR_ERRONEUSNICKNAME
			//     ERR_NICKNAMEINUSE           ERR_NICKCOLLISION
			//     ERR_UNAVAILRESOURCE         ERR_RESTRICTED
			//
			// Examples:
			//
			// NICK Wiz             ; Introducing new nick "Wiz" if session is
			// 			still unregistered, or user changing his
			// 			nickname to "Wiz"
			//
			// :WiZ!jto@tolsun.oulu.fi NICK Kilroy
			//			; Server telling that WiZ changed his
			//                      nickname to Kilroy.
			Test{
				name:    "NICK message sent",
				section: "3.1.2 Nick message",
				err:     "No NICK message was sent",

				criteria: func(msg string) bool {
					return true
				},
			},

			// 3.1.3 User message

			// Command: USER
			// Parameters: <user> <mode> <unused> <realname>

			// The USER command is used at the beginning of connection to specify
			// the username, hostname and realname of a new user.

			// The <mode> parameter should be a numeric, and can be used to
			// automatically set user modes when registering with the server.  This
			// parameter is a bitmask, with only 2 bits having any signification: if
			// the bit 2 is set, the user mode 'w' will be set and if the bit 3 is
			// set, the user mode 'i' will be set.  (See Section 3.1.5 "User
			// Modes").

			// The <realname> may contain space characters.

			// Numeric Replies:

			// ERR_NEEDMOREPARAMS    ERR_ALREADYREGISTRED

			// Example:

			// USER guest 0 * :Ronnie Reagan   ; User registering themselves with a
			// username of "guest" and real name
			// "Ronnie Reagan".

			// USER guest 8 * :Ronnie Reagan   ; User registering themselves with a
			// username of "guest" and real name
			// "Ronnie Reagan", and asking to be set
			// invisible.
			Test{
				name:    "USER message sent",
				section: "3.1.3 Nick message",
				err:     "No USER message sent",

				criteria: func(msg string) bool {
					return true
				},
			},
		},
	}
}

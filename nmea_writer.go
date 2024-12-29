package nmea

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

func Sentence(n nmea_sentence) string {
	message := fmt.Sprintf("%s,%s", n.Header(), n.Values())
	return fmt.Sprintf("$%s*%02X", message, Checksum(message))
}

func Checksum(s string) byte {
	var result byte
	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	return result
}

func WriteMessage(b NMEABoat, seq []string) string {
	var message strings.Builder
	log.Println("Building message for boat at: ", b.Latitude, b.Longitude)
	// Always mark data as from a virtual origin
	message.WriteString("$SOL\n")

	// Instanciate and print Sentences
	for i := 0; i < len(seq); i++ {
		s := strings.Split(seq[i], ".")

		// Ignore unsupported types
		mtype, ok := MessageTypes[s[0]]
		if ok {

			v := reflect.New(mtype.Elem()).Interface()
			n := v.(nmea_sentence)

			// Set option for message, if any
			if len(s) > 1 {
				n.SetOpt(s[1][0])
			}

			n.FromBoat(b)

			message.WriteString(Sentence(n))
			message.WriteString("\n")
		}
	}

	return message.String()
}

## AIS

AIVDM : VDM phrase, AI is the talker ID for "mobile AIS station"

Hardest part will be to encode data in the 6-bit format. Otherwise it's straight forward NMEA protocol.

2 Messages types :
Position (type 1)
Static (type 5, vessel name)

Have to find a way to link both messages, in real world it's done using the MMSI

https://www.navcen.uscg.gov/sites/default/files/pdf/MMSIAllotments/MMSI%20Summary.pdf

Maybe use 999 as MDI then sailaway ID (up to 1 million sa unique boats)

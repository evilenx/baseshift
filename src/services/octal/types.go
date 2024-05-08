package octal

var octalToHexadecimalMap = map[string]string{
	"0000": "0",
	"0001": "1",
	"0010": "2",
	"0011": "3",
	"0100": "4",
	"0101": "5",
	"0110": "6",
	"0111": "7",
	"1000": "8",
	"1001": "9",
	"1010": "A",
	"1011": "B",
	"1100": "C",
	"1101": "D",
	"1110": "E",
	"1111": "F",
}

var octalSystemMap = map[int]string{
	0: "000",
	1: "001",
	2: "010",
	3: "011",
	4: "100",
	5: "101",
	6: "110",
	7: "111",
}
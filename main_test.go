package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin_AllPositiveNumbers(t *testing.T) {
	r := min(4, 7)
	assert.Equal(t, 4, r, `Expected 4 as minumum of 4 and 7`)
}

func TestMin_OneNegativeNumber(t *testing.T) {
	r := min(-14, 7)
	assert.Equal(t, -14, r, `Expected -14 as minumum of -14 and 7`)
}

func TestProcessFile_FullTable(t *testing.T) {
	ProcessFile("a.raw")
	assert.Nil(t, EqualFiles("a_answer.xcbr", "a.xcbr"))
}

func TestProcessFile_MissingFields(t *testing.T) {
	ProcessFile("b.raw")
	assert.Nil(t, EqualFiles("b_answer.xcbr", "b.xcbr"))
}

func TestProcessFile_RealExample(t *testing.T) {
	ProcessFile("testcases/AF6HO-20231017-191821-471.raw")
	assert.Nil(t, EqualFiles("testcases/AF6HO-20231017-191821-471_answer.xcbr", "testcases/AF6HO-20231017-191821-471.xcbr"))
}

func TestIdentifyTableColumns_RightALignedColumns(t *testing.T) {
	newQsoLines := make([][]byte, 0)
	for i := 0; i < len(qsoLines); i++ {
		newQsoLines = append(newQsoLines, []byte(qsoLines[i]))
	}
	result := identifyTableColumns(newQsoLines)
	expectedResult := []int{0, 5, 11, 14, 25, 30, 43, 46, 51, 62, 67, 76}
	fmt.Printf("%v\n", result)
	assert.Equal(t, expectedResult, result)
}

func EqualFiles(f1, f2 string) error {
	f1Content, err := CabRead(f1)
	if err != nil {
		return fmt.Errorf("Failed to read f1 file")
	}
	f2Content, err := CabRead(f2)
	if err != nil {
		return fmt.Errorf("Failed to read f2 file")
	}

	return CabEqual(f1Content, f2Content)
}

var qsoLines = []string{
	//   0....5...10...15...20...25...30...35...40...45...50...55...60...65...70...75...80
	"QSO: 28405 PH 2023-10-07 1655 AF6HO         1 SCLA K5TR         60 TEXAS",
	"QSO: 21329 PH 2023-10-07 1656 AF6HO         2 SCLA K5TR         63 TEXAS",
	"QSO: 28400 PH 2023-10-07 1706 AF6HO         0 SCLA IK4GR           DX",
	"QSO: 28425 PH 2023-10-07 1707 AF6HO         3 SCLA K4RO         93 SCLA",
	"QSO: 28437 PH 2023-10-07 1708 AF6HO         5 SCLA N6TV        130 SCLA",
	"QSO: 21366 PH 2023-10-07 1713 AF6HO         0 SCLA W1AW/5          NM",
	"QSO: 21361 PH 2023-10-07 1716 AF6HO         7 SCLA K2KR         38 CO",
	"QSO: 21293 PH 2023-10-07 1721 AF6HO         8 SCLA KI6QDH       28 WY",
	"QSO: 21305 PH 2023-10-08 1905 AF6HO         9 SCLA OM2VL       524 DX",
	"QSO: 21351 PH 2023-10-08 1907 AF6HO        10 SCLA K6XX       2416 SCRU",
	"QSO: 21371 PH 2023-10-08 1908 AF6HO        11 SCLA W0BH        610 KANSAS",
	//   0....5...10...15...20...25...30...35...40...45...50...55...60...65...70...75...80
	"QSO: 28410 PH 2023-10-08 1913 AF6HO        12 SCLA NS4X        183 TENNESSEE",
	"QSO: 28448 PH 2023-10-08 1914 AF6HO        13 SCLA K6MM       1030 SCLA",
	"QSO: 28455 PH 2023-10-08 1916 AF6HO        14 SCLA W5CW        544 OKLAHOMA",
	"QSO: 28460 PH 2023-10-08 1917 AF6HO        15 SCLA K6MMM      1489 SCRU",
	"QSO: 28491 PH 2023-10-08 1918 AF6HO        16 SCLA K0EJ        806 TENNESSEE",
	"QSO: 28444 PH 2023-10-08 1923 AF6HO        17 SCLA W4KW        302 TENNESSEE",
	"QSO: 28440 PH 2023-10-08 1924 AF6HO        18 SCLA NW0M        603 MO",
	"QSO: 28422 PH 2023-10-08 1926 AF6HO        19 SCLA W6TCP      3480 ALAM",
}

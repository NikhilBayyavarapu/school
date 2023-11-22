package students

import (
	"strconv"

	"log"
)

type Student struct {
	SID         int
	Fname       string
	Lname       string
	Parent      string
	Acadyear    string
	Class       int
	Section     string
	Busfee      string
	Tutionfee   string
	Totalfee    string
	Totalmonths int
	Montharray  []string
	Remfee      string
}

func NewStudent(SID int, fname string, lname string, parent string, acadyear string, class int, section string, busfee string, tutionfee string, totalmonths int) *Student {
	busfeeFloat, err := strconv.ParseFloat(busfee, 64)
	if err != nil {
		log.Fatal(err)
	}

	tutionfeeFloat, err := strconv.ParseFloat(tutionfee, 64)
	if err != nil {
		log.Fatal(err)
	}

	totalfeeFloat := busfeeFloat + tutionfeeFloat
	totalfee := strconv.FormatFloat(totalfeeFloat, 'f', -1, 64)

	montharray := make([]string, totalmonths)

	permonth := float64(totalfeeFloat) / float64(totalmonths)

	permonthString := strconv.FormatFloat(float64(permonth), 'f', -1, 64)

	for i := range montharray {
		montharray[i] = permonthString
	}

	return &Student{
		SID:         SID,
		Fname:       fname,
		Lname:       lname,
		Parent:      parent,
		Acadyear:    acadyear,
		Class:       class,
		Section:     section,
		Tutionfee:   tutionfee,
		Totalfee:    totalfee,
		Totalmonths: totalmonths,
		Montharray:  montharray,
		Remfee:      totalfee,
	}
}

func (st *Student) PayFee(amount string) Student {

	i := 0
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatal(err)
	}

	for amountFloat >= 0 && i < st.Totalmonths {
		if st.Montharray[i] == "" {
			continue
		}
		monthArrayFloat, err := strconv.ParseFloat(st.Montharray[i], 64)
		if err != nil {
			log.Fatal(err)
		}
		if amountFloat > monthArrayFloat {
			amountFloat -= monthArrayFloat
			st.Montharray[i] = ""
			i++
		} else {
			monthArrayFloat -= amountFloat
			st.Montharray[i] = strconv.FormatFloat(monthArrayFloat, 'f', -1, 64)
			break
		}
	}

	var floatValFinal float64
	for _, val := range st.Montharray {

		if val == "" {
			continue
		} else {
			floatVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			floatValFinal += floatVal
		}
	}

	st.Remfee = strconv.FormatFloat(floatValFinal, 'f', -1, 64)

	return *st
}

// RR3TuipRLnX3M5pY
// root

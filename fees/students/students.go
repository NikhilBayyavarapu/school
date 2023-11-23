package students

import (
	"strconv"

	"log"
)

type Student struct {
	SID         int      `bson:"SID"`
	Fname       string   `bson:"Fname"`
	Lname       string   `bson:"Lname"`
	Parent      string   `bson:"Parent"`
	Contact     string   `bson:"Contact"`
	Acadyear    string   `bson:"Acadyear"`
	Class       int      `bson:"Class"`
	Section     string   `bson:"Section"`
	Busfee      string   `bson:"Busfee"`
	Tutionfee   string   `bson:"Tutionfee"`
	Totalfee    string   `bson:"Totalfee"`
	Totalmonths int      `bson:"Totalmonths"`
	Montharray  []string `bson:"Montharray"`
	Remfee      string   `bson:"Remfee"`
}

func NewStudent(SID int, fname string, lname string, parent string, contact string, acadyear string, class int, section string, busfee string, tutionfee string, totalmonths int) *Student {
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
		Contact:     contact,
		Acadyear:    acadyear,
		Class:       class,
		Section:     section,
		Busfee:      busfee,
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

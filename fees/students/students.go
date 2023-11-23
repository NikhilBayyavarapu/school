package students

type Student struct {
	SID         int       `bson:"SID"`
	Fname       string    `bson:"Fname"`
	Lname       string    `bson:"Lname"`
	Parent      string    `bson:"Parent"`
	Contact     string    `bson:"Contact"`
	Acadyear    string    `bson:"Acadyear"`
	Class       int       `bson:"Class"`
	Section     string    `bson:"Section"`
	Busfee      float32   `bson:"Busfee"`
	Tutionfee   float32   `bson:"Tutionfee"`
	Totalfee    float32   `bson:"Totalfee"`
	Totalmonths int       `bson:"Totalmonths"`
	Montharray  []float32 `bson:"Montharray"`
	Remfee      float32   `bson:"Remfee"`
}

func NewStudent(SID int, fname string, lname string, parent string, contact string, acadyear string, class int, section string, busfee float32, tutionfee float32, totalmonths int) *Student {

	totalfee := busfee + tutionfee

	montharray := make([]float32, totalmonths)

	permonth := float32(totalfee) / float32(totalmonths)

	for i := range montharray {
		montharray[i] = permonth
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

func (st *Student) PayFee(amount float32) Student {

	i := 0
	// amountFloat, err := strconv.ParseFloat(amount, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for amount > 0 && i < st.Totalmonths {
		if st.Montharray[i] <= 0 {
			continue
		}
		// monthArrayFloat, err := strconv.ParseFloat(st.Montharray[i], 64)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		if amount > st.Montharray[i] {
			amount -= st.Montharray[i]
			st.Montharray[i] = 0
			i++
		} else {
			st.Montharray[i] -= amount
			break
		}
	}

	var floatValFinal float32
	for _, val := range st.Montharray {

		if val <= 0 {
			continue
		} else {
			floatValFinal += val
		}
	}

	st.Remfee = floatValFinal

	return *st
}

// RR3TuipRLnX3M5pY
// root

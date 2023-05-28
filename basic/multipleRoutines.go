package basic

import (
	"fmt"
	"net/http"
	"time"
)

type rss struct{
	Name string
	Url string
}

func (r *rss) read(){
	t1 := time.Now()
	
	res, err := http.Get(r.Url)
	if(err != nil){
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	t2 := time.Since(t1).Seconds() // ---> Calcula en segundos cuanto tiempo pasó desde que se inició t1 hasta que se llegó acá
	fmt.Printf("Transcurred time: %v. Rss: %v \n", t2, r.Name)
}

func MultipleRoutines(){
	espn := new(rss)
	espn.Name = "ESPN"
	espn.Url = "http://www.espn.com/espn/rss/news"

	nyTimes := new(rss)
	nyTimes.Name = "NY Times"
	nyTimes.Url = "http://rss.nytimes.com/services/xml/rss/nyt/World.xml"
	rssList := [...]*rss{espn, nyTimes} // ---> slice
	

	for _, v := range rssList {
		go v.read()
	}
	time.Sleep(time.Second * 5)
}
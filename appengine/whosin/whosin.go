//App ID:	673350062694474
//App Secret:	736400b3b54625da28e9a250500e5ca8

package whosin

import (
	"appengine"
	"appengine/datastore"
	//	"appengine/user"

	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"
	//"time"
)

// type Inviter struct {
// 	Name string `json:"name"`
// }

type Joiner struct {
	JoinerGender             string `json:"joiner_gender"`
	JoinerRelationshipStatus string `json:"joiner_relationship_status"`
	JoinerName               string `json:"joiner_name"`
	JoinerLink               string `json:"joiner_link"`
	JoinerBirthday           string `json:"joiner_birthday"`
}

type Outing struct {
	Name           string `json:"name"`
	Location       string `json:"location"`
	LocationLat    string `json:"location_lat"`     //float64
	LocationLong   string `json:"location_long"`    //float64
	TimeFrameStart string `json:"time_frame_start"` //time.Time
	TimeFrameEnd   string `json:"time_frame_end"`   //time.Time
	OutingType     string `json:"outing_type"`
	JoinerAgeStart string `json:"joiner_age_start"` //int
	JoinerAgeEnd   string `json:"joiner_age_end"`   //int
	JoinCutOff     string `json:"join_cut_off"`     //time.Time
	//Inviter        Inviter   `json:"inviter"`
	Joiners []Joiner `json:"joiners"`

	InviterGender             string `json:"inviter_gender"`
	InviterRelationshipStatus string `json:"inviter_relationship_status"`
	InviterName               string `json:"inviter_name"`
	InviterLink               string `json:"inviter_link"`
	InviterBirthday           string `json:"inviter_birthday"`
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login/facebook", handler)
	http.HandleFunc("/outings", get_outings)
	http.HandleFunc("/outing/new", post_outing)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func get_outings(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Outing").Limit(100)
	outings := make([]Outing, 0, 100)
	if _, err := q.GetAll(c, &outings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(outings)
	fmt.Fprintf(w, "%s", json)
	return
}

func post_outing(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	outing := Outing{
		InviterBirthday:           r.FormValue("inviter_birthday"),
		InviterGender:             r.FormValue("inviter_gender"),
		InviterLink:               r.FormValue("inviter_link"),
		InviterName:               r.FormValue("inviter_name"),
		InviterRelationshipStatus: r.FormValue("inviter_relationship_status"),
		JoinerAgeStart:            r.FormValue("joiner_age_start"), //.(int),
		JoinerAgeEnd:              r.FormValue("joiner_age_end"),   //.(int),
		Location:                  r.FormValue("location"),
		LocationLat:               r.FormValue("location_lat"),  //.(float64),
		LocationLong:              r.FormValue("location_long"), //.(float64),
		OutingType:                r.FormValue("outing_type"),
		TimeFrameStart:            r.FormValue("time_frame_start"),
		TimeFrameEnd:              r.FormValue("time_frame_end"),
	}

	c.Debugf("outing received: %v", outing)

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Outing", nil), &outing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{}")
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"models"
	"HighFive"
	"log"
	"strconv"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
	"github.com/satori/go.uuid"
)
// Structs created for the Meeting, Races and Competitors which
// are used to store and parse data to and from the API via the web page.
type Competitor struct {
	PosNum int `json:"position"`
	Type   string `json:"type"`
	Name string `json:"name"`
}

type Race struct {
	Location string `json:"location"`
	Competitors []Competitor `json:"competitors"`
	Close int `json:"close"`
	CloseDate time.Time `json:"closeDate"`
	RaceType string `json:"racetype"`
	RaceID string `json:"id"`
}

type Meeting struct {
	Location string `json:"location"`
	Race []Race `json:"races"`
}

// Loads main page
func Index(c * gin.Context) {
	c.HTML(200, "index.html", gin.H{
	})
}

type ResponseData struct {
	Message string `json:"message"`
}

type FiveData struct {
	Message []string `json:"message"`
}

// Gets the next five races from the database and sends the data to the web page
func NextFive(five HighFive.HighFive) func (c * gin.Context) {
	return func(c *gin.Context) {
		// Retrieves next 5 races with the shortest time to close from the current time. Does not retrieve races already completed.
		allRaces, err := models.Races(five.DB, qm.Where("closing_time > ?", time.Now().UTC()), qm.OrderBy("closing_time"), qm.Limit(5)).All()

		if err != nil {
			c.JSON(200, "No races")
			return
		}

		var races []Race

		// Sets up races slice which will be sent to web page
		for i := 0; i < 5; i++ {
			nanos := allRaces[i].ClosingTime.UnixNano()
			millis := int(nanos/1000000)

			race1 := Race{
				RaceType: allRaces[i].RaceType,
				CloseDate: allRaces[i].ClosingTime,
				Close:       millis,
				RaceID: allRaces[i].RaceID,
			}
			races = append(races, race1)
		}

		// Creates meeting which has location and contains the races for the location
		meets := Meeting{
			Location: "Gold Coast",
			Race:     races,
		}

		c.JSON(200, meets)
	}
}

// Loads a web page containing the searched for race data
func ShowRace(c * gin.Context) {
	id := c.Param("id")
	log.Println(id)
	c.HTML(200, "race_data.html", gin.H{})
}

// Retrieves the information for the required race
func GetRace(five HighFive.HighFive) func (c *gin.Context) {
	return func(c *gin.Context) {
		form, _ := c.GetPostForm("search_key")

		raceUUID, _ := uuid.FromString(form)

		// Queries the database for the race that is being searched
		// Gets the location for that race, the race times and type, and all the competitors of that race
		raceDetails, _ := models.Races(five.DB, qm.Where("race_id =?", raceUUID)).All()
		meetingDetails, _ := models.Meetings(five.DB, qm.Where("meeting_id =?", raceDetails[0].MeetingID)).All()
		allCompetitors, _ := models.Competitors(five.DB, qm.Where("race_id = ?", form)).All()
		compCount, _ := models.Competitors(five.DB, qm.Where("race_id = ?", form)).Count()

		// Sets up time remaining until race closes
		nanos := raceDetails[0].ClosingTime.UnixNano()
		millis := int(nanos/1000000)

		var compArr []Competitor

		// Set up competitor slice for the race
		for i := 0; i < int(compCount); i++{
			comp1 := Competitor {
				PosNum: allCompetitors[i].Position,
				Type: allCompetitors[i].Type,
				Name: "Racer " + strconv.Itoa(i),
			}

			compArr = append(compArr, comp1)
		}

		// Creates race object to send race details and the competitors slice
		race1 := Race {
			Location : meetingDetails[0].MeetingName,
			Competitors: compArr,
			CloseDate: raceDetails[0].ClosingTime,
			Close: millis,
		}

		c.JSON(200, race1)
	}
}
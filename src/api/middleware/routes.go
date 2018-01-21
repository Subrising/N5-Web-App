package middleware

import (
//	"time"

	"github.com/gin-gonic/gin"
	"models"
	"net/http"
	"HighFive"
	"fmt"
	//"os"
	"log"
	//"io"
	"strconv"
	//"math/rand"
	//"path/filepath"
	//"math/rand"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
	"github.com/satori/go.uuid"
)

type Competitor struct {
	PosNum int `json:"position"`
	Type   string `json:"type"`
	Name string `json:"name"`
}

type Race struct {
	Location string `json:"location"`
	Competitors []Competitor `json:"competitors"`
	Close int `json:"close"`
	RaceType string `json:"racetype"`
	RaceID string `json:"id"`
}

type Meeting struct {
	Location string `json:"location"`
	Race []Race `json:"races"`
}

// Loads index upload page
func Index(c * gin.Context) {
	c.HTML(200, "index.html", gin.H{
	})
}

// Test page to return changed values in index page
func GetFive(five HighFive.HighFive) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main website",
			"owners": models.Meetings(five.DB).OneP(),
		})
	}
}

type ResponseData struct {
	Message string `json:"message"`
}

type FiveData struct {
	Message []string `json:"message"`
}

func ReceiveAjax(five HighFive.HighFive) func (c * gin.Context) {
	return func(c *gin.Context) {
		form := c.PostForm("ajax_post_data")
		log.Println(form)
		fmt.Println("Receive ajax post data string ", form)

		//var races []&models.Race
		//races, err := models.Races().All()

		now := time.Now().UTC()

		//timeAdd := time.Second * 60 * 1

		//now.Add(timeAdd)
		log.Println("Current time = ", now)

		allRaces, err := models.Races(five.DB, qm.Where("closing_time > ?", time.Now().UTC()), qm.OrderBy("closing_time"), qm.Limit(5)).All()

		log.Println("Query works")

		if err != nil {
			c.JSON(200, "No races")
			return
		}


		// THIS IS HOW YOU FUCKING GET THE DATA FROM THE ROWSS!!! DO NOT FORGET!!!!
		log.Println(allRaces[0].RaceID)
		log.Println(allRaces[0].RaceType)
		log.Println(allRaces[0].ClosingTime)

		/* firstClose := time.Now();
		nanos := firstClose.UnixNano()
		millis := int(nanos/1000000) + (rand.Intn(100000) + 1000)

		comp1 := Competitor{
			PosNum: 1,
			Type:   "Greyhound",
		}

		compArr := []Competitor{comp1}

		race1 := Race{
			Location:    "Location 1",
			Competitors: compArr,
			Close:       millis,
		}

		secondClose := time.Now();
		nanos = secondClose.UnixNano()
		millis = int(nanos/1000000) + (rand.Intn(100000) + 1000)

		race2 := Race{
			Location:    "Location 2",
			Competitors: compArr,
			Close:       millis,
		}*/

		var races []Race

		for i := 0; i < 5; i++ {
			nanos := allRaces[i].ClosingTime.UnixNano()
			millis := int(nanos/1000000)

			race1 := Race{
				RaceType: allRaces[i].RaceType,
				Close:       millis,
				RaceID: allRaces[i].RaceID,
			}
			races = append(races, race1)
		}

		meets := Meeting{
			Location: "Gold Coast",
			Race:     races,
		}

		c.JSON(200, meets)
	}
}

func SendResults(five HighFive.HighFive) func (c *gin.Context) {
	return func (c * gin.Context) {
		form := c.PostForm("search_key")
		log.Println("In send results")
		log.Println(form)
		fmt.Println("Receive get query ", form)
		strArr := []string{"Hello1", "Hello2", "Hello3"}
		resp := FiveData {
			Message: strArr,
		}
		c.JSON(200, resp)
	}
}

func ShowRace(c * gin.Context) {
	id := c.Param("id")
	log.Println(id)
	c.HTML(200, "race_data.html", gin.H{})
}

func GetRace(five HighFive.HighFive) func (c *gin.Context) {
	return func(c *gin.Context) {
		form, _ := c.GetPostForm("search_key")
		log.Println("In send results")
		log.Println("Received form ", form)
		//fmt.Println("Receive get query ", form)

		raceUUID, _ := uuid.FromString(form)

		log.Println(form)

		raceDetails, _ := models.Races(five.DB, qm.Where("race_id =?", raceUUID)).All()

		log.Println("Got race details")

		log.Println(raceDetails)

		meetingDetails, _ := models.Meetings(five.DB, qm.Where("meeting_id =?", raceDetails[0].MeetingID)).All()

		log.Println("Got meeting details")

		allCompetitors, _ := models.Competitors(five.DB, qm.Where("race_id = ?", form)).All()

		log.Println("Got competitor details")

		compCount, _ := models.Competitors(five.DB, qm.Where("race_id = ?", form)).Count()

		log.Println("Got comp count")

		fmt.Println("Finished queries")

		nanos := raceDetails[0].ClosingTime.UnixNano()
		millis := int(nanos/1000000)

		var compArr []Competitor

		for i := 0; i < int(compCount); i++{
			comp1 := Competitor {
				PosNum: allCompetitors[i].Position,
				Type: allCompetitors[i].Type,
				Name: "Racer " + strconv.Itoa(i),
			}

			compArr = append(compArr, comp1)
		}

		/*comp1 := Competitor{
			PosNum: 1,
			Type: "Greyhound",
		}

		compArr := []Competitor{comp1}
*/
		race1 := Race {
			Location : meetingDetails[0].MeetingName,
			Competitors: compArr,
			Close: millis,
		}

		c.JSON(200, race1)
	}
}

/*func SendFiles(dex Dex.Dex) func (c *gin.Context) {
	return *gin.Context{}
}*/

/*
make own ginErr
form file give it  a request, parameter, extension, directory path
same with custom parameter values
*/

/*
Check sha256 of uploaded file to sha256 in document database
If not exists, then create
If exists, do not upload
Remember to set all database statuses for documents and pages
 */

/*
Use jQuery Ajax for html page templating functions
 */

/*
Store things to return to user html page via JSON structs
 */

/*
Do document type checks with % for the highest type return
 */

/*
jQuery stuff, have table updated with document and owner
if user clicks on document, they are taken to new page with document page data
 */
package middleware

import (
"time"

"github.com/gin-gonic/gin"
"models"
"net/http"
"HighFive"
"fmt"
//"os"
"log"
//"io"
//"strconv"
//"math/rand"
//"path/filepath"
)

// Loads index upload page
func Index(c * gin.Context) {
	c.HTML(200, "index.html", gin.H{
	})
}

// Loads test index upload page
func Basic(c *gin.Context) {
	c.HTML(200, "index.tmpl.html", gin.H{
	})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	if len(nick) < 2 {
		nick = ""
	}
	if len(nick) > 13 {
		nick = nick[0:12] + "..."
	}
	c.HTML(200, "room_login.templ.html", gin.H{
		"roomid":    roomid,
		"nick":      nick,
		"timestamp": time.Now().Unix(),
	})
}
/*
func roomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	message := c.PostForm("message")
	message = strings.TrimSpace(message)

	validMessage := len(message) > 1 && len(message) < 200
	validNick := len(nick) > 1 && len(nick) < 14
	if !validMessage || !validNick {
		c.JSON(400, gin.H{
			"status": "failed",
			"error":  "the message or nickname is too long",
		})
		return
	}

	post := gin.H{
		"nick":    html.EscapeString(nick),
		"message": html.EscapeString(message),
	}
	messages.Add("inbound", 1)
	room(roomid).Submit(post)
	c.JSON(200, post)
}
*/

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

func ReceiveAjax(c * gin.Context) {
	form := c.PostForm("ajax_post_data")
	log.Println(form)
	fmt.Println("Receive ajax post data string ", form)

	resp := ResponseData{
		Message: "This is a test response",
	}

	c.JSON(200, resp)
}

func SendResults(five HighFive.HighFive) func (c *gin.Context) {
	return func (c * gin.Context) {
		form := c.PostForm("search_key")
		log.Println("In send results")
		log.Println(form)
		fmt.Println("Receive get query ", form)

		resp := ResponseData{
			Message: "This is a test get",
		}

		c.JSON(200, resp)
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
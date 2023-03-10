package app

import (
	"stash-scrapers/common/sqlite"
	"stash-scrapers/services/log"
	"time"
)

func getPerformerList() []*Performers {
	conn := sqlite.Conn()
	var list []*Performers
	err := conn.Select(&list, "select id,name from performers where gender ISNULL")
	if err != nil {
		log.Println("get performer list fail:", err)
		return nil
	}
	return list
}

func getNoneImagePerformerList() []*Performers {
	conn := sqlite.Conn()
	var list []*Performers
	err := conn.Select(&list, "select id,name from performers p LEFT JOIN performers_image pi ON pi.performer_id = p.id where pi.image ISNULL")
	if err != nil {
		log.Println("get performer list fail:", err)
		return nil
	}
	return list
}

// update performer metadata
func updatePerformer(performer *Performers) {
	conn := sqlite.Conn()
	_, err := conn.Exec("update performers set gender = ?,twitter = ?,birthdate = ?,ethnicity = ?,country = ?,height = ?,measurements = ?,career_length = ?,aliases = ?,updated_at = ? where id = ?",
		performer.Gender, performer.Twitter, performer.Birthdate, performer.Ethnicity, performer.Country, performer.Height, performer.Measurements, performer.CareerLength, performer.Aliases, time.Now().Local().Format("2006-01-02T15:04:05+08:00"), performer.ID)
	if err != nil {
		log.Println("insert performer metadata fail:", err)
		return
	}
}

// update performer updated_at
func updatePerformerDate(id int) {
	conn := sqlite.Conn()
	_, err := conn.Exec("update performers set updated_at = ? where id = ?", time.Now().Local().Format("2006-01-02T15:04:05+08:00"), id)
	if err != nil {
		log.Println("insert performer metadata fail:", err)
		return
	}
}

// save performer Image
func savePerformerImage(performerImage *PerformersImage) {
	if len(performerImage.Image) < 10 {
		return
	}
	conn := sqlite.Conn()
	_, err := conn.Exec("INSERT INTO performers_image (performer_id,image)VALUES(?,?)",
		performerImage.PerformerID, performerImage.Image)
	if err != nil {
		log.Println("insert performer image fail:", err)
		return
	}
}

func updatePerformerImage(performerImage *PerformersImage) {
	if len(performerImage.Image) < 10 {
		return
	}
	conn := sqlite.Conn()
	_, err := conn.Exec("update performers_image set image = ? where performer_id = ?",
		performerImage.Image, performerImage.PerformerID)
	if err != nil {
		log.Println("update performer image fail:", err)
		return
	}
}

func checkPerformerImage(performerImage *PerformersImage) int {
	conn := sqlite.Conn()
	var counts int
	err := conn.Get(&counts, "select Count(1) from performers_image where performer_id = ?", performerImage.PerformerID)
	if err != nil {
		log.Println("check performer image fail:", err)
		return 0
	}
	return counts
}

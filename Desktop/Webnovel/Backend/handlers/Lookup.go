package handlers

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	DictEN = "en"
	DictVI = "vi"
)

func Lookup(c *gin.Context) {
	dict := c.DefaultQuery("dict", DictEN)
	words, exists := c.GetQuery("words")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'words' is missing"})
		return
	}

	db, err := getDBConnection(dict)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	entry, err := queryEntry(db, words, dict)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func getDBConnection(dict string) (*sql.DB, error) {
	switch dict {
	case DictEN:
		return models.EnDictDB, nil
	case DictVI:
		return models.ViDictDB, nil
	default:
		return nil, fmt.Errorf("unsupported dictionary: %s", dict)
	}
}

func queryEntry(db *sql.DB, word string, dict string) (models.DictEntry, error) {
	var entry models.DictEntry

	query := buildQuery(dict)
	row := db.QueryRow(query, word)

	var (
		compound sql.NullString
		measure  sql.NullString
		snym     sql.NullString
		content  sql.NullString
		cnvi     sql.NullString
		topic    sql.NullString
	)

	if dict == DictVI {
		if err := row.Scan(&entry.Word, &entry.Pinyin, &entry.Phonetic, &compound, &entry.Kind, &measure, &snym, &content, &topic, &cnvi); err != nil {
			return entry, err
		}
	} else {
		if err := row.Scan(&entry.Word, &entry.Pinyin, &entry.Phonetic, &compound, &entry.Kind, &measure, &snym, &content, &topic); err != nil {
			return entry, err
		}
	}

	json.Unmarshal([]byte(compound.String), &entry.Compound)
	json.Unmarshal([]byte(measure.String), &entry.Measure)
	json.Unmarshal([]byte(snym.String), &entry.Snym)
	json.Unmarshal([]byte(content.String), &entry.Content)
	entry.CNVI = cnvi.String
	if topic.Valid {
		topics := strings.Split(topic.String, ",")
		for _, t := range topics {
			t = strings.TrimSpace(t)
			entry.Topics = append(entry.Topics, models.Topic{Id: t})
		}
	}
	return populateExamplesAndTopic(entry, db)
}

func buildQuery(dict string) string {
	if dict == DictVI {
		return "SELECT word, pinyin, phonetic, compound, kind, measure, snym, content, topic, cn_vi FROM cnvi WHERE word = $1"
	}
	return "SELECT word, pinyin, phonetic, compound, kind, measure, snym, content, topic FROM cnvi WHERE word = $1"
}

func populateExamplesAndTopic(entry models.DictEntry, db *sql.DB) (models.DictEntry, error) {
	for i := range entry.Content[0].Means {
		for j := range entry.Content[0].Means[i].Examples {
			row := db.QueryRow("SELECT id,e,m,p FROM e_cnvi WHERE id=$1", entry.Content[0].Means[i].Examples[j])
			if err := row.Err(); err != nil {
				return entry, err
			}

			var e models.Example
			row.Scan(&e.Id, &e.E, &e.M, &e.P)
			entry.Examples = append(entry.Examples, e)
		}
	}
	for i := range entry.Topics {
		row := db.QueryRow("SELECT path FROM topic WHERE id=$1", entry.Topics[i].Id)
		if err := row.Err(); err != nil {
			return entry, err
		}
		row.Scan(&entry.Topics[i].Name)
	}
	log.Printf("entry: %+v\n", entry)
	return entry, nil
}

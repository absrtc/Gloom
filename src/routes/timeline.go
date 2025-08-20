package routes

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Version struct {
	Season int
	Lobby string
}

func VersionInfo() Version {
	season := 10
	return Version{
		Season: season,
		Lobby: "LobbySeason" + strconv.Itoa(season),
	}
}

func FutureDate(hours int) string {
	return time.Now().Add(time.Duration(hours) * time.Hour).UTC().Format(time.RFC3339)
}

func HandleTimeline(c *gin.Context) {
	version := VersionInfo()
	future := FutureDate(24)
	now := time.Now().UTC().Format(time.RFC3339)
	activeEvents := []map[string]string{
		{
			"eventType":   "EventFlag.Season" + strconv.Itoa(version.Season),
			"activeSince": "2025-05-01T00:00:00.000Z",
			"activeUntil": future,
		},
		{
			"eventType":   "EventFlag." + version.Lobby,
			"activeSince": "2025-05-01T00:00:00.000Z",
			"activeUntil": future,
		},
	}

	season := "AthenaSeason:athenaseason" + strconv.Itoa(version.Season)

	c.JSON(200, gin.H{
		"channels": gin.H{
			"client-matchmaking": gin.H{
				"states":      []interface{}{},
				"cacheExpire": future,
			},
			"client-events": gin.H{
				"states": []gin.H{
					{
						"validFrom":    "0001-01-01T00:00:00.000Z",
						"activeEvents": activeEvents,
						"state": gin.H{
							"activeStorefronts":  []interface{}{},
							"eventNamedWeights":  gin.H{},
							"seasonNumber":       version.Season,
							"seasonTemplateId":   season,
							"matchXpBonusPoints": 0,
							"seasonBegin":        "2025-06-29T00:00:00.000Z",
							"seasonEnd":          future,
							"seasonDisplayedEnd": future,
							"weeklyStoreEnd":     future,
							"stwEventStoreEnd":   future,
							"stwWeeklyStoreEnd":  future,
							"sectionStoreEnds": gin.H{
								"Featured": future,
								"Daily":    future,
							},
							"dailyStoreEnd": future,
						},
					},
				},
				"cacheExpire": future,
			},
		},
		"eventsTimeOffsetHrs": 0,
		"cacheIntervalMins":   10,
		"currentTime": now,
	})
}
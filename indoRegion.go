package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Region struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ParentID int    `json:"parent_id"`
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:asdfQWER789@localhost/regions?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/regions", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, type, parent_id FROM regions ORDER BY type ASC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		regions := make(map[int]Region)
		for rows.Next() {
			var region Region
			if err := rows.Scan(&region.ID, &region.Name, &region.Type, &region.ParentID); err != nil {
				log.Fatal(err)
			}
			regions[region.ID] = region
		}
		for _, region := range regions {
			if region.Type != "country" {
				parentRegion := regions[region.ParentID]
				parentRegionChildren := parentRegion.Children()
				parentRegionChildren = append(parentRegionChildren, region)
				parentRegion.SetChildren(parentRegionChildren)
				regions[parentRegion.ID] = parentRegion
			}
		}
		country, _ := regions[1].MarshalJSON()
		w.Header().Set("Content-Type", "application/json")
		w.Write(country)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (r Region) Children() []Region {
	var children []Region
	return children
}

func (r *Region) SetChildren(children []Region) {
	r.Children = children
}

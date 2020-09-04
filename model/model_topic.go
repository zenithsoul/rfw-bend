package model

import (
	"context"
	"encoding/json"

	dbConn "rfw-bend/config/database"

	"github.com/arangodb/go-driver"
)

// Topic -
type Topic struct {
	TopicContent string `json:"topic"`
	Replys       []Reply
}

// Reply -
type Reply struct {
	ReplyContent string `json:"reply"`
	Subreplys    []Subreply
}

// Subreply -
type Subreply struct {
	SubreplyContent string `json:"subreply"`
}

func getTopic() string {

	conn := dbConn.ArangoDBConnect()
	getCursor, _, _ := conn.NewQuery(`
		FOR d IN d_content

			FILTER d.type == "topic"
			
			LET d_reply = 
			(
				FOR vector_reply, egde_reply, path_reply IN OUTBOUND d._id e_content
				FILTER path_reply.edge[*].label ALL == "reply"
				
				LET d_subreply = 
				(
					FOR vector_subreply, egde_subreply, path_subreply IN OUTBOUND vector_reply._id e_content
					FILTER egde_subreply.edge[*].label ALL == "subreply"
					
					RETURN {subreply : vector_subreply.content}
				)
				
				RETURN {reply :vector_reply.content , subreplys : d_subreply}
			)
		
		RETURN { topic : d.content , replys : d_reply}
	`)

	defer getCursor.Close()

	result := []Topic{}

	for {
		getData := Topic{}
		_, err := getCursor.ReadDocument(context.Background(), &getData)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}

		result = append(result, getData)
	}

	jsonResult, _ := json.Marshal(result)

	return string(jsonResult)
}

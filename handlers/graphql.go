package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"net/http"
	"restApiProject/database"
	"restApiProject/models"
)

// GraphqlGetItem POST /graphql
func GraphqlGetItem(db database.ItemsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		rootQuery := graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getItem": getItem(db),
			},
		})
		rootMutation := graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"insertItem": insertItem(db),
			},
		})
		schema, _ := graphql.NewSchema(graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		})

		// get query string
		requestString := c.Query("q")
		// get request body
		if requestString == "" {
			var body map[string]interface{}
			err := c.BindJSON(&body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			requestString = body["query"].(string)
		}

		res := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: requestString,
		})

		c.JSON(http.StatusOK, res)
	}
}

var itemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemType",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"sku":  &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"type": &graphql.Field{Type: graphql.String},
		"cost": &graphql.Field{Type: graphql.Float},
	},
})

func getItem(db database.ItemsInterface) *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.Int},
	}
	return &graphql.Field{
		Name:        "Item",
		Description: "get item by id",
		Type:        itemType,
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return db.GetItemById(id)
		},
	}
}

func insertItem(db database.ItemsInterface) *graphql.Field {
	args := graphql.FieldConfigArgument{
		"sku": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"type": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"cost": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	}
	return &graphql.Field{
		Name:        "Insert item",
		Type:        itemType,
		Description: "insert item",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			item := models.Item{}
			body, err := json.Marshal(p.Args)
			if err != nil {
				return "", err
			}
			err = json.Unmarshal(body, &item)
			if err != nil {
				return "", err
			}
			return db.CreateItem(item)
		},
	}
}

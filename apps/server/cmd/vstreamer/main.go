package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/altierawr/vstreamer"
	"github.com/altierawr/vstreamer/ent"
	"github.com/altierawr/vstreamer/ent/migrate"
	"github.com/altierawr/vstreamer/hooks"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Register hooks
	hooks.RegisterLibraryHooks(ctx, client)
	hooks.RegisterPlaySessionHooks(ctx, client)

	router := gin.Default()
	router.GET("/session/:sessionid/stream/:streamid/:file", func(c *gin.Context) {
		sessionId := c.Param("sessionid")
		streamId := c.Param("streamid")
		file := c.Param("file")

		fmt.Printf("%s, %s, %s\n", sessionId, streamId, file)
		c.String(http.StatusOK, "yello")
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(vstreamer.NewSchema(client))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.Use(extension.Introspection{})
	srv.Use(entgql.Transactioner{TxOpener: client})
	http.Handle("/",
		playground.Handler("VStreamer", "/query"),
	)
	http.Handle("/query", c.Handler(srv))
	go func() {
		router.Run(":8080")
	}()
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

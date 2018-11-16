package main

import (
	"context"
	"log"

	"github.com/foosball-hq/foosball.com/internal"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/foosball-hq/foosball.com/foosball"
	"github.com/jinzhu/gorm"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

var db *gorm.DB
var sess session.Session
var svc kmsiface.KMSAPI

func init() {
	sess := session.Must(session.NewSession())
	svc = kms.New(sess)

	var err error
	db, err = internal.CreateDatabaseConnection(svc, internal.DatabaseConnectionOptions{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}

func handle(ctx context.Context) error {
	return foosball.CreateUser(db)
}

func main() {
	lambda.Start(handle)
}

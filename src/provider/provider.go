package provider

import (
	"context"

	"github.com/aravindanve/gomeet-server/src/client"
	"github.com/aravindanve/gomeet-server/src/config"
	"github.com/aravindanve/gomeet-server/src/resource"
	"go.mongodb.org/mongo-driver/mongo"
)

type Provider interface {
	config.Config
	client.MongoClientProvider
	client.GoogleOAuth2ClientProvider
	resource.UserCollectionProvider
	resource.AuthCollectionProvider
	resource.MeetingCollectionProvider
	resource.ParticipantCollectionProvider
	Release(ctx context.Context)
}

type provider struct {
	config.Config
	mongoClient           *mongo.Client
	googleOAuth2Client    client.GoogleOAuth2Client
	authCollection        *resource.AuthCollection
	userCollection        *resource.UserCollection
	meetingCollection     *resource.MeetingCollection
	participantCollection *resource.ParticipantCollection
}

func NewProvider(ctx context.Context) Provider {
	cf := config.NewConfig()

	mongoClient := client.NewMongoClient(ctx, cf)
	mongoDB := client.GetMongoDatabaseDefault(mongoClient, cf)

	return &provider{
		Config:                cf,
		mongoClient:           mongoClient,
		googleOAuth2Client:    client.NewGoogleOAuth2Client(cf),
		authCollection:        resource.NewAuthCollection(ctx, mongoDB),
		userCollection:        resource.NewUserCollection(ctx, mongoDB),
		meetingCollection:     resource.NewMeetingCollection(ctx, mongoDB),
		participantCollection: resource.NewParticipantCollection(ctx, mongoDB),
	}
}

func (p *provider) Release(ctx context.Context) {
	p.mongoClient.Disconnect(ctx)
}

func (p *provider) MongoClient() *mongo.Client {
	return p.mongoClient
}

func (p *provider) GoogleOAuth2Client() client.GoogleOAuth2Client {
	return p.googleOAuth2Client
}

func (p *provider) AuthCollection() *resource.AuthCollection {
	return p.authCollection
}

func (p *provider) UserCollection() *resource.UserCollection {
	return p.userCollection
}

func (p *provider) MeetingCollection() *resource.MeetingCollection {
	return p.meetingCollection
}

func (p *provider) ParticipantCollection() *resource.ParticipantCollection {
	return p.participantCollection
}

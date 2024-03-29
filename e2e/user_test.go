package main_test

import (
	"context"
	"fmt"
	"sync"

	"github.com/aravindanve/livemeet-server/src/provider"
	"github.com/aravindanve/livemeet-server/src/resource"
)

var mockUser *resource.User
var mockUserMut sync.Mutex

func getMockUser() resource.User {
	mockUserMut.Lock()
	if mockUser == nil {
		// create mock user
		ctx := context.Background()
		p := provider.NewProvider(ctx)
		defer p.Release(ctx)

		mockUserImageURL := "https://example.com/image.jpg"
		mockUser = &resource.User{
			Name:               "Mock User",
			ImageURL:           &mockUserImageURL,
			Provider:           resource.UserProvider_Google,
			ProviderResourceID: "some-google-id",
		}

		err := p.UserCollection().Save(ctx, mockUser)
		if err != nil {
			panic(fmt.Sprintf("error saving user: %s", err.Error()))
		}
	}
	mockUserMut.Unlock()
	return *mockUser
}

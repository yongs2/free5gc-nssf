/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 */

package client

import (
	"context"
	"encoding/json"
    "fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

    . "free5gc-nssf/model"
	"free5gc-nssf/test"
	"free5gc-nssf/util/http2"
)

var testingNotification = test.TestingNssaiavailability{
	ConfigFile:             test.ConfigFileFromArgs,
    MuteLogInd:             test.MuteLogIndFromArgs,
	NfNssaiAvailabilityUri: "https://localhost:8080/notification",
}

func generateNotificationRequest() NssfEventNotification {
	const jsonRequest = `
        {
            "subscriptionId": "1",
            "authorizedNssaiAvailabilityData": [
                {
                    "tai": {
                        "plmnId": {
                            "mcc": "466",
                            "mnc": "92"
                        },
                        "tac": "33456"
                    },
                    "supportedSnssaiList": [
                        {
                            "sst": 1
                        },
                        {
                            "sst": 1,
                            "sd": "1"
                        },
                        {
                            "sst": 1,
                            "sd": "2"
                        },
                        {
                            "sst": 2
                        }
                    ]
                }
            ]
        }
    `

	var n NssfEventNotification
	if err := json.NewDecoder(strings.NewReader(jsonRequest)).Decode(&n); err != nil {
		fmt.Printf("Decode error: %v", err)
	}

	return n
}

func TestNotificationPost(t *testing.T) {
	var (
		requestBody string
	)

	// Create a server to accept testing requests
	router := gin.Default()
	router.POST("/notification", func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			t.Errorf(err.Error())
		}
		// Remove NL line feed, new line character
		requestBody = string(buf[:len(buf)-1])

		c.JSON(http.StatusNoContent, gin.H{})
	})

	srv, err := http2.NewServer(":8080", "../../nssfsslkey.log", router)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		err := srv.ListenAndServeTLS("../../support/tls/nssf.pem", "../../support/tls/nssf.key")
		if err != nil && err != http.ErrServerClosed {
			t.Fatal(err)
		}
	}()

	configuration := NewConfiguration()
	configuration.SetBasePathNoGroup(testingNotification.NfNssaiAvailabilityUri)
	apiClient := NewAPIClient(configuration)

	subtests := []struct {
		name                string
		generateRequestBody func() NssfEventNotification
		expectRequestBody   string
	}{
		{
			name:                "Notify",
			generateRequestBody: generateNotificationRequest,
			expectRequestBody: `{"subscriptionId":"1","authorizedNssaiAvailabilityData":[{` +
				`"tai":{"plmnId":{"mcc":"466","mnc":"92"},"tac":"33456"},` +
				`"supportedSnssaiList":[{"sst":1},{"sst":1,"sd":"1"},{"sst":1,"sd":"2"},{"sst":2}]}]}`,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			var (
				n    NssfEventNotification
				resp *http.Response
			)

			// Start to generate and send notification request after channel is closed
			if subtest.generateRequestBody != nil {
				n = subtest.generateRequestBody()
			}

			resp, err := apiClient.NotificationApi.NotificationPost(context.Background(), n)

			if err != nil {
				t.Errorf(err.Error())
			}

			if resp.StatusCode != http.StatusNoContent {
				t.Errorf("Incorrect status code: expected %d, got %d", http.StatusNoContent, resp.StatusCode)
			}

			if requestBody != subtest.expectRequestBody {
				t.Errorf("Incorrect request body:\nexpected\n%s\n, got\n%s", subtest.expectRequestBody, requestBody)
			}

			err = srv.Shutdown(context.Background())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

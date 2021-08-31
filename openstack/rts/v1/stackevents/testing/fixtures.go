package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/rts/v1/stackevents"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

// ListExpected represents the expected object from a List request.
var ListExpected = []stackevents.Event{
	{
		ResourceName: "hello_world",
		Time:         time.Date(2015, 2, 5, 21, 33, 11, 0, time.UTC),
		Links: []golangsdk.Link{
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world/events/06feb26f-9298-4a9b-8749-9d770e5d577a",
				Rel:  "self",
			},
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
				Rel:  "resource",
			},
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
				Rel:  "stack",
			},
		},
		LogicalResourceID:    "hello_world",
		ResourceStatusReason: "state changed",
		ResourceStatus:       "CREATE_IN_PROGRESS",
		PhysicalResourceID:   "",
		ID:                   "06feb26f-9298-4a9b-8749-9d770e5d577a",
	},
	{
		ResourceName: "hello_world",
		Time:         time.Date(2015, 2, 5, 21, 33, 27, 0, time.UTC),
		Links: []golangsdk.Link{
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world/events/93940999-7d40-44ae-8de4-19624e7b8d18",
				Rel:  "self",
			},
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
				Rel:  "resource",
			},
			{
				Href: "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
				Rel:  "stack",
			},
		},
		LogicalResourceID:    "hello_world",
		ResourceStatusReason: "state changed",
		ResourceStatus:       "CREATE_COMPLETE",
		PhysicalResourceID:   "49181cd6-169a-4130-9455-31185bbfc5bf",
		ID:                   "93940999-7d40-44ae-8de4-19624e7b8d18",
	},
}

// ListOutput represents the response body from a List request.
const ListOutput = `
{
  "events": [
  {
    "resource_name": "hello_world",
    "event_time": "2015-02-05T21:33:11",
    "links": [
    {
      "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world/events/06feb26f-9298-4a9b-8749-9d770e5d577a",
      "rel": "self"
    },
    {
      "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
      "rel": "resource"
    },
    {
      "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
      "rel": "stack"
    }
    ],
    "logical_resource_id": "hello_world",
    "resource_status_reason": "state changed",
    "resource_status": "CREATE_IN_PROGRESS",
    "physical_resource_id": null,
    "id": "06feb26f-9298-4a9b-8749-9d770e5d577a"
    },
    {
      "resource_name": "hello_world",
      "event_time": "2015-02-05T21:33:27",
      "links": [
      {
        "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world/events/93940999-7d40-44ae-8de4-19624e7b8d18",
        "rel": "self"
      },
      {
        "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b/resources/hello_world",
        "rel": "resource"
      },
      {
        "href": "http://166.78.160.107:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/5f57cff9-93fc-424e-9f78-df0515e7f48b",
        "rel": "stack"
      }
      ],
      "logical_resource_id": "hello_world",
      "resource_status_reason": "state changed",
      "resource_status": "CREATE_COMPLETE",
      "physical_resource_id": "49181cd6-169a-4130-9455-31185bbfc5bf",
      "id": "93940999-7d40-44ae-8de4-19624e7b8d18"
    }
  ]
}`

// HandleListSuccessfully creates an HTTP handler at `/stacks/hello_world/49181cd6-169a-4130-9455-31185bbfc5bf/events`
// on the test handler mux that responds with a `List` response.
func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/stacks/hello_world/49181cd6-169a-4130-9455-31185bbfc5bf/events", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, output)
		case "93940999-7d40-44ae-8de4-19624e7b8d18":
			fmt.Fprintf(w, `{"events":[]}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

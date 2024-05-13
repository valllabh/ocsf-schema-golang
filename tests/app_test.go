package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/events/application"
	appEnums "github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/events/application/enums"
	"github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/objects"
	objectEnums "github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/objects/enums"
	"github.com/valllabh/ocsf-schema-golang/ocsfjson"
	ocsfvalidate_1_2 "github.com/valllabh/ocsf-schema-golang/ocsfvalidate/1.2"
)

func TestAppplicationWebResourcesActivity(t *testing.T) {
	e := &application.WebResourcesActivity{
		Time:        time.Now().UTC().UnixMilli(),
		CategoryUid: appEnums.WEB_RESOURCES_ACTIVITY_CATEGORY_UID_WEB_RESOURCES_ACTIVITY_CATEGORY_UID_APPLICATION_ACTIVITY,
		ClassUid:    appEnums.WEB_RESOURCES_ACTIVITY_CLASS_UID_WEB_RESOURCES_ACTIVITY_CLASS_UID_WEB_RESOURCES_ACTIVITY,
		ActivityId:  appEnums.WEB_RESOURCES_ACTIVITY_ACTIVITY_ID_WEB_RESOURCES_ACTIVITY_ACTIVITY_ID_CREATE,
		SeverityId:  appEnums.WEB_RESOURCES_ACTIVITY_SEVERITY_ID_WEB_RESOURCES_ACTIVITY_SEVERITY_ID_INFORMATIONAL,
		ActionId:    appEnums.WEB_RESOURCES_ACTIVITY_ACTION_ID_WEB_RESOURCES_ACTIVITY_ACTION_ID_ALLOWED,
		TypeUid:     appEnums.WEB_RESOURCES_ACTIVITY_TYPE_UID_WEB_RESOURCES_ACTIVITY_TYPE_UID_WEB_RESOURCES_ACTIVITY_CREATE,
		WebResources: []*objects.WebResource{
			{
				UrlString: "https://api.example.com/v1/users",
			},
		},

		Actor: &objects.Actor{
			User: &objects.User{
				Uid:       "5431",
				Name:      "Alice",
				TypeId:    objectEnums.USER_TYPE_ID_USER_TYPE_ID_USER,
				Type:      "user",
				EmailAddr: "alice@example.com",
			},
		},
		HttpRequest: &objects.HttpRequest{
			HttpMethod: "POST",
			Url: &objects.Url{
				Scheme:   "https",
				Hostname: "api.example.com",
				Path:     "/v1/users",
			},
		},
		HttpResponse: &objects.HttpResponse{
			ContentType: "application/json",
			Code:        200,
			Status:      http.StatusText(http.StatusOK),
		},
		Cloud: &objects.Cloud{
			Account: &objects.Account{
				Uid:    "1234",
				Name:   "Test Account",
				Type:   "test",
				TypeId: objectEnums.ACCOUNT_TYPE_ID_ACCOUNT_TYPE_ID_OTHER,
			},
			Provider: "ExampleCloud",
		},
		SrcEndpoint: &objects.NetworkEndpoint{
			Ip: "1.2.3.4",
		},
		Metadata: &objects.Metadata{
			Version: "1.2.0",
			Product: &objects.Product{
				VendorName: "Example",
			},
		},
	}

	vedata, err := ocsfjson.Marshal(e)
	require.NoError(t, err)
	err = ocsfvalidate_1_2.Validate("web_resources_activity", vedata)
	require.NoError(t, err)
	rt := &application.WebResourcesActivity{}
	err = ocsfjson.Unmarshal(vedata, rt)
	require.NoError(t, err)
	require.EqualExportedValues(t, e, rt)
}

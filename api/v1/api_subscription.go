package v1

import (
	"github.com/gin-gonic/gin"
	"ockham-api/api/v1/middleware"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"strconv"
	"time"
)

// ViewSubscription
// @Summary View Subscription
// @Description Subscribe your nodes by link
// @Tags subscription
// @Success 201 {object} util.Pack
// @Failure 401,409,500 {object} util.Pack
// @Param client query string true "Client Type"
// @Param token query string true "Subscription Token"
// @Param id path uint true "Subscription ID"
// @Router /v1/subscriptions/{id}/subscribe [POST]
func ViewSubscription(c *gin.Context) {
	_ = c.Param("id")
}

type Traffic struct {
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Enabled           bool      `json:"enabled"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	TotalTrafficBytes int64     `json:"total_traffic_bytes"`
	UsedTrafficBytes  int64     `json:"used_traffic_bytes"`
	SystemPriority    int       `json:"system_priority"`
	UserPriority      int       `json:"user_priority"`
	AdminPriority     int       `json:"admin_priority"`
}

type Subscription struct {
	Title                          string    `json:"title"`
	Description                    string    `json:"description"`
	Enabled                        bool      `json:"enabled"`
	StartTime                      time.Time `json:"start_time"`
	EndTime                        time.Time `json:"end_time"`
	BundledTrafficSubscription     Traffic   `json:"bundled_traffic_subscription"`
	AdditionalTrafficSubscriptions []Traffic `json:"additional_traffic_subscriptions"`
}

// GetSubscription
// @Summary Get Subscription
// @Description List subscriptions
// @Tags subscription
// @Success 201 {object} util.Pack
// @Failure 401,409,500 {object} util.Pack
// @Param id path uint true "Subscription ID"
// @Router /v1/subscriptions/{id} [POST]
func GetSubscription(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	sub := model.ServicePlanSubscription{}
	database.Get[model.ServicePlanSubscription](uint(id), &sub)

	bt := model.TrafficPlanSubscription{}
	database.DBConn.First(&bt, &model.TrafficPlanSubscription{
		Bundled:                   true,
		ServicePlanSubscriptionID: sub.ID,
	})

	ts := make([]model.TrafficPlanSubscription, 0)
	database.DBConn.Find(&ts, &model.TrafficPlanSubscription{
		Bundled:                   false,
		ServicePlanSubscriptionID: sub.ID,
	})

	additionalTrafficBytes := make([]Traffic, 0)
	for _, t := range ts {
		additionalTrafficBytes = append(additionalTrafficBytes, Traffic{
			Title:             t.SubscriptionTitle,
			Description:       t.SubscriptionDescription,
			Enabled:           t.SubscriptionEnabled,
			StartTime:         t.SubscriptionStartTime,
			EndTime:           t.SubscriptionEndTime,
			TotalTrafficBytes: t.TotalTrafficBytes,
			UsedTrafficBytes:  t.UsedTrafficBytes,
			SystemPriority:    t.SystemPriority,
			UserPriority:      t.UserPriority,
			AdminPriority:     t.AdminPriority,
		})
	}

	util.SuccessPack(c).WithData(
		Subscription{
			Title:       sub.SubscriptionTitle,
			Description: sub.SubscriptionDescription,
			Enabled:     sub.SubscriptionEnabled,
			StartTime:   sub.SubscriptionStartTime,
			EndTime:     sub.SubscriptionEndTime,
			BundledTrafficSubscription: Traffic{
				Title:             bt.SubscriptionTitle,
				Description:       bt.SubscriptionDescription,
				Enabled:           bt.SubscriptionEnabled,
				StartTime:         bt.SubscriptionStartTime,
				EndTime:           bt.SubscriptionEndTime,
				TotalTrafficBytes: bt.TotalTrafficBytes,
				UsedTrafficBytes:  bt.UsedTrafficBytes,
				SystemPriority:    bt.SystemPriority,
				UserPriority:      bt.UserPriority,
				AdminPriority:     bt.AdminPriority,
			},
			AdditionalTrafficSubscriptions: additionalTrafficBytes,
		},
	).Responds()
}

// ListSubscriptions
// @Summary List Subscriptions
// @Description List subscriptions
// @Tags subscription
// @Success 201 {object} util.Pack
// @Failure 401,409,500 {object} util.Pack
// @Router /v1/users/me/subscriptions [POST]
func ListSubscriptions(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)

	subs := make([]model.ServicePlanSubscription, 0)
	database.DBConn.Find(&subs, model.ServicePlanSubscription{
		UserID: currentUser.ID,
	})

	respSubs := make([]Subscription, 0)

	for _, sub := range subs {
		bt := model.TrafficPlanSubscription{}
		database.DBConn.First(&bt, &model.TrafficPlanSubscription{
			Bundled:                   true,
			ServicePlanSubscriptionID: sub.ID,
		})

		ts := make([]model.TrafficPlanSubscription, 0)
		database.DBConn.Find(&ts, &model.TrafficPlanSubscription{
			Bundled:                   false,
			ServicePlanSubscriptionID: sub.ID,
		})

		additionalTrafficBytes := make([]Traffic, 0)
		for _, t := range ts {
			additionalTrafficBytes = append(additionalTrafficBytes, Traffic{
				Title:             t.SubscriptionTitle,
				Description:       t.SubscriptionDescription,
				Enabled:           t.SubscriptionEnabled,
				StartTime:         t.SubscriptionStartTime,
				EndTime:           t.SubscriptionEndTime,
				TotalTrafficBytes: t.TotalTrafficBytes,
				UsedTrafficBytes:  t.UsedTrafficBytes,
				SystemPriority:    t.SystemPriority,
				UserPriority:      t.UserPriority,
				AdminPriority:     t.AdminPriority,
			})
		}

		respSubs = append(respSubs, Subscription{
			Title:       sub.SubscriptionTitle,
			Description: sub.SubscriptionDescription,
			Enabled:     sub.SubscriptionEnabled,
			StartTime:   sub.SubscriptionStartTime,
			EndTime:     sub.SubscriptionEndTime,
			BundledTrafficSubscription: Traffic{
				Title:             bt.SubscriptionTitle,
				Description:       bt.SubscriptionDescription,
				Enabled:           bt.SubscriptionEnabled,
				StartTime:         bt.SubscriptionStartTime,
				EndTime:           bt.SubscriptionEndTime,
				TotalTrafficBytes: bt.TotalTrafficBytes,
				UsedTrafficBytes:  bt.UsedTrafficBytes,
				SystemPriority:    bt.SystemPriority,
				UserPriority:      bt.UserPriority,
				AdminPriority:     bt.AdminPriority,
			},
			AdditionalTrafficSubscriptions: additionalTrafficBytes,
		})
	}

	util.SuccessPack(c).WithData(respSubs).Responds()
}

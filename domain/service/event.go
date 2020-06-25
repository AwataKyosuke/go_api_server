package service

import "github.com/AwataKyosuke/go_api_server/domain/model"

// Filtered 開催方法で絞り込み
func Filtered(events []*model.Event, online bool, offline bool) []*model.Event {

	var ret []*model.Event

	for _, event := range events {

		if offline {
			if event.Position.Lat > 0 && event.Position.Lon > 0 {
				// オフライン開催のイベントは除外
				ret = append(ret, event)
			}
		}

		if online {
			if event.Position.Lat == 0 && event.Position.Lon == 0 {
				// オンライン開催のイベントは除外
				ret = append(ret, event)
			}
		}
	}

	return ret
}

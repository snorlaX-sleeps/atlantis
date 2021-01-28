// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	slack "github.com/nlopes/slack"
)

func AnySliceOfSlackChannel() []slack.Channel {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]slack.Channel))(nil)).Elem()))
	var nullValue []slack.Channel
	return nullValue
}

func EqSliceOfSlackChannel(value []slack.Channel) []slack.Channel {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []slack.Channel
	return nullValue
}

func NotEqSliceOfSlackChannel(value []slack.Channel) []slack.Channel {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue []slack.Channel
	return nullValue
}

func SliceOfSlackChannelThat(matcher pegomock.ArgumentMatcher) []slack.Channel {
	pegomock.RegisterMatcher(matcher)
	var nullValue []slack.Channel
	return nullValue
}

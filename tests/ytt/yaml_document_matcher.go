package ytt

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type YAMLDocumentMatcher struct {
	name string
	matcher  types.GomegaMatcher
	rendered string
}

func WithDocument(name string, matcher types.GomegaMatcher) *YAMLDocumentMatcher {
	return &YAMLDocumentMatcher{name, matcher, ""}
}

func (matcher *YAMLDocumentMatcher) Match(actual interface{}) (bool, error) {
	docsMap, ok := actual.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("YAMLDocument must be passed a map[string]interface{}. Got\n%s", format.Object(actual, 1))
	}

	doc := docsMap[matcher.name]
	return matcher.matcher.Match(doc)
}

func (matcher *YAMLDocumentMatcher) FailureMessage(actual interface{}) string {
	msg := fmt.Sprintf(
		"FailureMessage: There is a problem with this YAML:\n\n%s\n\n%s",
		matcher.rendered,
		matcher.matcher.FailureMessage(actual),
	)
	return msg
}

func (matcher *YAMLDocumentMatcher) NegatedFailureMessage(actual interface{}) string {
	msg := fmt.Sprintf(
		"NegatedFailureMessage: There is a problem with this YAML:\n\n%s\n\n%s",
		matcher.rendered,
		matcher.matcher.NegatedFailureMessage(actual),
	)
	return msg
}
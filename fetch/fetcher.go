package internal

import (
	graphql "github.com/utr1903/newrelic-tracker-internal/graphql"
)

const (
	FETCHER_GRAPHQL_HAS_RETURNED_ERRORS = "graphql has returned errors"
)

func Fetch(
	gqlc graphql.IGraphQlClient,
	qv any,
	res any,
) error {
	err := gqlc.Execute(qv, res)
	if err != nil {
		return err
	}
	return nil
}

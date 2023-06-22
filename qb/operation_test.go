package qb

import "testing"

type operationTest struct {
	operation *Operation
	expected  string
}

var operationTests = []operationTest{
	// General query building
	{
		NewQuery("").Fields(
			NewField("hero").Fields(
				NewField("name"),
			),
		),
		"query {hero {name}}",
	},
	{
		NewQuery("").Fields(
			NewField("hero").Fields(
				NewField("name"),
				NewField("friends").Fields(
					NewField("name"),
				),
			),
		),
		"query {hero {name friends {name}}}",
	},

	// Arguments test
	{
		NewQuery("").Fields(
			NewField("human").Args(NewArg("id", "\"1000\"")).Fields(
				NewField("name"),
				NewField("height"),
			),
		),
		"query {human(id: \"1000\") {name height}}",
	},
	{
		NewQuery("").Fields(
			NewField("human").Args(NewArg("id", "1000").Wrap()).Fields(
				NewField("name"),
				NewField("height"),
			),
		),
		"query {human(id: \"1000\") {name height}}",
	},
	{
		NewQuery("").Fields(
			NewField("human").Args(NewArg("id", "1000").Wrap()).Fields(
				NewField("name"),
				NewField("height").Args(NewArg("unit", "FOOT")),
			),
		),
		"query {human(id: \"1000\") {name height(unit: FOOT)}}",
	},

	// Aliases
	{
		NewQuery("").Fields(
			NewField("hero").Args(NewArg("episode", "EMPIRE")).Alias("empireHero").Fields(
				NewField("name"),
			),
			NewField("hero").Args(NewArg("episode", "JEDI")).Alias("jediHero").Fields(
				NewField("name"),
			),
		),
		"query {empireHero: hero(episode: EMPIRE) {name} jediHero: hero(episode: JEDI) {name}}",
	},
	{
		NewQuery("").Fields(
			NewField("hero").Alias("empireHero").Args(NewArg("episode", "EMPIRE")).Fields(
				NewField("name"),
			),
			NewField("hero").Alias("jediHero").Args(NewArg("episode", "JEDI")).Fields(
				NewField("name"),
			),
		),
		"query {empireHero: hero(episode: EMPIRE) {name} jediHero: hero(episode: JEDI) {name}}",
	},

	// Fragments
	{
		func() *Operation {
			fragment := NewFragment("comparisonFields", "Character").Fields(
				NewField("name"),
				NewField("appearsIn"),
				NewField("friends").Fields(
					NewField("name"),
				),
			)

			query := NewQuery("").Fields(
				NewField("hero").Args(NewArg("episode", "EMPIRE")).Alias("leftComparison").Fields(
					fragment.ToField(),
				),
				NewField("hero").Args(NewArg("episode", "JEDI")).Alias("rightComparison").Fields(
					fragment.ToField(),
				),
			)

			query.Fragments(
				fragment,
			)

			return query
		}(),
		"query {leftComparison: hero(episode: EMPIRE) {...comparisonFields} rightComparison: hero(episode: JEDI) {...comparisonFields}} fragment comparisonFields on Character {name appearsIn friends {name}}",
	},

	// Variables
	{
		func() *Operation {
			variable := NewVar("$first", "Int").Default("3")

			fragment := NewFragment("comparisonFields", "Character").Fields(
				NewField("name"),
				NewField("friendsConnection").Args(variable.ToArg("first")).Fields(
					NewField("totalCount"),
				),
			)

			query := NewQuery("HeroComparison")

			query.Vars(
				variable,
			)

			query.Fields(
				NewField("hero").Args(NewArg("episode", "EMPIRE")).Alias("leftComparison").Fields(
					fragment.ToField(),
				),
				NewField("hero").Args(NewArg("episode", "JEDI")).Alias("rightComparison").Fields(
					fragment.ToField(),
				),
			)

			query.Fragments(
				fragment,
			)

			return query
		}(),
		"query HeroComparison($first: Int = 3) {leftComparison: hero(episode: EMPIRE) {...comparisonFields} rightComparison: hero(episode: JEDI) {...comparisonFields}} fragment comparisonFields on Character {name friendsConnection(first: $first) {totalCount}}",
	},
	{
		func() *Operation {
			variable := NewVar("$episode", "Episode")

			query := NewQuery("HeroNameAndFriends")

			query.Vars(
				variable,
			)

			query.Fields(
				NewField("hero").Args(variable.ToArg("episode")).Fields(
					NewField("name"),
					NewField("friends").Fields(
						NewField("name"),
					),
				),
			)

			return query
		}(),
		"query HeroNameAndFriends($episode: Episode) {hero(episode: $episode) {name friends {name}}}",
	},
	{
		NewQuery("HeroNameAndFriends").Vars(
			NewVar("$episode", "Episode"),
		).Fields(
			NewField("hero").Args(NewArg("episode", "$episode")).Fields(
				NewField("name"),
				NewField("friends").Fields(
					NewField("name"),
				),
			),
		),
		"query HeroNameAndFriends($episode: Episode) {hero(episode: $episode) {name friends {name}}}",
	},

	// Directives
	{
		func() *Operation {
			episode := NewVar("$episode", "Episode")
			withFriends := NewVar("$withFriends", "Boolean!")

			query := NewQuery("Hero")

			query.Vars(
				episode,
				withFriends,
			)

			query.Fields(
				NewField("hero").Args(episode.ToArg("episode")).Fields(
					NewField("name"),
					NewField("friends").IncludeIf("$withFriends").Fields(
						NewField("name"),
					),
				),
			)

			return query
		}(),
		"query Hero($episode: Episode, $withFriends: Boolean!) {hero(episode: $episode) {name friends @include(if: $withFriends) {name}}}",
	},

	// Cycles
	{
		func() *Operation {
			name := NewField("name")

			return NewQuery("Hero").Fields(
				name.Fields(
					NewField("friends").Fields(
						name,
					),
				),
			)
		}(),
		ErrCycle.Error(),
	},
	{
		func() *Operation {
			name := NewField("name")

			return NewQuery("Hero").Fields(
				NewField("hero").Fields(
					name.Fields(
						name,
					),
				),
			)
		}(),
		ErrCycle.Error(),
	},
}

func TestOperationString(t *testing.T) {
	for _, test := range operationTests {
		s, err := test.operation.String()

		if err != nil && err.Error() == test.expected {
			continue
		}

		if s != test.expected {
			t.Errorf("Output \"%v\" not equal to expected \"%v\"", s, test.expected)
		}
	}
}

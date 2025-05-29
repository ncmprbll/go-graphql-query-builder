# GraphQL Query Builder
This is a GraphQL query builder which utilizes method chaining.
# Why?
Because all the known Golang GraphQL query builders looked "cumbersome" to me or were missing some of the functionality.

This library was partly inspired by [udacity/graphb](https://github.com/udacity/graphb). It offers 3 different ways of constructing a query and it also probably does some things better than my query builder does, so you might also want to check it out if you find it more suitable.

# Features
* Fields
  * Cycle detection
* Arguments
* Aliases
* Fragments
* Operation Name
* Variables
* Directives
* Mutations
* Inline Fragments

# Examples
## Example #1 - General query building

```graphql
query {
  me {
    name
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

func main() {
	query := qb.NewQuery("")

	query.Fields(
		qb.NewField("me").Fields(
			qb.NewField("name"),
		),
	)

	fmt.Println(query.PrettyString(2))
}
```
## Example #2 - Aliases and arguments

```graphql
query {
  empireHero: hero(episode: EMPIRE) {
    name
  }
  jediHero: hero(episode: JEDI) {
    name
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

func main() {
	query := qb.NewQuery("")

	query.Fields(
		qb.NewField("hero").Alias("empireHero").Args(qb.NewArg("episode", "EMPIRE")).Fields(
			qb.NewField("name"),
		),
		qb.NewField("hero").Alias("jediHero").Args(qb.NewArg("episode", "JEDI")).Fields(
			qb.NewField("name"),
		),
	)

	fmt.Println(query.PrettyString(2))
}
```

## Example #3 - Fragments

```graphql
query {
  leftComparison: hero(episode: EMPIRE) {
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI) {
    ...comparisonFields
  }
}

fragment comparisonFields on Character {
  name
  appearsIn
  friends {
    name
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

func main() {
	fragment := qb.NewFragment("comparisonFields", "Character").Fields(
		qb.NewField("name"),
		qb.NewField("appearsIn"),
		qb.NewField("friends").Fields(
			qb.NewField("name"),
		),
	)

	query := qb.NewQuery("").Fields(
		qb.NewField("hero").Args(qb.NewArg("episode", "EMPIRE")).Alias("leftComparison").Fields(
			fragment.ToField(),
		),
		qb.NewField("hero").Args(qb.NewArg("episode", "JEDI")).Alias("rightComparison").Fields(
			fragment.ToField(),
		),
	)

	query.Fragments(
		fragment,
	)

	fmt.Println(query.PrettyString(2))
}
```

## Example #4 - Variables

```graphql
query HeroComparison($first: Int = 3) {
  leftComparison: hero(episode: EMPIRE) {
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI) {
    ...comparisonFields
  }
}

fragment comparisonFields on Character {
  name
  friendsConnection(first: $first) {
    totalCount
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

func main() {
	variable := qb.NewVar("$first", "Int").Default("3")

	fragment := qb.NewFragment("comparisonFields", "Character").Fields(
		qb.NewField("name"),
		qb.NewField("friendsConnection").Args(variable.ToArg("first")).Fields(
			qb.NewField("totalCount"),
		),
	)

	query := qb.NewQuery("HeroComparison")

	query.Vars(
		variable,
	)

	query.Fields(
		qb.NewField("hero").Args(qb.NewArg("episode", "EMPIRE")).Alias("leftComparison").Fields(
			fragment.ToField(),
		),
		qb.NewField("hero").Args(qb.NewArg("episode", "JEDI")).Alias("rightComparison").Fields(
			fragment.ToField(),
		),
	)

	query.Fragments(
		fragment,
	)

	fmt.Println(query.PrettyString(2))
}
```

## Example #5 - Directives

```graphql
query Hero($episode: Episode, $withFriends: Boolean!) {
  hero(episode: $episode) {
    name
    friends @include(if: $withFriends) {
      name
    }
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

func main() {
	episode := qb.NewVar("$episode", "Episode")
	withFriends := qb.NewVar("$withFriends", "Boolean!")

	query := qb.NewQuery("Hero")

	query.Vars(
		episode,
		withFriends,
	)

	query.Fields(
		qb.NewField("hero").Args(episode.ToArg("episode")).Fields(
			qb.NewField("name"),
			qb.NewField("friends").IncludeIf("$withFriends").Fields(
				qb.NewField("name"),
			),
		),
	)

	fmt.Println(query.PrettyString(2))
}
```

## Example #6 - Cycles

```graphql
query {
  hero {
    friends {
      nemesis {
        friends
      }
    }
  }
}
```

```go
import (
	"fmt"
	"github.com/ncmprbll/go-graphql-query-builder/qb"
)

// Good
func main() {
	query := qb.NewQuery("")

	query.Fields(
		qb.NewField("hero").Fields(
			qb.NewField("friends").Fields(
				qb.NewField("nemesis").Fields(
					qb.NewField("friends"),
				),
			),
		),
	)

	s, err := query.PrettyString(2)

	if err != nil {
		// ...
	}

	// ...
}

// Also good
func main() {
	query := qb.NewQuery("")

	heroFriends := qb.NewField("friends")
	nemesisFriends := qb.NewField("friends")

	query.Fields(
		qb.NewField("hero").Fields(
			heroFriends.Fields(
				qb.NewField("nemesis").Fields(
					nemesisFriends,
				),
			),
		),
	)

	s, err := query.PrettyString(2)

	if err != nil {
		// ...
	}

	// ...
}

// Not good, but the program will not panic and instead just tell you that the cycle has been detected
func main() {
	query := qb.NewQuery("")

	friends := qb.NewField("friends")

	query.Fields(
		qb.NewField("hero").Fields(
			friends.Fields(
				qb.NewField("nemesis").Fields(
					friends, // Cycle!
				),
			),
		),
	)

	s, err := query.PrettyString(2)

	if err != nil {
		// ...
	}

	// ...
}
```

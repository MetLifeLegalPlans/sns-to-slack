package main

type Event struct {
	Name string

	Records []struct {
		Sns struct {
			Message string
		}
	}
}

func (e Event) Message() string {
	return e.Records[0].Sns.Message
}

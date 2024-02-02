package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pipeline mongo.Pipeline

func WithOrganisationLookup() func(*Pipeline) {
	return func(p *Pipeline) {
		*p = append(*p, organisationLookup()...)
	}
}

func WithSlotLookUp() func(*Pipeline) {
	return func(p *Pipeline) {
		*p = append(*p, slotLookup())
	}
}

func WithFilterShiftStatus(status string) func(*Pipeline) {
	return func(p *Pipeline) {
		if status != "" {
			*p = append(*p, filterShiftStatus(status)...)
		}
	}
}

func WithUsersLookup() func(*Pipeline) {
	return func(p *Pipeline) {
		*p = append(*p, usersLookUp()...)
	}
}

func organisationLookup() []bson.D {
	return []bson.D{
		{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "orgas"},
					{Key: "localField", Value: "organisationId"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "organisation"},
				},
			},
		},
		{{Key: "$unwind", Value: "$organisation"}},
	}
}

func createListTasks() bson.D {
	return bson.D{
		{Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: "$_id"},
				{Key: "name", Value: "$alias"},
				{Key: "organisation",
					Value: bson.D{
						{Key: "name", Value: "$organisation.name"},
						{Key: "address", Value: "$organisation.address"},
						{Key: "pictureurl", Value: "$organisation.logoUrl"},
					},
				},
				{Key: "slotsdata",
					Value: bson.D{
						{Key: "filled",
							Value: bson.D{
								{Key: "$sum",
									Value: bson.D{
										{Key: "$size",
											Value: bson.D{
												{Key: "$filter",
													Value: bson.D{
														{Key: "input", Value: "$slots"},
														{Key: "as", Value: "el"},
														{Key: "cond",
															Value: bson.D{
																{Key: "$eq",
																	Value: bson.A{
																		"$$el.status",
																		"filled",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						{Key: "total", Value: bson.D{{Key: "$size", Value: "$slots"}}},
					},
				},
			},
		},
	}
}

func createListTasks2() bson.D {
	return bson.D{
		{Key: "$project",
			Value: bson.D{
				{Key: "_id", Value: "$_id"},
				{Key: "name", Value: "$alias"},
				{Key: "userdata",
					Value: bson.D{
						{Key: "firstname", Value: "$users.profile.FirstName"},
						{Key: "lastname", Value: "$users.profile.LastName"},
					},
				},
			},
		},
	}
}

func slotLookup() bson.D {
	return bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "slots"},
				{Key: "localField", Value: "shiftIds"},
				{Key: "foreignField", Value: "shiftId"},
				{Key: "as", Value: "slots"},
			},
		},
	}
}

func usersLookUp() []bson.D {
	return []bson.D{
		{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "assigneeId"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "users"},
				},
			},
		},
		{{Key: "$unwind", Value: "$users"}},
	}
}

func shiftLookUp() bson.D {
	return bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "shifts"},
				{Key: "localField", Value: "shiftIds"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "shifts"},
			},
		},
	}
}

func filterShiftStatus(status string) []bson.D {
	ret := []bson.D{shiftLookUp()}
	var filter bson.D
	switch status {
	case "ongoing":
		filter = bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "shifts.time.startDate", Value: bson.D{{Key: "$gte", Value: "new Date()"}}},
				},
			}}
	case "upcoming":
		filter = bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "$or",
						Value: bson.A{
							bson.D{{Key: "shifts.time.startDate", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "gte", Value: "new Date()"}}}}}},
							bson.D{{Key: "shifts.time.endDate", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "lte", Value: "new Date()"}}}}}},
						},
					},
				},
			},
		}
	case "done":
		filter = bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "shifts.time.endDate", Value: bson.D{{Key: "$lte", Value: "new Date()"}}},
			},
		}}
	}

	ret = append(ret, filter)
	return ret
}

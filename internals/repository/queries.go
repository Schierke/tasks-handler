package repository

import "go.mongodb.org/mongo-driver/bson"

func organisationLookup() bson.D {
	return bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "orgas"},
				{Key: "localField", Value: "organisationId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "organisation"},
			},
		},
	}
}

func createListTasks() bson.D {
	return bson.D{
		{"$project",
			bson.D{
				{"_id", "$_id"},
				{"organisation",
					bson.D{
						{"name", "$organisation.name"},
						{"address", "$organisation.address"},
						{"pictureurl", "$organisation.logoUrl"},
					},
				},
				{"slotsdata",
					bson.D{
						{"filled",
							bson.D{
								{"$sum",
									bson.D{
										{"$size",
											bson.D{
												{"$filter",
													bson.D{
														{"input", "$slots"},
														{"as", "el"},
														{"cond",
															bson.D{
																{"$eq",
																	bson.A{
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
						{"total", bson.D{{"$size", "$slots"}}},
					},
				},
			},
		},
	}
}

func slotLookup() bson.D {
	return bson.D{
		{"$lookup",
			bson.D{
				{"from", "slots"},
				{"localField", "shiftIds"},
				{"foreignField", "shiftId"},
				{"as", "slots"},
			},
		},
	}
}

func filterShiftStatus(status string) bson.D {
	switch status {
	case "ongoing":
		return bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "shifts.time.startDate", Value: bson.D{{Key: "$gte", Value: "new Date()"}}},
				},
			}}
	case "upcoming":
		return bson.D{
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
		return bson.D{{
			Key: "$match",
			Value: bson.D{
				{Key: "shifts.time.endDate", Value: bson.D{{Key: "$lte", Value: "new Date()"}}},
			},
		}}
	}

	return bson.D{}
}

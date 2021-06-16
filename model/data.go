package model

var (
	coverCaption1 = "contoh image cover caption"
	statusDraft   = "draft"

	DataRequest = []ArticleRequest{
		{
			ID:                 1,
			Title:              "contoh title",
			Lead:               "contoh lead",
			Content:            "contoh content",
			ImageCoverURL:      "contoh image cover url",
			ImageCoverCaption:  &coverCaption1,
			Status:             &statusDraft,
			Tag1ID:             10,
			Tag1Name:           "nama tag 10",
			CategoryDetailID:   20,
			CategoryDetailName: "nama category 20",
		},
		{
			ID:                 3,
			Title:              "contoh title kedua",
			Lead:               "contoh leads",
			Content:            "contoh content kedua",
			ImageCoverURL:      "contoh image cover url",
			ImageCoverCaption:  nil,
			Status:             nil,
			Tag1ID:             11,
			Tag1Name:           "nama tag 11",
			CategoryDetailID:   21,
			CategoryDetailName: "nama category 21",
		},
	}
)

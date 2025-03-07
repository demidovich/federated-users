package person

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type usecase struct {
	Find findUsecase
	Save saveUsecase
}

func Usecase(db *sqlx.DB) usecase {
	return usecase{
		Find: findUsecase{db: db},
		Save: saveUsecase{db: db},
	}
}

type Person struct {
	Uuid           string `db:"uuid"`
	FederationUuid string `db:"federation_uuid"`
	Attrs          *Attrs `db:"attrs"`
}

type Attrs struct {
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         string `json:"age,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	Country     string `json:"country,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Lang        string `json:"lang,omitempty"`
	Address     string `json:"address,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Tag1        string `json:"tag_1,omitempty"`
	Tag2        string `json:"tag_2,omitempty"`
	Tag3        string `json:"tag_3,omitempty"`
	Tag4        string `json:"tag_4,omitempty"`
	Tag5        string `json:"tag_5,omitempty"`
	Tag6        string `json:"tag_6,omitempty"`
	Tag7        string `json:"tag_7,omitempty"`
	Tag8        string `json:"tag_8,omitempty"`
	Tag9        string `json:"tag_9,omitempty"`
	SocialId    string `json:"social_id,omitempty"`
	SocialId1   string `json:"social_id_1,omitempty"`
	SocialId2   string `json:"social_id_2,omitempty"`
	SocialId3   string `json:"social_id_3,omitempty"`
	SocialId4   string `json:"social_id_4,omitempty"`
	SocialId5   string `json:"social_id_5,omitempty"`
	SocialId6   string `json:"social_id_6,omitempty"`
	SocialId7   string `json:"social_id_7,omitempty"`
	SocialId8   string `json:"social_id_8,omitempty"`
	SocialId9   string `json:"social_id_9,omitempty"`
	Attribute1  string `json:" attribute_1,omitempty"`
	Attribute2  string `json:" attribute_2,omitempty"`
	Attribute3  string `json:" attribute_3,omitempty"`
	Attribute4  string `json:" attribute_4,omitempty"`
	Attribute5  string `json:" attribute_5,omitempty"`
	Attribute6  string `json:" attribute_6,omitempty"`
	Attribute7  string `json:" attribute_7,omitempty"`
	Attribute8  string `json:" attribute_8,omitempty"`
	Attribute9  string `json:" attribute_9,omitempty"`
	Attribute10 string `json:" attribute_10,omitempty"`
	Attribute11 string `json:" attribute_11,omitempty"`
	Attribute12 string `json:" attribute_12,omitempty"`
	Attribute13 string `json:" attribute_13,omitempty"`
	Attribute14 string `json:" attribute_14,omitempty"`
	Attribute15 string `json:" attribute_15,omitempty"`
	Attribute16 string `json:" attribute_16,omitempty"`
	Attribute17 string `json:" attribute_17,omitempty"`
	Attribute18 string `json:" attribute_18,omitempty"`
	Attribute19 string `json:" attribute_19,omitempty"`
	Attribute20 string `json:" attribute_20,omitempty"`
	Attribute21 string `json:" attribute_21,omitempty"`
	Attribute22 string `json:" attribute_22,omitempty"`
	Attribute23 string `json:" attribute_23,omitempty"`
	Attribute24 string `json:" attribute_24,omitempty"`
	Attribute25 string `json:" attribute_25,omitempty"`
	Attribute26 string `json:" attribute_26,omitempty"`
	Attribute27 string `json:" attribute_27,omitempty"`
	Attribute28 string `json:" attribute_28,omitempty"`
	Attribute29 string `json:" attribute_29,omitempty"`
	Attribute30 string `json:" attribute_30,omitempty"`
	Attribute31 string `json:" attribute_31,omitempty"`
	Attribute32 string `json:" attribute_32,omitempty"`
	Attribute33 string `json:" attribute_33,omitempty"`
	Attribute34 string `json:" attribute_34,omitempty"`
	Attribute35 string `json:" attribute_35,omitempty"`
	Attribute36 string `json:" attribute_36,omitempty"`
	Attribute37 string `json:" attribute_37,omitempty"`
	Attribute38 string `json:" attribute_38,omitempty"`
	Attribute39 string `json:" attribute_39,omitempty"`
	Attribute40 string `json:" attribute_40,omitempty"`
	Attribute41 string `json:" attribute_41,omitempty"`
	Attribute42 string `json:" attribute_42,omitempty"`
	Attribute43 string `json:" attribute_43,omitempty"`
	Attribute44 string `json:" attribute_44,omitempty"`
	Attribute45 string `json:" attribute_45,omitempty"`
	Attribute46 string `json:" attribute_46,omitempty"`
	Attribute47 string `json:" attribute_47,omitempty"`
	Attribute48 string `json:" attribute_48,omitempty"`
	Attribute49 string `json:" attribute_49,omitempty"`
	Attribute50 string `json:" attribute_50,omitempty"`
	Attribute51 string `json:" attribute_51,omitempty"`
	Attribute52 string `json:" attribute_52,omitempty"`
	Attribute53 string `json:" attribute_53,omitempty"`
	Attribute54 string `json:" attribute_54,omitempty"`
	Attribute55 string `json:" attribute_55,omitempty"`
	Attribute56 string `json:" attribute_56,omitempty"`
	Attribute57 string `json:" attribute_57,omitempty"`
	Attribute58 string `json:" attribute_58,omitempty"`
	Attribute59 string `json:" attribute_59,omitempty"`
	Attribute60 string `json:" attribute_60,omitempty"`
	Attribute61 string `json:" attribute_61,omitempty"`
	Attribute62 string `json:" attribute_62,omitempty"`
	Attribute63 string `json:" attribute_63,omitempty"`
	Attribute64 string `json:" attribute_64,omitempty"`
	Attribute65 string `json:" attribute_65,omitempty"`
	Attribute66 string `json:" attribute_66,omitempty"`
	Attribute67 string `json:" attribute_67,omitempty"`
	Attribute68 string `json:" attribute_68,omitempty"`
	Attribute69 string `json:" attribute_69,omitempty"`
	Attribute70 string `json:" attribute_70,omitempty"`
}

func (a Attrs) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

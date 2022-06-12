package plexapi

import (
	"time"
)

type MediaContainer struct {
	Size               int         `xml:"size,attr,omitempty"`
	Identifier         string      `xml:"identifier,attr,omitempty"`
	Title1             string      `xml:"title1,attr,omitempty"`
	Title2             string      `xml:"title2,attr,omitempty"`
	AllowSync          bool        `xml:"allowSync,attr,omitempty"`
	Content            string      `xml:"content,attr,omitempty"`
	LibrarySectionID   string      `xml:"librarySectionID,attr,omitempty"`
	LibrarySectionUUID string      `xml:"librarySectionUUID,attr,omitempty"`
	ViewGroup          string      `xml:"viewGroup,attr,omitempty"`
	ViewMode           int         `xml:"viewMode,attr,omitempty"`
	Directories        []Directory `xml:"Directory,omitempty"`
	Videos             []Video     `xml:"Video,omitempty"`
}

type Directory struct {
	Key              string     `xml:"key,attr,omitempty"`
	Title            string     `xml:"title,attr,omitempty"`
	Type             string     `xml:"type,attr,omitempty"`
	Agent            string     `xml:"agent,attr,omitempty"`
	Scanner          string     `xml:"scanner,attr,omitempty"`
	Language         string     `xml:"language,attr,omitempty"`
	UUID             string     `xml:"uuid,attr,omitempty"`
	UpdatedAt        int64      `xml:"updatedAt,attr,omitempty"`
	CreatedAt        int64      `xml:"createdAt,attr,omitempty"`
	ScannedAt        int64      `xml:"scannedAt,attr,omitempty"`
	Content          bool       `xml:"content,attr,omitempty"`
	Directory        bool       `xml:"directory,attr,omitempty"`
	ContentChangedAt int        `xml:"contentChangedAt,attr,omitempty"`
	Hidden           bool       `xml:"hidden,attr,omitempty"`
	Secondary        bool       `xml:"secondary,attr,omitempty"`
	Prompt           string     `xml:"prompt,attr,omitempty"`
	Search           bool       `xml:"search,attr,omitempty"`
	Location         []Location `xml:"location,omitempty"`
}

type Location struct {
	ID   string `xml:"id,attr,omitempty"`
	Path string `xml:"path,attr,omitempty"`
}

type Video struct {
	RatingKey              string  `xml:"ratingKey,attr,omitempty"`
	Key                    string  `xml:"key,attr,omitempty"`
	GUID                   string  `xml:"guid,attr,omitempty"`
	Studio                 string  `xml:"studio,attr,omitempty"`
	Type                   string  `xml:"type,attr,omitempty"`
	Title                  string  `xml:"title,attr,omitempty"`
	ContentRating          string  `xml:"contentRating,attr,omitempty"`
	Summary                string  `xml:"summary,attr,omitempty"`
	Rating                 float32 `xml:"rating,attr,omitempty"`
	AudienceRating         float32 `xml:"audienceRating,attr,omitempty"`
	Year                   int     `xml:"year,attr,omitempty"`
	Tagline                string  `xml:"tagline,attr,omitempty"`
	Duration               int64   `xml:"duration,attr,omitempty"`
	OriginallyAvailableAt  string  `xml:"originallyAvailableAt,attr,omitempty"`
	AddedAt                int64   `xml:"addedAt,attr,omitempty"`
	UpdatedAt              int64   `xml:"updatedAt,attr,omitempty"`
	ChapterSource          string  `xml:"chapterSource,attr,omitempty"`
	HasPremiumExtras       bool    `xml:"hasPremiumExtras,attr,omitempty"`
	HasPremiumPrimaryExtra bool    `xml:"hasPremiumPrimaryExtra,attr,omitempty"`
	Media                  []Media `xml:"Media,omitempty"`
	Genres                 []Tag   `xml:"Genre,omitempty"`
	Directors              []Tag   `xml:"Director,omitempty"`
	Writers                []Tag   `xml:"Writer,omitempty"`
	Countries              []Tag   `xml:"Country,omitempty"`
	Roles                  []Tag   `xml:"Role,omitempty"`
}

type Media struct {
	ID              string `xml:"id,attr,omitempty"`
	Duration        string `xml:"duration,attr,omitempty"`
	Bitrate         string `xml:"bitrate,attr,omitempty"`
	Width           string `xml:"width,attr,omitempty"`
	Height          string `xml:"height,attr,omitempty"`
	AspectRatio     string `xml:"aspectRatio,attr,omitempty"`
	AudioChannels   string `xml:"audioChannels,attr,omitempty"`
	AudioCodec      string `xml:"audioCodec,attr,omitempty"`
	VideoCodec      string `xml:"videoCodec,attr,omitempty"`
	VideoResolution string `xml:"videoResolution,attr,omitempty"`
	Container       string `xml:"container,attr,omitempty"`
	VideoFrameRate  string `xml:"videoFrameRate,attr,omitempty"`
	VideoProfile    string `xml:"videoProfile,attr,omitempty"`
	Parts           []Part `xml:"Part,omitempty"`
}

type Part struct {
	ID           string        `xml:"id,attr,omitempty"`
	Key          string        `xml:"key,attr,omitempty"`
	Duration     time.Duration `xml:"duration,attr,omitempty"`
	File         string        `xml:"file,attr,omitempty"`
	Size         int64         `xml:"size,attr,omitempty"`
	Container    string        `xml:"container,attr,omitempty"`
	VideoProfile string        `xml:"videoProfile,attr,omitempty"`
}

type Tag struct {
	Tag string `xml:"tag,attr,omitempty"`
}

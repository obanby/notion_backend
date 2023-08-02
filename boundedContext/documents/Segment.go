package documents

/* TODO:
HeaderSegment
TextSegment
CallOutSegment
QuoteSegment
ListSegment
ImageSegment --> Block storage, Multi-Part http requests
TableOfContentsSegment

// Behavioural Segments
TableSegment
MentionSegment
DatabaseSegment

SyncedSegment
	* One way sync view
	* Two way editing view

Next Steps:
[ ] - Create Header Segment
[ ] - Marshall Header Segment
[ ] - Save Segment to Repo
*/

type ISegment interface {
	GetSegmentId() string
	GetSegmentType() string
	Marshall() string
}

type Segment struct {
	Id   ID
	Type string
}

type HeaderSegment struct {
	Text  string // Max length
	Level int    // 1 through 3
}

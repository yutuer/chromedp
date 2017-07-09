package page

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ResourceType resource type as it was perceived by the rendering engine.
type ResourceType string

// String returns the ResourceType as string value.
func (t ResourceType) String() string {
	return string(t)
}

// ResourceType values.
const (
	ResourceTypeDocument    ResourceType = "Document"
	ResourceTypeStylesheet  ResourceType = "Stylesheet"
	ResourceTypeImage       ResourceType = "Image"
	ResourceTypeMedia       ResourceType = "Media"
	ResourceTypeFont        ResourceType = "Font"
	ResourceTypeScript      ResourceType = "Script"
	ResourceTypeTextTrack   ResourceType = "TextTrack"
	ResourceTypeXHR         ResourceType = "XHR"
	ResourceTypeFetch       ResourceType = "Fetch"
	ResourceTypeEventSource ResourceType = "EventSource"
	ResourceTypeWebSocket   ResourceType = "WebSocket"
	ResourceTypeManifest    ResourceType = "Manifest"
	ResourceTypeOther       ResourceType = "Other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ResourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ResourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ResourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ResourceType(in.String()) {
	case ResourceTypeDocument:
		*t = ResourceTypeDocument
	case ResourceTypeStylesheet:
		*t = ResourceTypeStylesheet
	case ResourceTypeImage:
		*t = ResourceTypeImage
	case ResourceTypeMedia:
		*t = ResourceTypeMedia
	case ResourceTypeFont:
		*t = ResourceTypeFont
	case ResourceTypeScript:
		*t = ResourceTypeScript
	case ResourceTypeTextTrack:
		*t = ResourceTypeTextTrack
	case ResourceTypeXHR:
		*t = ResourceTypeXHR
	case ResourceTypeFetch:
		*t = ResourceTypeFetch
	case ResourceTypeEventSource:
		*t = ResourceTypeEventSource
	case ResourceTypeWebSocket:
		*t = ResourceTypeWebSocket
	case ResourceTypeManifest:
		*t = ResourceTypeManifest
	case ResourceTypeOther:
		*t = ResourceTypeOther

	default:
		in.AddError(errors.New("unknown ResourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ResourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// FrameResource information about the Resource on the page.
type FrameResource struct {
	URL          string              `json:"url"`                    // Resource URL.
	Type         ResourceType        `json:"type"`                   // Type of this resource.
	MimeType     string              `json:"mimeType"`               // Resource mimeType as determined by the browser.
	LastModified *cdp.TimeSinceEpoch `json:"lastModified,omitempty"` // last-modified timestamp as reported by server.
	ContentSize  float64             `json:"contentSize,omitempty"`  // Resource content size.
	Failed       bool                `json:"failed,omitempty"`       // True if the resource failed to load.
	Canceled     bool                `json:"canceled,omitempty"`     // True if the resource was canceled during loading.
}

// FrameResourceTree information about the Frame hierarchy along with their
// cached resources.
type FrameResourceTree struct {
	Frame       *cdp.Frame           `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []*FrameResource     `json:"resources"`             // Information about frame resources.
}

// ScriptIdentifier unique script identifier.
type ScriptIdentifier string

// String returns the ScriptIdentifier as string value.
func (t ScriptIdentifier) String() string {
	return string(t)
}

// TransitionType transition type.
type TransitionType string

// String returns the TransitionType as string value.
func (t TransitionType) String() string {
	return string(t)
}

// TransitionType values.
const (
	TransitionTypeLink             TransitionType = "link"
	TransitionTypeTyped            TransitionType = "typed"
	TransitionTypeAutoBookmark     TransitionType = "auto_bookmark"
	TransitionTypeAutoSubframe     TransitionType = "auto_subframe"
	TransitionTypeManualSubframe   TransitionType = "manual_subframe"
	TransitionTypeGenerated        TransitionType = "generated"
	TransitionTypeAutoToplevel     TransitionType = "auto_toplevel"
	TransitionTypeFormSubmit       TransitionType = "form_submit"
	TransitionTypeReload           TransitionType = "reload"
	TransitionTypeKeyword          TransitionType = "keyword"
	TransitionTypeKeywordGenerated TransitionType = "keyword_generated"
	TransitionTypeOther            TransitionType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TransitionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t TransitionType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TransitionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch TransitionType(in.String()) {
	case TransitionTypeLink:
		*t = TransitionTypeLink
	case TransitionTypeTyped:
		*t = TransitionTypeTyped
	case TransitionTypeAutoBookmark:
		*t = TransitionTypeAutoBookmark
	case TransitionTypeAutoSubframe:
		*t = TransitionTypeAutoSubframe
	case TransitionTypeManualSubframe:
		*t = TransitionTypeManualSubframe
	case TransitionTypeGenerated:
		*t = TransitionTypeGenerated
	case TransitionTypeAutoToplevel:
		*t = TransitionTypeAutoToplevel
	case TransitionTypeFormSubmit:
		*t = TransitionTypeFormSubmit
	case TransitionTypeReload:
		*t = TransitionTypeReload
	case TransitionTypeKeyword:
		*t = TransitionTypeKeyword
	case TransitionTypeKeywordGenerated:
		*t = TransitionTypeKeywordGenerated
	case TransitionTypeOther:
		*t = TransitionTypeOther

	default:
		in.AddError(errors.New("unknown TransitionType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TransitionType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// NavigationEntry navigation history entry.
type NavigationEntry struct {
	ID             int64          `json:"id"`             // Unique id of the navigation history entry.
	URL            string         `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string         `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string         `json:"title"`          // Title of the navigation history entry.
	TransitionType TransitionType `json:"transitionType"` // Transition type.
}

// ScreencastFrameMetadata screencast frame metadata.
type ScreencastFrameMetadata struct {
	OffsetTop       float64             `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64             `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64             `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64             `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64             `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64             `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       *cdp.TimeSinceEpoch `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// DialogType javascript dialog type.
type DialogType string

// String returns the DialogType as string value.
func (t DialogType) String() string {
	return string(t)
}

// DialogType values.
const (
	DialogTypeAlert        DialogType = "alert"
	DialogTypeConfirm      DialogType = "confirm"
	DialogTypePrompt       DialogType = "prompt"
	DialogTypeBeforeunload DialogType = "beforeunload"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DialogType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DialogType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DialogType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch DialogType(in.String()) {
	case DialogTypeAlert:
		*t = DialogTypeAlert
	case DialogTypeConfirm:
		*t = DialogTypeConfirm
	case DialogTypePrompt:
		*t = DialogTypePrompt
	case DialogTypeBeforeunload:
		*t = DialogTypeBeforeunload

	default:
		in.AddError(errors.New("unknown DialogType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DialogType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AppManifestError error while paring app manifest.
type AppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int64  `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int64  `json:"line"`     // Error line.
	Column   int64  `json:"column"`   // Error column.
}

// NavigationResponse proceed: allow the navigation; Cancel: cancel the
// navigation; CancelAndIgnore: cancels the navigation and makes the requester
// of the navigation acts like the request was never made.
type NavigationResponse string

// String returns the NavigationResponse as string value.
func (t NavigationResponse) String() string {
	return string(t)
}

// NavigationResponse values.
const (
	NavigationResponseProceed         NavigationResponse = "Proceed"
	NavigationResponseCancel          NavigationResponse = "Cancel"
	NavigationResponseCancelAndIgnore NavigationResponse = "CancelAndIgnore"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t NavigationResponse) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t NavigationResponse) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *NavigationResponse) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch NavigationResponse(in.String()) {
	case NavigationResponseProceed:
		*t = NavigationResponseProceed
	case NavigationResponseCancel:
		*t = NavigationResponseCancel
	case NavigationResponseCancelAndIgnore:
		*t = NavigationResponseCancelAndIgnore

	default:
		in.AddError(errors.New("unknown NavigationResponse value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *NavigationResponse) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// LayoutViewport layout viewport position and dimensions.
type LayoutViewport struct {
	PageX        int64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// VisualViewport visual viewport position, dimensions, and scale.
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

// CaptureScreenshotFormat image compression format (defaults to png).
type CaptureScreenshotFormat string

// String returns the CaptureScreenshotFormat as string value.
func (t CaptureScreenshotFormat) String() string {
	return string(t)
}

// CaptureScreenshotFormat values.
const (
	CaptureScreenshotFormatJpeg CaptureScreenshotFormat = "jpeg"
	CaptureScreenshotFormatPng  CaptureScreenshotFormat = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t CaptureScreenshotFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t CaptureScreenshotFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *CaptureScreenshotFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch CaptureScreenshotFormat(in.String()) {
	case CaptureScreenshotFormatJpeg:
		*t = CaptureScreenshotFormatJpeg
	case CaptureScreenshotFormatPng:
		*t = CaptureScreenshotFormatPng

	default:
		in.AddError(errors.New("unknown CaptureScreenshotFormat value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *CaptureScreenshotFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ScreencastFormat image compression format.
type ScreencastFormat string

// String returns the ScreencastFormat as string value.
func (t ScreencastFormat) String() string {
	return string(t)
}

// ScreencastFormat values.
const (
	ScreencastFormatJpeg ScreencastFormat = "jpeg"
	ScreencastFormatPng  ScreencastFormat = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ScreencastFormat) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ScreencastFormat) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ScreencastFormat) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ScreencastFormat(in.String()) {
	case ScreencastFormatJpeg:
		*t = ScreencastFormatJpeg
	case ScreencastFormatPng:
		*t = ScreencastFormatPng

	default:
		in.AddError(errors.New("unknown ScreencastFormat value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ScreencastFormat) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

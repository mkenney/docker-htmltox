package CSS

import (
	DOM "app/chrome/dom"
	Page "app/chrome/page"
	"fmt"
)

/*
GetMatchedStylesForNodeParams represents CSS.getMatchedStylesForNode parameters.
*/
type GetMatchedStylesForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetInlineStylesForNodeParams represents CSS.getInlineStylesForNode parameters.
*/
type GetInlineStylesForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetComputedStyleForNodeParams represents CSS.getComputedStyleForNode parameters.
*/
type GetComputedStyleForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetPlatformFontsForNodeParams represents CSS.getPlatformFontsForNode parameters.
*/
type GetPlatformFontsForNodeParams struct {
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
GetStyleSheetTextParams represents CSS.getStyleSheetText parameters.
*/
type GetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
CollectClassNamesParams represents CSS.collectClassNames parameters.
*/
type CollectClassNamesParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

/*
SetStyleSheetTextParams represents CSS.setStyleSheetText parameters.
*/
type SetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Text         string       `json:"text"`
}

/*
SetRuleSelectorParams represents CSS.setRuleSelector parameters.
*/
type SetRuleSelectorParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        SourceRange  `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetKeyframeKeyParams represents CSS.setKeyframeKey parameters.
*/
type SetKeyframeKeyParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        SourceRange  `json:"range"`
	Selector     string       `json:"selector"`
}

/*
SetStyleTextsParams represents CSS.setStyleTexts parameters.
*/
type SetStyleTextsParams struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

/*
SetMediaTextParams represents CSS.setMediaText parameters.
*/
type SetMediaTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        SourceRange  `json:"range"`
	Text         string       `json:"text"`
}

/*
CreateStyleSheetParams represents CSS.createStyleSheet parameters.
*/
type CreateStyleSheetParams struct {
	FrameID Page.FrameID `json:"frameId"`
}

/*
AddRuleParams represents CSS.addRule parameters.
*/
type AddRuleParams struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// The text of a new rule.
	RuleText string `json:"ruleText"`

	// Text position of a new rule in the target style sheet.
	Location SourceRange `json:"location"`
}

/*
ForcePseudoStateParams represents CSS.forcePseudoState parameters.
*/
type ForcePseudoStateParams struct {
	// The element ID for which to force the pseudo state.
	NodeID DOM.NodeID `json:"nodeId"`

	// Element pseudo classes to force when computing the element's style.
	// Allowed values: active, focus, hover, visited.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

/*
GetMediaQueriesParams represents CSS.getMediaQueries parameters.
*/
type GetMediaQueriesParams struct {
	Medias []*CSSMedia `json:"medias"`
}

/*
SetEffectivePropertyValueForNodeParams represents CSS.setEffectivePropertyValueForNode parameters.
*/
type SetEffectivePropertyValueForNodeParams struct {
	// The element id for which to set property.
	NodeID       DOM.NodeID `json:"nodeId"`
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

/*
GetBackgroundColorsParams represents CSS.getBackgroundColors parameters.
*/
type GetBackgroundColorsParams struct {
	// Id of the node to get background colors for.
	NodeID DOM.NodeID `json:"nodeId"`
}

/*
MediaQueryResultChangedEvent represents CSS.mediaQueryResultChanged event data.
*/
type MediaQueryResultChangedEvent struct{}

/*
FontsUpdatedEvent represents CSS.fontsUpdated event data.
*/
type FontsUpdatedEvent struct{}

/*
StyleSheetChangedEvent represents CSS.styleSheetChanged event data.
*/
type StyleSheetChangedEvent struct{}

/*
StyleSheetAddedEvent represents CSS.styleSheetAdded event data.
*/
type StyleSheetAddedEvent struct{}

/*
StyleSheetID is the ID of a stylesheet
*/
type StyleSheetID string

/*
StyleSheetOrigin is an enum defining:
	- "injected" for stylesheets injected via extension
	- "user-agent" for user-agent stylesheets
	- "inspector" for stylesheets created by the inspector (i.e. those holding the "via inspector"
	  rules)
	- "regular" for regular stylesheets.
*/
type StyleSheetOrigin int

const (
	_injected StyleSheetOrigin = iota
	_userAgent
	_inspector
	_regular
)

func (a StyleSheetOrigin) String() string {
	if a == 0 {
		return "injected"
	}
	if a == 1 {
		return "user-agent"
	}
	if a == 2 {
		return "inspector"
	}
	if a == 3 {
		return "regular"
	}
	panic(fmt.Errorf("Invalid StyleSheetOrigin %d", a))
}

/*
PseudoElementMatches is a CSS rule collection for a single pseudo style.
*/
type PseudoElementMatches struct {
	// Pseudo element type.
	PseudoType *DOM.PseudoType `json:"pseudoType"`

	// Matches of CSS rules applicable to the pseudo style.
	Matches []*RuleMatch `json:"matches"`
}

/*
InheritedStyleEntry is an inherited CSS rule collection from ancestor node.
*/
type InheritedStyleEntry struct {
	// Optional. The ancestor node's inline style, if any, in the style inheritance chain.
	InlineStyle CSSStyle `json:"inlineStyle,omitempty"`

	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`
}

/*
RuleMatch is match data for a CSS rule.
*/
type RuleMatch struct {
	// CSS rule in the match.
	Rule CSSRule `json:"rule"`

	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors []int `json:"matchingSelectors"`
}

/*
Value represents data for a simple selector (these are delimited by commas in a selector list).
*/
type Value struct {
	// Value text.
	Text string `json:"text"`

	// Optional. Value range in the underlying resource (if available).
	Range SourceRange `json:"range,omitempty"`
}

/*
SelectorList represents selector list data.
*/
type SelectorList struct {
	// Selectors in the list.
	Selectors []*Value `json:"selectors"`

	// Rule selector text.
	Text string `json:"text"`
}

/*
CSSStyleSheetHeader holds CSS stylesheet metainformation.
*/
type CSSStyleSheetHeader struct {
	// The stylesheet identifier.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Owner frame identifier.
	FrameID Page.FrameID `json:"frameId"`

	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`

	// Optional. URL of source map associated with the stylesheet (if any).
	SourceMapURL string `json:"sourceMapURL,omitempty"`

	// Stylesheet origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Stylesheet title.
	Title string `json:"title"`

	// Optional. The backend ID for the owner node of the stylesheet.
	OwnerNode DOM.BackendNodeID `json:"ownerNode,omitempty"`

	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`

	// Optional. Whether the sourceURL field value comes from the sourceURL comment.
	HasSourceURL bool `json:"hasSourceURL,omitempty"`

	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for
	// document.written STYLE tags.
	IsInline bool `json:"isInline"`

	// Line offset of the stylesheet within the resource (zero based).
	StartLine float64 `json:"startLine"`

	// Column offset of the stylesheet within the resource (zero based).
	StartColumn float64 `json:"startColumn"`

	// Size of the content (in characters).
	Length float64 `json:"length"`
}

/*
CSSRule is a CSS rule representation.
*/
type CSSRule struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Rule selector data.
	SelectorList SelectorList `json:"selectorList"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated style declaration.
	Style CSSStyle `json:"style"`

	// Optional. Media list array (for rules involving media queries). The array enumerates media
	// queries starting with the innermost one, going outwards.
	Media []*CSSMedia `json:"media,omitempty"`
}

/*
RuleUsage holds CSS coverage information.
*/
type RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset float64 `json:"startOffset"`

	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset float64 `json:"endOffset"`

	// Indicates whether the rule was actually used by some element in the DOM.
	Used bool `json:"used"`
}

/*
SourceRange is a text range within a resource. All numbers are zero-based.
*/
type SourceRange struct {
	// Start line of range.
	StartLine int `json:"startLine"`

	// Start column of range (inclusive).
	StartColumn int `json:"startColumn"`

	// End line of range.
	SndLine int `json:"endLine"`

	// End column of range (exclusive).
	SndColumn int `json:"endColumn"`
}

/*
ShorthandEntry holds a CSS shorthand entry
*/
type ShorthandEntry struct {
	// Shorthand name.
	Name string `json:"name"`

	// Shorthand value.
	Value string `json:"value"`

	// Optional. Whether the property has "!important" annotation (implies false if absent).
	Important bool `json:"important,omitempty"`
}

/*
CSSComputedStyleProperty holds a computed style property
*/
type CSSComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`

	// Computed style property value.
	Value string `json:"value"`
}

/*
CSSStyle is a CSS style representation.
*/
type CSSStyle struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// CSS properties in the style.
	CSSProperties []*CSSProperty `json:"cssProperties"`

	// Computed values for all shorthands found in the style.
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`

	// Style declaration text (if available).
	CSSText string `json:"cssText"`

	// Optional. Style declaration range in the enclosing stylesheet (if available).
	Range SourceRange `json:"range,omitempty"`
}

/*
CSSProperty holds CSS property declaration data.
*/
type CSSProperty struct {
	// The property name.
	Name string `json:"name"`

	// The property value.
	Value string `json:"value"`

	// Optional. Whether the property has "!important" annotation (implies false if absent).
	Important bool `json:"important,omitempty"`

	// Optional. Whether the property is implicit (implies false if absent).
	Implicit bool `json:"implicit,omitempty"`

	// Optional. The full property text as specified in the style.
	Text string `json:"text,omitempty"`

	// Optional. Whether the property is understood by the browser (implies true
	// if absent).
	ParsedOk bool `json:"parsedOk,omitempty"`

	// Optional. Whether the property is disabled by the user (present for
	// source-based properties only).
	Disabled bool `json:"disabled,omitempty"`

	// Optional. The entire property range in the enclosing style declaration (if available).
	Range *SourceRange `json:"range,omitempty"`
}

/*
CSSMedia holds a CSS media rule descriptor.
*/
type CSSMedia struct {
	// Media query text.
	Text string `json:"text"`

	// Source of the media query:
	//	- "mediaRule" if specified by a @media rule
	//	- "importRule" if specified by an @import rule
	//	- "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag
	//	- "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	//	  Allowed values: mediaRule, importRule, linkedSheet, inlineSheet.
	Source string `json:"source"`

	// Optional. URL of the document containing the media query description.
	SourceURL string `json:"sourceURL,omitempty"`

	// Optional. The associated rule (@media or @import) header range in the enclosing stylesheet
	// (if available).
	Range *SourceRange `json:"range,omitempty"`

	// Optional. Identifier of the stylesheet containing this object (if exists).
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Optional. Array of media queries.
	MediaList []*MediaQuery `json:"mediaList,omitempty"`
}

/*
MediaQuery is a media query descriptor.
*/
type MediaQuery struct {
	// Array of media query expressions.
	Expressions []*MediaQueryExpression `json:"expressions"`

	// Whether the media query condition is satisfied.
	Active bool `json:"active"`
}

/*
MediaQueryExpression is a media query expression descriptor.
*/
type MediaQueryExpression struct {
	// Media query expression value.
	Value float64 `json:"value"`

	// Media query expression units.
	Unit string `json:"unit"`

	// Media query expression feature.
	Feature string `json:"feature"`

	// Optional. The associated range of the value text in the enclosing stylesheet (if available).
	ValueRange *SourceRange `json:"valueRange,omitempty"`

	// Optional. Computed length of media query expression (if applicable).
	ComputedLength float64 `json:"computedLength,omitempty"`
}

/*
PlatformFontUsage holds information about the amount of glyphs that were rendered with given font.
*/
type PlatformFontUsage struct {
	// Font's family name reported by platform.
	FamilyName string `json:"familyName"`

	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont bool `json:"isCustomFont"`

	// Amount of glyphs that were rendered with this font.
	GlyphCount float64 `json:"glyphCount"`
}

/*
CSSKeyframesRule is a CSS keyframes rule representation.
*/
type CSSKeyframesRule struct {
	// Animation name.
	AnimationName *Value `json:"animationName"`

	// List of keyframes.
	Keyframes []*CSSKeyframeRule `json:"keyframes"`
}

/*
CSSKeyframeRule is a CSS keyframe rule representation.
*/
type CSSKeyframeRule struct {
	// Optional. The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"`

	// Parent stylesheet's origin.
	Origin StyleSheetOrigin `json:"origin"`

	// Associated key text.
	KeyText Value `json:"keyText"`

	// Associated style declaration.
	Style CSSStyle `json:"style"`
}

/*
StyleDeclarationEdit is a descriptor of operation to mutate style declaration text.
*/
type StyleDeclarationEdit struct {
	// The css style sheet identifier.
	StyleSheetID StyleSheetID `json:"styleSheetId"`

	// The range of the style text in the enclosing stylesheet.
	Range *SourceRange `json:"range"`

	// New style text.
	Text string `json:"text"`
}

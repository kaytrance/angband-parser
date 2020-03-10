package parser

/**
This is a single generic spell structure
*/
type Spell struct {
	SpellColor  string
	Damage      int
	Description string
}

/**
Here we define colors. It can also be extracted to a structure of Term colors, but lets just use bare color values.
 */
var SpellTypeColors = map[string]string{
	"GF_ACID": "#65ba3b", //TERM_SLATE
	// ..
	// more to add here
	// ..
	"DEFAULT": "#ffffff", // TERM_WHITE
}

/**
All possible spells defined here. To make it more universal instead of defining all spells here they shall be
populated from external text file file that follows certain structure. For now, for demonstration purposes
let's leave it as is.
*/
var SPELLS = map[string]Spell{
	"BO_ACID": {SpellColor: SpellTypeColors["GF_ACID"], Damage: 9000, Description: "Ball of Acid"},
	// ..
	// more to add here
	// ..
}

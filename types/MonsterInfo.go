package parser

// https://github.com/sulkasormi/frogcomposband/blob/master/lib/edit/r_info.txt
// === Understanding monster.txt ===
// N: serial number : monster name
// G: symbol : color
// I: speed : hit points : vision : armor class : alertness
// W: depth : rarity : unused (always 0) : experience for kill
// B: attack method : attack effect : damage
// S: spell frequency |
// S: spell type | spell type | etc
// F: flag | flag | etc
// D: Description

type MonsterInfo struct {
	SerialNumber int    // taken from 'N'
	Name         string // taken from 'N'
	B            Blow
	D            Description
	F            Flags
	G            Graphics
	I            Information
	S            Spells
	W            MoreInfo
}

// 'B' is for blows - method of attack, effect of attack, and damage from attack.
// There may be up to four of these lines; effect and damage are optional.
type Blow struct {
	Method string // todo: change to Method struct, must be filled manually
	Effect string // todo: change to Effect struct, must be filled manually
	Damage string
}

type Blows []Blow

// 'D' is for description.
type Description string

// 'F' is for Flags
type Flags []strings // todo: change to Flag struct, lots of crap here I imagine

// 'G' is for graphics - symbol and color (ex. G:@:w). There are 16 colors, as follows:
//  D - Dark Gray    w - White          s - Gray          o - Orange
//  r - Red          g - Green          b - Blue          u - Brown
//  d - Black        W - Light Gray     v - Violet        y - Yellow
//  R - Light Red    G - Light Green    B - Light Blue    U - Light Brown
type Graphics struct {
	Symbol string
	Color  string
}

// 'I' is for Information
// Has speed, health, vision in tens of feet, armor class, and alertness.
// 110 is normal speed.
// Alertness ranges from 0 (ever vigilant for intruders) to 255 (prefers to ignore intruders).
type Information struct {
	Speed      int
	Health     string
	Vision     int // in tens of feet
	ArmorClass int
	Alertness  int // ranges from 0 (ever vigilant for intruders) to 255 (prefers to ignore intruders
}

// 'S' is for Spells
// first S: spell frequency |
// second S: spell type | spell type | etc
type Spells struct {
	Frequency int     // treated as 1_IN_<Frequency>
	Spells    []Flags // todo: change to Spell predefined struct, must be filled manually ? colors https://github.com/angbandplus/AngbandPlus/blob/10670ef50d7454d0a5d076b5dd529c281b940e05/src/spells1.c#L336
}

// 'W' is for more information - level, rarity, and experience for killing. 4 slots, third slot is unused.
type MoreInfo struct {
	Level      int
	Rarity     int
	Experience int
}

// essa entity é como se fosse um ENUM em ts
// são enums baseados em strings que em go são tipos

package entity

type FragranceFamily string
type Intensity string
type Season string
type Occasion string

type FragranceNotes struct {
	Top   []string
	Heart []string
	Base  []string
}

type Fragrance struct {
	ID          string
	Name        string
	Family      FragranceFamily
	Notes       FragranceNotes
	Intensity   Intensity
	Seasons     []Season
	Occasions   []Occasion
	Description string
}

// aqui a const criada é tipada pelo type criado acima, seria algo como em ts:
// const FamilyFlora: FragranceFamily = "floral"
const (
	FamilyFloral   FragranceFamily = "floral"
	FamilyWoody    FragranceFamily = "woody"
	FamilyCitrus   FragranceFamily = "citrus"
	FamilyOriental FragranceFamily = "oriental"
	FamilyFresh    FragranceFamily = "fresh"
	FamilyAquatic  FragranceFamily = "aquatic"
	FamilyGourmand FragranceFamily = "gourmand"
)

const (
	IntensityLight    Intensity = "light"
	IntensityModerate Intensity = "moderate"
	IntensityStrong   Intensity = "strong"
	IntensityIntense  Intensity = "intense"
)

const (
	SeasonSpring Season = "spring"
	SeasonSummer Season = "summer"
	SeasonAutumn Season = "autumn"
	SeasonWinter Season = "winter"
)

const (
	OccasionCasual   Occasion = "casual"
	OccasionFormal   Occasion = "formal"
	OccasionRomantic Occasion = "romantic"
	OccasionSport    Occasion = "sport"
	OccasionWork     Occasion = "work"
	OccasionNight    Occasion = "night"
)
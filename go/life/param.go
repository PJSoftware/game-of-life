package life

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Param provides interface to values from ini file
type Param struct {
	Width, Height int
	RuleValues    map[string]map[string]bool
	SpawnPercent  int
	WrapX, WrapY  bool
	ini           INI
}

// Init sets values to those from ini file, or to defaults if an error occurs
func (s *Param) Init(board string, rules string) {
	s.RuleValues = make(map[string]map[string]bool)
	s.RuleValues["b"] = make(map[string]bool)
	s.RuleValues["s"] = make(map[string]bool)

	s.setDefaults()
	err := s.ini.Parse("../game-of-life.ini")
	if err != nil {
		log.Print("world.Init: Error reading INI file: " + err.Error())
		log.Print("world.Init: using Default parameters instead")
		return
	}

	section := "world-" + board
	_, err = s.ini.Section(section)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	s.readResolution(section)
	s.readWrap(section)
	s.readSpawn(section)
	s.readRules(rules)
}

func (s *Param) readResolution(sectName string) {
	val, err := s.ini.Value(sectName, "resolution")
	if err == nil {
		s.setResolution(val)
	} else {
		log.Print(err.Error())
		log.Print("Param: resolution not found, using default")
	}
}

func (s *Param) readWrap(sectName string) {
	val1, err1 := s.ini.Value(sectName, "wrapx")
	val2, err2 := s.ini.Value(sectName, "wrapy")
	if err1 == nil && err2 == nil {
		s.setWrap(val1 == "yes", val2 == "yes")
	} else {
		log.Print("Param: unable to read wrap values, using defaults")
	}
}

func (s *Param) readSpawn(sectName string) {
	val, err1 := s.ini.Value(sectName, "spawn_percent")
	if err1 == nil {
		percent, err2 := strconv.Atoi(val)
		if err2 == nil {
			s.setSpawn(percent)
		} else {
			log.Printf("Param: spawn percentage '%s' not numeric, using default", val)
		}
	} else {
		log.Print("Param: unable to read spawn percentage, using default")
	}
}

func (s *Param) readRules(valName string) {
	val, err := s.ini.Value("rules", valName)
	if err == nil {
		s.setRules(val)
	} else {
		log.Printf("Param: could not load '%s' rules, using default (classic Conway)", valName)
	}
}

func (s *Param) setDefaults() {
	// world configuration "default", AKA my VZ-200 resolution
	s.setResolution("128x64")
	s.setWrap(true, true)
	s.setSpawn(20)

	// default rules, the "Conway" classics
	s.setRules("b3/s23")
}

func (s *Param) setResolution(resolution string) {
	reRes, err := regexp.Compile(`^(\d+)x(\d+)$`)
	if err != nil {
		log.Fatal(err)
	}
	s.Width, err = strconv.Atoi(reRes.FindStringSubmatch(resolution)[1])
	if err != nil {
		log.Fatal(err)
	}
	s.Height, err = strconv.Atoi(reRes.FindStringSubmatch(resolution)[2])
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Param) setWrap(wrapX, wrapY bool) {
	s.WrapX = wrapX
	s.WrapY = wrapY
}

func (s *Param) setSpawn(percentage int) {
	s.SpawnPercent = percentage
}

func (s *Param) setRules(simRules string) {
	reRules, err := regexp.Compile(`b(\d*)/s(\d*)`)
	if err != nil {
		log.Fatal(err)
	}
	bDigits := reRules.FindStringSubmatch(simRules)[1]
	sDigits := reRules.FindStringSubmatch(simRules)[2]
	s.parseRuleDigits("b", bDigits)
	s.parseRuleDigits("s", sDigits)
}

func (s *Param) parseRuleDigits(rule string, digits string) {
	digitSlice := strings.Split(digits, "")
	for _, digit := range digitSlice {
		s.RuleValues[rule][digit] = true
	}
}

package life

type Setup struct {
	Width, Height              int
	BirthValues, SurviveValues map[string]bool
	SpawnPercent               int
	WrapX, WrapY               bool
}

func (s *Setup) Init() {
	s.BirthValues = make(map[string]bool)
	s.SurviveValues = make(map[string]bool)

	s.setDefaults()
	var ini INI
	err := ini.Parse("..\\simulations.ini")
	if err != nil {
		s.setDefaults()
	} else {
		err := ini.Section("default")
		if err != nil {
			s.setDefaults()
		} else {
			//res, err := ini.Value("resolution")
		}
	}

}

func (s *Setup) setDefaults() {
	s.Width = 128
	s.Height = 64
	s.BirthValues["b3"] = true
	s.SurviveValues["s2"] = true
	s.SurviveValues["s3"] = true
	s.SpawnPercent = 20
	s.WrapX = true
	s.WrapY = true
}

/*
class ErrorProcessingINI extends Exception { }

class Setup
{
    public $width;
    public $height;
    public $birth_values;
    public $survive_values;
    public $spawn_percent;
    public $wrapx;
    public $wrapy;

    function __construct($section = "default")
    {
        $sim_ini = "..\simulations.ini";
        if (!file_exists($sim_ini)) {
            throw new ErrorProcessingINI("File not found: '$sim_ini'");
        }
        $ini = parse_ini_file($sim_ini, true);

        if (!array_key_exists($section, $ini)) {
            throw new ErrorProcessingINI("Section not found: '$section'");
        }
        $setup = $ini[$section];

        $keys = array("resolution", "rules", "wrapx", "wrapy", "spawn_percent");    // "grid" not used yet
        foreach ($keys as $key) {
            if (!array_key_exists($key, $setup)) {
                throw new ErrorProcessingINI("Value '$key' not found in $section section");
            }
        }

        if (preg_match("/^(\d+)x(\d+)$/",$setup["resolution"],$res)) {
            $this->width = $res[1];
            $this->height = $res[2];
        } else {
            throw new ErrorProcessingINI("Expected ##x## in 'resolution'");
        }

        if (preg_match("/^b(\d+)\//",$setup["rules"],$birth)) {
            foreach (preg_split("//",$birth[1]) as $val) {
                IF ($val <> "") {
                    $this->birth_values["b$val"] = 1;
                }
            }
        } else {
            throw new ErrorProcessingINI("Birth rule not specified; expecting 'b#/...'");
        }
        if (preg_match("/\/s(\d+)$/",$setup["rules"],$survive)) {
            foreach (preg_split("//",$survive[1]) as $val) {
                if ($val <> "") {
                    $this->survive_values["s$val"] = 1;
                }
            }
        } else {
            throw new ErrorProcessingINI("Survival rule not specified; expecting '.../s#'");
        }

        $this->wrapx = ($setup["wrapx"] === "yes");
        $this->wrapy = ($setup["wrapy"] === "yes");
        $this->spawn_percent = intval($setup["spawn_percent"]);
        if ($this->spawn_percent < 5 || $this->spawn_percent > 95) {
            throw new ErrorProcessingINI("Expected positive numeric spawn_percent (between 5 and 95)");
        }

    }
}
*/

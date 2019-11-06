<?php

error_reporting(E_ALL);

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

    function __construct($section = "default", $rules = "conway")
    {
        try {
            $this->parse_ini($section, $rules);
        } catch (ErrorProcessingINI $e) {
            echo "Unable to read ini file: ".$e->getMessage()."\n";
            $this->width = 40;
            $this->height = 20;
            $this->birth_values = array();
            $this->birth_values["b3"] = 1;
            $this->survive_values = array();
            $this->survive_values["s2"] = 1;
            $this->survive_values["s3"] = 1;
            $this->spawn_percent = 20;
            $this->wrapx = true;
            $this->wrapy = false;
        }
    }

    private function parse_ini($section, $rules) {
        $sim_ini = "../game-of-life.ini";
        if (!file_exists($sim_ini)) {
            throw new ErrorProcessingINI("File not found: '$sim_ini'");
        }
        $ini = parse_ini_file($sim_ini, true);

        $section = "world-$section";
        if (!array_key_exists($section, $ini)) {
            throw new ErrorProcessingINI("Section not found: '$section'");
        }
        if (!array_key_exists("rules", $ini)) {
            throw new ErrorProcessingINI("Section not found: 'rules'");
        }
        $setup = $ini[$section];
        $ruleset = $ini["rules"];

        $keys = array("resolution", "wrapx", "wrapy", "spawn_percent");    // "grid" not used yet
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

        $this->wrapx = ($setup["wrapx"] === "yes");
        $this->wrapy = ($setup["wrapy"] === "yes");
        $this->spawn_percent = intval($setup["spawn_percent"]);
        if ($this->spawn_percent < 5 || $this->spawn_percent > 95) {
            throw new ErrorProcessingINI("Expected positive numeric spawn_percent (between 5 and 95)");
        }

        if (preg_match("/^b(\d+)\//",$ruleset[$rules],$birth)) {
            foreach (preg_split("//",$birth[1]) as $val) {
                IF ($val <> "") {
                    $this->birth_values["b$val"] = 1;
                }
            }
        } else {
            throw new ErrorProcessingINI("Birth rule not specified; expecting 'b#/...'");
        }
        if (preg_match("/\/s(\d+)$/",$ruleset[$rules],$survive)) {
            foreach (preg_split("//",$survive[1]) as $val) {
                if ($val <> "") {
                    $this->survive_values["s$val"] = 1;
                }
            }
        } else {
            throw new ErrorProcessingINI("Survival rule not specified; expecting '.../s#'");
        }

    }
}

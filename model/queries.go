package model

import (
	"errors"
	"strings"

	"github.com/trezbit/csf-kn-corpus-generator/utils"
)

// create a type with type, label, querycontent
type CorpusQuery struct {
	QType    string `json:"type"`
	RPath    string `json:"label"`
	QContent string `json:"querycontent"`
}

// define string variables that store parameterized query content
const (
	_C_CSF_STD    = "NIST CSF version 2.0"
	_Q_CSF_ASSESS = "What are some example questions to assess compliance in terms of implementation of  ___STANDARD___ control ___CTRL___?"
	_Q_CSF_CORE   = "What are the key concepts that the NIST CSF version 2.0 control ___CTRL2___ cover, in connection with, or expanding on the NIST CSF version 1.1 control ___CTRL1___, where the primary focus is on the verification of if '___CHECK___'? In identifying these key concepts, consider specific examples, implementation and policy development guidance provided by NIST."
	_Q_CSF_MAP    = "What are the key concepts that the ___STANDARD___ security control ___CTRL___ cover, in specific examples, and in terms of applicable policies and procedures? Also emphasize potential overlaps with NIST CSF version 1.1 and version 2.0 controls and control families where applicable."
)

// define a function that returns a corpus query slice
func BuildCorpusQueries(qtype string, inpath string, outpath string) ([]CorpusQuery, error) {

	var mapped bool
	var query string

	if qtype == "assess" {
		mapped = false
		query = _Q_CSF_ASSESS
	} else if qtype == "csf" {
		mapped = true
		query = _Q_CSF_CORE
	} else if qtype == "map" {
		mapped = true
		query = _Q_CSF_MAP
	} else {
		return nil, errors.New("invalid query type")
	}

	lines, err := utils.ReadLines(inpath)
	utils.CheckError(err)

	var queries []CorpusQuery

	for _, line := range lines {
		var queryline string
		var path string
		if !mapped {
			queryline = strings.Replace(query, "___CTRL___", line, 1)
			queryline = strings.Replace(queryline, "___STANDARD___", _C_CSF_STD, 1)
			path = outpath + "/" + line + ".txt"
		} else {
			res := strings.Split(line, "|")
			if qtype == "csf" {
				// DE.AE-02|DE.AE-2|Potentially adverse events are analyzed to better understand associated activities
				queryline = strings.Replace(query, "___CTRL2___", res[0], 1)
				queryline = strings.Replace(queryline, "___CTRL1___", res[1], 1)
				queryline = strings.Replace(queryline, "___CHECK___", res[2], 1)
				path = outpath + "/" + res[0] + "+" + res[1] + ".txt"
			} else {
				// 62443-2-1:2009|4.3.3.3.9
				queryline = strings.Replace(query, "___CTRL___", res[1], 1)
				queryline = strings.Replace(queryline, "___STANDARD___", res[0], 1)
				var stdcode string
				stdcode = strings.Replace(res[0], " ", "_", -1)
				stdcode = strings.Replace(stdcode, ":", "-", -1)
				stdcode = strings.Replace(stdcode, "/", "-", -1)
				path = outpath + "/" + stdcode + "." + strings.Replace(res[1], " ", "-", -1) + ".txt"
			}
		}
		queries = append(queries, CorpusQuery{QType: qtype, RPath: path, QContent: queryline})
	}

	return queries, nil
}

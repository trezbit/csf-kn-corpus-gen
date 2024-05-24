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
	_Q_CSF_KEY    = `Task: Generate a JSON object capturing the key concepts for the NIST CSF version 2.0 control ___CTRL2___ cover, in connection with, or expanding on the NIST CSF version 1.1___CTRL1___. With each key concept, provide a description that provides context for a better understanding of what the key concept means.
	Instructions:
	In identifying these key concepts, consider specific examples, implementation and policy development guidance provided by NIST.
	Do not use references to actual Control or Standard labels and acronyms within the captured key concepts.
	Do not go beyond 10 key concepts as the result set.
	For each key concept identified, make sure it is stated concisely, not going beyond 3-4 phrases.
	
	Note: Do not include any explanations or apologies in your responses.
	Do not include any text except the generated JSON. 
	Do not include any JSON marker, block/code indicator.
	
	Examples:  Sample generated JSON for this would look like: 
	{ "results": [
	{"key_concept": "Automated Response Mechanisms", "description": "A next-generation CSF might advocate for not only detecting events through correlation but also automatically responding to certain types of detected events based on pre-established criteria and playbooks. This would help in accelerating the response to common or known types of incidents."}, 
	{"key_concept": "Anomaly Detection Evolution", "description": "With continuously evolving cyber threats, an updated control could encourage the ongoing development and refinement of anomaly detection capabilities, ensuring they remain effective against the latest attack methodologies."}
	]}"`
	_Q_ASSESSQ = `Task: Generate a JSON object capturing the core questions to be able to assess an entity's (organization, small business etc.) compliance within the scope of the ___STANDARD___ control ___CTRL___. These questions should target various applicable implementation aspects: related policies & procedures in place, process life-cycle management, expanding on the earlier NIST CSF version 1.1. if needed.
	With each question, provide a brief rationale capturing the essential concepts and considerations as well as a high-level scope of the question.
	Instructions:
	In identifying these assessment questions, consider specific examples, implementation and policy development guidance provided by NIST.
	Do not use references to actual Control or Standard labels and acronyms within the captured key concepts.
	Do not go beyond 25 core questions scoped under 5-6 key areas (with 2-4 questions under each scoped area) as the result set.
	For each question identified, make sure it is stated concisely.
	
	Note: Do not include any explanations or apologies in your responses.
	Do not include any text except the generated JSON. 
	Do not include any JSON marker, block/code indicator.
	
	Examples:  Sample generated JSON for this would look like: 
	{ "results": [
	{"scope": "Detection Capability", "question": "Can you describe the process and tools used to detect anomalies and events in your network?", "rationale": "This question helps in understanding the organization's capability to detect potential cybersecurity events and anomalies in a timely manner."},
	{"scope": "Detection Capability", "question": "How do you ensure that your detection capabilities are effective in identifying potential cybersecurity events?", "rationale": "This question helps in assessing the effectiveness of the organization's detection capabilities in identifying potential cybersecurity events."},
	{"scope": "Alerting Mechanisms", "question": "How are detected events and anomalies communicated to the appropriate personnel?", "rationale": "This question helps in understanding the organization's process for alerting the appropriate personnel about detected events and anomalies."},
	{"scope": "Alerting Mechanisms", "question": "What is the process for escalating alerts that may indicate a potential cybersecurity incident?", "rationale": "This question helps in assessing the organization's process for escalating alerts that may indicate a potential cybersecurity incident."}
	]}"`
)

// define a constant _Q_CSF_CORE that stores a query string with newline characters

// define a function that returns a corpus query slice
func BuildCorpusQueries(qtype string, inpath string, outpath string) ([]CorpusQuery, error) {

	var mapped bool
	var query string

	if qtype == "assess" {
		mapped = false
		query = _Q_CSF_ASSESS
	} else if qtype == "assessq" {
		mapped = false
		query = _Q_ASSESSQ
	} else if qtype == "csf" {
		mapped = true
		query = _Q_CSF_CORE
	} else if qtype == "csfkey" {
		mapped = true
		query = _Q_CSF_KEY
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
			if qtype == "assessq" {
				path = outpath + "/questions/" + line + ".json"
			} else {
				path = outpath + "/" + line + ".txt"
			}
		} else {
			res := strings.Split(line, "|")
			if qtype == "csf" {
				// DE.AE-02|DE.AE-2|Potentially adverse events are analyzed to better understand associated activities
				queryline = strings.Replace(query, "___CTRL2___", res[0], 1)
				queryline = strings.Replace(queryline, "___CTRL1___", res[1], 1)
				queryline = strings.Replace(queryline, "___CHECK___", res[2], 1)
				path = outpath + "/" + res[0] + "+" + res[1] + ".txt"
			} else if qtype == "csfkey" {
				queryline = strings.Replace(query, "___CTRL2___", res[0], 1)
				if res[1] == "NOMAP" {
					queryline = strings.Replace(queryline, "___CTRL1___", "", 1)
				} else {
					queryline = strings.Replace(queryline, "___CTRL1___", " control "+res[1], 1)
				}

				path = outpath + "/keys/" + res[0] + "+" + res[1] + ".json"
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

package expand_arpeggio

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"asdfgh/src/constants"
	"asdfgh/src/modifier"
	"asdfgh/src/music"

	log "github.com/schollz/logger"
)

const ARPEGGIO_UP = "u"
const ARPEGGIO_DOWN = "d"

func ExpandArpeggio(s string) (result string, err error) {
	newS, arpString := modifier.Pop(s, constants.MODIFIER_ARPEGGIO)
	if arpString == "" || newS == "" {
		return s, nil
	}
	modified := modifier.Split(newS)
	newS = modified.Unmodified
	log.Tracef("%s --> %s, %s", s, result, arpString)

	tokenized := tokenizeLetterNumberes(arpString)
	log.Tracef("tokenized: %v", tokenized)
	notei := 0
	octave := 0
	notes := []music.Note{}
	originalNotes, _ := music.Parse(newS, 60)
	log.Tracef("originalNotes: %v", originalNotes)

	re := regexp.MustCompile(`\d+`)
	for i, token := range tokenized {
		// get the number using regex
		match := re.FindString(token)
		if match == "" {
			err = fmt.Errorf("invalid arpeggio token: %s", token)
			return
		}
		var steps int
		steps, err = strconv.Atoi(match)
		if err != nil {
			return
		}
		log.Tracef("steps: %d, notei: %d", steps, notei)
		direction := 0
		if string(token[0]) == ARPEGGIO_UP {
			direction = 1
		} else if string(token[0]) == ARPEGGIO_DOWN {
			direction = -1
		} else {
			err = fmt.Errorf("invalid arpeggio direction: %s", token)
			return
		}

		log.Tracef("direction: %d", direction)
		for j := 0; j < steps; j++ {
			if !(i == 0 && j == 0) {
				notei += direction
			}
			for notei < 0 {
				notei += len(originalNotes)
				octave--
			}
			for notei >= len(originalNotes) {
				notei -= len(originalNotes)
				octave++
			}
			newNote := originalNotes[notei].Add(octave * 12)
			log.Tracef("notei: %d (%s)", notei, newNote.NameSharp)
			notes = append(notes, newNote)
		}
	}

	noteString := ""
	for _, note := range notes {
		noteString += modifier.Merge(note.NameSharp, modified.Modifiers) + " "
	}
	result = strings.TrimSpace(noteString)

	log.Tracef("result: %s", modified)

	return
}

func tokenizeLetterNumberes(s string) (tokens []string) {
	// take a string with letters and numbers and split it into tokens
	// where the letter determines the start of a new token
	for i := 0; i < len(s); i++ {
		if i == 0 {
			tokens = append(tokens, string(s[i]))
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			tokens[len(tokens)-1] += string(s[i])
		} else {
			tokens = append(tokens, string(s[i]))
		}
	}
	return
}
